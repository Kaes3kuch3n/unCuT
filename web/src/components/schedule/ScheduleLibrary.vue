<script setup lang="ts">
import { GetStoredSchedules } from "@wails/go/gui/App";
import { onMounted, ref } from "vue";
import { ScheduleTemplate } from "@/types/schedule";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useScheduleStore } from "@/stores/schedule";
import { useConfirmStore } from "@/stores/confirm";

const templates = ref<ScheduleTemplate[]>([]);
const schedule = useScheduleStore();
const { prompt } = useConfirmStore();

onMounted(async () => {
  (await GetStoredSchedules())
    .map((json) => JSON.parse(json))
    .forEach((template) => templates.value.push(template));
});

async function load(template: ScheduleTemplate) {
  const confirm = await prompt(
    `Do you really want to load the schedule template "${template.name}"? The current schedule will be discarded.`
  );
  if (!confirm) return;
  schedule.set(template.scheduleTemplate);
}

function clone(template: ScheduleTemplate) {
  console.log(`Cloning template ${template.name}...`);
  // TODO: Implement
}

function deleteTemplate(template: ScheduleTemplate) {
  console.log(`Deleting template ${template.name}...`);
  // TODO: Implement
}
</script>

<template>
  <section class="column">
    <h2>Saved Schedules</h2>
    <ul>
      <li
        v-for="template in templates"
        :key="template.name"
        @click="load(template)"
      >
        <span>{{ template.name }}</span>
        <div class="actions">
          <button title="Clone" @click.stop="clone(template)">
            <font-awesome-icon icon="fa-solid fa-clone" />
          </button>
          <button title="Delete" @click.stop="deleteTemplate(template)">
            <font-awesome-icon icon="fa-solid fa-trash" />
          </button>
        </div>
      </li>
    </ul>
  </section>
</template>

<style scoped>
li {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.5em;
  border-radius: 0.2em;

  &:hover {
    background: var(--highlight-light);
  }
}

.actions {
  display: flex;
  gap: 0.5em;
}
</style>
