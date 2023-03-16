<script setup lang="ts">
import { useDraggable } from "@/composables/useDragDrop";
import { useContentStore } from "@/stores/content";
import { Resource } from "@/types/resource";
import { computed } from "vue";

const props = defineProps<{
  item: Resource;
}>();

const { drag } = useDraggable(props.item);
const { containsResource } = useContentStore();

const isUsed = computed(() => containsResource(props.item));
</script>

<template>
  <li draggable="true" @dragstart="drag">
    <font-awesome-icon :icon="`fa-solid fa-${props.item.icon}`" />
    <span>{{ props.item.name }}</span>
    <font-awesome-icon v-show="isUsed" icon="fa-solid fa-check" />
  </li>
</template>

<style scoped>
li {
  display: flex;
  gap: 0.5em;
  align-items: center;
  padding: 0.5em;
}

li span {
  flex: 1;
}

li:hover {
  background: var(--highlight-light);
  border-radius: 0.2em;
}
</style>
