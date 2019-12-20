<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item label="中文转写" prop="chinese">
      <el-input v-model="formData.chinese" placeholder="请输入中文转写" clearable />
    </el-form-item>

    <el-form-item label="英文转写" prop="english">
      <el-input v-model="formData.english" placeholder="请输入英文转写" clearable />
    </el-form-item>

    <el-form-item label="词性" prop="pos">
      <word-pos-selector v-model="formData.pos" />
    </el-form-item>

    <el-form-item label="音序" prop="initial">
      <word-initial-selector v-model="formData.initial" />
    </el-form-item>
  </el-form>
</template>

<script>
import formMixin from "./formMixin";
import WordPosSelector from "@/components/form/WordPosSelector";
import WordInitialSelector from "@/components/form/WordInitialSelector";
export default {
  name: "LexiconForm",
  components: {
    WordPosSelector,
    WordInitialSelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        chinese: [
          { required: true, message: "请输入中文转写", trigger: "blur" }
        ],
        english: [{ required: true, message: "请输入英文", trigger: "blur" }],
        pos: [{ required: true, message: "请选择词性", trigger: "blur" }],
        initial: [{ required: true, message: "请选择音序", trigger: "blur" }]
      }
    };
  }
};
</script>