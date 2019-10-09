<template>
  <el-card class="news-list">
    <div slot="header" class="list-header">
      <svg-icon :icon-class="newsTypes[params.column].icon" />
      <h2>{{ newsTypes[params.column].name }}</h2>
    </div>

    <div class="list-content">
      <div v-for="item in list" :key="item.id">
        <i class="dot">&bull;</i>
        <el-link class="news-title" @click="showDetail(item)">{{ item.title }}</el-link>
        <span class="news-date">{{ $d(new Date(item.date),'short') }}</span>
      </div>
    </div>

    <el-pagination
      v-if="total>params.limit"
      background
      layout="total,prev, pager, next"
      :total="total"
      :page-size.sync="params.limit"
      :current-page.sync="params.page"
    />
  </el-card>
</template>

<script>
import { mapGetters } from "vuex";
import { GetNewsList } from "@/api/news";
export default {
  name: "NewsList",
  data() {
    return {
      list: [],
      params: {
        limit: 10,
        page: 1,
        column: ""
      },
      total: 0
    };
  },
  computed: { ...mapGetters(["newsTypes"]) },
  watch: {
    "params.page"() {
      this.getList();
    }
  },
  created() {
    this.getList();
  },
  methods: {
    getList() {
      const column = this.$route.params.column;
      this.params.column = column;
      GetNewsList(this.params).then(res => {
        this.list = res.data;
        this.total = res.total;
        console.log(this.list);
      });
    },
    showDetail(item) {
      if (item.type === "link") {
        window.open(item.text);
      } else if (item.type === "document") {
        this.$router.push({ path: `/news-detail/${item.id}` });
      }
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.news-list {
  max-width: 900px;
  margin: 20px auto;
  .list-header {
    text-align: center;
    color: $--color-primary;
    svg {
      font-size: 25px;
    }
  }
  .list-content {
    .dot {
      color: $--color-primary;
      padding: 0.2rem;
    }
    .el-link {
      display: inline-block;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      max-width: 500px;
      font-size: 1.1rem;
    }

    .news-date {
      float: right;
      color: #a0a0a0;
      padding-top: 0.2rem;
    }
  }
  .el-pagination {
    text-align: center;
  }
}
</style>