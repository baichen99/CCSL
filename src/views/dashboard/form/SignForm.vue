<template>
  <el-form
    label-position="left"
    label-width="100px"
    :model="formData"
  >
    <el-form-item
      label="手形名称"
      required
    >
      <el-input
        v-model="formData.name"
        placeholder="请输入手形名称"
      />
    </el-form-item>
    <el-form-item
      label="手形图片（SVG）"
      required
    >
      <image-uploader
        v-model="formData.image"
        type="svg"
        dir="signs"
      />
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
  mixins: [formMixin]
};
</script>