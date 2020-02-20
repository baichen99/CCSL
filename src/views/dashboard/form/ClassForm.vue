<template>
  <el-form ref="form" label-position="left" :model="formData" :rules="rules">
    <el-form-item :label="$t('ClassName')" prop="name">
      <el-input v-model="formData.name" :placeholder="$t('nameTip')" clearable />
    </el-form-item>
    <el-tabs type="border-card">
      <el-tab-pane label="班级详情">
        <el-form-item>
          <rich-text-editor v-model="formData.details" height="500" dir="classes" />
        </el-form-item>
      </el-tab-pane>
      <el-tab-pane label="班级资源">
        <el-form-item>
          <rich-text-editor v-model="formData.resources" height="500" dir="classes" />
        </el-form-item>
      </el-tab-pane>
    </el-tabs>
  </el-form>
</template>

<i18n>
{
    "zh-CN": {
        "nameTip":"请输入班级名称"
    },
    "en-US": {
        "nameTip":"Input class name"
    }
}
</i18n>

<script>
import formMixin from "./formMixin";
import RichTextEditor from "@/components/form/RichTextEditor";
export default {
  name: "ClassForm",
  components: {
    RichTextEditor
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {}
    };
  }
};
</script>