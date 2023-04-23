<script setup lang="ts">
import { useConfirmStore } from "@/stores/confirm";
import DialogBox from "@/components/utils/DialogBox.vue";
import { storeToRefs } from "pinia";
import { useI18n } from "@/stores/i18n";

const confirm = useConfirmStore();
const { show, promptText } = storeToRefs(confirm);

const { t } = useI18n();
</script>

<template>
  <DialogBox v-if="show">
    <p>{{ promptText }}</p>
    <template #buttons>
      <button @click="confirm.reply(true)">{{ t("confirm.yes") }}</button>
      <button @click="confirm.reply(false)">{{ t("confirm.no") }}</button>
    </template>
  </DialogBox>
</template>

<style scoped></style>
