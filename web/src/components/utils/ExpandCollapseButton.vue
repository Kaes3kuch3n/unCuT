<script setup lang="ts">
import { computed } from "vue";

const props = defineProps<{
  type: "expand" | "collapse";
}>();

const emit = defineEmits<{
  (e: "click"): void;
}>();

const title = computed(
  () => `${props.type === "expand" ? "Expand" : "Collapse"} all`
);
</script>

<template>
  <button @click="emit('click')" :title="title" :class="props.type">
    <font-awesome-icon
      v-if="props.type === 'expand'"
      icon="fa-solid fa-chevron-up"
    />
    <font-awesome-icon icon="fa-solid fa-chevron-down" />
    <font-awesome-icon
      v-if="props.type === 'collapse'"
      icon="fa-solid fa-chevron-up"
    />
  </button>
</template>

<!--suppress CssUnusedSymbol -->
<style scoped>
button {
  flex-direction: column;
  font-size: 0.5em;
  width: 2em;
  height: 2em;

  &:hover {
    background: var(--highlight-light);
  }
}

.expand {
  & svg:first-child {
    margin-bottom: -0.2em;
  }

  & svg:last-child {
    margin-top: -0.2em;
  }
}

.collapse {
  & svg:first-child {
    margin-bottom: -0.1em;
  }

  & svg:last-child {
    margin-top: -0.1em;
  }
}
</style>
