<script setup lang="ts">
import { dtos } from "@wails/go/models";
import { useGeneratorStore } from "@/stores/generator";
import Cinema = dtos.Cinema;

const { selectCinema } = useGeneratorStore();

const props = defineProps<{
  cinema: Cinema;
}>();
</script>

<template>
  <li>
    <button class="cinema" @click="selectCinema(cinema)">
      <span>{{ props.cinema.name }}</span>
      <span>Planned Screenings: {{ props.cinema.screenings.length }}</span>
    </button>
  </li>
</template>

<style scoped>
button {
  flex-direction: column;
  align-items: flex-start;
  text-align: start;
  width: 100%;
  padding: 0.5em;

  &:hover {
    background: var(--highlight-light);
  }
}

span + span {
  margin-block-start: 0.5em;
  font-size: 0.9em;
}
</style>
