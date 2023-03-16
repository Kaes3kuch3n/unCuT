<script setup lang="ts">
import { useGeneratorStore } from "@/stores/generator";
import { storeToRefs } from "pinia";

const generatorStore = useGeneratorStore();
const { selectCinema } = generatorStore;
const { selectedCinema } = storeToRefs(generatorStore);
</script>

<template>
  <div>
    Selected cinema:&nbsp;
    <span v-if="!selectedCinema">None</span>
    <div v-if="selectedCinema" class="cinema-container">
      <div class="cinema">
        <p>{{ selectedCinema.name }}</p>
        <p>Planned Screenings: {{ selectedCinema.screenings.length }}</p>
      </div>
      <button @click="selectCinema(null)" title="Deselect">
        <font-awesome-icon icon="fa-solid fa-x" />
      </button>
    </div>
  </div>
</template>

<style scoped>
div {
  margin-block-end: 1em;
}

.cinema-container {
  display: flex;
  margin-block-start: 0.5em;
  padding: 0.5em;
  border-radius: 0.2em;
  background: var(--secondary-dark);
}

.cinema {
  flex: 1;
  margin: 0;
}

p {
  margin: 0;
}

p + p {
  margin-block-start: 0.5em;
  font-size: 0.9em;
}

button {
  padding-inline: 0.75em;
}
</style>
