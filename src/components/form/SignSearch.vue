<template>
  <el-card
    class="sign-search"
    shadow="hover"
  >
    <h3 class="title">手形检索</h3>
    <div class="sign">
      <div
        v-for="item in signs"
        :key="item.id"
        class="sign-box"
      >
        <el-tooltip
          effect="dark"
          :content="item.name"
          placement="top"
        >
          <img
            style="width:100%"
            :src="'https://ccsl.shu.edu.cn/public/'+item.image"
            :alt="item.name"
            @click="onSignSelected(item.id)"
          >
        </el-tooltip>
      </div>
    </div>
  </el-card>
</template>

<script>
import { getSigns } from "@/api/signs";
export default {
  name: "SignSearch",
  data() {
    return {
      signs: []
    };
  },
  created() {
    getSigns({ limit: 0 }).then(res => {
      this.signs = res.data;
    });
  },
  methods: {
    onSignSelected(sign) {
      this.$emit("sign-selected", sign);
    }
  }
};
</script>

<style lang="scss" scoped>
.sign-search {
  width: 18%;
  min-height: 600px;
  max-height: 80vh;
  overflow: scroll;
  .title {
    text-align: center;
  }
  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  overflow: -moz-scrollbars-none;

  .sign {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    .sign-box {
      width: 33%;
      padding: 5px;
      img {
        cursor: pointer;
        max-height: 60px;
      }
    }
  }
}
</style>