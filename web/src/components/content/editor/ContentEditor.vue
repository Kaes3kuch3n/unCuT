<script setup lang="ts">
import { useScheduleStore } from "@/stores/schedule";
import EditableContentSlot from "@/components/content/editor/EditableContentSlot.vue";
import ContentSlot from "@/components/content/editor/ContentSlot.vue";
import { SlotTypes } from "@/types/scheduleSlot";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const { slots } = useScheduleStore();
</script>

<template>
  <section class="editor">
    <h2>{{ t("content.scheduled") }}</h2>
    <ul>
      <!--suppress TypeScriptValidateTypes -->
      <Component
        v-for="slot in slots"
        :key="slot.id"
        :is="slot.type === SlotTypes.TRAILER ? ContentSlot : EditableContentSlot"
        :scheduleSlot="slot"
      />
    </ul>
  </section>
</template>

<style scoped>
.editor {
  flex-direction: column;
  align-items: center;
  flex: 1;
  overflow-y: scroll;
}

ul {
  margin-block-end: 0.5em;
}
</style>
