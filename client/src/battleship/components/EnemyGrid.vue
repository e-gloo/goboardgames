<template>
  <div>
    <div class="text-xl">
      <span v-if="yourTurn">A vous de jouer</span>
      <span v-else>...</span>
    </div>
    <div class="flex justify-center items-center">
      <div class="bg-blue-400 rounded-lg border-4" :class="[yourTurn ? 'border-green-400' : 'border-red-400']">
        <div v-for="(_, h) in height" :key="h" class="flex justify-center">
            <div @click="attack(h * width + w)" v-for="(_, w) in width" :key="h * width + w" class="flex justify-center items-center w-10 h-10" :class="{'hover:bg-blue-300 hover:bg-opacity-50 hover:rounded-md': selectable && getCellType(h * width + w) !== 0}">
              <div v-if="getCellType(h * width + w) === 2" class="w-full h-full bg-gray-700 rounded-lg"></div>
              <div v-else-if="getCellType(h * width + w) === 0" class="rounded-full w-1.5 h-1.5 bg-gray-700"></div>
              <div v-else-if="!canAttack(h * width + w)" class="w-full h-full bg-red-300"></div>
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
  props: {
    selectable: Boolean,
    playerNb: Number,
  },
  setup(props) {
    const game = useBattleship();
    const { enemyBoards, mapWidth, mapHeight, yourTurn, playerTurn } = game;
    const board = enemyBoards[props.playerNb]

    const getCellType = ((cell: number) => {
      return board.value.hits[cell]
    })

    return {
      height: mapHeight,
      width: mapWidth,
      attack: (pos: number) => {
        if (yourTurn.value && enemyBoards[props.playerNb].value.canAttack(pos)) {
          game.attack(pos)
        }
      },
      canAttack: (pos: number) => enemyBoards[props.playerNb].value.canAttack(pos),
      getCellType,
      yourTurn,
      playerTurn,
    };
  },
};
</script>

<style></style>
