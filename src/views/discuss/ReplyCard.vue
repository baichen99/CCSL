<template>
  <el-card class="ReplyCard">
    <div class="user-info">
      <div>{{ reply.creator.name }}</div>
      <el-button v-if="showDelete" type="danger" plain @click="$emit('delete', reply.id)">{{ $t('Delete') }}</el-button>
    </div>
    <div class="content"> {{ reply.content }} </div>
  </el-card>
</template>

<script>
export default {
  name: "ReplyCard",
  props: {
    reply: {
      type: Object,
      default: ()=>({})
    }
  },
  computed: {
    IsSuper() {
      let roles = this.$store.getters.roles;
      if (roles.indexOf('super') > -1) {
        return true;
      }
      return false;
    },
    showDelete() {
      let id = this.$store.getters.id;
      return this.reply.creator.id == id || this.IsSuper;
    }
  },
}
</script>

<style lang="scss" scoped>
  .ReplyCard {
    width: 70%;
    .user-info {
      font-size: 15px;
      font-weight: 600;
      color: #444;
      display: flex;
      justify-content: space-between;
    }
    .content {
      padding: 10px;
    }
  }

</style>