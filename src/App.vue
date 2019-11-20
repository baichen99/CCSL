<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
  watch: {
    "$i18n.locale"() {
      this.onLanguageChanged();
    }
  },
  mounted() {
    this.initData();
  },
  methods: {
    initData() {
      this.$store.dispatch("user/refreshToken");
      this.onLanguageChanged();
    },
    onLanguageChanged() {
      const lang = this.$i18n.locale;
      switch (lang) {
        case "en-US":
          document.title = "Center for CSL and Deaf Studies";
          break;

        default:
          document.title = "中国手语及聋人研究中心";
          break;
      }
    }
  }
};
</script>

<style lang="scss">
html,
body {
  height: 100%;
  width: 100%;
  padding: 0;
  margin: 0;
}

#app {
  height: 100%;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
