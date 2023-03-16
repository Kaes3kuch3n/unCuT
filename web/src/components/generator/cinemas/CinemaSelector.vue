<script setup lang="ts">
import { useDefaultsStore } from "@/stores/defaults";
import { ref, watchEffect } from "vue";
import { dtos } from "@wails/go/models";
import { GetCinemas } from "@wails/go/gui/App";
import { storeToRefs } from "pinia";
import SelectedCinema from "@/components/generator/cinemas/SelectedCinema.vue";
import CinemaItem from "@/components/generator/cinemas/CinemaItem.vue";
import Cinema = dtos.Cinema;

const { maxCinemaSearchResults } = storeToRefs(useDefaultsStore());

const cinemaSearch = ref("");
const cinemas = ref<Cinema[]>([]);

watchEffect(async () => {
  cinemas.value = await GetCinemas(
    cinemaSearch.value,
    maxCinemaSearchResults.value
  );
});
</script>

<template>
  <section class="column">
    <h2>Cinemas</h2>
    <SelectedCinema />
    <label>
      Search:
      <input
        type="text"
        name="cinemaSearch"
        id="cinemaSearch"
        v-model="cinemaSearch"
      />
    </label>
    <ul>
      <CinemaItem
        v-for="cinema in cinemas"
        :key="cinema.name"
        :cinema="cinema"
      />
    </ul>
  </section>
</template>

<style scoped>
label {
  display: flex;
  align-items: center;
  gap: 0.5em;

  & input {
    flex: 1;
    min-width: 0;
  }
}

ul {
  margin-block: 1em;
}
</style>
