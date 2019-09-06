import Vue from "vue";
import ElementUI from "element-ui";
import App from "./App";
import store from "@/store";
import router from "@/router";
import i18n from "@/locale"
import settings from "./settings";
import VueVideoPlayer from "vue-video-player";
import locale from "element-ui/lib/locale/lang/zh-CN"; // lang i18n
import "video.js/dist/video-js.css";
import "normalize.css/normalize.css"; // CSS resets
import "@/styles/element-variables.scss";
import "@/styles/index.scss"; // element css
import "@/icons"; // icon

Vue.use(ElementUI, {
  zIndex: 3000,
  locale
});

Vue.use(VueVideoPlayer, settings.videojsOptions);
Vue.config.productionTip = false;

new Vue({
  el: "#app",
  router,
  store,
  i18n,
  render: h => h(App)
});
