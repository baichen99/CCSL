<template>
  <el-select v-model="data" placeholder="请选择性别" clearable @clear="$emit('clear')">
    <el-option v-for="(item,index) in options" :key="index" :label="item" :value="item" />
  </el-select>
</template>

<script>
export default {
  name: "GenderSelector",
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
      options: ["男", "女"]
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