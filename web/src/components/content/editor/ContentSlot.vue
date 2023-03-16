<script setup lang="ts">
import { prettySlotType } from "@/utils";
import { ScheduleSlot } from "@/types/scheduleSlot";

const props = defineProps<{
  scheduleSlot: ScheduleSlot;
}>();
</script>

<template>
  <li>
    <header>
      <span>{{ prettySlotType(props.scheduleSlot.type) }}</span>
      <hr />
      <h3>{{ props.scheduleSlot.name }}</h3>
    </header>
    <slot></slot>
  </li>
</template>

<style scoped>
li {
  border: 1px solid var(--primary-light);
  border-radius: 0.5em;
  margin-block-end: 0.5em;

  &:last-child {
    margin-block-end: 0;
  }
}

header {
  padding-inline: 1em;
  display: flex;
  align-items: center;

  & hr {
    margin-inline: 1em;
    color: var(--primary-light);
    height: 2em;
  }

  & h3 {
    flex: 1;
    margin: 0;
  }
}
</style>
