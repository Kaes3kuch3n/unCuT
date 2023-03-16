<script setup lang="ts">
import { SlotTypes } from "@/types/scheduleSlot";

const emit = defineEmits<{
  (e: "add-slot", type: SlotTypes): void;
}>();

function addSlot(type: SlotTypes) {
  emit("add-slot", type);
}
</script>

<template>
  <div id="addMenu">
    <button @click="addSlot(SlotTypes.TRAILER)">+ Add Trailer Slot</button>
    <button @click="addSlot(SlotTypes.AD_BIN)">+ Add Ad Slot</button>
    <button @click="addSlot(SlotTypes.CUSTOM)">+ Add Custom Slot</button>
  </div>
</template>

<style scoped>
#addMenu {
  display: flex;
  gap: 1em;
  justify-content: center;
  margin-block: 1em;
}

button {
  background: var(--primary-light);
  color: var(--primary-dark);
  font-size: 1em;
  border-radius: 2em;
  padding: 0.5em 1em;
}
</style>
