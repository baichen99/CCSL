<template>
  <el-select
    v-model="data"
    :disabled="disabled"
    :placeholder="$t('tip')"
    clearable
    @clear="$emit('clear')"
  >
    <el-option v-for="(item, key) in userState" :key="key" :label="$t(item.name)" :value="key" />
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择用户状态"
  },
  "en-US": {
    "tip": "Select state"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "UserStateSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: String,
      default: () => ""
    },
    disabled: {
      type: Boolean,
      default: false
    }
  },
  computed: {
    ...mapGetters(["userState"]),
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