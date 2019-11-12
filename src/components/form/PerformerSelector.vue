<template>
  <el-select
    v-model="data"
    filterable
    remote
    clearable
    placeholder="请输入姓名"
    :remote-method="filterRemote"
    :loading="loading"
    @clear="filterRemote"
  >
    <el-option
      v-for="item in options"
      :key="item.id"
      :value="item.id"
      :label="item.name"
    >{{ item.name }} - {{ $t(genderTypes[item.gender].name) }} - {{ item.region.name }}</el-option>
  </el-select>
</template>

<script>
import { mapGetters } from "vuex";
import { GetPerformersList } from "@/api/performers";
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
    this.filterRemote();
  },
  methods: {
    filterRemote(query) {
      if (query !== "") {
        this.loading = true;
        GetPerformersList({ limit: 0, name: query })
          .then(res => {
            this.options = res.data;
            this.loading = false;
          })
          .catch(() => {
            this.loading = false;
          });
      } else {
        this.options = [];
      }
    }
  }
};
</script>