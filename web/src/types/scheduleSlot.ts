export enum SlotTypes {
  TRAILER = "trailer",
  AD_BIN = "ad",
  CUSTOM = "custom",
}

export interface ScheduleSlot {
  id: string;
  name?: string;
  type: string;
}
