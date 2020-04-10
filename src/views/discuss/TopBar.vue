<template>
  <div class="container">
    <div class="topBar">
      <el-input v-model="params.title" :placeholder="$t('searchPlaceholder')">
        <i
          slot="suffix"
          class="el-input__icon el-icon-search"
          style="cursor:pointer;"
          @click="$emit('search')"
        />
      </el-input>
      <div class="btns">
        <el-button type="primary" round @click="toggleForm">{{ $t('createPost') }}</el-button>
      </div>
    </div>
    <el-collapse-transition>
      <div v-show="showForm" class="PostCreateForm">
        <el-input v-model="form.title" class="title" :placeholder="$t('titlePlaceholder')" />
        <RichTextEditor v-model="form.content" />
        <div class="btns">
          <el-button type="primary" @click="submit">{{ $t('Publish') }}</el-button>
          <el-button type="warning" @click="toggleForm">{{ $t('Cancel') }}</el-button>
        </div>
      </div>
    </el-collapse-transition>
  </div>
</template>

<i18n>
{
  "zh-CN": {
    "searchPlaceholder": "搜索话题",
    "titlePlaceholder": "标题",
    "createPost": "创建帖子"
  },
  "en-US": {
    "searchPlaceholder": "Search topic",
    "titlePlaceholder": "title",
    "createPost": "Create Post"
  }
}
</i18n>

<script>
import RichTextEditor from "@/components/form/RichTextEditor"
import { CreatePost } from '@/api/posts'
export default {
  components: {
    RichTextEditor
  },
  model: {
    prop: 'params',
    event: 'update:params'
  },
  props: {
    params: {
      type: Object,
      default: ()=>({})
    }
  },
  data() {
    return {
      searchPlaceholder: "搜索话题",
      showForm: false,
      form: {
        title: "",
        content: ""
      }
    }
  },
  methods: {
    toggleForm() {
      this.showForm = !this.showForm;
    },
    async submit() {
      this.loading = true;
      try {
        await CreatePost(this.form).
          catch((err) => {
            this.$notify.info({
              title: this.$t(err),
              duration: 2000
            });
          });
        this.$emit('createPost', this.form)
        this.$notify.info({
          title: this.$t("SuccessfulOperation"),
          duration: 2000
        });
      } finally {
        this.loading = false;
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.container {
  width: 100%;
  .topBar {
      width: 80%;
      height: 60px;
      border-bottom: 1px solid #000;
      display: flex;
      align-items: center;
      justify-content: space-between;
      .el-input {
          width: 30%;
      }
  }
  .PostCreateForm {
    margin-top: 20px;
    width: 80%;
    display: flex;
    flex-direction: column;
    .title {
      width: 300px;
      margin-bottom: 10px;
    }
    .btns {
      align-self: flex-end;
    }
  }
}

</style>
