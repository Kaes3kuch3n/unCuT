import { defineStore } from "pinia";
import ScheduleView from "../components/schedule/ScheduleView.vue";
import ContentView from "../components/content/ContentView.vue";
import { Component, markRaw } from "vue";
import GenerateView from "../components/generator/GenerateView.vue";

interface Route {
  name: string;
  component: Component;
}

const routes: Route[] = [
  {
    name: "Schedule",
    component: markRaw(ScheduleView),
  },
  {
    name: "Content",
    component: markRaw(ContentView),
  },
  {
    name: "Generate",
    component: markRaw(GenerateView),
  },
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
