<template>
  <el-card
    class="sidebar"
    shadow="hover"
  >
    <h3 class="title">音序检索</h3>
    <div class="initial">
      <el-collapse accordion>
        <el-collapse-item
          v-for="(value, key) in initWords"
          :key="key"
          :title="key"
        >
          <div
            v-for="word in value"
            :key="word.id"
          >
            <el-link
              type="primary"
              @click="onWordSelected(word.id)"
            >{{ word.chinese }}</el-link>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </el-card>
</template>

<script>
export default {
  name: "WordSelector",
  props: {
    start: {
      type: String,
      default: "A"
    },
    end: {
      type: String,
      default: "Z"
    }
  },
  computed: {
    initWords() {
      let dict = {};
      const words = this.$store.getters.words;
      for (let key in words) {
        if (this.start <= key && key <= this.end) {
          dict[key] = words[key];
        }
      }
      return dict;
    }
  },
  methods: {
    onWordSelected(wordID) {
      this.$emit("word-selected", wordID);
    }
  }
};
</script>

<style lang="scss" scoped>
.sidebar {
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
  .initial {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;
    .el-collapse {
      width: 100%;
    }
  }
}
</style>