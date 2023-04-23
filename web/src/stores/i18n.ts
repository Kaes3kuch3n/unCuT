import { defineStore } from "pinia";
import en from "@/assets/locales/en.json";
import de from "@/assets/locales/de.json";

function format(text: string, params: string[]) {
  return text.replace(/{(\d+)}/g, (match, number) => params[number] ?? match);
}

export const useI18n = defineStore("i18n", {
  state: () => ({
    currentLocale: "de",
    localizations: { de, en },
  }),
  getters: {
    locales: (state) => Object.keys(state.localizations),
    localeName: (state) => {
      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      return (locale: string): string => state.localizations[locale].locale
    },
    hasLocale: (state) =>
      (locale: string): boolean => Object.keys(state.localizations).includes(locale),
    t: (state) => {
      return (key: string, ...params: string[]): string => {
        const keys = key.split(".");
        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        let localization = state.localizations[state.currentLocale];
        if (!localization) return key;

        while (keys.length > 0) {
          const k = keys.shift() as string;
          localization = localization[k];
          if (!localization) return key;
        }

        if (params.length > 0)
          localization = format(localization, params);

        return localization;
      };
    }
  },
  actions: {
    switchLocale(locale: string) {
      if (!(locale in this.localizations))
        throw new Error(`Invalid locale specified. Locale "${locale}" does not exist.`);
      this.currentLocale = locale;
    },
  },
});
