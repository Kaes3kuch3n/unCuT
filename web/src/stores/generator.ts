import { defineStore } from "pinia";
import { dtos } from "@wails/go/models";
import Cinema = dtos.Cinema;
import Screening = dtos.Screening;

export const useGeneratorStore = defineStore("generator", {
  state: () => ({
    selectedCinema: null as Cinema | null,
    selectedScreenings: [] as Screening[],
  }),
  getters: {
    currentCinema: (state) => state.selectedCinema,
    isSelected(state): (screening: Screening) => boolean {
      return (screening: Screening) => {
        return state.selectedScreenings.includes(screening);
      };
    },
  },
  actions: {
    selectCinema(cinema: Cinema | null) {
      this.selectedCinema = cinema;
      this.selectedScreenings = [];
    },
    selectScreening(screening: Screening) {
      this.selectedScreenings.push(screening);
    },
    deselectScreening(screening: Screening) {
      const index = this.selectedScreenings.findIndex(
        (s) => s.id === screening.id
      );
      this.selectedScreenings.splice(index, 1);
    },
  },
});
