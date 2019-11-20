<template>
  <div>
    <el-tag
      v-for="tag in data"
      :key="tag"
      closable
      :disable-transitions="false"
      @close="handleClose(tag)"
    >{{ tag }}</el-tag>
    <el-input
      v-if="inputVisible"
      ref="saveTagInput"
      v-model="inputValue"
      class="input-new-tag"
      size="small"
      @keyup.enter.native="handleInputConfirm"
      @blur="handleInputConfirm"
    />
    <el-button v-else class="button-new-tag" plain size="small" type="primary" @click="showInput">+ {{ $t("New") }}</el-button>
  </div>
</template>
<script>
export default {
  name: "MorphemesPicker",
  model: {
    prop: "value",
    event: "update"
  },
  props: {
    value: {
      type: Array,
      default: () => []
    }
  },
  data() {
    return {
      inputVisible: false,
      inputValue: ""
    };
  },
  computed: {
    data: {
      get() {
        return this.value;
      },
      set(val) {
        this.$emit("update", val);
      }
    }
  },
  methods: {
    handleClose(tag) {
      this.data.splice(this.data.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        if (this.data) {
          this.data.push(inputValue);
        } else {
          this.data = [inputValue];
        }
        this.$emit("update", this.data);
      }
      this.inputVisible = false;
      this.inputValue = "";
    }
  }
};
</script>

<style lang="scss" scoped>
.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding: 0 10px;
}
.input-new-tag {
  width: 90px !important;
  margin-left: 10px;
}
</style>