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
        :default-active="defaultActive"
        mode="horizontal"
        :router="true"
        active-text-color="#2363C3"
      >
        <el-menu-item index="/">首页</el-menu-item>
        <el-menu-item index="/introduction">中心简介</el-menu-item>
        <el-menu-item index="/team">研究团队</el-menu-item>
        <el-submenu index="/research">
          <template slot="title">研究成果</template>
          <el-menu-item
            v-for="item in researches"
            :key="item.url"
            :index="item.url"
            style="text-align:center"
          >
            {{ item.title }}
          </el-menu-item>
        </el-submenu>
        <el-submenu index="/database">
          <template slot="title">数据库</template>
          <el-menu-item
            v-for="item in databases"
            :key="item.url"
            :index="item.url"
            style="text-align:center;"
          >
            {{ item.title }}
          </el-menu-item>
        </el-submenu>
        <el-menu-item index="/learning-platform">学习平台</el-menu-item>
        <el-menu-item index="/contact">联系我们</el-menu-item>
      </el-menu>
    </div>
  </header>
</template>

<script>
export default {
  name: "Header",

  data: () => ({
    defaultActive: "/",
    researches: [
      {
        url: "/projects",
        title: "科研项目"
      },
      {
        url: "/papers",
        title: "论文发表"
      },
      {
        url: "/books",
        title: "专著出版"
      }
    ],
    databases: [
      {
        url: "/universal-contrast",
        title: "国家通用手语比对语料库"
      },
      {
        url: "/shanghai-verb",
        title: "上海手语动词语料库"
      },
      {
        url: "/proper-noun",
        title: "专有名词语料库"
      },
      {
        url: "/signlang-article",
        title: "手语语篇数据库"
      },
      {
        url: "/research-papers",
        title: "手语研究文献数据库"
      },
      {
        url: "signlang-term",
        title: "手语研究术语库"
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
  created() {
    this.setActiveMenu();
  },
  methods: {
    setActiveMenu() {
      this.defaultActive = this.$route.path;
    }
  }
};
</script>

<style lang="scss" scoped>
header {
  height: 60px;

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
