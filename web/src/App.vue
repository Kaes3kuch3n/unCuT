<script lang="ts" setup>
import MenuBar from "./components/MenuBar.vue";
import { useRouter } from "./stores/router";
import ConfirmDialog from "@/components/utils/ConfirmDialog.vue";
import { onBeforeMount } from "vue";
import { useResourcesStore } from "@/stores/resources";
import ErrorDialog from "@/components/utils/ErrorDialog.vue";
import { useI18n } from "@/stores/i18n";

const { hasLocale, switchLocale } = useI18n();
const router = useRouter();

const { loadResources } = useResourcesStore();

function getUserLocale() {
  if (navigator.languages && navigator.languages.length > 0) {
    for (const language of navigator.languages) {
      const locale = language.split("-")[0];
      if (hasLocale(locale))
        return locale;
    }
  } else if (navigator.language) {
    const locale = navigator.language.split("-")[0];
    if (hasLocale(locale))
      return locale;
  }
  return "de";
}

onBeforeMount(async () => {
  switchLocale(getUserLocale());
  await loadResources();
});
</script>

<template>
  <MenuBar/>
  <main>
    <Component :is="router.activeRoute.component"/>
  </main>
  <ConfirmDialog/>
  <ErrorDialog/>
</template>

<style>
main {
  margin-inline: 1em;
  flex: 1;
  display: flex;
  gap: 1em;
  overflow: hidden;
}
</style>
