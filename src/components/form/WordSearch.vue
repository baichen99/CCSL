<template>
  <el-card
    class="sidebar"
    shadow="hover"
  >
    <h3 class="title">音序检索</h3>
    <div class="initial">
      <el-collapse accordion>
        <el-collapse-item
          v-for="(value, key) in words"
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
import { GetLexicalWordsList } from "@/api/words";
export default {
  name: "WordSearch",
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
  data() {
    return {
      words: {}
    };
  },
  created() {
    GetLexicalWordsList({ limit: 0 }).then(res => {
      let wordsDict = {};
      const wordsArray = res.data;
      wordsArray.map(item => {
        const initial = item.initial;
        if (wordsDict[initial]) {
          wordsDict[initial].push(item);
        } else {
          wordsDict[initial] = [item];
        }
      });
      let dict = {};
      for (let key in wordsDict) {
        if (this.start <= key && key <= this.end) {
          dict[key] = wordsDict[key];
        }
      }
      this.words = dict;
    });
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