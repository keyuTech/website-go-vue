import { createI18n } from "vue-i18n";
import zh from "./zh";
import en from "./en";
import ja from "./jp";

/* 这里必须是messages名称 */
const messages = {
  "zh-CN": zh,
  "en-US": en,
  "ja-JP": ja,
};

const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: "zh-CN",
  messages,
});

export { i18n };
