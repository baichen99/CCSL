<template>
  <el-cascader
    v-model="value"
    :options="cities"
    placeholder="请选择地区"
    :props="{ expandTrigger: 'hover' }"
    clearable
    @change="changeRegion"
  />
</template>

<script>
import cities from "@/assets/cities";
export default {
  name: "CitySelector",
  model: {
    prop: "city",
    event: "update"
  },
  props: {
    city: {
      type: Number,
      default: undefined
    }
  },
  data() {
    return {
      value: ["31", "3101"],
      cities
    };
  },
  created() {
    const reg = /^\d{6}$/;
    if (this.city && reg.test(this.city)) {
      const districtCode = String(this.city);
      const cityCode = districtCode.substring(0, 4);
      const provinceCode = districtCode.substring(0, 2);
      this.value = [provinceCode, cityCode, districtCode];
    }
  },
  methods: {
    changeRegion() {
      if (this.value[2]) {
        this.$emit("update", Number(this.value[2]));
      } else {
        this.$emit("update", undefined);
      }
    }
  }
};
</script>

<style lang="scss" scoped>
</style>