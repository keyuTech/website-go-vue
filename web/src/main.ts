import { createApp } from "vue";
import "./index.css";
import App from "./App.vue";
import { i18n } from "./locales/i18n";
import { Select, Layout } from "ant-design-vue";
import Icon from "ant-design-vue";

const app = createApp(App);

app.use(i18n).use(Select).use(Layout).use(Icon).mount("#app");
