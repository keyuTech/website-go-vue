<template>
  <h1>{{ $t("message.hello") }}</h1>
  <a-button @click="handleClick"> 点击 </a-button>
  <div>{{ id }}</div>
</template>

<script lang="ts" setup>
import { useI18n } from "vue-i18n";
import { Ref, ref } from "vue";
import { getSettings } from "@/api/settings/settings";

const id: Ref<string> = ref("111");

const handleClick = async () => {
  try {
    const { data: res } = await getSettings();
    console.log("res");
    console.log(res);
    id.value = res.id;
  } catch (error) {
    throw error;
  }
};
</script>
