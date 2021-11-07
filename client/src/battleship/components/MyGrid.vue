<template>
  <div class="flex justify-center items-center">
    <div class="bg-blue-400 rounded-lg border-2 border-white">
      <div v-for="(_, h) in height" :key="h" class="flex justify-center">
        <div
          v-for="(_, w) in width"
          :key="h * width + w"
          class="flex justify-center items-center w-10 h-10 border-white"
          :class="{ 'border-b': h < 9, 'border-r': w < 9 }"
        >
          <div
            v-if="getCellType(h * width + w) === 1 || getCellType(h * width + w) === 2"
            class="w-full h-full bg-red-700"
          ></div>
          <div v-else-if="isBoatHere(h * width + w)" class="w-full h-full bg-red-300"></div>
          <div
            v-else-if="getCellType(h * width + w) === 0"
            class="rounded-full w-1.5 h-1.5 bg-gray-700"
          ></div>
          <div v-else class="rounded-full w-1.5 h-1.5 bg-gray-700 opacity-20"></div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { useBattleship } from "../hooks/useBattleship";

export default {
  setup() {
    const { board, mapWidth, mapHeight } = useBattleship();

    const getCellType = ((cell: number) => {
      return board.value.hits[cell]
    })

    return {
      height: mapHeight,
      width: mapWidth,
      getCellType,
      isBoatHere: (pos: number) => board.value.isBoatHere(pos),
    };
  },
};
</script>

<style>
</style>
