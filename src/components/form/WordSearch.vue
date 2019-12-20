<template>
  <el-card class="sidebar" shadow="hover">
    <h3 class="title">{{ $t("title") }}</h3>
    <div class="initial">
      <el-collapse accordion>
        <el-collapse-item v-for="(value, key) in words" :key="key" :title="key">
          <div v-for="word in value" :key="word.id">
            <el-link type="primary" @click="onWordSelected(word.id)">
              <span v-if="$i18n.locale==='en-US'">{{ word.english }}</span>
              <span v-else class="word-sup" v-html="$options.filters.addNumberSup(word.chinese)"></span>
            </el-link>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </el-card>
</template>

<i18n>
{
  "zh-CN": {
    "title": "音序检索"
  },
  "en-US": {
    "title": "Search by Initial"
  }
}
</i18n>

<script>
import { GetWordsList } from "@/api/lexicons";
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
  watch: {
    "$i18n.locale"() {
      this.getData();
    }
  },
  created() {
    this.getData();
  },
  methods: {
    onWordSelected(wordID) {
      this.$emit("word-selected", wordID);
    },
    getData() {
      GetWordsList({ limit: 0 }).then(res => {
        let wordsDict = {};
        const wordsArray = res.data;
        wordsArray.map(item => {
          let initial;
          if (this.$i18n.locale == "en-US") {
            initial = item.english[0].toUpperCase();
          } else {
            initial = item.initial;
          }
          if (wordsDict[initial]) {
            wordsDict[initial].push(item);
          } else {
            wordsDict[initial] = [item];
          }
        });
        let dict = {};
        Object.keys(wordsDict)
          .sort()
          .map(key => {
            if (this.start <= key && key <= this.end) {
              dict[key] = wordsDict[key];
            }
          });
        this.words = dict;
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.sidebar {
  width: 18%;
  min-height: 600px;
  max-height: 80vh;
  overflow: auto;
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