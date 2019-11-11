import Vue from "vue";
import VueI18n from "vue-i18n";
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

const browserLanguage = navigator.language;

let locale;
switch (browserLanguage) {
  case "en-US":
    locale = "en-US";
    break;

  default:
    locale = "zh-CN";
    break;
}

const messages = {
  "zh-CN": zh,
  "en-US": en
};

Vue.use(VueI18n);

const i18n = new VueI18n({
  messages,
  locale,
  dateTimeFormats
});

export default i18n;
