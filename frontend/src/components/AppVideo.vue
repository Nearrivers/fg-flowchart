<template>
  <main class="h-full w-full">
    <header class="mb-4 py-2 text-center text-sm opacity-65">
      {{ path }}
    </header>
    <div class="flex w-full justify-center px-2">
      <video
        controls
        preload="auto"
        alt="Fichier vidéo"
        class="block max-w-full"
        :src="src"
      ></video>
    </div>
  </main>
</template>

<script setup lang="ts">
import { OpenMedia, OpenMediaConc } from '$/file_handler/FileHandler';
import { useShowErrorToast } from '@/composables/useShowErrorToast';
import { computed, ref, watchEffect } from 'vue';
import { useRoute } from 'vue-router';

const src = ref('');
const route = useRoute();
const { showToast } = useShowErrorToast();

const path = computed(() =>
  (route.params.path as string).replaceAll('/', ' / '),
);

watchEffect(async () => {
  try {
    console.time('open media');
    src.value = await OpenMedia(route.params.path as string);
    console.timeEnd('open media');
  } catch (error) {
    showToast(error);
  }
});
</script>
