<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Name')" prop="name">
      <el-input v-model="formData.name" :placeholder="$t('tipName')" />
    </el-form-item>

    <el-form-item :label="$t('UserRole')" prop="roles">
      <multiple-selector v-model="formData.roles" :options="userRoles" />
    </el-form-item>

    <el-form-item :label="$t('Account')" prop="username">
      <el-input v-model="formData.username" :placeholder="$t('tipAccount')" />
    </el-form-item>

    <el-form-item v-if="mode==='create'" :label="$t('Password')" prop="password">
      <el-input v-model="formData.password" type="password" :placeholder="$t('tipPassword')" />
    </el-form-item>

    <el-form-item :label="$t('AccountState')" prop="state">
      <simple-selector v-model="formData.state" :options="userState" />
    </el-form-item>
  </el-form>
</template>

<i18n>
{
  "zh-CN": {
    "tipAccount": "请输入学工号或者邮箱",
    "tipName": "请输入姓名",
    "tipPassword": "请输入密码（校内用户请留空）"
  },
  "en-US": {
    "tipAccount": "Input account",
    "tipName": "Input name",
    "tipPassword": "Input password(Leave it empty if is SHU user)"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import MultipleSelector from "@/components/form/MultipleSelector";
import SimpleSelector from "@/components/form/SimpleSelector";
import formMixin from "./formMixin";
export default {
  name: "UserForm",
  components: {
    MultipleSelector,
    SimpleSelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        name: [{ required: true, message: "请输入用户姓名", trigger: "blur" }],
        roles: [{ required: true, message: "请选择用户角色", trigger: "blur" }],
        username: [
          {
            required: true,
            message: "请输入用户学工号或者邮箱",
            trigger: "blur"
          },
          {
            pattern: /(^[0-9]{8}$)|(^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$)/,
            message: "请检查学工号或者邮箱",
            trigger: "blur"
          }
        ]
      }
    };
  },
  computed: {
    ...mapGetters(["userRoles", "userState"])
  }
};
</script>