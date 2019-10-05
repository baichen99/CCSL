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
export default {
  name: "SignForm",
  components: {
    ImageUploader
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