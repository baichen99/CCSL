<template>
  <el-select v-model="data" filterable clearable :placeholder="$t('tip')" :loading="loading">
    <el-option
      v-for="item in options"
      :key="item.id"
      :value="item.id"
      :label="item.name"
    >{{ item.name }} - {{ $t(genderTypes[item.gender].name) }} - {{ item.regionID | getRegionName }}</el-option>
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请输入姓名"
  },
  "en-US": {
    "tip": "Input name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "PerformerSelector",
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
    ...mapGetters(["genderTypes"]),
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
      await this.$store.dispatch("data/getPerformers");
      this.options = Object.values(this.$store.state.data.performers);
      this.loading = false;
    });
  }
};
</script>