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
    <template #toolbar="{params, handleSearch}">
      <el-input
        v-model="params.chinese"
        prefix-icon="el-icon-search"
        placeholder="请输入中文"
        clearable
        @keyup.enter="handleSearch"
        @change="handleSearch"
      />
      <el-input
        v-model="params.english"
        prefix-icon="el-icon-search"
        placeholder="请输入英文"
        clearable
        @change="handleSearch"
      />
      <pos-selector v-model="params.pos" @update="handleSearch" />
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
        <span>{{ $t(partOfSpeech[value].name) }}</span>
      </el-tag>
    </template>

    <template #action="{row,handleDeleteItem}">
      <el-button
        type="danger"
        size="mini"
        plain
        @click.stop="handleDeleteItem(row.id)"
      >{{ $t("Delete") }}</el-button>
    </template>
  </list-view>
</template>

<script>
import { mapGetters } from "vuex";
import LexiconForm from "@/views/dashboard/form/LexiconForm";
import PosSelector from "@/components/form/PosSelector";
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
    PosSelector
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
          width: "200px"
        },
        {
          slot: "action",
          label: this.$t("Action"),
          width: "90px",
          fixed: "right"
        }
      ]
    };
  },
  computed: {
    ...mapGetters(["wordInitial", "partOfSpeech"])
  },
  created() {
    this.wordInitial.map(item => {
      this.columns[0].filters.push({
        text: item,
        value: item
      });
    });
  },
  methods: {
    exportConfig(item) {
      const partOfSpeech = [];
      item.pos.map(k => {
        const v = this.$t(this.partOfSpeech[k].name);
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