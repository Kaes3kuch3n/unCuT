<script lang="ts" setup>
import MenuBar from "./components/MenuBar.vue";
import { useRouter } from "./stores/router";
import ConfirmDialog from "@/components/utils/ConfirmDialog.vue";
import { onBeforeMount } from "vue";
import { useResourcesStore } from "@/stores/resources";
import ErrorDialog from "@/components/utils/ErrorDialog.vue";

const router = useRouter();

const { loadResources } = useResourcesStore();

onBeforeMount(async () => {
  await loadResources();
});
</script>

<template>
  <MenuBar />
  <main>
    <Component :is="router.activeRoute.component" />
  </main>
  <ConfirmDialog />
  <ErrorDialog />
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
