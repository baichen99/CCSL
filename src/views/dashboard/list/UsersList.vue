<template>
  <list-view
    :get-list-method="GetUsersList"
    :create-item-method="CreateUser"
    :update-item-method="UpdateUser"
    :delete-item-method="DeleteUser"
    :list-form-component="UserForm"
    :columns="columns"
  >
    <template #toolbar-button="{selection,handleUpdateItems}">
      <el-button
        :disabled="!showButton(selection,'inactive')"
        size="mini"
        type="success"
        plain
        @click="handleUpdateItems({state:'active'})"
      >{{ $t("Enable") }}</el-button>
      <el-button
        :disabled="!showButton(selection,'active')"
        size="mini"
        type="warning"
        plain
        @click="handleUpdateItems({state:'inactive'})"
      >{{ $t("Disable") }}</el-button>
    </template>

    <template #toolbar-search="{params,handleSearch}">
      <search-input
        v-model="params.username"
        :placeholder="$t('tipAccount')"
        @update="handleSearch"
      />
      <search-input v-model="params.name" :placeholder="$t('tipName')" @update="handleSearch" />
    </template>

    <template #state="{row}">
      <el-tag
        size="small"
        :disable-transitions="true"
        :type="$options.filters.getObjectItem(userState, row.state).color"
      >{{ $t($options.filters.getObjectItem(userState, row.state).text) }}</el-tag>
    </template>
    <template #roles="{row}">
      <el-tag
        v-for="item in row.roles"
        :key="item"
        size="small"
        :disable-transitions="true"
        :type="$options.filters.getObjectItem(userRoles, item).color"
      >{{ $t($options.filters.getObjectItem(userRoles, item).text) }}</el-tag>
    </template>
  </list-view>
</template>

<i18n>
{
  "zh-CN": {
    "tipAccount": "请输入账号",
    "tipName": "请输入姓名"
  },
  "en-US": {
    "tipAccount": "Input account",
    "tipName": "Input name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import { GetUsersList, CreateUser, UpdateUser, DeleteUser } from "@/api/users";
import SearchInput from "@/components/form/SearchInput";
import UserForm from "@/views/dashboard/form/UserForm";
import ListView from "@/components/ListView";
export default {
  name: "UsersList",
  components: {
    ListView,
    SearchInput
  },
  data() {
    return {
      GetUsersList,
      CreateUser,
      UpdateUser,
      DeleteUser,
      UserForm,
      columns: [
        {
          prop: "createdAt",
          label: this.$t("CreatedAt"),
          width: "180px",
          formatter: row => this.$d(new Date(row.createdAt), "long")
        },
        {
          slot: "state",
          label: this.$t("AccountState"),
          width: "100px",
          filters: [
            { text: this.$t("Active"), value: "active" },
            { text: this.$t("Inactive"), value: "inactive" }
          ]
        },
        {
          slot: "roles",
          label: this.$t("UserRole"),
          width: "200px",
          filters: []
        },
        {
          prop: "username",
          label: this.$t("Account"),
          width: "150px"
        },
        {
          prop: "name",
          label: this.$t("Name"),
          width: "100px"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["userRoles", "userState"])
  },
  methods: {
    showButton(selection, state) {
      if (selection.length > 0) {
        const allowEnable = selection.filter(item => item.state === state);
        return allowEnable.length === selection.length;
      } else {
        return false;
      }
    }
  }
};
</script>


<style lang="scss" scoped>
.el-tag {
  margin: 3px;
}
</style>