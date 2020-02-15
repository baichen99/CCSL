<template>
  <el-input v-model="data" :size="size" :placeholder="placeholder" clearable />
</template>

<script>
export default {
  name: "SearchInput",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: String,
      default: () => ""
    },
    placeholder: {
      type: String,
      default: () => ""
    },
    size: {
      type: String,
      default: () => "mini"
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

<style scoped>
.el-input {
  width: 150px;
  margin-right: 5px;
}
</style>