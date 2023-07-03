<template>
  <div class="flex justify-between h-full items-center">
    <router-link to="/">
      <img src="@/assets/seele-logo.svg" class="w-16" />
    </router-link>

    <div class="flex">
      <a-button type="link">
        <router-link to="/blog">BLOG</router-link>
      </a-button>
      <a-button
        type="link"
        href="https://www.keyu.wang/seele-ui/#/"
        target="_black"
        >SEELE-UI</a-button
      >
      <a-button
        type="link"
        href="https://github.com/keyuTech"
        class="mr-8"
        target="_black"
        >Github</a-button
      >
      <a-select
        ref="select"
        v-model:value="$i18n.locale"
        class="w-32"
        :options="locales"
      ></a-select>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type { SelectProps } from "ant-design-vue";
import { ref } from "vue";

const locales = ref<SelectProps["options"]>([
  {
    value: "zh-CN",
    label: "中文",
  },
  {
    value: "en-US",
    label: "English",
  },
  {
    value: "ja-JP",
    label: "日本語",
  },
]);
</script>
