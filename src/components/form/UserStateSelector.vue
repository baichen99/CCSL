<template>
  <el-select v-model="data" placeholder="请选择用户状态" clearable @clear="$emit('clear')">
    <el-option v-for="(item, key) in userState" :key="key" :label="item.name" :value="key" />
  </el-select>
</template>

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