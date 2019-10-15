<template>
  <el-select
    v-model="data"
    multiple
    filterable
    remote
    clearable
    reserve-keyword
    placeholder="请输入手形名称"
    :remote-method="filterRemote"
    :loading="loading"
    @clear="filterRemote"
  >
    <el-option v-for="item in options" :key="item.id" :label="item.name" :value="item.id">
      <span>{{ item.name }}</span>
      <img :src="'https://ccsl.shu.edu.cn/public/'+item.image" :alt="item.name" />
    </el-option>
  </el-select>
</template>

<script>
import { GetSignsList } from "@/api/signs";
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
      GetSignsList({ limit: 0, name: query })
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