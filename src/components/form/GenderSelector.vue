<template>
  <el-select v-model="data" placeholder="请选择性别" clearable @clear="$emit('clear')">
    <el-option v-for="(item, key) in genderTypes" :key="key" :label="item.name" :value="item.value" />
  </el-select>
</template>

<script>
import { mapGetters } from "vuex";
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
  computed: {
    ...mapGetters(["genderTypes"]),
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