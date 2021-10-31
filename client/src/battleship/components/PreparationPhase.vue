<template>
  <div>
    <div class="flex justify-center space-x-4">
      <div :class="readyColor1" class="rounded-full w-4 h-4"></div>
      <div :class="readyColor2" class="rounded-full w-4 h-4"></div>
    </div>
    <ul>
      <li v-for="(ship, i) in board.fleet" :key="i">
        <span v-for="cell in ship.Cells" :key="`cell-${cell}`">
          {{ cell }}&nbsp;
        </span>
      </li>
    </ul>
    <GridVue />
    <button @click="getBoardFleet" class="mt-6 px-4 py-2 rounded-md bg-green-800 hover:bg-green-700 transition-all">Randomize</button>
    <button @click="ready" class="mt-6 px-4 py-2 rounded-md bg-green-800 hover:bg-green-700 transition-all">Ready</button>
  </div>
</template>

<script lang="ts">
import GridVue from './MyGrid.vue';
import { ready } from '../Api'
import { onMounted, computed, ref, watch } from 'vue';
import { useBattleship } from '../hooks/useBattleship';

export default {
  components: {
    GridVue,
  },
  setup() {
    const game = useBattleship();
    const { board, socket } = game;
    const isReady = ref(false);
    const isEnemyReady = ref(false);

    const readyColor1 = computed(() => {
      return isReady.value ? 'bg-green-500' : 'bg-red-500';
    })
    const readyColor2 = computed(() => {
      return isEnemyReady.value ? 'bg-green-500' : 'bg-red-500';
    })

    ready(socket, (playerNb) => {
      if (playerNb === game.playerNb) {
        isReady.value = true;
      } else {
        isEnemyReady.value = true;
      }
    })

    game.createBoard();

    onMounted(() => {
      game.getBoardFleet();
    });

    return {
      board,
      getBoardFleet: () => game.getBoardFleet(),
      ready: () => game.ready(),
      readyColor1,
      readyColor2,
    };
  },
};
</script>

<style></style>
