<template>
  <div>
    <div class="grid grid-cols-2">
      <GridVue />
      <EnemyGridVue v-for="player in players" :key="`player-${player}`" :selectable="true" :playerNb="+player" />
    </div>
  </div>
</template>

<script lang="ts">
import GridVue from './MyGrid.vue';
import EnemyGridVue from './EnemyGrid.vue';
import { useBattleship } from '../hooks/useBattleship';
import { onMounted, Ref } from '@vue/runtime-core';
import { attackResult, gameOver } from '../Api';
import { AttackResult } from '../types/AttackResult';
import { Board } from '../Board';

export default {
  components: {
    GridVue,
    EnemyGridVue,
  },
  setup() {
    const game = useBattleship();

    const players = Object.keys(game.enemyBoards);

    function handleResult(attackResult: AttackResult) {
      let board: Ref<Board> = null;
      if (attackResult.PlayerNb == game.playerNb) {
        board = game.board
      } else {
        board = game.enemyBoards[attackResult.PlayerNb]
      }

      board.value.hits[attackResult.Position] = attackResult.Result
      return;
    }

    onMounted(() => {
      const attackWrap = attackResult(game.socket, handleResult)
      gameOver(game.socket, attackWrap)
    })

    return {
      players
    };
  },
};
</script>

<style>
</style>
