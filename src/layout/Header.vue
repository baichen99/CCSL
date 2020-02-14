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
      <el-menu :default-active="activeMenu" mode="horizontal" router active-text-color="#2363C3">
        <el-menu-item index="/">{{ $t("HomeMenu") }}</el-menu-item>
        <el-menu-item v-permission="[StudentUser]" index="/learning-platform">{{ $t("LearnMenu") }}</el-menu-item>
        <el-submenu index="/research">
          <template slot="title">{{ $t("ResearchMenu") }}</template>
          <el-menu-item
            v-for="item in researches"
            :key="item.url"
            :index="item.url"
          >{{ $t(item.title) }}</el-menu-item>
        </el-submenu>
        <el-submenu v-permission="[DatabaseUser]" index="/database">
          <template slot="title">{{ $t("DatabaseMenu") }}</template>
          <el-menu-item
            v-for="item in databases"
            :key="item.url"
            :index="item.url"
          >{{ $t(item.title) }}</el-menu-item>
        </el-submenu>
        <el-submenu index="/about">
          <template slot="title">{{ $t("AboutMenu") }}</template>
          <el-menu-item v-for="item in about" :key="item.url" :index="item.url">{{ $t(item.title) }}</el-menu-item>
        </el-submenu>
        <el-submenu index="/language">
          <template slot="title">{{ languages[$i18n.locale].name }}</template>
          <el-menu-item v-for="(value, key) in languages" :key="key" @click="changeLanguage(key)">
            <svg-icon class="language-flag" :icon-class="key" />
            <span>{{ value.name }}</span>
          </el-menu-item>
        </el-submenu>
        <el-submenu v-if="login" index="/user-profile">
          <template slot="title">{{ $t("MyMenu") }}</template>
          <el-menu-item index="/user-profile">{{ $t("ProfileMenu") }}</el-menu-item>
          <el-menu-item @click="logout">{{ $t("Logout") }}</el-menu-item>
        </el-submenu>
        <el-menu-item v-else index="/login">{{ $t("Login") }}</el-menu-item>
      </el-menu>
    </div>
  </header>
</template>

<script>
import { DatabaseUser, StudentUser } from "@/router/routes";
export default {
  name: "Header",
  data: () => ({
    defaultActive: "/",
    researches: [
      { url: "/academic/projects", title: "ProjectsMenu" },
      { url: "/academic/papers", title: "PapersMenu" },
      { url: "/academic/literature", title: "PublicationsMenu" }
    ],
    databases: [
      { url: "/database/lexical-database", title: "LexicalDatabase" },
      { url: "/database/verb-corpus", title: "VerbCorpus" },
      { url: "/database/proper-nouns-corpus", title: "ProperNounsCorpus" },
      { url: "/database/text-corpus", title: "TextCorpus" },
      { url: "/database/literature-database", title: "LiteratureDatabase" },
      { url: "/database/term-database", title: "TermDatabase" }
    ],
    about: [
      { url: "/about/introduction", title: "IntroductionMenu" },
      { url: "/about/team", title: "TeamMenu" },
      { url: "/about/contact", title: "ContactMenu" }
    ],
    languages: {
      "zh-CN": { name: "简体中文" },
      "en-US": { name: "English(US)" }
    },
    DatabaseUser,
    StudentUser
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
    login() {
      return this.$store.getters.token;
    }
  },
  methods: {
    changeLanguage(lang) {
      this.$root.$i18n.locale = lang;
      this.$store.dispatch("app/setLocale", lang);
    },
    async logout() {
      await this.$store.dispatch("user/logout");
    }
  }
};
</script>

<style lang="scss">
.el-menu--popup {
  text-align: center;
}
.language-flag {
  font-size: 35px;
  float: left;
}
</style>

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
</style>
