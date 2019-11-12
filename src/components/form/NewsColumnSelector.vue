<template>
  <el-select v-model="data" placeholder="请选择新闻栏目" clearable @clear="$emit('clear')">
    <el-option v-for="(item,key) in newsColumns" :key="key" :label="$t(item.name)" :value="item.value" />
  </el-select>
</template>

<script>
import { mapGetters } from "vuex";
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
  computed: {
    ...mapGetters(["newsColumns"]),
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