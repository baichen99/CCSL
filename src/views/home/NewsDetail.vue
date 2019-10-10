<template>
  <div class="news-detail">
    <el-card shadow="hover">
      <div class="news-header">
        <h2 class="news-title">{{ data.title }}</h2>
        <el-divider>
          <svg-icon v-if="data.column" :icon-class="data.column" />
        </el-divider>
        <span v-if="data.date">发布日期： {{ $d(new Date(data.date),"short") }}</span>
        <span>发布人：{{ data.creator.name }}</span>
        <el-divider />
      </div>
      <div class="news-content" v-html="data.text"></div>
    </el-card>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
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
  computed: { ...mapGetters(["newsTypes"]) },
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
    font-size: 18px;
    line-height: 2rem;
  }
}
</style>