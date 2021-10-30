<template>
  <div>
    <ConnectionState :socket="socket" />
    <component
      :is="componentPhase"
      :socket="socket"
      :code="generatedCode"
      :api="api"
    ></component>
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
import { ref, onUnmounted, watch } from 'vue';
import io from 'socket.io-client';
import { SOCKET_URL } from '@/constants';
import ConnectionState from '@/components/ConnectionState.vue';
import { useRoute } from 'vue-router';
import Api from '@/battleship/Api';
import { GamePhase } from '@/battleship/enums/GamePhase';
import WaitingPhaseVue from '@/components/battleship/WaitingPhase.vue';
import PreparationPhaseVue from '@/components/battleship/PreparationPhase.vue';

export default {
  components: {
    ConnectionState,
    WaitingPhaseVue,
    PreparationPhaseVue,
  },
  setup() {
    const route = useRoute();
    const phase = ref<GamePhase>(null);
    const componentPhase = ref(null);
    const socket = io(`${SOCKET_URL}/battleship`, {
      reconnection: false,
    });
    const api = new Api(socket);

    const generatedCode = ref((route.query?.code as string) || '');
    socket.connect();

    socket.on('connect', () => {
      console.log('connected');
      if (generatedCode.value) {
        socket.emit('joinRoom', generatedCode.value);
      }
    });
    socket.on('disconnect', () => {
      console.log('disconnected');
    });

    socket.on('newCode', (data: string) => {
      generatedCode.value = data;
    });
    socket.on('gamePhaseUpdated', async (data: number) => {
      phase.value = data;
    });

    function playWithFriend() {
      phase.value = GamePhase.WAITING;
      socket?.emit('playWithFriend');
    }

    onUnmounted(() => {
      socket?.disconnect();
    });

    watch(phase, async (gamePhase) => {
      console.log('Gamephase: ', gamePhase);
      console.log('Preparation: ', GamePhase.PREPARATION);

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
      playWithFriend,
      componentPhase,
      api,
    };
  },
};
</script>

<style lang=""></style>
