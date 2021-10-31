<template>
  <div>
    <div class="text-lg text-white">Code: {{ generatedCode }}</div>
    <div>
      Share: <a :href="shareUrl" target="_blank">{{ shareUrl }}</a>
    </div>
  </div>
</template>

<script lang="ts">
import { computed } from 'vue';
import { useBattleship } from '../hooks/useBattleship';

export default {
  setup() {
    const { generatedCode } = useBattleship();

    const shareUrl = computed(() => {
      return `${document.location.origin}${document.location.pathname}?code=${generatedCode.value}`;
    });

    return {
      shareUrl,
      generatedCode
    };
  },
};
</script>
