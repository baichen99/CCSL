<template>
  <el-card class="news-list">
    <div slot="header" class="list-header">
      <svg-icon :icon-class="$options.filters.getObjectItem(newsColumns, params.column).icon" />
      <h2>{{ $t($options.filters.getObjectItem(newsColumns, params.column).text) }}</h2>
    </div>

    <div class="list-content">
      <div v-for="item in list" :key="item.id" class="news-item">
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
        column: "",
        language: "",
        order: "desc",
        state: "published"
      },
      total: 0
    };
  },
  computed: { ...mapGetters(["newsColumns"]) },
  watch: {
    "params.page"() {
      this.getList();
    },
    "$i18n.locale"() {
      this.getList();
    }
  },
  created() {
    this.getList();
  },
  methods: {
    async getList() {
      const lang = this.$i18n.locale;
      this.params.language = lang;
      const column = this.$route.params.column;
      this.params.column = column;
      const res = await GetNewsList(this.params);
      this.list = res.data;
      this.total = res.total;
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
    .news-item {
      line-height: 1.2rem;
      padding: 0.5rem;
      .dot {
        color: $--color-primary;
        padding: 0.2rem;
      }
      .el-link {
        display: inline-block;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        max-width: 650px;
        font-size: 1.1rem;
      }

      .news-date {
        float: right;
        color: #a0a0a0;
        padding-top: 0.2rem;
      }
    }
  }
  .el-pagination {
    text-align: center;
  }
}
</style>