<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item label="重要性" prop="importance">
      <el-rate v-model="formData.importance" />
      <el-button size="mini" @click="formData.importance=0">重置</el-button>
    </el-form-item>

    <el-form-item label="图片标题" prop="title" required>
      <el-input v-model="formData.title" placeholder="请输入图片标题" />
    </el-form-item>
    <el-form-item label="滚动图片（尺寸16:9）" prop="image" required>
      <image-uploader v-model="formData.image" dir="news" />
    </el-form-item>

    <el-form-item label="状态" prop="state" required>
      <news-state-selector v-model="formData.state" />
    </el-form-item>
  </el-form>
</template>

<script>
import ImageUploader from "@/components/form/ImageUploader";
import NewsStateSelector from "@/components/form/NewsStateSelector";
import formMixin from "./formMixin";
export default {
  name: "CarouselForm",
  components: {
    ImageUploader,
    NewsStateSelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        title: [{ required: true, message: "请输入标题", trigger: "blur" }],
        image: [{ required: true, message: "请上传图片", trigger: "blur" }],
        state: [{ required: true, message: "请选择状态", trigger: "blur" }]
      }
    };
  }
};
</script>