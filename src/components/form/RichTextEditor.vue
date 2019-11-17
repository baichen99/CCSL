<template>
  <div
    v-loading="loading"
    :class="{fullscreen:fullscreen}"
    class="tinymce-container"
    :style="{width:containerWidth}"
  >
    <textarea :id="tinymceId" class="tinymce-textarea" />
  </div>
</template>

<script>
const plugins = [
  "lists advlist anchor autolink code codesample directionality fullscreen hr image imagetools insertdatetime link media nonbreaking noneditable pagebreak paste preview print save searchreplace spellchecker tabfocus table template textpattern visualblocks visualchars wordcount"
];

const toolbar = [
  "searchreplace bold italic underline strikethrough alignleft aligncenter alignright outdent indent  blockquote undo redo removeformat subscript superscript code codesample hr bullist numlist link image charmap preview anchor pagebreak insertdatetime media table forecolor backcolor fullscreen"
];

const menubar = "file edit insert view format table";

const tinymceCDN = "https://cdn.jsdelivr.net/npm/tinymce@5.1.1/tinymce.min.js";

let callbacks = [];

import { UploadFile } from "@/api/files";

export default {
  name: "RichTextEditor",
  model: {
    prop: "value",
    event: "input"
  },
  props: {
    id: {
      type: String,
      default: function() {
        return (
          "vue-tinymce-" +
          +new Date() +
          ((Math.random() * 1000).toFixed(0) + "")
        );
      }
    },
    value: {
      type: String,
      default: ""
    },
    height: {
      type: [Number, String],
      required: false,
      default: 360
    },
    width: {
      type: [Number, String],
      required: false,
      default: "auto"
    }
  },
  data() {
    return {
      loading: false,
      hasChange: false,
      hasInit: false,
      tinymceId: this.id,
      fullscreen: false,
      languageTypeList: {
        "en-US": "en_US",
        "zh-CN": "zh_CN"
      },
      callbacks: []
    };
  },
  computed: {
    containerWidth() {
      const width = this.width;
      if (/^[\d]+(\.[\d]+)?$/.test(width)) {
        // matches `100`, `'100'`
        return `${width}px`;
      }
      return width;
    }
  },
  watch: {
    value(val) {
      if (!this.hasChange && this.hasInit) {
        this.$nextTick(() =>
          window.tinymce.get(this.tinymceId).setContent(val || "")
        );
      }
    }
  },
  mounted() {
    this.init();
  },
  activated() {
    if (window.tinymce) {
      this.initTinymce();
    }
  },
  deactivated() {
    this.destroyTinymce();
  },
  destroyed() {
    this.destroyTinymce();
  },
  methods: {
    dynamicLoadScript(src, callback) {
      const existingScript = document.getElementById(src);
      const cb = callback || function() {};
      if (!existingScript) {
        const script = document.createElement("script");
        script.src = src; // src url for the third-party library being loaded.
        script.id = src;
        document.body.appendChild(script);
        callbacks.push(cb);
        script.onload = function() {
          this.onerror = this.onload = null;
          for (const cb of callbacks) {
            cb(null, script);
          }
          callbacks = null;
        };
        script.onerror = function() {
          this.onerror = this.onload = null;
          cb(new Error("Failed to load " + src), script);
        };
      }
      if (existingScript && cb) {
        if (window.tinymce) {
          cb(null, existingScript);
        } else {
          callbacks.push(cb);
        }
      }
    },
    init() {
      // dynamic load tinymce from cdn
      this.loading = true;
      this.dynamicLoadScript(tinymceCDN, err => {
        if (err) {
          this.loading = false;
          this.$notify({ title: err, type: "error" });
          return;
        }
        this.loading = false;
        this.initTinymce();
      });
    },
    initTinymce() {
      const _this = this;
      window.tinymce.init({
        selector: `#${this.tinymceId}`,
        language: this.languageTypeList[this.$i18n.locale],
        height: this.height,
        body_class: "panel-body ",
        object_resizing: false,
        toolbar: toolbar,
        menubar: menubar,
        plugins: plugins,
        end_container_on_empty_block: true,
        powerpaste_word_import: "clean",
        code_dialog_height: 400,
        code_dialog_width: 1000,
        imagetools_cors_hosts: ["https://ccsl.shu.edu.cn"],
        default_link_target: "_blank",
        link_title: false,
        statusbar: false,
        language_url: "/tinymce/langs/zh_CN.js",
        nonbreaking_force_tab: true, // inserting nonbreaking space &nbsp; need Nonbreaking Space Plugin
        init_instance_callback: editor => {
          if (_this.value) {
            editor.setContent(_this.value);
          }
          _this.hasInit = true;
          editor.on("NodeChange Change KeyUp SetContent", () => {
            this.hasChange = true;
            this.$emit("input", editor.getContent());
          });
        },
        setup(editor) {
          editor.on("FullscreenStateChanged", e => {
            _this.fullscreen = e.state;
          });
        },
        images_upload_handler(blobInfo, success, failure, progress) {
          progress(0);
          const formData = new FormData();
          formData.append("file", blobInfo.blob());
          UploadFile(formData, "news")
            .then(res => {
              const url = _this.$store.getters.settings.publicURL + res.data;
              success(url);
              progress(100);
            })
            .catch(() => {
              failure("上传文件出错");
            });
        },
        file_picker_callback: function(callback) {
          const input = document.createElement("input");
          input.setAttribute("type", "file");
          input.onchange = function() {
            const file = this.files[0];
            const formData = new FormData();
            formData.append("file", file);
            UploadFile(formData, "news").then(res => {
              const url = _this.$store.getters.settings.publicURL + res.data;
              callback(url);
            });
          };
          input.click();
        }
      });
    },
    destroyTinymce() {
      const tinymce = window.tinymce.get(this.tinymceId);
      if (this.fullscreen) {
        tinymce.execCommand("mceFullScreen");
      }
      if (tinymce) {
        tinymce.setContent("");
        tinymce.destroy();
      }
    },
    setContent(value) {
      window.tinymce.get(this.tinymceId).setContent(value);
    },
    getContent() {
      window.tinymce.get(this.tinymceId).getContent();
    }
  }
};
</script>

<style lang="scss">
// .tinymce-container {
//   position: relative;
//   line-height: normal;
// }

.tox-tinymce {
  border-radius: 3px !important;
}
/* .tinymce-container >>> .mce-fullscreen {
  z-index: 10000;
}
.tinymce-textarea {
  visibility: hidden;
  z-index: -1;
}

.fullscreen .editor-custom-btn-container {
  z-index: 10000;
  position: fixed;
}
.editor-upload-btn {
  display: inline-block;
} */
</style>
