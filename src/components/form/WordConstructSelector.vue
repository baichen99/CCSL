<template>
  <el-select v-model="data" placeholder="请选择构词方式" clearable @clear="$emit('clear')">
    <el-option v-for="(item,key) in constructTypes" :key="key" :label="item.name" :value="key" />
  </el-select>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "WordConstructSelector",
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
    ...mapGetters(["constructTypes"]),
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