import Vue from "vue";
import ElementUI from "element-ui";
import App from "./App";
import store from "./store";
import router from "./router";
import settings from "./settings";
import VueVideoPlayer from "vue-video-player";
// import locale from "element-ui/lib/locale/lang/zh-CN"; // lang i18n
import "video.js/dist/video-js.css";
// import "normalize.css/normalize.css"; // CSS resets
// import "@/styles/index.scss"; // element css
import "@/styles/global.scss";
import "@/icons"; // icon
import "@/utils/permission"; // permission control

Vue.use(ElementUI, {
  zIndex: 3000
  // locale
});
Vue.use(VueVideoPlayer, settings.videojsOptions);
Vue.config.productionTip = false;

new Vue({
  el: "#app",
  router,
  store,
  render: h => h(App)
});
