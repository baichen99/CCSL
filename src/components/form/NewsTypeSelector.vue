<template>
  <el-select v-model="data" placeholder="请选择新闻类型" clearable @clear="$emit('clear')">
    <el-option v-for="item in options" :key="item.value" :label="item.name" :value="item.value" />
  </el-select>
</template>

<script>
export default {
  name: "NewsTypeSelector",
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
        { name: "文档", value: "document" },
        { name: "链接", value: "link" }
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