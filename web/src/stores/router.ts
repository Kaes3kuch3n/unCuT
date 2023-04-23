import { defineStore } from "pinia";
import ScheduleView from "../components/schedule/ScheduleView.vue";
import ContentView from "../components/content/ContentView.vue";
import { Component, markRaw } from "vue";
import GenerateView from "../components/generator/GenerateView.vue";
import SettingsView from "@/components/settings/SettingsView.vue";

interface Route {
  id: string;
  name?: string;
  icon?: string;
  component: Component;
}

const routes: Route[] = [
  {
    id: "schedule",
    name: "tabs.schedule",
    component: markRaw(ScheduleView),
  },
  {
    id: "content",
    name: "tabs.content",
    component: markRaw(ContentView),
  },
  {
    id: "generate",
    name: "tabs.generate",
    component: markRaw(GenerateView),
  },
  {
    id: "settings",
    icon: "cog",
    component: markRaw(SettingsView),
  }
];

export const useRouter = defineStore("route", {
  state: () => {
    return {
      routes: routes,
      activeRoute: routes[0],
    };
  },
  actions: {
    setActiveRoute(route: Route) {
      this.activeRoute = route;
    },
  },
});
