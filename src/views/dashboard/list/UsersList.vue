<template>
  <list-view
    :get-list-method="GetUsersList"
    :create-item-method="CreateUser"
    :update-item-method="UpdateUser"
    :delete-item-method="DeleteUser"
    :list-form-component="UserForm"
    :columns="columns"
  >
    <template #toolbar="{params,handleSearch}">
      <el-input
        v-model="params.username"
        :placeholder="$t('tipAccount')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <el-input
        v-model="params.name"
        :placeholder="$t('tipName')"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
    </template>
    <template #state="{row}">
      <el-tag
        size="small"
        :disable-transitions="true"
        :type="userState[row.state].color"
      >{{ $t(userState[row.state].name) }}</el-tag>
    </template>
    <template #roles="{row}">
      <el-tag
        v-for="item in row.roles"
        :key="item"
        size="small"
        :disable-transitions="true"
        :type="userRoles[item].color"
      >{{ $t(userRoles[item].name) }}</el-tag>
    </template>
    <template #action="{row,handleUpdateItem, handleDeleteItem}">
      <el-button
        v-if="row.state==='inactive'"
        :disabled="isSuperUser(row)"
        type="success"
        size="mini"
        plain
        @click.stop="handleUpdateItem(row.id, {
          state: 'active'
        })"
      >{{ $t("Enable") }}</el-button>
      <el-button
        v-if="row.state==='active'"
        :disabled="isSuperUser(row)"
        type="warning"
        size="mini"
        plain
        @click.stop="handleUpdateItem(row.id, {
          state: 'inactive'
        })"
      >{{ $t("Disable") }}</el-button>
      <el-button
        :disabled="isSuperUser(row)"
        type="danger"
        size="mini"
        plain
        @click.stop="handleDeleteItem(row.id)"
      >{{ $t("Delete") }}</el-button>
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
import UserForm from "@/views/dashboard/form/UserForm";
import ListView from "@/components/ListView";
export default {
  name: "UsersList",
  components: {
    ListView
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
          width: "150px",
          filters: [
            { text: this.$t("Active"), value: "active" },
            { text: this.$t("Inactive"), value: "inactive" }
          ]
        },
        {
          slot: "roles",
          label: this.$t("UserRole"),
          width: "150px",
          filters: [
            { text: this.$t("SuperAdmin"), value: "super" },
            { text: this.$t("Admin"), value: "admin" },
            { text: this.$t("Student"), value: "student" },
            { text: this.$t("Teacher"), value: "teacher" },
            { text: this.$t("DatabaseUser"), value: "dbuser" }
          ]
        },
        {
          prop: "username",
          label: this.$t("Account"),
          width: "200px"
        },
        {
          prop: "name",
          label: this.$t("Name"),
          width: "180px"
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "180px",
          fixed: "right"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["userRoles", "userState"])
  },
  methods: {
    isSuperUser(row) {
      return row.roles.includes("super");
    }
  }
};
</script>


<style lang="scss" scoped>
.el-tag {
  margin: 3px;
}
</style>