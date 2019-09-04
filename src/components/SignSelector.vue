<template>
  <el-card
    class="sign-search"
    shadow="hover"
  >
    <h3 class="title">右手手形检索</h3>
    <div class="sign">
      <div
        v-for="(item,index) in $store.getters.sign"
        :key="index"
        class="sign-box"
      >
        <el-tooltip
          effect="dark"
          :content="item"
          placement="top"
        >
          <img
            style="width:100%"
            :src="'sign/'+item+'.jpg'"
            :alt="item"
            @click="onSignSelected(item)"
          >
        </el-tooltip>
      </div>
    </div>
  </el-card>
</template>

<script>
export default {
  name: "SignSelector",
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
  height: 600px;
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
      width: 50%;
      img {
        cursor: pointer;
      }
    }
  }
}
</style>