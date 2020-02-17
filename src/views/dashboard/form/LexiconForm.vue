<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item label="中文转写" prop="chinese">
      <el-input v-model="formData.chinese" placeholder="请输入中文转写" clearable />
    </el-form-item>

    <el-form-item label="英文转写" prop="english">
      <el-input v-model="formData.english" placeholder="请输入英文转写" clearable />
    </el-form-item>

    <el-form-item label="词性" prop="pos">
      <multiple-selector v-model="formData.pos" :options="partOfSpeech" />
    </el-form-item>

    <el-form-item label="音序" prop="initial">
      <word-initial-selector v-model="formData.initial" />
    </el-form-item>
  </el-form>
</template>

<script>
import { mapGetters } from "vuex";
import formMixin from "./formMixin";
import MultipleSelector from "@/components/form/MultipleSelector";
import WordInitialSelector from "@/components/form/WordInitialSelector";
export default {
  name: "LexiconForm",
  components: {
    MultipleSelector,
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
  },
  computed: {
    ...mapGetters(["partOfSpeech"])
  }
};
</script>