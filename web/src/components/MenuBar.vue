<script setup lang="ts">
import { useRouter } from "@/stores/router";
import { useI18n } from "@/stores/i18n";

const router = useRouter();
const { t } = useI18n();
</script>

<template>
  <nav>
    <h1>{{ t("appName") }}</h1>
    <a
      href="#"
      v-for="route in router.routes"
      :key="route.id"
      :class="{ active: router.activeRoute === route }"
      @click.prevent="router.setActiveRoute(route)"
    >
      <font-awesome-icon v-if="route.icon" :icon="route.icon"/>
      <span v-if="route.name">{{ t(route.name) }}</span>
    </a>
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

svg:has(+span) {
  margin-inline-end: 0.5em;
}

a {
  margin-inline: 0.5em;
  padding: 1em;
  color: inherit;
  text-decoration: none;

  &.active {
    border-bottom: 1px solid var(--primary-light);
  }

  &:last-child {
    margin: 0;
    position: absolute;
    right: 0.5em;
  }
}
</style>
