<template>
  <div class="news-detail">
    <el-card shadow="hover">
      <div class="news-header">
        <h2 class="news-title">{{ data.title }}</h2>
        <el-divider>
          <svg-icon v-if="data.column" :icon-class="data.column" />
        </el-divider>
        <span v-if="data.date">{{ $t("date") }}： {{ $d(new Date(data.date),"short") }}</span>
        <span>{{ $t("Publisher") }}：{{ data.creator.name }}</span>
        <el-divider />
      </div>
      <div class="news-content" v-html="data.text"></div>
    </el-card>
    <el-button
      v-show="scroll > 200"
      type="primary"
      icon="el-icon-caret-top"
      class="back-to-top"
      circle
      plain
      @click="backToTop"
    />
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "date": "发布日期"
  },
  "en-US": {
    "date": "Publish Date"
  }
}
</i18n>

<script>
import { GetNews } from "@/api/news";
import Plyr from "plyr";
export default {
  name: "NewsDetail",
  data() {
    return {
      data: {
        title: "",
        date: "",
        creator: {},
        text: ""
      },
      scroll: 0
    };
  },
  async mounted() {
    await this.getData();
    window.addEventListener("scroll", this.getScroll);
    this.setupVideos();
  },
  methods: {
    async getData() {
      const newsID = this.$route.params.id;
      const res = await GetNews(newsID);
      this.data = res.data;
    },
    getScroll() {
      this.scroll =
        window.pageYOffset ||
        document.documentElement.scrollTop ||
        document.body.scrollTop;
    },
    backToTop() {
      const scrollToptimer = setInterval(() => {
        let top = document.body.scrollTop || document.documentElement.scrollTop;
        let speed = top / 5;
        if (document.body.scrollTop != 0) {
          document.body.scrollTop -= speed;
        } else {
          document.documentElement.scrollTop -= speed;
        }
        if (top == 0) {
          clearInterval(scrollToptimer);
        }
      }, 30);
    },
    setupVideos() {
      Plyr.setup("video", {
        controls: ["play-large", "progress", "fullscreen", "pip", "volume"]
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
    img {
      border-radius: 4px;
    }
  }
  .back-to-top {
    position: fixed;
    right: 40px;
    bottom: 40px;
  }
}
</style>