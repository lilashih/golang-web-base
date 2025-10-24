import helper from "../../js/helper/helper.js";
import notify from "../../js/helper/notify.js";
import dialog from "../../js/helper/dialog.js";

const template = /*html*/ `
<el-container>
  <el-header>
    <h1 v-text="title"></h1>
  </el-header>
  <el-main class="content">
    <el-table
      class="w-full"
      stripe
      border
      :data="items"
    >
      <el-table-column width="90" align="center" v-if="isOrderable" #default="scope">
        <el-button class="order-btn" type="warning" icon="ArrowUp" circle title="調整排序" @click="handleOrder(scope.$index, 'up')"></el-button>
        <el-button class="order-btn" type="warning" icon="ArrowDown" circle title="調整排序" @click="handleOrder(scope.$index, 'down')"></el-button>
      </el-table-column>

      <el-table-column v-if="isShowIdColumn" prop="id" label="ID" sortable width="100"></el-table-column>

      <slot name="column"></slot>

      <el-table-column width="135" align="center" v-if="isCreatable || isEditable || isDeletable">
        <template #header>
          <el-button  v-if="isCreatable" size="small" type="primary" @click="handleFormOpen('create', formDefault)">
            新增
          </el-button>
        </template>
        <template #default="scope">
          <el-button size="small" v-if="isEditable" @click="handleFormOpen('edit', scope.row)">編輯</el-button>
          <el-button size="small" v-if="isDeletable" type="danger" @click="handleDelete(scope.row)">刪除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div class="pagination-container">
      <el-pagination
        background
        layout="prev, pager, next, jumper"
        :hide-on-single-page="true"
        :total="pagination.total"
        :page-size="pagination.perPage"
        :current-page="pagination.currentPage"
        @current-change="handlePageChange"
      >
      </el-pagination>
    </div>
  </el-main>

  <el-dialog v-model="isFormVisible" :title="isCreating ? '新增' : '編輯'" width="700" :before-close="handleFormClose">
    <el-form label-position="right" label-width="25%">
      <el-form-item v-if="form.id > 0" label="ID">
        <el-input v-model="form.id" disabled />
      </el-form-item>
      <slot name="form" v-bind="{ form, errors, isCreating }"></slot>
    </el-form>

    <template #footer>
      <div class="dialog-footer">
        <el-button @click="handleFormClose">取消</el-button>
        <el-button type="primary" @click="handleFormSave">儲存</el-button>
      </div>
    </template>
  </el-dialog>
</el-container>
`;

export default {
    props: {
        // 表單預設資料
        formDefaultData: {
            type: Object,
            default: {},
        },
        // 主鍵欄位，預設為id
        pkColumn: {
            type: String,
            default: "id",
        },
        // 主鍵欄位名稱，預設為ID
        pkColumnName: {
            type: String,
            default: "ID",
        },
        // 是否顯示ID欄位
        isShowIdColumn: {
            type: Boolean,
            default: true,
        },
        // 是否可以編輯排序
        isOrderable: {
            type: Boolean,
            default: false,
        },
        // 是否可以新增
        isCreatable: {
            type: Boolean,
            default: false,
        },
        // 是否可以編輯
        isEditable: {
            type: Boolean,
            default: false,
        },
        // 是否可以刪除
        isDeletable: {
            type: Boolean,
            default: false,
        },
    },
    setup(props) {
        const route = useRoute();
        const api = route.meta.api;

        // 列表
        const items = ref([]);
        const pagination = ref({});

        // 表單
        const isFormVisible = ref(false);
        const formDefault = ref({ ...props.formDefaultData });
        const form = ref({ ...formDefault.value });
        const errors = ref({});
        const isCreating = ref(true);

        const handleTableData = async (page = 1) => {
            const response = await helper.api("get", api, { page });
            const data = response.data;

            items.value = data[api];
            pagination.value = data.pagination;
        };

        const handlePageChange = (page) => {
            handleTableData(page);
            helper.scrollToTop();
        };

        const handleFormOpen = (type, row) => {
            isCreating.value = type === "create";
            isFormVisible.value = true;
            form.value = { ...row };
            errors.value = {};
        };

        const handleFormClose = () => {
            isFormVisible.value = false;
        };

        const handleFormSave = () => {
            const method = isCreating.value ? `post` : `put`;
            const url = isCreating.value ? api : `${api}/${form.value[props.pkColumn]}`;

            helper
                .api(method, url, { ...form.value })
                .then((res) => {
                    notify.success(res.message);
                    handleTableData(pagination.value.page);
                    handleFormClose();
                })
                .catch((err) => {
                    if (err.code == 422) {
                        errors.value = err.data.errors;
                    }
                });
        };

        const handleDelete = (row) => {
            dialog
                .deleteConfirm(`確定刪除 ${props.pkColumnName} 為 ${row[props.pkColumn]} 的項目`)
                .then(() => helper.api("delete", `${api}/${row[props.pkColumn]}`))
                .then((res) => {
                    notify.success(res.message);
                    handleTableData(pagination.value.page);
                })
                .catch(() => {});
        };

        const handleOrder = (index, direction) => {
            let index2 = -1;

            if (direction === "up" && index > 0) {
                index2 = index - 1;
            } else if (direction === "down" && index < items.value.length - 1) {
                index2 = index + 1;
            }

            if (index2 > -1) {
                helper
                    .api(`post`, `${api}/order`, { id1: items.value[index].id, id2: items.value[index2].id })
                    .then((res) => {
                        notify.success(res.message);
                        [items.value[index2], items.value[index]] = [items.value[index], items.value[index2]];
                    })
                    .catch((err) => {
                        if (err.code == 422) {
                            notify.error(err.data.errors);
                        } else {
                            notify.error(err.message);
                        }
                    });
            }
        };

        handleTableData();

        return {
            title: `${route.meta.header} ${route.meta.description}`,
            items,
            formDefault,
            form,
            errors,
            pagination,
            isFormVisible,
            isCreating,
            handlePageChange,

            handleFormOpen,
            handleFormClose,
            handleFormSave,

            handleDelete,

            handleOrder,
        };
    },
    template,
};
