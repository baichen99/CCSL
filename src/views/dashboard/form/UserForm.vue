<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Name')" prop="name">
      <el-input v-model="formData.name" :placeholder="$t('tipName')" />
    </el-form-item>

    <el-form-item :label="$t('UserRole')" prop="userType">
      <user-type-selector v-model="formData.userType" />
    </el-form-item>

    <el-form-item :label="$t('Account')" prop="username">
      <el-input v-model="formData.username" :placeholder="$t('tipAccount')" />
    </el-form-item>

    <el-form-item :label="$t('AccountState')" prop="state">
      <user-state-selector v-model="formData.state" />
    </el-form-item>
  </el-form>
</template>

<i18n>
{
  "zh-CN": {
    "tipAccount": "请输入学号或者工号",
    "tipName": "请输入姓名"
  },
  "en-US": {
    "tipAccount": "Input account",
    "tipName": "Input name"
  }
}
</i18n>

<script>
import UserTypeSelector from "@/components/form/UserTypeSelector";
import UserStateSelector from "@/components/form/UserStateSelector";
import formMixin from "./formMixin";
export default {
  name: "UserForm",
  components: {
    UserTypeSelector,
    UserStateSelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        name: [{ required: true, message: "请输入用户姓名", trigger: "blur" }],
        userType: [
          { required: true, message: "请选择用户角色", trigger: "blur" },
          { type: "enum", enum: ["super", "admin", "user", "learner"] }
        ],
        username: [
          {
            required: true,
            message: "请输入用户学号或者工号",
            trigger: "blur"
          },
          {
            pattern: /^[0-9]{8}$/,
            message: "请检查学工号是否正确（应为8位数字）",
            trigger: "blur"
          }
        ]
      }
    };
  }
};
</script>