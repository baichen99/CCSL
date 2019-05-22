<template>
  <div class="container">
    <sl-header />
    <main class="index">
      <router-view />
    </main>
    <sl-footer />
  </div>
</template>

<script>
import SlHeader from "@/views/Header.vue";
import SlFooter from "@/views/Footer.vue";
export default {
  name: "Index",
  components: {
    SlHeader,
    SlFooter
  },
  methods: {}
};
</script>

<style lang="scss" scoped>
.container {
  display: flex;
  flex-direction: column;
  height: 100%;

  header {
    flex: 0 0 auto;
  }

  main {
    background-color: #e6e6e6;
    flex: 1 0 auto;
  }

  footer {
    flex: 0 0 auto;
  }
}
</style>
