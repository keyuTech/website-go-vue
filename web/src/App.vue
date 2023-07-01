<template>
  <a-layout class="layout">
    <a-layout-header> <AppHeader /></a-layout-header>
    <a-layout-content class="px-24 flex-1">
      <h1 class="text-3xl">111111</h1>
      <h1>{{ t("message.hello") }}</h1>
    </a-layout-content>
    <a-layout-footer class="text-center">
      ©2023 Created by Keyu.Wang. All rights reserved.
    </a-layout-footer>
  </a-layout>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import AppHeader from "@/components/header/index.vue";

const { t } = useI18n({ useScope: "global" });
</script>
