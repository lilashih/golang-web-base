import helper from "../../js/helper/helper.js";
import notify from "../../js/helper/notify.js";

const template = /*html*/ `
<el-container>
  <el-header>
    <h1 v-text="title"></h1>
  </el-header>
  <el-main class="content">
    <div class="setting-btn-group">
      <el-button v-if="!isEditing" @click="handleFormOpen()">編輯</el-button>

      <div v-if="isEditing">
        <el-button type="primary" @click="handleFormSave()">儲存</el-button>
        <el-button @click="handleFormClose()">取消</el-button>
      </div>
    </div>

    <div>
      <el-table
        class="w-full"
        stripe
        border
        :data="items"
      >

        <el-table-column prop="name" label="名稱" width="300"></el-table-column>

        <el-table-column label="設定" v-if="!isEditing" #default="scope">
          <span v-if="scope.row.option">{{ scope.row.option[scope.row.value] }}</span>
          <span v-else-if="scope.row.type === 'password'">{{ maskPassword(scope.row.value) }}</span>
          <template v-else>
            <span>{{ scope.row.value }}</span>
          </template>
        </el-table-column>

        <el-table-column label="設定" v-if="isEditing" #default="scope">
          <el-input v-if="scope.row.type == 'text'" v-model="scope.row.value" placeholder="輸入"/>
          <el-input v-if="scope.row.type == 'password'" type="password" v-model="scope.row.value" placeholder="輸入"/>
          <el-select v-if="scope.row.type == 'select'" v-model="scope.row.value" placeholder="選擇">
            <el-option v-for="(o, key) in scope.row.option" :key="key" :value="key" :label="o"></el-option>
          </el-select>
        </el-table-column>

      </el-table>
    </div>

  </el-main>

</el-container>
`;

export default {
    setup() {
        const route = useRoute();
        const api = route.meta.api;

        // 列表
        const items = ref([]);

        // 表單
        const isEditing = ref(false);
        const errors = ref({});

        const maskPassword = (text) => {
            return "*".repeat(text.length);
        };

        const handleTableData = async () => {
            const response = await helper.api("get", api);
            const data = response.data;

            items.value = data.settings;
        };

        const handleFormOpen = () => {
            isEditing.value = true;
        };

        const handleFormClose = () => {
            isEditing.value = false;
            handleTableData();
        };

        const handleFormSave = () => {
            helper
                .api("put", api, Object.values(items.value))
                .then((res) => {
                    notify.success(res.message);
                    handleFormClose();
                })
                .catch((err) => {
                    if (err.code == 422) {
                        errors.value = err.data.errors;
                    }
                });
        };

        handleTableData();

        return {
            title: route.meta.title,
            items,
            isEditing,
            errors,

            handleFormOpen,
            handleFormClose,
            handleFormSave,

            maskPassword,
        };
    },
    template,
};
