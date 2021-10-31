<template>
  <div>
    <div class="grid grid-cols-2">
      <GridVue />
      <EnemyGridVue :selectable="true" />
    </div>
  </div>
</template>

<script lang="ts">
import GridVue from './MyGrid.vue';
import EnemyGridVue from './EnemyGrid.vue';
import { onMounted, computed, ref, watch } from 'vue';
import { useBattleship } from '../hooks/useBattleship';

export default {
  components: {
    GridVue,
    EnemyGridVue,
  },
  setup() {
    const game = useBattleship();

    return {};
  },
};
</script>

<style>
</style>
