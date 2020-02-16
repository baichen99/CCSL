<template>
  <el-form ref="form" label-position="left" label-width="150px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Importance')" prop="importance">
      <el-rate v-model="formData.importance" />
      <el-button size="mini" @click="formData.importance=0">{{ $t("Reset") }}</el-button>
    </el-form-item>

    <el-form-item :label="$t('TitleZh')" prop="titleZh">
      <el-input v-model="formData.titleZh" :placeholder="$t('titleZhTip')" />
    </el-form-item>

    <el-form-item :label="$t('TitleEn')" prop="titleEn">
      <el-input v-model="formData.titleEn" :placeholder="$t('titleEnTip')" />
    </el-form-item>

    <el-form-item :label="$t('Carousel')" prop="image">
      <image-uploader v-model="formData.image" dir="news" />
    </el-form-item>

    <el-form-item :label="$t('State')" prop="state">
      <simple-selector v-model="formData.state" :options="newsState" />
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
import { mapGetters } from "vuex";
import ImageUploader from "@/components/form/ImageUploader";
import SimpleSelector from "@/components/form/SimpleSelector";
import formMixin from "./formMixin";
export default {
  name: "CarouselForm",
  components: {
    ImageUploader,
    SimpleSelector
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
  },
  computed: {
    ...mapGetters(["newsState"])
  }
};
</script>