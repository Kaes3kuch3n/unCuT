import { defineStore } from "pinia";
import { GetAdvertisements, GetScreenTypes } from "@wails/go/gui/App";
import { Resource } from "@/types/resource";

export const useResourcesStore = defineStore("resources", {
  state: () => ({
    advertisements: [] as [string, Resource[]][],
    screens: [] as Resource[],
  }),
  actions: {
    async loadResources() {
      const advertisers = new Map<string, Resource[]>();
      for (const ad of await GetAdvertisements()) {
        if (advertisers.has(ad.advertiser)) {
          advertisers.get(ad.advertiser)?.push(Resource.fromAdvertisement(ad));
        } else {
          advertisers.set(ad.advertiser, [Resource.fromAdvertisement(ad)]);
        }
      }
      this.advertisements = Array.from(advertisers).sort((a, b) =>
        a[0].localeCompare(b[0])
      );

      this.screens = (await GetScreenTypes()).map(Resource.fromScreen);
    },
  },
});
