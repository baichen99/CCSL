<template>
  <el-form
    label-position="left"
    label-width="100px"
    :model="formData"
  >
    <el-form-item
      label="图片标题"
      required
    >
      <el-input
        v-model="formData.title"
        placeholder="请输入图片标题"
      />
    </el-form-item>
    <el-form-item
      label="滚动图片（16:9）"
      required
    >
      <image-uploader
        v-model="formData.image"
        dir="news"
      />
    </el-form-item>

    <el-form-item
      label="状态"
      required
    >
      <news-state-selector v-model="formData.state" />
    </el-form-item>

    <el-form-item
      label="重要性"
      required
    >
      <el-rate v-model="formData.importance" />
      <el-button size="mini" @click="formData.importance=0">重置</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import ImageUploader from "@/components/form/ImageUploader";
import NewsStateSelector from "@/components/form/NewsStateSelector";
export default {
  name: "CarouselForm",
  components: {
    ImageUploader,
    NewsStateSelector
  },
  model: {
    prop: "data",
    event: "update"
  },
  props: {
    data: {
      type: Object,
      default: () => ({})
    }
  },
  computed: {
    formData: {
      get() {
        return this.data;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  }
};
</script>