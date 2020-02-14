<template>
  <list-view
    :get-list-method="GetMembersList"
    :create-item-method="CreateMember"
    :update-item-method="UpdateMember"
    :delete-item-method="DeleteMember"
    :list-form-component="MemberForm"
    :columns="columns"
  >
    <template #action="{row,handleDeleteItem}">
      <el-button
        type="danger"
        size="mini"
        plain
        @click.stop="handleDeleteItem(row.id)"
      >{{ $t("Delete") }}</el-button>
    </template>
  </list-view>
</template>

<script>
import { mapGetters } from "vuex";
import {
  GetMembersList,
  CreateMember,
  UpdateMember,
  DeleteMember
} from "@/api/members";
import ListView from "@/components/ListView";
import MemberForm from "@/views/dashboard/form/MemberForm";
export default {
  name: "MembersList",
  components: {
    ListView
  },
  data() {
    return {
      GetMembersList,
      CreateMember,
      UpdateMember,
      DeleteMember,
      MemberForm,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          prop: "name",
          label: this.$t("Name"),
          width: "150px",
          formatter: row =>
            this.$i18n.locale === "en-US" ? row.nameEn : row.nameZh
        },
        {
          prop: "type",
          label: this.$t("Type"),
          width: "200px",
          formatter: row => this.$t(this.memberTypes[row.type].name)
        },
        {
          prop: "employer",
          label: this.$t("Employer"),
          width: "300px",
          hideOverflow: true,
          formatter: row =>
            this.$i18n.locale === "en-US" ? row.employerEn : row.employerZh
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "90px",
          fixed: "right"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["memberTypes"])
  }
};
</script>

