<template>
  <list-view
    :get-list-method="GetWordsList"
    :create-item-method="CreateWord"
    :update-item-method="UpdateWord"
    :delete-item-method="DeleteWord"
    :list-form-component="LexiconForm"
    :export-list-config="exportConfig"
    :columns="columns"
    delete-warning="此操作会删除所有视频中该词语的标注，此操作将永久删除, 是否继续?"
    entity="lexicons"
    order-by="initial,id"
  >
    <template #toolbar-search="{params, handleSearch}">
      <search-input v-model="params.chinese" :placeholder="$t('tipZh')" @update="handleSearch" />
      <search-input v-model="params.english" :placeholder="$t('tipEn')" @update="handleSearch" />
    </template>

    <template #chinese="{row}">
      <span class="word-sup" v-html="$options.filters.addNumberSup(row.chinese) " />
    </template>

    <template #pos="{row}">
      <el-tag
        v-for="value in row.pos"
        :key="value"
        :disable-transitions="true"
        size="small"
        class="tags"
      >
        <span>{{ $t($options.filters.getObjectItem(partOfSpeech, value).text) }}</span>
      </el-tag>
    </template>
  </list-view>
</template>

<i18n>
{
  "zh-CN":{
    "tipZh": "请输入中文",
    "tipEn": "请输入英文"
  },
  "en-US":{
    "tipZh": "Input Chinese",
    "tipEn": "Input English"
  }
}
</i18n>

<script>
import { mapGetters } from "vuex";
import LexiconForm from "@/views/dashboard/form/LexiconForm";
import SearchInput from "@/components/form/SearchInput";
import ListView from "@/components/ListView";
import {
  GetWordsList,
  CreateWord,
  UpdateWord,
  DeleteWord
} from "@/api/lexicons";
export default {
  name: "LexiconsList",
  components: {
    ListView,
    SearchInput
  },
  data() {
    return {
      GetWordsList,
      CreateWord,
      UpdateWord,
      DeleteWord,
      LexiconForm,
      columns: [
        {
          prop: "initial",
          label: this.$t("Initial"),
          width: "80px",
          filters: []
        },
        {
          slot: "chinese",
          label: this.$t("Chinese"),
          width: "150px",
          hideOverflow: true
        },
        {
          prop: "english",
          label: this.$t("English"),
          width: "150px",
          hideOverflow: true
        },
        {
          slot: "pos",
          label: this.$t("PoS"),
          width: "200px",
          filters: []
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["wordInitial", "partOfSpeech"])
  },
  created() {
    this.wordInitial.data.map(item => {
      this.columns[0].filters.push({
        text: item.text,
        value: item.value
      });
    });
    this.partOfSpeech.data.map(item => {
      this.columns[3].filters.push({
        text: this.$t(item.text),
        value: item.value
      });
    });
  },
  methods: {
    exportConfig(item) {
      const partOfSpeech = [];
      item.pos.map(k => {
        const v = this.$t(
          this.$options.filters.getObjectItem(this.partOfSpeech, k).text
        );
        partOfSpeech.push(v);
      });
      return {
        [this.$t("CreatedAt")]: new Date(item.createdAt),
        [this.$t("UpdatedAt")]: new Date(item.updatedAt),
        [this.$t("Initial")]: item.initial,
        [this.$t("Chinese")]: item.chinese,
        [this.$t("English")]: item.english,
        [this.$t("PoS")]: partOfSpeech.join("/")
      };
    }
  }
};
</script>

<style lang="scss" scoped>
.tags {
  margin: 3px;
}
</style>