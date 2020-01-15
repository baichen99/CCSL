<template>
  <el-select v-model="data" filterable clearable :placeholder="$t('tip')" :loading="loading">
    <el-option v-for="item in options" :key="item.id" :label="item.chinese" :value="item.id">
      {{ item.initial }} -
      <span
        class="tag-value word-sup"
        v-html="$options.filters.addNumberSup(item.chinese)"
      ></span>
    </el-option>
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请输入中文转写"
  },
  "en-US": {
    "tip": "Input Chinese"
  }
}
</i18n>

<script>
export default {
  name: "WordSelector",
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
      loading: false,
      options: []
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
  },
  created() {
    this.$nextTick(async () => {
      this.loading = true;
      await this.$store.dispatch("data/getLexicons");
      this.options = Object.values(this.$store.state.data.lexicons);
      this.loading = false;
    });
  }
};
</script>