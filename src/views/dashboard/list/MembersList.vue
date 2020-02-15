<template>
  <list-view
    :get-list-method="GetMembersList"
    :create-item-method="CreateMember"
    :update-item-method="UpdateMember"
    :delete-item-method="DeleteMember"
    :list-form-component="MemberForm"
    :columns="columns"
  >
    <template #toolbar-search="{params, handleSearch}">
      <search-input v-model="params.nameZh" :placeholder="$t('tipName')" @update="handleSearch" />
    </template>
  </list-view>
</template>
<i18n>
{
  "zh-CN":{
    "tipName":"请输入姓名"
  },
  "en-US":{
    "tipName":"Input name"
  }
}
</i18n>
<script>
import { mapGetters } from "vuex";
import {
  GetMembersList,
  CreateMember,
  UpdateMember,
  DeleteMember
} from "@/api/members";
import ListView from "@/components/ListView";
import SearchInput from "@/components/form/SearchInput";
import MemberForm from "@/views/dashboard/form/MemberForm";
export default {
  name: "MembersList",
  components: {
    ListView,
    SearchInput
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
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["memberTypes"])
  }
};
</script>

