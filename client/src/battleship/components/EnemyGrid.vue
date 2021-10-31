<template>
    <div>
        <div v-for="(_, h) in height" :key="h" class="flex justify-center">
            <div v-for="(_, w) in width" :key="h * width + w" class="border border-white">
              <div v-if="!canAttack(h * width + w)" class="w-10 h-10 bg-red-300"></div>
              <div v-else :class="{'hover:bg-blue-400': selectable}" class="w-10 h-10 bg-blue-300"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { useBattleship } from "../hooks/useBattleship";

export default {
  props: {
    selectable: Boolean
  },
  setup() {
    const { enemyBoard, mapWidth, mapHeight } = useBattleship();

    return {
      height: mapHeight,
      width: mapWidth,
      canAttack: (pos: number) => enemyBoard.value.canAttack(pos)
    };
  },
};
</script>

<style></style>
