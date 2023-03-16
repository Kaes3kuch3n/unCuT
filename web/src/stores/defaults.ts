import { defineStore } from "pinia";

export const useDefaultsStore = defineStore("defaults", {
  state: () => ({
    imageDuration: 5,
    maxCinemaSearchResults: 5,
  }),
});
