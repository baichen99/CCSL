<template>
  <header>
    <div>
      <div>
        <svg-icon icon-class="logo" class="logo" />
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
        <el-menu-item index="/learning-platform">学习平台</el-menu-item>
        <el-submenu index="/research">
          <template slot="title">研究成果</template>
          <el-menu-item
            v-for="item in researches"
            :key="item.url"
            :index="item.url"
            style="text-align:center"
          >{{ item.title }}</el-menu-item>
        </el-submenu>
        <el-submenu index="/database">
          <template slot="title">数据库</template>
          <el-menu-item
            v-for="item in databases"
            :key="item.url"
            :index="item.url"
            style="text-align:center;"
          >{{ item.title }}</el-menu-item>
        </el-submenu>
        <el-submenu index="/about">
          <template slot="title">中心概况</template>
          <el-menu-item
            v-for="item in about"
            :key="item.url"
            :index="item.url"
            style="text-align:center;"
          >{{ item.title }}</el-menu-item>
        </el-submenu>
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
      { url: "/academic/projects", title: "科研项目" },
      { url: "/academic/papers", title: "论文发表" },
      { url: "/academic/literature", title: "专著出版" }
    ],
    databases: [
      { url: "/database/lexical-database", title: "国家通用手语比对语料库" },
      { url: "/database/verb-corpus", title: "上海手语动词语料库" },
      { url: "/database/proper-nouns-corpus", title: "专有名词语料库" },
      { url: "/database/text-corpus", title: "手语语篇数据库" },
      { url: "/database/literature-database", title: "手语研究文献数据库" },
      { url: "/database/term-database", title: "手语研究术语库" }
    ],
    about: [
      { url: "/about/introduction", title: "中心介绍" },
      { url: "/about/team", title: "研究团队" },
      { url: "/about/contact", title: "联系我们" }
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
@import "@/styles/element-variables.scss";
header {
  height: 60px;

  div {
    height: 100%;
    padding: 0 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    // img {
    //   @media only screen and (max-device-width: 480px) {
    //     display: none;
    //   }
    // }
    .logo {
      fill: $--color-primary !important;
      font-size: 70px;
    }

    .title {
      padding: 0;
      width: 220px;
      display: flex;
      flex-direction: column;
      justify-content: center;
      font-size: 18px;

      // @media only screen and (max-device-width: 480px) {
      //   display: none;
      // }

      .en {
        font-size: 12px;
      }
    }
  }
}
</style>
