<template>
  <div>
    <div class="text-lg text-white">Code: {{ code }}</div>
    <div>
      Share: <a :href="shareUrl" target="_blank">{{ shareUrl }}</a>
    </div>
  </div>
</template>

<script lang="ts">
import { computed } from 'vue';

export default {
  props: {
    code: String,
  },
  setup(props) {
    const shareUrl = computed(() => {
      return `${document.location.origin}${document.location.pathname}?code=${props.code}`;
    });

    return {
      shareUrl,
    };
  },
};
</script>
