<script setup lang="ts">
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { useGeneratorStore } from "@/stores/generator";
import { storeToRefs } from "pinia";
import { computed, onMounted, ref } from "vue";
import {
  GenerateLeads,
  GetDefaultOutputDirectory,
  GetSchedulingStrategies,
  SelectOutputDirectory,
} from "@wails/go/gui/App";
import { useScheduleStore } from "@/stores/schedule";
import { useContentStore } from "@/stores/content";
import { gui } from "@wails/go/models";
import { useI18n } from "@/stores/i18n";
import SelectWrapper from "@/components/utils/SelectWrapper.vue";
import adSchedulingStrategyDescriptor = gui.adSchedulingStrategyDescriptor;

const { t } = useI18n();

const generatorStore = useGeneratorStore();
const { selectedScreenings } = storeToRefs(generatorStore);
const { json: scheduleJson } = storeToRefs(useScheduleStore());
const { json: contentJson } = storeToRefs(useContentStore());

const outputDirectory = ref<string | null>(null);

async function selectOutputDirectory() {
  const outDir = await SelectOutputDirectory();
  if (outDir.trim() === "") return;
  outputDirectory.value = outDir;
}

enum State {
  NONE,
  GENERATING,
  DONE,
}

const state = ref<State>(State.NONE);
const schedulingStrategies = ref<adSchedulingStrategyDescriptor[]>([]);
const selectedStrategy = ref<adSchedulingStrategyDescriptor>();
const strategyParamRefs = ref<Record<string, HTMLInputElement>>({});
const schedulingStrategyJson = computed(() => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const options: any = {};
  for (const ref in strategyParamRefs.value) {
    const refType = selectedStrategy.value?.parameters?.find(
      (param) => param.id === ref
    )?.input.type;
    const value = strategyParamRefs.value[ref].value;
    options[ref] = refType === "number" ? parseFloat(value) : value;
  }

  return JSON.stringify({
    strategy: selectedStrategy.value?.id,
    options: options,
  });
});

const buttonText = computed(() => {
  switch (state.value) {
    case State.GENERATING:
      return t("generator.generating");
    default:
      return ("generator.generate");
  }
});

async function generate() {
  if (!outputDirectory.value) return;
  state.value = State.GENERATING;
  await GenerateLeads(
    scheduleJson.value,
    contentJson.value,
    selectedScreenings.value.map((s) => s.id),
    outputDirectory.value,
    schedulingStrategyJson.value
  );
  state.value = State.DONE;
}

onMounted(async () => {
  outputDirectory.value = await GetDefaultOutputDirectory();
  schedulingStrategies.value = await GetSchedulingStrategies();
  selectedStrategy.value = schedulingStrategies.value[0];
});
</script>

<template>
  <section>
    <h2>{{ t("generator.title") }}</h2>
    <h3>{{ t("generator.options.title") }}</h3>
    <div class="options">
      <h4>{{ t("generator.options.outputDirectory") }}</h4>
      <span>{{ outputDirectory }}</span>
      <button class="col-right" @click="selectOutputDirectory">
        <font-awesome-icon icon="fa-solid fa-folder-open"/>
        {{ t("generator.options.selectOutputDirectory") }}
      </button>
      <h4>{{ t("generator.options.adSchedulingStrategy") }}</h4>
      <span>{{ selectedStrategy?.description }}</span>
      <SelectWrapper class="col-right">
        <select
          name="schedulingStrategy"
          id="schedulingStrategy"
          v-model="selectedStrategy"
        >
          <option
            v-for="s in schedulingStrategies"
            :key="s.id"
            :value="s"
            :title="s.description"
          >
            {{ s.name }}
          </option>
        </select>
      </SelectWrapper>
      <template v-if="selectedStrategy?.parameters">
        <h4>Strategy Options:</h4>
        <template v-for="param in selectedStrategy.parameters" :key="param.name">
          <span>{{ param.name }}</span>
          <SelectWrapper class="col-right" v-if="param.input.type === 'select'">
            <select
              :name="param.name"
              :id="param.name"
              :value="param.input.defaultValue"
              :ref="
                // The following comment is required because TS doesn't recognize the element as input element
                // @ts-ignore
                (el) => (strategyParamRefs[param.id] = el)
              "
            >
              <option
                v-for="(option, id) in param.input.options"
                :key="id"
                :value="id"
              >
                {{ option }}
              </option>
            </select>
          </SelectWrapper>
          <input
            class="col-right"
            v-else-if="param.input.type === 'number'"
            :type="param.input.type"
            :min="param.input.min"
            :max="param.input.max"
            :step="param.input.step"
            :value="param.input.defaultValue"
            :ref="
              // The following comment is required because TS doesn't recognize the element as input element
              // @ts-ignore
              (el) => (strategyParamRefs[param.id] = el)
            "
          />
        </template>
      </template>
    </div>
    <div class="generate">
      <button
        :class="{
          generateButton: true,
          generating: state === State.GENERATING,
        }"
        :disabled="selectedScreenings.length === 0 || state === State.GENERATING"
        @click="generate"
      >
        <font-awesome-icon icon="fa-solid fa-clapperboard"/>
        <span>{{ buttonText }}</span>
      </button>
      <p v-show="state === State.DONE">{{ t("generator.done") }}</p>
    </div>
  </section>
</template>

<style scoped>
section {
  flex: 1;
}

h3 {
  margin-block: 1.5em 0.5em;
}

button {
  background: var(--secondary-dark);
  padding: 0.5em;

  & svg {
    margin-inline-end: 0.5em;
  }
}

.options {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.5em;
  align-items: center;
  justify-items: start;

  & h4 {
    grid-column: span 2;
  }

  & > :nth-child(even of :not(h4)) {
    justify-self: end;
  }

  & .col-right {
    justify-self: end;
  }
}

.generate {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  margin-block-start: 3em;
}

.generateButton {
  background: var(--secondary-dark);
  display: flex;
  gap: 0.5em;
  font-size: 1.5em;
  padding: 0.5em;

  &[disabled] {
    background: var(--highlight-light);
  }

  &.generating svg {
    animation: 1s ease-in-out infinite spin;
  }
}

@keyframes spin {
  0% {
    transform: rotateY(0);
  }

  80% {
    transform: rotateY(2turn);
  }

  81% {
    transform: rotateY(0);
  }
}
</style>
