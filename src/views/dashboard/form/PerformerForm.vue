<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Name')" prop="name">
      <el-input v-model="formData.name" placeholder="请输入被试姓名" clearable />
    </el-form-item>

    <el-form-item :label="$t('Region')" prop="regionID">
      <city-selector v-model="formData.regionID" />
    </el-form-item>

    <el-form-item :label="$t('Gender')" prop="gender">
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
        regionID: [{ required: true, message: "请选择地区", trigger: "blur" }],
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