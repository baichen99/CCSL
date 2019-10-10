<template>
  <el-date-picker
    v-model="value"
    type="date"
    placeholder="选择日期"
    format="yyyy 年 MM 月 dd 日"
    value-format="yyyy-MM-ddTHH:mm:ssZ"
    clearable
    @clear="$emit('clear')"
  />
</template>

<script>
// RFC3339 Formated Date - example: 2000-12-30T00:00:00Z
export default {
  name: "DatePicker",
  model: {
    prop: "date",
    event: "update"
  },
  props: {
    date: {
      type: String,
      default: ""
    }
  },
  computed: {
    value: {
      get() {
        return this.date;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  }
};
</script>