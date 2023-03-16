<script setup lang="ts">
import { useResourcesStore } from "@/stores/resources";
import LibraryFolder from "@/components/content/library/LibraryFolder.vue";
import { ref } from "vue";
import LibraryItem from "@/components/content/library/LibraryItem.vue";
import ExpandCollapseButton from "@/components/utils/ExpandCollapseButton.vue";
import { storeToRefs } from "pinia";

const { advertisements, screens } = storeToRefs(useResourcesStore());

const expanded = ref<boolean[]>(
  new Array(advertisements.value.length).fill(false)
);
const screensExpanded = ref(false);
const adsExpanded = ref(false);

function setExpanded(value: boolean) {
  expanded.value.forEach((_, i, a) => (a[i] = value));
}
</script>

<template>
  <section class="column">
    <h2>
      <span>Library</span>
      <ExpandCollapseButton type="expand" @click="setExpanded(true)" />
      <ExpandCollapseButton type="collapse" @click="setExpanded(false)" />
    </h2>
    <ul>
      <LibraryFolder v-model:expanded="screensExpanded">
        <template #name>
          <strong>Screens</strong>
        </template>
        <LibraryItem
          v-for="screen in screens"
          :key="screen.id"
          :item="screen"
        />
      </LibraryFolder>
      <LibraryFolder v-model:expanded="adsExpanded">
        <template #name>
          <strong>Ads</strong>
        </template>
        <LibraryFolder
          v-for="(entry, index) in advertisements"
          :key="entry[0]"
          v-model:expanded="expanded[index]"
        >
          <template #name>{{ entry[0] }}</template>
          <LibraryItem
            v-for="ad in entry[1].sort((a, b) => a.name.localeCompare(b.name))"
            :key="ad.id"
            :item="ad"
          />
        </LibraryFolder>
      </LibraryFolder>
    </ul>
  </section>
</template>

<style scoped>
h2 {
  display: flex;
  align-items: center;
}

span {
  flex: 1;
}
</style>
