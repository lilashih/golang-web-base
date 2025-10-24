import Section from "./core.template.js";

const template = /*html*/ `
<Section :is-creatable="true" :is-editable="true" :is-deletable="true" :is-orderable="true" :form-default-data="formDefaultData">
  <template #column>
    <el-table-column prop="name" label="名稱" sortable></el-table-column>
  </template>
  
  <template #form="{ form, errors }">
    <el-form-item label="名稱" :error="errors.name">
      <el-input v-model="form.name" placeholder="請輸入使用者名稱" maxlength="50" />
    </el-form-item>
  </template>
</Section>
`;

export default {
    components: {
        Section,
    },
    setup() {
        const formDefaultData = ref({ name: "" });

        return {
            formDefaultData,
        };
    },
    template,
};
