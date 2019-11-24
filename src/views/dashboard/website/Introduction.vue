<template>
  <div class="app-container">
    <el-tabs v-model="activeTab" type="card">
      <el-tab-pane :label="$t('Chinese')" name="zh-CN">
        <rich-text-editor v-model="$data['zh-CN']" height="600px" />
      </el-tab-pane>
      <el-tab-pane :label="$t('English')" name="en-US">
        <rich-text-editor v-model="$data['en-US']" height="600px" />
      </el-tab-pane>
      <div class="table-toolbar">
        <el-button v-if="checkDiff" type="primary" plain @click="handleSave">{{ $t("Save") }}</el-button>
      </div>
    </el-tabs>
  </div>
</template>

<script>
import { GetAppInfo, UpdateAppInfo } from "@/api/systems";
import RichTextEditor from "@/components/form/RichTextEditor";
export default {
  name: "Introduction",
  components: {
    RichTextEditor
  },
  data() {
    return {
      activeTab: "zh-CN",
      "zh-CN": "",
      "en-US": "",
      "original-zh-CN": "",
      "original-en-US": ""
    };
  },
  computed: {
    checkDiff() {
      const lang = this.activeTab;
      return this[lang] !== this[`original-${lang}`];
    }
  },
  mounted() {
    this.initData();
  },
  methods: {
    initData() {
      GetAppInfo("introduction-zh-CN").then(res => {
        this["zh-CN"] = res.data;
        this["original-zh-CN"] = res.data;
      });
      GetAppInfo("introduction-en-US").then(res => {
        this["en-US"] = res.data;
        this["original-en-US"] = res.data;
      });
    },
    handleSave() {
      const lang = this.activeTab;
      UpdateAppInfo(`introduction-${lang}`, this[lang]).then(() => {
        this.$notify({
          type: "success",
          title: this.$t("SuccessfulOperation")
        });
        this.initData();
      });
    }
  }
};
</script>

<style lang="scss" scoped>
.el-tabs {
  margin: 30px;
}
</style>