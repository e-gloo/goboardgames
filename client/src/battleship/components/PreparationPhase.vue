<template>
  <div>
    <ul>
      <li v-for="(ship, i) in board.fleet" :key="i">
        <span v-for="cell in ship.Cells" :key="`cell-${cell}`">
          {{ cell }}&nbsp;
        </span>
      </li>
    </ul>
    <GridVue />
    <button @click="getBoardFleet" class="mt-6 px-4 py-2 rounded-md bg-green-800 hover:bg-green-700 transition-all">Randomize</button>
  </div>
</template>

<script lang="ts">
import GridVue from './Grid.vue';
import { onMounted } from 'vue';
import { useBattleship } from '../hooks/useBattleship';

export default {
  components: {
    GridVue,
  },
  props: {
    api: Object,
  },
  setup() {
    const game = useBattleship();
    const { board } = game;

    game.createBoard();

    onMounted(() => {
      game.getBoardFleet();
    });

    return {
      board,
      getBoardFleet: () => game.getBoardFleet(),
    };
  },
};
</script>

<style></style>
