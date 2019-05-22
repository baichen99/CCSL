import Vue from "vue";
import App from "./App";
import store from "./store";
import router from "./router";
import ElementUI from "element-ui";
import locale from "element-ui/lib/locale/lang/zh-CN"; // lang i18n

import VueVideoPlayer from "vue-video-player";
import "video.js/dist/video-js.css";
import "normalize.css/normalize.css"; // CSS resets
import "element-ui/lib/theme-chalk/index.css";
import "@/styles/index.scss"; // global css
import "@/icons"; // icon
import "@/utils/permission"; // permission control

Vue.use(ElementUI, { locale });
Vue.use(VueVideoPlayer, {
  options: {
    controls: true,
    language: "zh-CN",
    autoplay: false,
    fluid: true,
    controlBar: {
      volumeBar: false,
      volumePanel: false,
      currentTimeDisplay: false
    },
    languages: {
      "zh-CN": {
        "The media could not be loaded, either because the server or network failed or because the format is not supported.":
          "暂无数据"
      }
    }
  }
});

Vue.config.productionTip = false;

new Vue({
  el: "#app",
  router,
  store,
  render: h => h(App)
});
