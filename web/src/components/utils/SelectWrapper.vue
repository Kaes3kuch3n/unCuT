<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
</script>

<template>
  <span class="select-wrapper">
    <slot/>
    <font-awesome-icon icon="caret-down"/>
  </span>
</template>

<style scoped>
.select-wrapper {
  margin-inline-start: 0.5em;
  background: var(--secondary-dark);
  color: var(--primary-light);
  position: relative;
  display: inline-block;
  border-radius: 0.2em;
}

.select-wrapper > svg {
  position: absolute;
  right: 0.4em;
  top: 50%;
  transform: translateY(-50%);
}
</style>
