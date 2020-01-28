<template>
  <el-select
    ref="dragableSelect"
    v-model="data"
    class="dragable-select"
    multiple
    filterable
    clearable
    reserve-keyword
    :placeholder="$t('tip')"
    :loading="loading"
    v-on="$listeners"
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
import Sortable from "sortablejs";
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
  },
  mounted() {
    this.setSort();
  },
  methods: {
    setSort() {
      const el = this.$refs.dragableSelect.$el.querySelectorAll(
        ".el-select__tags > span"
      )[0];
      this.sortable = Sortable.create(el, {
        setData: function(dataTransfer) {
          dataTransfer.setData("Text", "");
          // to avoid Firefox bug
          // Detail see : https://github.com/RubaXa/Sortable/issues/1012
        },
        onEnd: evt => {
          const targetRow = this.value.splice(evt.oldIndex, 1)[0];
          this.value.splice(evt.newIndex, 0, targetRow);
        }
      });
    }
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