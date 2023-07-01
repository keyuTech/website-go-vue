<template>
  <div class="flex justify-between h-full items-center">
    <img src="@/assets/seele-logo.svg" class="w-16" />
    <a-select
      ref="select"
      v-model:value="$i18n.locale"
      class="w-32"
      :options="locales"
    ></a-select>
  </div>
</template>

<script lang="ts" setup>
import type { SelectProps } from "ant-design-vue";
import { ref } from "vue";
import { useI18n } from "vue-i18n";

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
