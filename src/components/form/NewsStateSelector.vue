<template>
  <el-select
    v-model="data"
    placeholder="请选择状态"
    clearable
    @clear="$emit('clear')"
  >
    <el-option
      v-for="item in options"
      :key="item.value"
      :label="item.name"
      :value="item.value"
    />
  </el-select>
</template>

<script>
export default {
  name: "NewsStateSelector",
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
        { name: "发布", value: "published" },
        { name: "草稿", value: "draft" }
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