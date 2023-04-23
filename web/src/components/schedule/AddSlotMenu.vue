<script setup lang="ts">
import { SlotTypes } from "@/types/scheduleSlot";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const emit = defineEmits<{
  (e: "add-slot", type: SlotTypes): void;
}>();

function addSlot(type: SlotTypes) {
  emit("add-slot", type);
}
</script>

<template>
  <div id="addMenu">
    <button @click="addSlot(SlotTypes.TRAILER)">{{ t("schedule.add.trailer") }}</button>
    <button @click="addSlot(SlotTypes.AD_BIN)">{{ t("schedule.add.ad") }}</button>
    <button @click="addSlot(SlotTypes.CUSTOM)">{{ t("schedule.add.custom") }}</button>
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
