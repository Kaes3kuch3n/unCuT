<script setup lang="ts">
import { dtos } from "@wails/go/models";
import { computed } from "vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useGeneratorStore } from "@/stores/generator";
import Screening = dtos.Screening;

const props = defineProps<{
  screening: Screening;
}>();

const date = computed(() => {
  return new Date(props.screening.date);
});

const dateFormat = new Intl.DateTimeFormat(undefined, {
  year: "numeric",
  month: "2-digit",
  day: "2-digit",
  hour: "2-digit",
  minute: "2-digit",
});

const formattedDate = computed(() => {
  return dateFormat.format(date.value);
});

const generatorStore = useGeneratorStore();
const { selectScreening, deselectScreening, isSelected } = generatorStore;

function updateSelectedScreenings(event: Event) {
  const selected = (event.target as HTMLInputElement).checked;
  if (selected) selectScreening(props.screening);
  else deselectScreening(props.screening);
}
</script>

<template>
  <li>
    <label>
      <input
        type="checkbox"
        class="checkbox"
        :name="`includeScreening-${props.screening.id}`"
        :id="`includeScreening-${props.screening.id}`"
        :checked="isSelected(props.screening)"
        @change="updateSelectedScreenings"
      />
      <font-awesome-icon icon="fa-solid fa-check" class="checkbox" />
      <span>{{ props.screening.movie }}</span>
      <time :datetime="date.toISOString()">{{ formattedDate }}</time>
    </label>
  </li>
</template>

<style scoped>
li {
  margin-block-end: 0.5em;
}

label {
  display: grid;
  grid-template-areas:
    "checkbox name"
    "checkbox date";
  grid-template-columns: 2em auto;
}

.checkbox {
  grid-area: checkbox;
  visibility: hidden;

  &:checked + .checkbox {
    visibility: visible;
    justify-self: center;
    align-self: center;
    margin-inline-end: 0.25em;
    font-size: 1.2em;
  }
}
</style>
