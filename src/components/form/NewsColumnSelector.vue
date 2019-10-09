<template>
  <el-select v-model="data" placeholder="请选择新闻栏目" clearable @clear="$emit('clear')">
    <el-option v-for="item in options" :key="item.value" :label="item.name" :value="item.value" />
  </el-select>
</template>

<script>
export default {
  name: "NewsColumnSelector",
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
        { name: "新闻动态", value: "news" },
        { name: "研究成果", value: "research" },
        { name: "学术活动", value: "activity" },
        { name: "通知公告", value: "notice" }
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