<script setup lang="ts">
import FolderCaret from "@/components/utils/FolderCaret.vue";

const props = defineProps<{
  expanded: boolean;
}>();

const emit = defineEmits<{
  (e: "update:expanded", value: boolean): void;
}>();
</script>

<template>
  <li
    :class="{ expanded: props.expanded }"
    @click.stop="emit('update:expanded', !props.expanded)"
  >
    <div>
      <FolderCaret :expanded="props.expanded" />
      <slot name="name" />
    </div>
    <ul v-show="props.expanded">
      <slot />
    </ul>
  </li>
</template>

<style scoped>
div {
  display: flex;
  gap: 0.5em;
  align-items: center;
  padding: 0.5em;

  &:hover {
    background: var(--highlight-light);
    border-radius: 0.2em;
  }
}

.expanded ul {
  display: inherit;
}

ul {
  padding-inline-start: 1.5em;
  display: none;
}
</style>
