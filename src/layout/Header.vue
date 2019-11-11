<template>
  <header>
    <div>
      <div class="header-logo" @click="$router.push('/')">
        <svg-icon icon-class="logo" class="logo" />
        <div class="title">
          <span>中国手语及聋人研究中心</span>
          <span class="en">Center for CSL and Deaf Studies</span>
        </div>
      </div>
      <el-menu
        :default-active="activeMenu"
        mode="horizontal"
        router
        active-text-color="#2363C3"
      >
        <el-menu-item index="/">首页</el-menu-item>
        <el-menu-item v-permission="['admin','super','learner']" index="/learning-platform">学习平台</el-menu-item>
        <el-submenu index="/research">
          <template slot="title">研究成果</template>
          <el-menu-item
            v-for="item in researches"
            :key="item.url"
            :index="item.url"
            class="sub-menu-item"
          >{{ item.title }}</el-menu-item>
        </el-submenu>
        <el-submenu v-permission="['admin','super','user']" index="/database">
          <template slot="title">数据库</template>
          <el-menu-item
            v-for="item in databases"
            :key="item.url"
            :index="item.url"
            class="sub-menu-item"
          >{{ item.title }}</el-menu-item>
        </el-submenu>
        <el-submenu index="/about">
          <template slot="title">中心概况</template>
          <el-menu-item
            v-for="item in about"
            :key="item.url"
            :index="item.url"
            class="sub-menu-item"
          >{{ item.title }}</el-menu-item>
        </el-submenu>
        <el-submenu v-if="userToken" index="/profile">
          <template slot="title">我的</template>
          <el-menu-item index="/profile" class="sub-menu-item">个人中心</el-menu-item>
          <el-menu-item class="sub-menu-item" @click="logout">退出登录</el-menu-item>
        </el-submenu>
        <el-menu-item v-else index="/login">登录</el-menu-item>
      </el-menu>
    </div>
  </header>
</template>

<script>
import { getToken } from "@/utils/tools";
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
      const route = this.$route;
      const { meta, path } = route;
      if (meta.activeMenu) {
        return meta.activeMenu;
      }
      return path;
    },
    userToken() {
      return getToken() || this.$store.getters.token;
    }
  },
  methods: {
    async logout() {
      await this.$store.dispatch("user/logout");
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

    .header-logo {
      cursor: pointer;
    }

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

.sub-menu-item {
  text-align: center;
}
</style>
