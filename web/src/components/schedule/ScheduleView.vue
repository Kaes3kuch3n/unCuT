<script setup lang="ts">
import AddSlotMenu from "./AddSlotMenu.vue";
import ScheduleSlot from "./ScheduleSlot.vue";
import { useScheduleStore } from "@/stores/schedule";
import ScheduleLibrary from "./ScheduleLibrary.vue";
import { storeToRefs } from "pinia";

const schedule = useScheduleStore();
const { slots } = storeToRefs(schedule);
</script>

<template>
  <ScheduleLibrary />
  <section class="editor">
    <ScheduleSlot
      v-for="(slot, i) in slots"
      :key="slot.id"
      :item="slot"
      :has-add-menu="i !== slots.length - 1"
      @add="(type) => schedule.insertSlot(type, i + 1)"
      @delete="schedule.deleteSlot(i)"
      @set-name="(name) => (slot.name = name)"
    />
    <AddSlotMenu @add-slot="(type) => schedule.addSlot(type)" />
  </section>
</template>

<style scoped>
.editor {
  flex-direction: column;
  align-items: center;
  flex: 1;
  overflow-y: scroll;
}
</style>
