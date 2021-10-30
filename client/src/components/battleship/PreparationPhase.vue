<template>
  <div>
    <ul>
      <li v-for="(ship, i) in fleet" :key="i">
        <span v-for="cell in ship.Cells" :key="`cell-${cell}`">
          {{ cell }}&nbsp;
        </span>
      </li>
    </ul>
    <GridVue :fleet="fleet" />
    <button @click="randomizeFleet" class="mt-6 px-4 py-2 rounded-md bg-green-800 hover:bg-green-700 transition-all">Randomize</button>
  </div>
</template>

<script lang="ts">
import { Fleet } from '@/battleship/types/Fleet';
import GridVue from './Grid.vue';
import { ref, onMounted } from 'vue';

export default {
  components: {
    GridVue,
  },
  props: {
    api: Object,
  },
  setup(props) {
    const fleet = ref<Fleet>([]);

    onMounted(() => {
      randomizeFleet();
    });

    async function randomizeFleet() {
      fleet.value = await props.api.randomizeFleet();
      console.log('NewFleet', fleet);
    }

    return {
      fleet,
      randomizeFleet,
    };
  },
};
</script>

<style></style>
