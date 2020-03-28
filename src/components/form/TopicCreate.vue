<template>
  <el-form class="TopicCreate">
    <el-form-item class="title">
      <el-input v-model="params.title" placeholder="创建话题"></el-input>
    </el-form-item>
    <el-form-item class="content">
      <el-input v-model="params.content" type="textarea" :autosize="{ minRows: 4, maxRows: 6}"></el-input>
    </el-form-item>
    <el-form-item class="btns">
      <el-button type="primary">确定</el-button>
      <el-button type="warning">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  models: {
    prop: "params",
    event: "update"

  },
  props: {
    params: {
      type: Object,
      default: ()=>({})
    }
  },
  data() {
      return {
          form: {
              title: "",
              content: ""
          }
      }
  }
}
</script>

<style lang="scss" scoped>
  .el-form {
    display: flex;
    flex-direction: column;
    .title {
        width: 300px;
    }
    .content {
        width: 100%;
    }
    .btns {
        align-self: flex-end;
    }
  }

</style>