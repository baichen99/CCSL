<template>
  <div class="introduction flex-column">
    <h2 class="title">{{ $t("IntroductionMenu") }}</h2>
    <el-card>
      <div class="content" v-html="introduction"></div>
    </el-card>
  </div>
</template>

<script>
import { GetAppInfo } from "@/api/systems";
export default {
  name: "Introduction",
  data() {
    return {
      introduction: ""
    };
  },
  watch: {
    "$i18n.locale"() {
      this.initData();
    }
  },
  created() {
    this.initData();
  },
  methods: {
    initData() {
      const lang = this.$i18n.locale;
      GetAppInfo(`introduction-${lang}`).then(res => {
        this.introduction = res.data;
      });
    }
  }
};
</script>


<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.introduction {
  align-items: center;

  .title {
    color: $--color-primary;
    margin: 20px;
  }
  .el-card {
    width: 900px;
    margin: 1rem auto;
    padding: 2rem;
    .content {
      text-indent: 2rem;
      padding: 10px 30px;
      text-indent: 2em;
      line-height: 2rem;
    }
  }
}
</style>
