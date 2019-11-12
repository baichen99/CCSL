import Vue from "vue";
import ElementUI from "element-ui";
import App from "./App";
import store from "@/store";
import router from "@/router";
import i18n from "@/locale";
import settings from "./settings";
import VueVideoPlayer from "vue-video-player";
import "@/directives";
import "@/utils/error";
import "video.js/dist/video-js.css";
import "normalize.css/normalize.css"; // CSS resets
import "@/styles/element-variables.scss";
import "@/styles/index.scss"; // element css
import "@/icons"; // icon

Vue.use(ElementUI, {
  zIndex: 1000,
  i18n: (key, value) => i18n.t(key, value)
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
