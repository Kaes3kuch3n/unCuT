<script setup lang="ts">
import { useConfirmStore } from "@/stores/confirm";
import DialogBox from "@/components/utils/DialogBox.vue";
import { storeToRefs } from "pinia";

const confirm = useConfirmStore();
const { show, promptText } = storeToRefs(confirm);
</script>

<template>
  <DialogBox v-if="show">
    <p>{{ promptText }}</p>
    <template #buttons>
      <button @click="confirm.reply(true)">Yes</button>
      <button @click="confirm.reply(false)">No</button>
    </template>
  </DialogBox>
</template>

<style scoped></style>
