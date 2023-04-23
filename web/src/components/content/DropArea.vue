<script setup lang="ts">
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const props = defineProps<{
  isEmpty: boolean;
}>();
</script>

<template>
  <div :class="{ 'drop-area': true, empty: props.isEmpty }">
    <p v-if="props.isEmpty">
      {{ t("content.dropAreaDescription") }}
    </p>
    <ul v-else>
      <slot/>
    </ul>
  </div>
</template>

<style scoped>
.drop-area {
  margin: 1em;
  margin-block-start: 0.5em;
  border-radius: 0.5em;
  display: flex;
  flex-direction: column;
}

.drop-area.empty {
  border: 2px dashed var(--highlight-light);
  justify-content: center;
}

.drop-area p {
  margin: 1em 20%;
  text-align: center;
  color: var(--highlight-light);
}
</style>
