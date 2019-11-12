<template>
  <el-select v-model="data" :placeholder="$t('tip')" clearable @clear="$emit('clear')">
    <el-option v-for="item in options" :key="item.value" :label="$t(item.name)" :value="item.value" />
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择语言"
  },
  "en-US": {
    "tip": "Select language"
  }
}
</i18n>

<script>
export default {
  name: "LanguageSelector",
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
        { name: "Chinese", value: "zh-CN" },
        { name: "English", value: "en-US" }
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