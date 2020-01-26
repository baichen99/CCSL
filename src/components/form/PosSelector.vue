<template>
  <el-select v-model="data" clearable :placeholder="$t('tip')" @clear="$emit('clear')">
    <el-option v-for="(item,key) in partOfSpeech" :key="key" :label="$t(item.name)" :value="key" />
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
  name: "PosSelector",
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
    ...mapGetters(["wordPosTypes", "partOfSpeech"]),
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