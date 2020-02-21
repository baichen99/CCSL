<template>
  <el-select v-model="data" filterable clearable :loading="loading" @clear="$emit('clear')">
    <el-option v-for="item in options" :key="item.id" :value="item.id" :label="formatter(item)">
      <slot :item="item"></slot>
    </el-option>
  </el-select>
</template>

<script>
export default {
  name: "RemoteSelector",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: String,
      default: () => ""
    },
    getListMethod: {
      type: Function,
      required: true
    },
    formatter: {
      type: Function,
      required: true
    },
    params: {
      type: Object,
      default: () => ({
        limit: 0
      })
    }
  },
  data() {
    return {
      options: [],
      loading: false
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
  async created() {
    this.$nextTick(async () => {
      this.loading = true;
      try {
        const res = await this.getListMethod(this.params);
        this.options = res.data;
      } catch (err) {
        console.error(err);
      } finally {
        this.loading = false;
      }
    });
  }
};
</script>