<script setup lang="ts">
import { ScheduleSlot, SlotTypes } from "@/types/scheduleSlot";
import AddSlotMenu from "./AddSlotMenu.vue";
import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
import { computed, nextTick, onMounted, ref } from "vue";
import { prettySlotType } from "@/utils";
import { useI18n } from "@/stores/i18n";

const { t } = useI18n();

const props = defineProps<{
  item: ScheduleSlot;
  hasAddMenu: boolean;
}>();

const emit = defineEmits<{
  (e: "delete"): void;
  (e: "add", type: SlotTypes): void;
  (e: "set-name", name: string): void;
}>();

const slotType = computed(() => prettySlotType(props.item.type));

const name = computed(
  () => props.item.name ?? t("schedule.slot.unnamed")
);

const newName = ref(props.item.name ?? "");
const input = ref<HTMLInputElement>();
const editMode = ref(false);

async function changeName() {
  editMode.value = true;
  await nextTick();
  input.value?.focus();
  input.value?.setSelectionRange(0, input.value?.value.length);
}

function updateName() {
  editMode.value = false;

  const name: string | undefined = newName.value.trim();
  if (name.length > 0) {
    emit("set-name", newName.value);
  } else {
    newName.value = props.item.name ?? "";
  }
}

onMounted(() => {
  if (props.item.name === undefined) changeName();
  console.log(props.item.type, props.hasAddMenu);
});
</script>

<template>
  <div class="container">
    <div class="slot">
      <span>{{ slotType }}</span>
      <hr/>
      <input
        type="text"
        :name="`name-${props.item.id}`"
        :id="`name-${props.item.id}`"
        :placeholder="t('schedule.slot.namePlaceholder')"
        v-if="editMode"
        v-model="newName"
        ref="input"
        @blur="updateName"
        @keydown.enter.prevent="updateName"
      />
      <h3
        v-else
        :class="{ placeholder: props.item.name === undefined }"
        @click="changeName"
      >
        {{ name }}
      </h3>
      <button
        @click="emit('delete')"
        :title="t('schedule.slot.delete')"
      >
        <font-awesome-icon icon="fa-solid fa-trash"/>
      </button>
    </div>
    <div v-if="props.hasAddMenu" class="menu">
      <AddSlotMenu @add-slot="(type) => emit('add', type)"/>
    </div>
  </div>
</template>

<style scoped>
.container {
  width: 100%;

  &:hover > .menu {
    max-height: 4em;
  }
}

.slot {
  background: var(--secondary-dark);
  border-radius: 0.5em;
  padding-inline: 1em;
  display: flex;
  justify-content: space-between;
  align-items: center;

  & hr {
    margin-inline: 1em;
    color: var(--primary-light);
    height: 2em;
  }

  & h3,
  input {
    margin: 0;
    flex: 1;
    font-size: 1.2em;
    color: inherit;
    background: inherit;
    border: none;
    outline: none;
  }
}

.placeholder {
  color: var(--secondary-light);
  font-style: italic;
}

button {
  font-size: 1em;
  padding-block: 1em;
}

.menu {
  max-height: 1em;
  overflow: hidden;
  transition: 250ms;
}
</style>
