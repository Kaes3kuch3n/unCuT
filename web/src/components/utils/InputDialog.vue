<script setup lang="ts">
import { ref, watchEffect } from "vue";
import DialogBox from "@/components/utils/DialogBox.vue";

const props = defineProps<{
  show: boolean;
  value: string | number;
  label: string;
  type: string;
}>();

const emit = defineEmits<{
  (e: "update:value", value: string): void;
  (e: "update:show", value: boolean): void;
}>();

const input = ref<HTMLInputElement>();

function updateValue() {
  if (!input.value) return;
  emit("update:value", input.value?.value);
  emit("update:show", false);
}

watchEffect(() => {
  if (!props.show) return;
  input.value?.focus();
});
</script>

<template>
  <DialogBox v-if="props.show">
    <p>
      <slot />
    </p>
    <label>
      {{ label }}:
      <input
        :type="props.type"
        ref="input"
        :value="props.value"
        @keydown.enter.prevent="updateValue"
        @keydown.esc.prevent="emit('update:show', false)"
      />
    </label>
    <template #buttons>
      <button @click="updateValue">Save</button>
      <button @click="emit('update:show', false)">Cancel</button>
    </template>
  </DialogBox>
</template>

<style scoped>
label {
  display: flex;
  align-items: center;
  gap: 0.5em;
}

input {
  width: 100%;
}
</style>
