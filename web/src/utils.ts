import { SlotTypes } from "@/types/scheduleSlot";

export function prettySlotType(type: string): string {
  switch (type) {
    case SlotTypes.TRAILER:
      return "Trailer";
    case SlotTypes.AD_BIN:
      return "Advertisement";
    case SlotTypes.CUSTOM:
      return "Custom";
    default:
      return "undefined";
  }
}
