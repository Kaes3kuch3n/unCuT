<script setup lang="ts">
import { storeToRefs } from "pinia";
import { useGeneratorStore } from "@/stores/generator";
import ScreeningItem from "@/components/generator/screenings/ScreeningItem.vue";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const generatorStore = useGeneratorStore();
const { selectedCinema } = storeToRefs(generatorStore);
</script>

<template>
  <section class="column">
    <h2>{{ t("generator.cinemas.screenings") }}</h2>
    <p v-if="!selectedCinema" class="centered">
      {{ t("generator.cinemas.screeningsPlaceholder") }}
    </p>
    <ul v-else>
      <ScreeningItem
        v-for="screening in selectedCinema.screenings"
        :key="screening.id"
        :screening="screening"
      />
    </ul>
  </section>
</template>

<style scoped>
section {
  border-inline-end: none;
  position: relative;
}

.centered {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 80%;
  text-align: center;
  color: var(--highlight-light);
}
</style>
