<template>
  <el-select
    v-model="data"
    multiple
    filterable
    clearable
    reserve-keyword
    :placeholder="$t('tip')"
    :loading="loading"
  >
    <el-option v-for="item in options" :key="item.id" :label="item.name" :value="item.id">
      <span>{{ item.name }}</span>
      <img :src="settings.publicURL + item.image" :alt="item.name" />
    </el-option>
  </el-select>
</template>

<i18n>
{
  "zh-CN": {
    "tip": "请输入手形名称"
  },
  "en-US": {
    "tip": "Input sign name"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "MultiSignSelector",
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
  data() {
    return {
      loading: false,
      options: []
    };
  },
  computed: {
    ...mapGetters(["settings"]),
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
      await this.$store.dispatch("data/getSigns");
      const data = this.$store.state.data.signs;
      this.options = Object.values(data);
    });
  }
};
</script>

<style lang="scss" scoped>
.el-select-dropdown__item {
  display: flex;
  justify-content: space-around;
  align-items: center;

  height: 60px;
  img {
    max-height: 100%;
    max-width: 70px;
    padding: 5px;
  }
}
</style>