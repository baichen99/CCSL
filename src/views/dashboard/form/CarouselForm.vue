<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Importance')" prop="importance">
      <el-rate v-model="formData.importance" />
      <el-button size="mini" @click="formData.importance=0">{{ $t("Reset") }}</el-button>
    </el-form-item>

    <el-form-item :label="$t('TitleZh')" prop="titleZh" required>
      <el-input v-model="formData.titleZh" :placeholder="$t('titleZhTip')" />
    </el-form-item>

    <el-form-item :label="$t('TitleEn')" prop="titleEn" required>
      <el-input v-model="formData.titleEn" :placeholder="$t('titleEnTip')" />
    </el-form-item>

    <el-form-item :label="$t('Carousel')" prop="image" required>
      <image-uploader v-model="formData.image" dir="news" />
    </el-form-item>

    <el-form-item :label="$t('State')" prop="state" required>
      <news-state-selector v-model="formData.state" />
    </el-form-item>
  </el-form>
</template>

<i18n>
{
  "zh-CN": {
    "titleZhTip":"请输入中文图片标题",
    "titleEnTip": "请输入英文图片标题"
  },
  "en-US": {
    "titleZhTip":"Input image title(Chinese)",
    "titleEnTip": "Input image title(English)"
  }
}
</i18n>

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
        titleZh: [
          { required: true, message: "请输入中文标题", trigger: "blur" }
        ],
        titleEn: [
          { required: true, message: "请输入英文标题", trigger: "blur" }
        ],
        image: [{ required: true, message: "请上传图片", trigger: "blur" }],
        state: [{ required: true, message: "请选择状态", trigger: "blur" }]
      }
    };
  }
};
</script>