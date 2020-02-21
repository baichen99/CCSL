<template>
  <el-select
    ref="dragableSelect"
    v-model="data"
    class="dragable-select"
    :placeholder="options.placeholder[$i18n.locale]"
    :size="size"
    filterable
    multiple
    @clear="$emit('clear')"
    v-on="$listeners"
  >
    <el-option
      v-for="item in options.data"
      :key="item.value"
      :label="$t(item.text)"
      :value="item.value"
    >
      <slot name="option"></slot>
    </el-option>
  </el-select>
</template>

<script>
import Sortable from "sortablejs";
export default {
  name: "MultipleSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: Array,
      default: () => []
    },
    size: {
      type: String,
      default: () => ""
    },
    options: {
      type: Object,
      default: () => ({})
    }
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