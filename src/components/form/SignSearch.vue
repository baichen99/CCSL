<template>
  <el-card class="sign-search" shadow="hover">
    <h3 class="title">{{ $t("title") }}</h3>
    <div class="sign">
      <div v-for="item in signs" :key="item.id" class="sign-box">
        <el-popover trigger="hover" placement="top">
          <div class="sign-name-popover" style>{{ item.name }}</div>
          <img width="200px" :src="settings.publicURL + item.image" :alt="item.name" />
          <img
            slot="reference"
            style="width:100%"
            :src="settings.publicURL + item.image"
            :alt="item.name"
            @click="onSignSelected(item.id)"
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
    "title": "Search by Sign"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import { GetSignsList } from "@/api/signs";
export default {
  name: "SignSearch",
  data() {
    return {
      signs: []
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  created() {
    GetSignsList({ limit: 0 }).then(res => {
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

<style lang="scss">
.sign-name-popover {
  text-align: center;
  font-size: 20px;
  font-weight: bold;
  margin: 5px;
}
</style>

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