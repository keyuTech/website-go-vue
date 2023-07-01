import { createApp } from "vue";
import "./index.css";
import App from "./App.vue";
import { i18n } from "./locales/i18n";
import { Select, Layout, Button } from "ant-design-vue";
import Icon from "ant-design-vue";
import router from "@/router/index";

const app = createApp(App);

app
  .use(router)
  .use(i18n)
  .use(Select)
  .use(Layout)
  .use(Icon)
  .use(Button)
  .mount("#app");
