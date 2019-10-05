<template>
  <div class="app-container">
    <div class="table-toolbar">
      <el-input />
      <el-button
        class="new"
        type="primary"
        @click="create"
      >增加</el-button>
    </div>

  </div>
</template>

<script>
export default {
  name: "AllUsers",
  methods: {
    create() {}
  }
};
</script>

<style lang="scss" scoped>
</style>

