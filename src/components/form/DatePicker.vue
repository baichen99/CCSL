<template>
  <el-date-picker
    v-model="value"
    type="date"
    :placeholder="$t('tip')"
    format="yyyy / MM / dd"
    value-format="yyyy-MM-ddTHH:mm:ssZ"
    clearable
    @clear="$emit('clear')"
  />
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请选择日期"
  },
  "en-US": {
    "tip": "Select date"
  }
}
</i18n>


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