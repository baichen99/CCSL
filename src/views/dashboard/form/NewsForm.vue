<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Importance')" prop="importance">
      <el-rate v-model="formData.importance" />
      <el-button size="mini" @click="formData.importance=0">{{ $t("Reset") }}</el-button>
    </el-form-item>

    <el-form-item :label="$t('Language')" prop="language">
      <language-selector v-model="formData.language" />
    </el-form-item>

    <el-form-item :label="$t('Title')" prop="title">
      <el-input v-model="formData.title" :placeholder="$t('titleTip')" clearable />
    </el-form-item>

    <el-form-item :label="$t('State')" prop="state">
      <news-state-selector v-model="formData.state" />
    </el-form-item>

    <el-form-item :label="$t('Date')" prop="date">
      <date-picker v-model="formData.date" />
    </el-form-item>

    <el-form-item :label="$t('Column')" prop="column">
      <news-column-selector v-model="formData.column" />
    </el-form-item>

    <el-form-item :label="$t('Type')" prop="type">
      <news-type-selector v-model="formData.type" />
    </el-form-item>

    <el-form-item v-if="formData.type==='document'" :label="$t('Content')">
      <rich-text-editor v-model="formData.text" />
    </el-form-item>

    <el-form-item v-if="formData.type==='link'" :label="$t('Link')" prop="text">
      <el-input v-model="formData.text" :placeholder="$t('linkTip')" />
    </el-form-item>
  </el-form>
</template>

<i18n>
{
  "zh-CN": {
    "titleTip":"请输入新闻标题",
    "linkTip": "请输入新闻链接"
  },
  "en-US": {
    "titleTip":"Input news title",
    "linkTip": "Input news link"
  }
}
</i18n>

<script>
import NewsStateSelector from "@/components/form/NewsStateSelector";
import NewsTypeSelector from "@/components/form/NewsTypeSelector";
import NewsColumnSelector from "@/components/form/NewsColumnSelector";
import LanguageSelector from "@/components/form/LanguageSelector";
import RichTextEditor from "@/components/form/RichTextEditor";
import DatePicker from "@/components/form/DatePicker";

import formMixin from "./formMixin";
export default {
  name: "NewsForm",
  components: {
    NewsStateSelector,
    NewsTypeSelector,
    NewsColumnSelector,
    LanguageSelector,
    RichTextEditor,
    DatePicker
  },
  mixins: [formMixin],
  data() {
    const that = this;
    return {
      rules: {
        title: [{ required: true, message: "请输入新闻标题", trigger: "blur" }],
        state: [{ required: true, message: "请选择状态", trigger: "blur" }],
        date: [{ required: true, message: "请选择日期", trigger: "blur" }],
        column: [{ required: true, message: "请选择栏目", trigger: "blur" }],
        type: [{ required: true, message: "请选择类型", trigger: "blur" }],
        text: [
          { required: true, message: "请输入新闻内容", trigger: "blur" },
          {
            validator(rule, value, callback) {
              if (
                that.formData.type === "link" &&
                !new RegExp(
                  /http(s)?:\/\/([\w-]+\.)+[\w-]+(\/[\w- .\/?%&=]*)?/
                ).test(value)
              ) {
                callback(
                  new Error(
                    "请输入有效的链接（HTTP链接应该以http://或者https://开头）"
                  )
                );
              } else {
                callback();
              }
            }
          }
        ]
      }
    };
  }
};
</script>