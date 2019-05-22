<template>
  <header>
    <div>
      <div>
        <img
          src="@/assets/logo.svg"
          alt="logo"
          height="35px"
        >
        <div class="title">
          <span>中国手语及聋人研究中心</span>
          <span class="en">Center for CSL and Deaf Studies</span>
        </div>
      </div>
      <el-menu
        :default-active="'/'"
        mode="horizontal"
        :router="true"
        active-text-color="#2363C3"
      >
        <el-menu-item index="/">首页</el-menu-item>
        <el-menu-item index="/dictionary">语料库</el-menu-item>
        <el-submenu index="/academic">
          <template slot="title">学术</template>
          <el-menu-item
            v-for="item in academy"
            :key="item.url"
            :index="item.url"
            :route="item.route"
          >
            {{ item.title }}
          </el-menu-item>
        </el-submenu>
        <el-submenu index="/about">
          <template slot="title">关于</template>
          <el-menu-item
            v-for="item in about"
            :key="item.url"
            :index="item.url"
            :route="item.route"
          >
            {{ item.title }}
          </el-menu-item>
        </el-submenu>
      </el-menu>
    </div>
  </header>
</template>

<script>
export default {
  name: "Header",
  data: () => ({
    academy: [
      {
        url: "/team",
        title: "研究团队",
        route: {
          name: "news-detail",
          query: { data: "研究团队", type: "introduction" }
        }
      },
      {
        url: "/research",
        title: "研究成果",
        route: { name: "news-list", query: { newsType: "研究成果" } }
      },
      {
        url: "/activity",
        title: "学术活动",
        route: { name: "news-list", query: { newsType: "学术活动" } }
      }
    ],
    about: [
      {
        url: "/introduction",
        title: "中心简介",
        route: {
          name: "news-detail",
          query: { data: "中心简介", type: "introduction" }
        }
      },
      {
        url: "/inform",
        title: "通知公告",
        route: { name: "news-list", query: { newsType: "通知公告" } }
      },
      {
        url: "/download",
        title: "资料下载",
        route: { name: "news-list", query: { newsType: "资料下载" } }
      },
      {
        url: "/contact",
        title: "联系我们",
        route: {
          name: "news-detail",
          query: { data: "联系我们", type: "introduction" }
        }
      }
    ]
  }),
  computed: {
    activeMenu() {
      if (this.$route.matched.length > 1) {
        return this.$route.matched[0].path;
      } else {
        return this.$route.path;
      }
    }
  },
  methods: {
    getNews(newsType) {
      this.$router.push({
        name: "news-list",
        params: { newsType }
      });
    }
  }
};
</script>

<style lang="scss">
header {
  height: 62px;

  div {
    height: 100%;
    padding: 0 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    img {
      @media only screen and (max-device-width: 480px) {
        display: none;
      }
    }

    .title {
      padding: 0;
      width: 220px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      font-size: 18px;

      @media only screen and (max-device-width: 480px) {
        display: none;
      }

      .en {
        font-size: 12px;
      }
    }
  }
}
</style>
