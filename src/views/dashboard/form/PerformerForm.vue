<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item label="姓名" prop="name" required>
      <el-input v-model="formData.name" placeholder="请输入被试姓名" clearable />
    </el-form-item>

    <el-form-item label="地区" prop="regionID" required>
      <city-selector v-model="formData.regionID" />
    </el-form-item>

    <el-form-item label="性别" prop="gender" required>
      <gender-selector v-model="formData.gender" />
    </el-form-item>
  </el-form>
</template>

<script>
import GenderSelector from "@/components/form/GenderSelector";
import CitySelector from "@/components/form/CitySelector";
import formMixin from "./formMixin";
export default {
  name: "PerformerForm",
  components: {
    GenderSelector,
    CitySelector
  },
  mixins: [formMixin],
  data() {
    return {
      rules: {
        name: [{ required: true, message: "请输入姓名", trigger: "blur" }],
        region: [{ required: true, message: "请选择地区", trigger: "blur" }],
        gender: [{ required: true, message: "请选择性别", trigger: "blur" }]
      }
    };
  }
};
</script>

<style lang="scss" scoped>
.el-cascader {
  width: 100%;
}
</style>