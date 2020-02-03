<template>
  <el-card class="handshape-search" shadow="hover">
    <h3 class="title">{{ $t("title") }}</h3>
    <div class="handshape">
      <div v-for="item in handshapes" :key="item.id" class="handshape-box">
        <el-popover trigger="hover" placement="top">
          <div class="handshape-name-popover" style>{{ item.name }}</div>
          <img width="200px" :src="settings.publicURL + item.image" :alt="item.name" />
          <img
            slot="reference"
            style="width:100%"
            :src="settings.publicURL + item.image"
            :alt="item.name"
            @click="onHandshapeSelected(item.id)"
          />
        </el-popover>
      </div>
    </div>
  </el-card>
</template>

<i18n>
{
  "zh-CN": {
    "title": "手形检索"
  },
  "en-US": {
    "title": "Search by Handshape"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
export default {
  name: "HandshapeSearch",
  data() {
    return {
      handshapes: []
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  created() {
    this.$nextTick(async () => {
      await this.$store.dispatch("data/getHandshapes");
      const data = this.$store.state.data.handshapes;
      this.handshapes = Object.values(data);
    });
  },
  methods: {
    onHandshapeSelected(handshape) {
      this.$emit("handshape-selected", handshape);
    }
  }
};
</script>

<style lang="scss">
.handshape-name-popover {
  text-align: center;
  font-size: 20px;
  font-weight: bold;
  margin: 5px;
}
</style>

<style lang="scss" scoped>
.handshape-search {
  width: 18%;
  min-height: 600px;
  max-height: 80vh;
  overflow: auto;
  .title {
    text-align: center;
  }
  &::-webkit-scrollbar {
    display: none;
  }
  -ms-overflow-style: none;
  overflow: -moz-scrollbars-none;

  .handshape {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    .handshape-box {
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