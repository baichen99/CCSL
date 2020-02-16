<template>
  <div class="team flex-column">
    <h2 class="title">{{ $t("TeamMenu") }}</h2>
    <el-card shadow="hover">
      <div slot="header">
        <div>
          <el-link type="primary" class="back" @click="$router.push('/about/team')">{{ $t("Back") }}</el-link>
        </div>
      </div>
      <div class="profile">
        <img v-if="data.profile" :src="settings.publicURL + data.profile" alt="profile" />
        <h3>
          <span v-if="$i18n.locale==='en-US'">{{ data.nameEn }}</span>
          <span v-else>{{ data.nameZh }}</span>
        </h3>
        <div v-if="$i18n.locale==='en-US'" class="description" v-html="data.descriptionEn"></div>
        <div v-else class="description" v-html="data.descriptionZh"></div>
      </div>
    </el-card>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { GetMember } from "@/api/members";
export default {
  name: "TeamDetail",
  data() {
    return {
      data: {}
    };
  },
  computed: {
    ...mapGetters(["settings"])
  },
  created() {
    this.getData();
  },
  methods: {
    async getData() {
      const id = this.$route.params.id;
      const res = await GetMember(id);
      this.data = res.data;
    }
  }
};
</script>

<style lang="scss" scoped>
@import "@/styles/element-variables.scss";
.team {
  align-items: center;

  .title {
    color: $--color-primary;
    margin: 20px;
  }

  .el-card {
    width: 900px;
    margin: 1rem auto;
    padding: 1rem;

    .profile {
      text-align: center;
      padding: 10px 30px;
      img {
        max-width: 400px;
        border-radius: 5px;
      }
      .description {
        text-align: left;
        text-indent: 2em;
        line-height: 2rem;
      }
    }
  }
}
</style>