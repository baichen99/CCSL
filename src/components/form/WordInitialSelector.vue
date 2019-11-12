<template>
  <el-select v-model="data" :placeholder="$t('tip')" clearable @clear="$emit('clear')">
    <el-option v-for="(item,index) in options" :key="index" :label="item" :value="item" />
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择音序"
  },
  "en-US": {
    "tip": "Select initial"
  }
}
</i18n>

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
  data() {
    return {
      options: [
        "A",
        "B",
        "C",
        "D",
        "E",
        "F",
        "G",
        "H",
        "I",
        "J",
        "K",
        "L",
        "M",
        "N",
        "O",
        "P",
        "Q",
        "R",
        "S",
        "T",
        "W",
        "X",
        "Y",
        "Z"
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