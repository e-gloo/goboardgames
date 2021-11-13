<template>
  <div>
    <div class="text-xl">
      <span v-if="!yourTurn && playerTurn > 0">Au joueur {{ playerTurn }} de jouer</span>
      <span v-else>...</span>
    </div>
    <div class="flex justify-center items-center">
      <div class="bg-blue-400 rounded-lg border-4 border-white" :class="[!yourTurn ? 'border-green-400' : 'border-red-400']">
        <div v-for="(_, h) in height" :key="h" class="flex justify-center">
          <div
            v-for="(_, w) in width"
            :key="h * width + w"
            class="flex justify-center items-center w-10 h-10"
          >
            <div v-if="getCellType(h * width + w) === 1" class="w-full h-full bg-red-700"></div>
            <div v-else-if="getCellType(h * width + w) === 2" class="w-full h-full bg-gray-700 rounded-lg"></div>
            <div v-else-if="isBoatHere(h * width + w)" class="w-full h-full bg-yellow-300 rounded-lg"></div>
            <div v-else-if="getCellType(h * width + w) === 0" class="rounded-full w-1.5 h-1.5 bg-gray-700"></div>
            <div v-else class="rounded-full w-1.5 h-1.5 bg-gray-700 opacity-20"></div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { useBattleship } from "../hooks/useBattleship";

export default {
  setup() {
    const { board, mapWidth, mapHeight, yourTurn, playerTurn } = useBattleship();

    const getCellType = ((cell: number) => {
      return board.value.hits[cell]
    })

    return {
      height: mapHeight,
      width: mapWidth,
      getCellType,
      yourTurn,
      playerTurn,
      isBoatHere: (pos: number) => board.value.isBoatHere(pos),
    };
  },
};
</script>

<style>
</style>
