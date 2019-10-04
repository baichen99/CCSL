<template>
  <div
    id="app"
    class="loading-area"
  >
    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
  mounted() {
    this.initData();
  },
  methods: {
    initData() {
      this.$store.dispatch("user/refreshToken");
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
  -ms-overflow-style: none;
  overflow: -moz-scrollbars-none;
  &::-webkit-scrollbar {
    display: none;
  }
}

#app {
  height: 100%;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
