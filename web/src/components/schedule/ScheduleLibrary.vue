<script setup lang="ts">
import { GetStoredSchedules } from "@wails/go/gui/App";
import { onMounted, ref } from "vue";
import { ScheduleTemplate } from "@/types/schedule";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useScheduleStore } from "@/stores/schedule";
import { useConfirmStore } from "@/stores/confirm";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const templates = ref<ScheduleTemplate[]>([]);
const schedule = useScheduleStore();
const { prompt } = useConfirmStore();

onMounted(async () => {
  (await GetStoredSchedules())
    .map((json) => JSON.parse(json))
    .forEach((template) => templates.value.push(template));
});

async function load(template: ScheduleTemplate) {
  const confirm = await prompt(t("schedule.library.loadConfirm", template.name));
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
    <h2>{{ t("schedule.library.title") }}</h2>
    <ul>
      <li
        v-for="template in templates"
        :key="template.name"
        @click="load(template)"
      >
        <span>{{ template.name }}</span>
        <div class="actions">
          <button :title="t('schedule.library.clone')" @click.stop="clone(template)">
            <font-awesome-icon icon="fa-solid fa-clone"/>
          </button>
          <button :title="t('schedule.library.delete')" @click.stop="deleteTemplate(template)">
            <font-awesome-icon icon="fa-solid fa-trash"/>
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
