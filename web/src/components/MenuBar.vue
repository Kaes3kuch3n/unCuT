<script setup lang="ts">
import { useRouter } from "@/stores/router";

const router = useRouter();
</script>

<template>
  <nav>
    <h1>unCuT</h1>
    <a
      href="#"
      v-for="route in router.routes"
      :key="route.name"
      :class="{ active: router.activeRoute === route }"
      @click.prevent="router.setActiveRoute(route)"
      >{{ route.name }}</a
    >
  </nav>
</template>

<style scoped>
nav {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 1em;
}

h1 {
  margin: 0;
  position: absolute;
  left: 0.5em;
}

a {
  margin-inline: 0.5em;
  padding: 1em;
  color: inherit;
  text-decoration: none;

  &.active {
    border-bottom: 1px solid var(--primary-light);
  }
}
</style>
