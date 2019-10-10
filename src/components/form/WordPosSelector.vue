<template>
  <el-select v-model="data" placeholder="请选择词性" clearable @clear="$emit('clear')">
    <el-option
      v-for="(item,index) in $store.getters.wordTypes"
      :key="index"
      :label="item"
      :value="item"
    />
  </el-select>
</template>

<script>
export default {
  name: "WordPosSelector",
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