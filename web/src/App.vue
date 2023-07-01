<template>
  <a-layout class="layout">
    <a-layout-header> <AppHeader /></a-layout-header>
    <a-layout-content class="px-24 flex-1">
      <router-view></router-view>
    </a-layout-content>
    <a-layout-footer class="text-center">
      ©2023 Created by Keyu.Wang. All rights reserved.
    </a-layout-footer>
  </a-layout>
</template>

<script setup lang="ts">
import AppHeader from "@/components/header/index.vue";

// const { t } = useI18n({ useScope: "global" });
</script>
