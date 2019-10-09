<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item label="重要性" prop="importance">
      <el-rate v-model="formData.importance" />
      <el-button size="mini" @click="formData.importance=0">重置</el-button>
    </el-form-item>

    <el-form-item label="新闻语言" prop="language">
      <language-selector v-model="formData.language" placeholder="请选择新闻语言" />
    </el-form-item>

    <el-form-item label="新闻标题" prop="title" required>
      <el-input v-model="formData.title" placeholder="请输入新闻标题" clearable />
    </el-form-item>

    <el-form-item label="新闻状态" prop="state" required>
      <news-state-selector v-model="formData.state" />
    </el-form-item>

    <el-form-item label="日期" prop="date" required>
      <date-picker v-model="formData.date" />
    </el-form-item>

    <el-form-item label="新闻类型" prop="type" required>
      <news-type-selector v-model="formData.type" />
    </el-form-item>

    <el-form-item label="新闻栏目" prop="column" required>
      <news-column-selector v-model="formData.column" />
    </el-form-item>

    <el-form-item v-if="formData.type==='document'" label="新闻内容" required>
      <rich-text-editor v-model="formData.text" />
    </el-form-item>

    <el-form-item v-if="formData.type==='link'" label="新闻链接" prop="text" required>
      <el-input v-model="formData.text" placeholder="请输入新闻链接" />
    </el-form-item>
  </el-form>
</template>

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
  // watch: {
  //   "formData.type"() {
  //     if (this.formData.text) {
  //       this.formData.text = "";
  //     }
  //   }
  // }
};
</script>