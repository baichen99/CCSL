<template>
  <el-card class="sidebar">
    <h3 class="title">音序检索</h3>
    <div class="initial">
      <el-collapse accordion>
        <el-collapse-item
          v-for="(value, key) in $store.getters.words"
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
import { mapGetters } from "vuex";
export default {
  name: "WordSelector",
  computed: {
    ...mapGetters(["getWordsByInitial"])
  },
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
  methods: {
    onWordSelected(wordID) {
      console.log(wordID);
    }
  }
};
</script>

<style lang="scss" scoped>
.sidebar {
  width: 18%;
  height: 600px;
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