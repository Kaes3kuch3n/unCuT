<script setup lang="ts">
const props = defineProps<{
  expanded: boolean;
}>();
</script>

<template>
  <font-awesome-icon v-if="props.expanded" icon="fa-solid fa-caret-down" />
  <font-awesome-icon v-else icon="fa-solid fa-caret-right" />
</template>

<style scoped>
svg {
  width: 0.5em;
}
</style>
