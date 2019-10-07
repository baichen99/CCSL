<template>
  <el-form
    label-position="left"
    label-width="100px"
    :model="formData"
  >

    <el-form-item
      label="重要性"
      required
    >
      <el-rate v-model="formData.importance" />
      <el-button
        size="mini"
        @click="formData.importance=0"
      >重置</el-button>
    </el-form-item>

    <el-form-item
      label="新闻标题"
      required
    >
      <el-input
        v-model="formData.title"
        placeholder="请输入新闻标题"
      />
    </el-form-item>

    <el-form-item
      label="状态"
      required
    >
      <news-state-selector v-model="formData.state" />
    </el-form-item>

    <el-form-item
      label="日期"
      required
    >
      <date-picker v-model="formData.date" />
    </el-form-item>

    <el-form-item
      label="新闻类型"
      required
    >
      <news-type-selector v-model="formData.type" />
    </el-form-item>

    <el-form-item
      label="新闻栏目"
      required
    >
      <news-column-selector v-model="formData.column" />
    </el-form-item>

    <el-form-item
      label="新闻内容"
      required
    >
      <rich-text-editor />
    </el-form-item>

  </el-form>
</template>

<script>
import NewsStateSelector from "@/components/form/NewsStateSelector";
import NewsTypeSelector from "@/components/form/NewsTypeSelector";
import NewsColumnSelector from "@/components/form/NewsColumnSelector";
import RichTextEditor from "@/components/form/RichTextEditor";
import DatePicker from "@/components/form/DatePicker";

import formMixin from "./formMixin";
export default {
  name: "NewsForm",
  components: {
    NewsStateSelector,
    NewsTypeSelector,
    NewsColumnSelector,
    RichTextEditor,
    DatePicker
  },
  mixins: [formMixin]
};
</script>