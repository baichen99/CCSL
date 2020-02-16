<template>
  <div class="home flex-column">
    <carousel-card />
    <div class="news-container">
      <news-card
        v-for="item in newsColumns.data"
        :key="item.value"
        :icon="item.icon"
        :column="item.value"
        :title="item.text"
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
  computed: { ...mapGetters(["newsColumns"]) }
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
