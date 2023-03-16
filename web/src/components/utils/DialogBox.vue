<script setup lang="ts"></script>

<template>
  <div class="overlay">
    <div class="box dialog-box">
      <div class="info">
        <slot name="icon">
          <font-awesome-icon icon="fa-solid fa-question-circle" />
        </slot>
        <div>
          <slot />
        </div>
      </div>
      <div class="buttons">
        <slot name="buttons" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.overlay {
  position: absolute;
  inset: 0;
  background: var(--highlight-dark);
}

.box {
  background: var(--secondary-light);
  color: var(--primary-dark);
  border-radius: 0.5em;
  max-width: 500px;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 1em 2em;
  box-shadow: 0 0 1em var(--highlight-dark);
}

.info {
  display: flex;
  margin-bottom: 1em;
}

.buttons {
  display: flex;
  justify-content: flex-end;
  gap: 1em;
}
</style>
