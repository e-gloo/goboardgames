<template>
    <div>
        <div v-for="(_, h) in height" :key="h" class="flex justify-center">
            <div v-for="(_, w) in width" :key="h * width + w" class="border border-white">
              <div v-if="isBoatHere(h * width + w)" class="w-10 h-10 bg-red-300"></div>
              <div v-else class="w-10 h-10 bg-blue-300"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { useBattleship } from "../hooks/useBattleship";

export default {
  setup() {
    const { board, mapWidth, mapHeight } = useBattleship();

    return {
      height: mapHeight,
      width: mapWidth,
      isBoatHere: (pos: number) => board.value.isBoatHere(pos),
    };
  },
};
</script>

<style></style>
