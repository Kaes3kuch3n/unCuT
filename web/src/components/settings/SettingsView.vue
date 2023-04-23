<script setup lang="ts">
import { useI18n } from "@/stores/i18n";
import SelectWrapper from "@/components/utils/SelectWrapper.vue";

const { locales, localeName, switchLocale, currentLocale, t } = useI18n();

function setLocale(event: Event) {
  switchLocale((event.target as HTMLInputElement).value);
}
</script>

<template>
  <section>
    <h2>{{ t("settings.title") }}</h2>
    <div class="settings">
      <label for="language">
        {{ t("settings.language") }}
      </label>
      <SelectWrapper>
        <select name="language" id="language" @change="setLocale">
          <option v-for="locale in locales" :key="locale" :value="locale" :selected="locale === currentLocale">
            {{ localeName(locale) }}
          </option>
        </select>
      </SelectWrapper>
    </div>
  </section>
</template>

<style scoped>
section {
  flex: 1;
  max-width: 30em;
  margin: 0 auto;
}

h2 {
  text-align: center;
}

.settings {
  display: grid;
  grid-template-columns: 1fr 1fr;
}
</style>
