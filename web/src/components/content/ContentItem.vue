<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { Resource, ResourceTypes } from "@/types/resource";
import { computed, onMounted, ref, watch } from "vue";
import InputDialog from "@/components/utils/InputDialog.vue";
import { useDefaultsStore } from "@/stores/defaults";
import { storeToRefs } from "pinia";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const props = defineProps<{
  resource: Resource;
  duration?: number;
}>();

const emit = defineEmits<{
  (e: "remove"): void;
  (e: "update:duration", duration: number): void;
}>();

const { imageDuration } = storeToRefs(useDefaultsStore());

const isImage = computed(() => props.resource.type === ResourceTypes.IMAGE);

const showOptionsDialog = ref(false);
const duration = ref(props.duration ?? imageDuration.value);

function durationToNumber(value: string | number): number {
  return typeof value === "string" ? parseFloat(value) : value;
}

watch(duration, (newDuration) => {
  const duration = durationToNumber(newDuration);
  emit("update:duration", duration);
});

onMounted(() => {
  if (!isImage.value) return;
  if (props.duration) return;
  // If resource is an image and the duration is not initialized, initialize it with a default value
  emit("update:duration", durationToNumber(duration.value));
  showOptionsDialog.value = true;
});
</script>

<template>
  <li>
    <font-awesome-icon :icon="`fa-solid fa-${props.resource.icon}`"/>
    <h3>{{ props.resource.name }}</h3>
    <button
      v-if="isImage"
      @click="showOptionsDialog = true"
      :title="t('content.editOptions')"
    >
      <font-awesome-icon icon="fa-solid fa-cog"/>
    </button>
    <button @click="emit('remove')" :title="t('content.remove')">
      <font-awesome-icon icon="fa-solid fa-x"/>
    </button>
    <InputDialog
      v-model:show="showOptionsDialog"
      v-model:value="duration"
      :label="t('content.duration')"
      type="number"
    >
      {{ t("content.durationDialog") }}
    </InputDialog>
  </li>
</template>

<style scoped>
li {
  background: var(--secondary-dark);
  border-radius: 0.2em;
  margin-block-end: 0.5em;
  padding: 0.5em 1em;
  display: flex;
  gap: 0.5em;
  align-items: center;

  &:last-child {
    margin-block-end: 0;
  }
}

h3 {
  margin-block: 0.2em;
  flex: 1;
}

button:has(svg) {
  background: none;
  border: none;
  border-radius: 0.2em;
  color: inherit;
  font-size: 0.8em;
  padding: 0.4em 0.5em;
  display: flex;
  align-items: center;

  &:hover {
    background: var(--highlight-light);
  }
}
</style>
