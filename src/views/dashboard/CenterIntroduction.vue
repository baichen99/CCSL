<template>
  <div class="app-container">
    <el-tabs v-model="activeTab" type="border-card">
      <el-tab-pane :label="$t('Chinese')" name="zh-CN">
        <rich-text-editor v-model="$data['zh-CN']" height="600px" />
      </el-tab-pane>
      <el-tab-pane :label="$t('English')" name="en-US">
        <rich-text-editor v-model="$data['en-US']" height="600px" />
      </el-tab-pane>
      <div class="save">
        <el-button v-if="checkDiff" type="primary" plain @click="handleSave">{{ $t("Save") }}</el-button>
      </div>
    </el-tabs>
  </div>
</template>

<script>
import { GetAppInfo, UpdateAppInfo } from "@/api/systems";
import RichTextEditor from "@/components/form/RichTextEditor";
export default {
  name: "CenterIntroduction",
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
  created() {
    this.$nextTick(() => {
      this.initData();
    });
  },
  methods: {
    async initData() {
      const resZh = await GetAppInfo("introduction-zh-CN");
      this["zh-CN"] = resZh.data;
      this["original-zh-CN"] = resZh.data;
      const resEn = await GetAppInfo("introduction-en-US");
      this["en-US"] = resEn.data;
      this["original-en-US"] = resEn.data;
    },
    async handleSave() {
      const lang = this.activeTab;
      await UpdateAppInfo(`introduction-${lang}`, this[lang]);
      this.$notify({
        type: "success",
        title: this.$t("SuccessfulOperation")
      });
      this.initData();
    }
  }
};
</script>

<style lang="scss" scoped>
.el-tabs {
  margin: 30px;
}

.save {
  margin: 20px auto;
  text-align: center;
  .el-button {
    width: 200px;
  }
}
</style>