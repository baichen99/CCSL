<template>
  <el-input v-model="data" clearable :placeholder="$t('tip')" @keyup.enter="$emit('enter')" />
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请输入复合词构词语素"
  },
  "en-US": {
    "tip": "Input morpheme(Chinese)"
  }
}
</i18n>

<script>
export default {
  name: "MorphemeInput",
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

