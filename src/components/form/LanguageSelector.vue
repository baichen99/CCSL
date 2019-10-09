<template>
  <el-select v-model="data" placeholder="请选择语言" clearable @clear="$emit('clear')">
    <el-option v-for="item in options" :key="item.value" :label="item.name" :value="item.value" />
  </el-select>
</template>

<script>
export default {
  name: "Languageelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: String,
      default: () => ""
    }
  },
  data() {
    return {
      options: [
        { name: "中文", value: "zh-CN" },
        { name: "英文", value: "en-US" }
      ]
    };
  },
  computed: {
    data: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  }
};
</script>