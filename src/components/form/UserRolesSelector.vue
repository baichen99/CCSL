<template>
  <el-select v-model="data" :placeholder="$t('tip')" clearable multiple @clear="$emit('clear')">
    <el-option v-for="(item, key) in userRoles" :key="key" :label="$t(item.name)" :value="key" />
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择用户角色"
  },
  "en-US": {
    "tip": "Select role"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "UserRolesSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: Array,
      default: () => []
    }
  },
  computed: {
    ...mapGetters(["userRoles"]),
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