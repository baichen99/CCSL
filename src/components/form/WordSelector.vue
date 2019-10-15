<template>
  <el-select
    v-model="data"
    filterable
    clearable
    remote
    placeholder="请输入中文转写"
    :remote-method="filterRemote"
    :loading="loading"
    @clear="filterRemote"
  >
    <el-option
      v-for="item in options"
      :key="item.id"
      :label="item.chinese"
      :value="item.id"
    >{{ item.initial }} - {{ item.chinese }}</el-option>
  </el-select>
</template>

<script>
import { GetLexicalWordsList } from "@/api/words";
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
    this.filterRemote();
  },
  methods: {
    filterRemote(query) {
      this.loading = true;
      GetLexicalWordsList({
        limit: 0,
        chinese: query,
        orderBy: "initial, id"
      })
        .then(res => {
          this.options = res.data;
          this.loading = false;
        })
        .catch(() => {
          this.loading = false;
        });
    }
  }
};
</script>