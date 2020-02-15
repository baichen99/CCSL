<template>
  <el-cascader
    v-model="value"
    :size="size"
    :options="cities"
    :placeholder="$t('tip')"
    :props="{ expandTrigger: 'hover' }"
    clearable
    @change="changeRegion"
  />
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择地区"
  },
  "en-US": {
    "tip": "Select region"
  }
}
</i18n>

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
    },
    size: {
      type: String,
      default: ""
    }
  },
  data() {
    return {
      value: [31, 3101, undefined],
      cities
    };
  },
  watch: {
    city() {
      this.updateValue();
    }
  },
  created() {
    this.updateValue();
  },
  methods: {
    changeRegion() {
      if (this.value[2]) {
        this.$emit("update", this.value[2]);
      } else {
        this.$emit("update", undefined);
      }
    },
    updateValue() {
      const reg = /^\d{6}$/;
      if (this.city && reg.test(this.city)) {
        const districtCode = String(this.city);
        const cityCode = districtCode.substring(0, 4);
        const provinceCode = districtCode.substring(0, 2);
        this.value = [
          Number(provinceCode),
          Number(cityCode),
          Number(districtCode)
        ];
      }
    }
  }
};
</script>