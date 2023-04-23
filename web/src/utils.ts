import { SlotTypes } from "@/types/scheduleSlot";
import { useI18n } from "@/stores/i18n";

type T = (key: string, ...params: string[]) => string;

let t = null as T | null;

function getT(): T {
  if (t) return t;
  t = useI18n().t;
  return t;
}

export function prettySlotType(type: string): string {
  const t = getT();
  
  switch (type) {
    case SlotTypes.TRAILER:
      return t("contentTypes.trailer");
    case SlotTypes.AD_BIN:
      return t("contentTypes.ad");
    case SlotTypes.CUSTOM:
      return t("contentTypes.custom");
    default:
      return "undefined";
  }
}
