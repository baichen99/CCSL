<template>
  <el-select v-model="data" placeholder="请选择用户角色" clearable @clear="$emit('clear')">
    <el-option v-for="(item, key) in userTypes" :key="key" :label="item.name" :value="key" />
  </el-select>
</template>

<script>
import { mapGetters } from "vuex";
export default {
  name: "UserTypeSelector",
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
    ...mapGetters(["userTypes"]),
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