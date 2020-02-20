<template>
  <list-view
    :get-list-method="GetClassesList"
    :create-item-method="CreateClass"
    :update-item-method="UpdateClass"
    :delete-item-method="DeleteClass"
    :list-form-component="ClassForm"
    :columns="columns"
  ></list-view>
</template>

<script>
import ListView from "@/components/ListView";
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
    ListView
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
        }
      ]
    };
  }
};
</script>