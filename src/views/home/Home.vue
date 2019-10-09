<template>
  <div class="home flex-column">
    <carousel-card />
    <div class="news-container">
      <news-card
        v-for="(value, key) in newsTypes"
        :key="key"
        :icon="value.icon"
        :column="value.icon"
        :title="value.name"
      />
    </div>
  </div>
</template>
<script>
import { mapGetters } from "vuex";
import CarouselCard from "@/views/home/CarouselCard.vue";
import NewsCard from "@/views/home/NewsCard.vue";
export default {
  name: "Home",
  components: {
    CarouselCard,
    NewsCard
  },
  computed: { ...mapGetters(["newsTypes"]) }
};
</script>


<style lang="scss" scoped>
.home {
  align-items: center;
  .news-container {
    width: 1050px;
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    justify-content: space-around;
  }
}
</style>
