<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules">
    <el-form-item :label="$t('Name')" prop="name">
      <el-input v-model="formData.name" placeholder="请输入被试姓名" clearable />
    </el-form-item>

    <el-form-item :label="$t('Region')" prop="regionID">
      <city-selector v-model="formData.regionID" />
    </el-form-item>

    <el-form-item :label="$t('Gender')" prop="gender">
      <simple-selector v-model="formData.gender" :options="genderTypes" />
    </el-form-item>
  </el-form>
</template>

<script>
import { mapGetters } from "vuex";
import SimpleSelector from "@/components/form/SimpleSelector";
import CitySelector from "@/components/form/CitySelector";
import formMixin from "./formMixin";
export default {
  name: "PerformerForm",
  components: {
    SimpleSelector,
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
  },
  computed: {
    ...mapGetters(["genderTypes"])
  }
};
</script>

<style lang="scss" scoped>
.el-cascader {
  width: 100%;
}
</style>