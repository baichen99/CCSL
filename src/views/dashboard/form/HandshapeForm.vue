<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('HandshapeName')" prop="name">
      <el-input v-model="formData.name" placeholder="请输入手形名称" />
    </el-form-item>

    <el-form-item :label="$t('HandshapeGlyph')" prop="glyph">
      <el-input v-model="formData.glyph" maxlength="1" placeholder="请输入手形字体符号" />
    </el-form-item>

    <el-form-item :label="$t('HandshapeImage')+'（SVG）'" prop="image">
      <image-uploader v-model="formData.image" type="svg" dir="handshapes" height="300px" />
    </el-form-item>
  </el-form>
</template>


<script>
import ImageUploader from "@/components/form/ImageUploader";
import formMixin from "./formMixin";
export default {
  name: "HandshapeForm",
  components: {
    ImageUploader
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        name: [{ required: true, message: "请输入手形名称" }],
        image: [{ required: true, message: "请上传手形图片" }],
        glyph: [{ required: true, message: "请输入手形字体符号" }]
      }
    };
  }
};
</script>