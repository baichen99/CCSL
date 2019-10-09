<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item label="手形名称" prop="name" required>
      <el-input v-model="formData.name" placeholder="请输入手形名称" />
    </el-form-item>

    <el-form-item label="手形图片（SVG）" prop="image" required>
      <image-uploader v-model="formData.image" type="svg" dir="signs" height="300px" />
    </el-form-item>
  </el-form>
</template>


<script>
import ImageUploader from "@/components/form/ImageUploader";
import formMixin from "./formMixin";
export default {
  name: "SignForm",
  components: {
    ImageUploader
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        name: [{ required: true, message: "请输入手形名称" }],
        image: [{ required: true, message: "请上传手形图片" }]
      }
    };
  }
};
</script>