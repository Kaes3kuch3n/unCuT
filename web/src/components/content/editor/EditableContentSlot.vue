<script setup lang="ts">
import { useContentStore } from "@/stores/content";
import { useDropArea } from "@/composables/useDragDrop";
import { ScheduleSlot } from "@/types/scheduleSlot";
import ContentItem from "@/components/content/ContentItem.vue";
import ContentSlot from "@/components/content/editor/ContentSlot.vue";
import { ContentType, Resource } from "@/types/resource";
import { computed } from "vue";
import DropArea from "@/components/content/DropArea.vue";

const props = defineProps<{
  scheduleSlot: ScheduleSlot;
}>();

const {
  getScheduledScreens,
  getScheduledAds,
  addScheduled,
  removeScheduledScreen,
  removeScheduledAd,
} = useContentStore();
const { dragEnter, dragLeave, drop } = useDropArea(getAllowedContentType);

function getAllowedContentType(): ContentType | null {
  // Only allow dropping if only resources of the same type exist in this slot
  if (getScheduledScreens(props.scheduleSlot.id).length > 0)
    return ContentType.SCREEN;
  if (getScheduledAds(props.scheduleSlot.id).length > 0)
    return ContentType.ADVERTISEMENT;
  return null;
}

function dropResource(event: DragEvent, slot: string) {
  const data = drop(event);
  if (!data) return;
  const resource = Resource.fromDropData(data);
  addScheduled(resource, slot);
}

const isEmpty = computed(
  () =>
    getScheduledAds(props.scheduleSlot.id).length +
      getScheduledScreens(props.scheduleSlot.id).length ===
    0
);
</script>

<template>
  <ContentSlot
    :schedule-slot="scheduleSlot"
    @dragenter="dragEnter"
    @dragleave="dragLeave"
    @drop="(e) => dropResource(e, props.scheduleSlot.id)"
  >
    <DropArea :is-empty="isEmpty">
      <ContentItem
        v-for="(item, index) in getScheduledScreens(props.scheduleSlot.id)"
        :key="item.resource.id"
        :resource="item.resource"
        v-model:duration="item.duration"
        @remove="removeScheduledScreen(props.scheduleSlot.id, index)"
      />
      <ContentItem
        v-for="(item, index) in getScheduledAds(props.scheduleSlot.id)"
        :key="item.resource.id"
        :resource="item.resource"
        v-model:duration="item.duration"
        @remove="removeScheduledAd(props.scheduleSlot.id, index)"
      />
    </DropArea>
  </ContentSlot>
</template>

<style scoped></style>
