import { defineStore } from "pinia";
import { v4 as uuid } from "uuid";
import { ScheduleSlot, SlotTypes } from "@/types/scheduleSlot";

export const useScheduleStore = defineStore("schedule", {
  state: () => ({
    slots: [] as ScheduleSlot[],
  }),
  getters: {
    json(state): string {
      return JSON.stringify(
        state.slots.map((slot) => ({
          id: slot.id,
          type: slot.type,
        }))
      );
    },
  },
  actions: {
    set(slots: ScheduleSlot[]) {
      this.slots = slots;
    },
    addSlot(type: SlotTypes) {
      this.slots.push({
        id: uuid(),
        type,
      });
    },
    insertSlot(type: SlotTypes, index: number) {
      this.slots.splice(index, 0, {
        id: uuid(),
        type,
      });
    },
    deleteSlot(index: number) {
      this.slots.splice(index, 1);
    },
    empty() {
      this.slots = [];
    },
  },
});
