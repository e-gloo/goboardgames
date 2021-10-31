<template>
  <div>
    <ConnectionState :socket="socket" />
    <component :is="componentPhase"></component>
    <button
      v-if="componentPhase === null"
      class="px-6 py-2 rounded-md bg-green-300 text-gray-800"
      @click="playWithFriend"
    >
      Jouer avec un ami
    </button>
  </div>
</template>

<script lang="ts">
import { ref, onUnmounted, onMounted, watch, provide } from 'vue';
import ConnectionState from '@/components/ConnectionState.vue';
import { useRoute } from 'vue-router';
import { GamePhase } from '@/battleship/enums/GamePhase';
import WaitingPhaseVue from '@/battleship/components/WaitingPhase.vue';
import PreparationPhaseVue from '@/battleship/components/PreparationPhase.vue';
import { BATTLESHIP_KEY } from '@/battleship/components/constants';
import { Game } from '@/battleship/Game';

export default {
  components: {
    ConnectionState,
    WaitingPhaseVue,
    PreparationPhaseVue,
  },
  setup() {
    const game = new Game();
    const {phase, generatedCode, socket } = game;
    provide(BATTLESHIP_KEY, game);
    const route = useRoute();
    const componentPhase = ref(null);

    onMounted(() => {
      generatedCode.value = route.query?.code as string || '';
    });

    onUnmounted(() => {
      game.disconnect()
    })

    watch(phase, async (gamePhase) => {
      console.log('PHASE:', GamePhase[gamePhase]);

      switch (gamePhase) {
        case GamePhase.WAITING:
          componentPhase.value = 'WaitingPhaseVue';
          break;
        case GamePhase.PREPARATION:
          componentPhase.value = 'PreparationPhaseVue';
          break;
        default:
          break;
      }
    });

    return {
      socket,
      phase,
      generatedCode,
      playWithFriend: () => game.playWithFriend(),
      componentPhase,
    };
  },
};
</script>

<style lang=""></style>
