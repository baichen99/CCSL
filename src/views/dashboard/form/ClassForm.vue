<template>
  <el-form ref="form" v-loading="loading" :model="formData" :rules="rules">
    <el-form-item :label="$t('ClassName')" prop="name">
      <el-input v-model="formData.name" :placeholder="$t('nameTip')" clearable />
    </el-form-item>
    <el-tabs type="border-card">
      <el-tab-pane :label="$t('ClassDetails')">
        <el-form-item>
          <rich-text-editor v-model="formData.details" height="600" dir="classes" />
        </el-form-item>
      </el-tab-pane>
      <el-tab-pane :label="$t('ClassResources')">
        <el-form-item>
          <rich-text-editor v-model="formData.resources" height="600" dir="classes" />
        </el-form-item>
      </el-tab-pane>
      <el-tab-pane v-if="mode==='edit'" label="教师管理"></el-tab-pane>
      <el-tab-pane v-if="mode==='edit'" label="学生管理"></el-tab-pane>
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
import { GetClass } from "@/api/classes";
export default {
  name: "ClassForm",
  components: {
    RichTextEditor
  },
  mixins: [formMixin],

  props: {
    mode: {
      type: String,
      default: () => ""
    }
  },
  data() {
    return {
      rules: {
        name: [{ required: true, message: "请输入班级名称" }]
      },
      loading: false,
      teachers: [],
      students: []
    };
  },
  created() {
    this.$nextTick(() => {
      if (this.mode === "edit") {
        this.getData();
      }
    });
  },
  methods: {
    async getData() {
      const id = this.formData.id;
      this.loading = true;
      try {
        const res = await GetClass(id);
        this.teachers = res.data.teachers;
        this.students = res.data.students;
        console.log(res.data);
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    }
  }
};
</script>