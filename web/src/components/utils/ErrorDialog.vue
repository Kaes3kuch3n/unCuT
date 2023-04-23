<script setup lang="ts">
import DialogBox from "@/components/utils/DialogBox.vue";
import { onMounted, ref } from "vue";
import { EventsOn } from "@wails/runtime";
import { useI18n } from "@/stores/i18n";

const show = ref(false);
const errorMessage = ref("");

const { t } = useI18n();

onMounted(() =>
  EventsOn("error", (msg) => {
    errorMessage.value = msg;
    show.value = true;
  })
);
</script>

<template>
  <DialogBox v-if="show">
    <template #icon>
      <font-awesome-icon icon="fa-solid fa-circle-xmark"/>
    </template>
    <p>{{ t("error.occurred") }}</p>
    <p>{{ errorMessage }}</p>
    <template #buttons>
      <button @click="show = false">{{ t("error.dismiss") }}</button>
    </template>
  </DialogBox>
</template>

<style scoped></style>
