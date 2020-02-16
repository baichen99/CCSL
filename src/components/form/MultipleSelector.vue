<template>
  <el-select
    v-model="data"
    :placeholder="options.placeholder[$i18n.locale]"
    :size="size"
    clearable
    multiple
    @clear="$emit('clear')"
  >
    <el-option
      v-for="item in options.data"
      :key="item.value"
      :label="$t(item.text)"
      :value="item.value"
    />
  </el-select>
</template>

<script>
export default {
  name: "MultipleSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: Array,
      default: () => []
    },
    size: {
      type: String,
      default: () => ""
    },
    options: {
      type: Object,
      default: () => ({})
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