<template>
  <el-select v-model="data" :placeholder="$t('tip')" clearable @clear="$emit('clear')">
    <el-option v-for="(item,index) in wordPosTypes" :key="index" :label="item" :value="item" />
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择词性"
  },
  "en-US": {
    "tip": "Select part of speech"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
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
    ...mapGetters(["wordPosTypes"]),
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