<template>
  <el-select v-model="data" placeholder="请选择用户角色" clearable @clear="$emit('clear')">
    <el-option v-for="item in options" :key="item.value" :label="item.name" :value="item.value" />
  </el-select>
</template>

<script>
export default {
  name: "UserTypeSelector",
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
        { name: "普通管理员", value: "admin" },
        { name: "学习平台用户", value: "learner" },
        { name: "语料库用户", value: "user" },
        { name: "超级管理员", value: "super" }
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