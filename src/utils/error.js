import settings from "@/settings";
import Vue from "vue";
import { SaveJsError } from "@/api/systems";
const { errorLog: needErrorLog } = settings;

function checkNeed() {
  const env = process.env.NODE_ENV;
  if (env === needErrorLog) {
    return true;
  }
  return false;
}

if (checkNeed()) {
  Vue.config.errorHandler = function(err, vm, info) {
    Vue.nextTick(() => {
      const state = vm.$store.state;
      const jsError = {
        err: {
          message: err.message,
          stack: err.stack
        },
        store: {
          app: state.app,
          user: state.user
        },
        info,
        url: window.location.href
      };
      SaveJsError(jsError);
      console.error(err, info);
    });
  };
}
