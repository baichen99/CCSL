import Vue from "vue";
import VueI18n from "vue-i18n";
import store from "@/store";
import zh from "@/locale/zh-CN";
import en from "@/locale/en-US";

const dateTimeFormats = {
  "en-US": {
    short: {
      year: "numeric",
      month: "short",
      day: "numeric"
    },
    long: {
      year: "numeric",
      month: "short",
      day: "numeric",
      hour: "numeric",
      minute: "numeric"
    }
  },
  "zh-CN": {
    short: {
      year: "numeric",
      month: "short",
      day: "numeric"
    },
    long: {
      year: "numeric",
      month: "short",
      day: "numeric",
      hour: "numeric",
      minute: "numeric",
      hour12: false
    }
  }
};

let locale;

const browserLanguage = navigator.language;
const savedLanguage = store.state.app.locale;

if (savedLanguage) {
  locale = savedLanguage;
} else {
  const langPrefix = browserLanguage.split("-")[0];
  switch (langPrefix) {
    case "en":
      locale = "en-US";
      break;

    default:
      locale = "zh-CN";
      break;
  }
  store.dispatch("app/setLocale", locale);
}

const messages = {
  "zh-CN": zh,
  "en-US": en
};

Vue.use(VueI18n);

const i18n = new VueI18n({
  messages,
  locale,
  dateTimeFormats,
  silentTranslationWarn: true
});

export default i18n;
