import { defineStore } from "pinia";
import { ContentType, Resource } from "@/types/resource";

type Content = { resource: Resource; duration?: number };

export const useContentStore = defineStore("content", {
  state: () => ({
    screens: new Map<string, Content[]>(),
    scheduledAds: new Map<string, Content[]>(),
    unscheduledAds: [] as Content[],
  }),
  getters: {
    containsResource(): (resource: Resource) => boolean {
      return (resource: Resource): boolean => {
        switch (resource.contentType) {
          case ContentType.ADVERTISEMENT:
            return this.containsAd(resource.id);
          case ContentType.SCREEN:
            return this.containsScreen(resource.id);
        }
      };
    },
    containsAd(state): (id: number) => boolean {
      return (id: number): boolean => {
        const unscheduledContains = state.unscheduledAds.find(
          (item) => item.resource.id === id
        );
        if (unscheduledContains) return true;
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        for (const [_, ads] of state.scheduledAds) {
          if (ads.find((item) => item.resource.id === id)) return true;
        }
        return false;
      };
    },
    containsScreen(state): (id: number) => boolean {
      return (id: number): boolean => {
        // eslint-disable-next-line @typescript-eslint/no-unused-vars
        for (const [_, screens] of state.screens) {
          if (screens.find((item) => item.resource.id === id)) return true;
        }
        return false;
      };
    },
    getScheduledScreens(state): (slot: string) => Content[] {
      return (slot: string): Content[] => state.screens.get(slot) ?? [];
    },
    getScheduledAds(state): (slot: string) => Content[] {
      return (slot: string): Content[] => state.scheduledAds.get(slot) ?? [];
    },
    json(state): string {
      function mapToArray(
        map: Map<string, Content[]>
      ): { id: number; duration: number | undefined; slot: string }[] {
        const array = [];
        for (const [slot, content] of map) {
          for (const screen of content) {
            array.push({
              id: screen.resource.id,
              duration: screen.duration,
              slot: slot,
            });
          }
        }
        return array;
      }

      return JSON.stringify({
        screens: mapToArray(state.screens),
        ads: {
          scheduled: mapToArray(state.scheduledAds),
          unscheduled: state.unscheduledAds.map((ad) => ({
            id: ad.resource.id,
            duration: ad.duration,
          })),
        },
      });
    },
  },
  actions: {
    addUnscheduled(
      resource: Resource,
      duration: number | undefined = undefined
    ) {
      this.unscheduledAds.push({ resource, duration });
    },
    removeUnscheduled(index: number) {
      this.unscheduledAds.splice(index, 1);
    },
    addScheduled(
      resource: Resource,
      slot: string,
      duration: number | undefined = undefined
    ) {
      switch (resource.contentType) {
        case ContentType.SCREEN:
          if (this.screens.has(slot)) {
            this.screens.get(slot)?.push({ resource, duration });
          } else {
            this.screens.set(slot, [{ resource, duration }]);
          }
          break;
        case ContentType.ADVERTISEMENT:
          if (this.scheduledAds.has(slot)) {
            this.scheduledAds.get(slot)?.push({ resource, duration });
          } else {
            this.scheduledAds.set(slot, [{ resource, duration }]);
          }
          break;
      }
    },
    removeScheduledScreen(slot: string, index: number) {
      this.screens.get(slot)?.splice(index, 1);
    },
    removeScheduledAd(slot: string, index: number) {
      this.scheduledAds.get(slot)?.splice(index, 1);
    },
  },
});
