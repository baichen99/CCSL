<template>
  <el-select v-model="data" :placeholder="$t('tip')" clearable @clear="$emit('clear')">
    <el-option v-for="(item,key) in newsState" :key="key" :label="$t(item.name)" :value="key" />
  </el-select>
</template>


<i18n>
{
  "zh-CN": {
    "tip": "请选择状态"
  },
  "en-US": {
    "tip": "Select state"
  }
}
</i18n>


<script>
import { mapGetters } from "vuex";
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
  computed: {
    ...mapGetters(["newsState"]),
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