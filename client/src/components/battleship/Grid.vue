<template>
    <div id="board">
        <div v-for="(_, h) in height" :key="h" class="flex justify-center">
            <div v-for="(_, w) in width" :key="h * width + w" class="border border-white">
              <div v-if="isBoatHere(h * width + w)" class="w-10 h-10 bg-red-300"></div>
              <div v-else class="w-10 h-10 bg-blue-300"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { MAP_HEIGHT, MAP_WIDTH } from '@/constants';

export default {
  props: {
    fleet: Array
  },
  setup(props) {
    const isBoatHere = (pos: number) => {
      return props.fleet.filter(f => f.Cells.includes(pos)).length
    }
    return {
      isBoatHere,
      height: MAP_HEIGHT,
      width: MAP_WIDTH
    };
  },
};
</script>

<style></style>
