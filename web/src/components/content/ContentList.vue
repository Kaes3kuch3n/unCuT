<script setup lang="ts">
import { useDropArea } from "@/composables/useDragDrop";
import ContentItem from "@/components/content/ContentItem.vue";
import { useContentStore } from "@/stores/content";
import { ContentType, Resource } from "@/types/resource";
import DropArea from "@/components/content/DropArea.vue";
import { storeToRefs } from "pinia";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const { dragEnter, dragLeave, drop } = useDropArea(ContentType.ADVERTISEMENT);
const contentStore = useContentStore();
const { addUnscheduled, removeUnscheduled } = contentStore;
const { unscheduledAds } = storeToRefs(contentStore);

function dropContent(event: DragEvent) {
  const data = drop(event);
  if (!data) return;
  const resource = Resource.fromDropData(data);
  addUnscheduled(resource);
}
</script>

<template>
  <section class="column">
    <h2>{{ t("content.unscheduled") }}</h2>
    <DropArea
      class="content-list"
      :is-empty="unscheduledAds.length === 0"
      @dragenter="dragEnter"
      @dragleave="dragLeave"
      @drop="dropContent"
    >
      <ContentItem
        v-for="(item, index) in unscheduledAds"
        :key="item.resource.id"
        :resource="item.resource"
        v-model:duration="item.duration"
        @remove="removeUnscheduled(index)"
      />
    </DropArea>
  </section>
</template>

<style scoped>
section {
  display: flex;
  flex-direction: column;
  border-inline-end: none;
}

.content-list {
  flex: 1;
  margin-inline: 0;
  margin-block-start: 0;
}
</style>
