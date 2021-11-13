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
import { AttackResultEnum } from '../enums/AttackResultEnum';

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

    function handleSunk(cells: string, playerNb: number) {
      try {
        let cellsNb = Array.from(atob(cells)).map((v) => v.charCodeAt(0));
        let board: Ref<Board> = null;
        if (playerNb == game.playerNb) {
          board = game.board
        } else {
          board = game.enemyBoards[playerNb]
        }

        for (let cell of cellsNb) {
          board.value.hits[cell] = AttackResultEnum.SUNK
        }
      } catch (e) {
        console.error(e.message);
      }
    }

    onMounted(() => {
      const attackWrap = attackResult(game.socket, handleResult, handleSunk)
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
