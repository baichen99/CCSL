<template>
  <div class="news-detail">
    <el-card shadow="hover">
      <div class="news-header">
        <h2 class="news-title">{{ data.title }}</h2>
        <el-divider>
          <svg-icon v-if="data.column" :icon-class="data.column" />
        </el-divider>
        <span v-if="data.date">{{ $t("date") }}： {{ $d(new Date(data.date),"short") }}</span>
        <span>{{ $t("publisher") }}：{{ data.creator.name }}</span>
        <el-divider />
      </div>
      <div class="news-content" v-html="data.text"></div>
    </el-card>
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "date": "发布日期",
    "publisher": "发布人"
  },
  "en-US": {
    "date": "Publish Date",
    "publisher": "Publisher"
  }
}
</i18n>

<script>
import { GetNews } from "@/api/news";
export default {
  name: "NewsDetail",
  data() {
    return {
      data: {
        title: "",
        date: "",
        creator: {},
        text: ""
      }
    };
  },
  created() {
    this.getData();
  },
  methods: {
    getData() {
      const newsID = this.$route.params.id;
      GetNews(newsID).then(res => {
        this.data = res.data;
      });
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.news-detail {
  max-width: 900px;
  margin: 20px auto;
  .news-header {
    text-align: center;
    .news-title {
      color: $--color-primary;
      margin: 20px 100px;
    }
    svg {
      font-size: 25px;
      color: $--color-primary;
    }
    span {
      margin: 0 20px;
    }
  }
  .news-content {
    padding: 10px 30px;
    text-indent: 2em;
    line-height: 2rem;
    p {
      text-indent: 2rem;
    }
  }
}
</style>