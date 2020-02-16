<template>
  <div class="news-card">
    <el-card shadow="hover">
      <div slot="header">
        <svg-icon :icon-class="icon" />
        <span class="news-header">{{ $t(title) }}</span>
        <el-link class="news-more" type="primary" @click="showList(column)">{{ $t("more") }}</el-link>
      </div>
      <div v-for="item in news" :key="item.id" class="news-content">
        <i class="dot">&bull;</i>
        <el-tooltip effect="light" placement="top" :content="item.title">
          <el-link class="news-title" @click="showDetail(item)">{{ item.title }}</el-link>
        </el-tooltip>
        <div class="news-date">{{ $d(new Date(item.date),'short') }}</div>
      </div>
    </el-card>
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "more": "更多"
  },
  "en-US": {
    "more": "More"
  }
}
</i18n>

<script>
import { GetNewsList } from "@/api/news";
export default {
  name: "NewsCard",
  props: {
    icon: {
      type: String,
      required: true
    },
    column: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      news: []
    };
  },
  watch: {
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
      const params = {
        column: this.column,
        language: lang,
        order: "desc",
        state: "published",
        limit: 5
      };
      const res = await GetNewsList(params);
      this.news = res.data;
    },
    showDetail(item) {
      if (item.type === "link") {
        window.open(item.text);
      } else if (item.type === "document") {
        this.$router.push({ path: `/news-detail/${item.id}` });
      }
    },
    showList(column) {
      this.$router.push({ path: `/news-list/${column}` });
    }
  }
};
</script>



<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.news-card {
  .el-card {
    margin: 10px;
    min-width: 500px;
    .news-header {
      font-weight: bolder;
      color: $--color-primary;
    }
    .news-more {
      float: right;
      padding: 3px 0;
    }
    .svg-icon {
      fill: $--color-primary;
      margin-right: 1.2rem;
      font-size: 1.2rem;
    }

    .news-content {
      padding: 0.3rem;
    }
    .dot {
      color: #2363c3;
      padding: 0.2rem;
    }
    .news-title {
      display: inline-block;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      max-width: 300px;
      font-size: 1rem;
      // font-weight: normal;

      // @media only screen and (max-device-width: 480px) {
      //   width: 50vw;
      // }
    }
    .news-date {
      float: right;
      color: #a0a0a0;
      height: 100%;
      padding-top: 0.2rem;
    }
  }
}
</style>
