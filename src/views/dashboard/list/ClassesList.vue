<template>
  <list-view
    :get-list-method="GetClassesList"
    :create-item-method="CreateClass"
    :update-item-method="UpdateClass"
    :delete-item-method="DeleteClass"
    :list-form-component="ClassForm"
    :columns="columns"
  >
    <template #toolbar-search="{params, handleSearch}">
      <search-input v-model="params.name" :placeholder="$t('tip')" @update="handleSearch" />
    </template>
  </list-view>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请输入班级名称"
  },
  "en-US": {
    "tip": "Input class name"
  }
}
</i18n>

<script>
import ListView from "@/components/ListView";
import SearchInput from "@/components/form/SearchInput";
import ClassForm from "@/views/dashboard/form/ClassForm";
import {
  GetClassesList,
  CreateClass,
  UpdateClass,
  DeleteClass
} from "@/api/classes";
export default {
  name: "ClassesList",
  components: {
    ListView,
    SearchInput
  },
  data() {
    return {
      GetClassesList,
      CreateClass,
      UpdateClass,
      DeleteClass,
      ClassForm,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          prop: "updatedAt",
          label: this.$t("UpdatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.updatedAt), "long")
        },
        {
          prop: "name",
          label: this.$t("ClassName"),
          width: "300px",
          hideOverflow: true
        }
      ]
    };
  }
};
</script>