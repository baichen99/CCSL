<template>
  <el-form ref="form" label-position="left" label-width="100px" :model="formData" :rules="rules"></el-form>
</template>

<script>
import formMixin from "./formMixin";
export default {
  name: "LexicalVideoForm",
  mixins: [formMixin],
  data() {
    return {
      rules: {
        performerID: [],
        lexicalWordID: [],
        videoPath: [],
        constructType: [],
        constructWords: [],
        leftSigns: [],
        rightSigns: []
      }
    };
  }
};
</script>