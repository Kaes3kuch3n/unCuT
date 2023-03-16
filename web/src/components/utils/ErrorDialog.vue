<script setup lang="ts">
import DialogBox from "@/components/utils/DialogBox.vue";
import { onMounted, ref } from "vue";
import { EventsOn } from "@wails/runtime";

const show = ref(false);
const errorMessage = ref("");

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
      <font-awesome-icon icon="fa-solid fa-circle-xmark" />
    </template>
    <p>An error occurred:</p>
    <p>{{ errorMessage }}</p>
    <template #buttons>
      <button @click="show = false">Okay</button>
    </template>
  </DialogBox>
</template>

<style scoped></style>
