<template>
  <div class="home">
    <carousel-card />
    <el-rows
      class="news"
      :gutter="12"
    >
      <el-col
        v-for="(value, key) in news"
        :key="key"
        :md="12"
        :sm="24"
      >

      </el-col>
    </el-rows>
  </div>
</template>
<script>
import CarouselCard from "@/views/home/CarouselCard.vue";
export default {
  name: "Home",
  components: {
    CarouselCard
  },
  data: () => ({
    news: []
  })
};
</script>


<style lang="scss" scoped>
.home {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
</style>
