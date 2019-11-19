<template>
  <el-select v-model="data" :placeholder="$t('tip')" clearable @clear="$emit('clear')">
    <el-option v-for="(item,key) in newsTypes" :key="key" :label="$t(item.name)" :value="key" />
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择类型"
  },
  "en-US": {
    "tip": "Select type"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "NewsTypeSelector",
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
        { name: "文档", value: "document" },
        { name: "链接", value: "link" }
      ]
    };
  },
  computed: {
    ...mapGetters(["newsTypes"]),
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