<template lang="">
  <div>
    <ConnectionState :socket="socket"/>
    <template v-if="generatedCode">
      <div class="text-lg text-white">Code: {{ generatedCode }}</div>
      <div>Share: <a :href="shareUrl" target="_blank">{{ shareUrl }}</a></div>
    </template>
    <button class="px-6 py-2 rounded-md bg-green-300 text-gray-800" @click="playWithFriend">Jouer avec un ami</button>
  </div>
</template>

<script lang="ts">
import { Vue, Options } from "vue-class-component";
import io, { Socket } from "socket.io-client";
import { SOCKET_URL } from '@/constants';
import ConnectionState from '@/components/ConnectionState.vue';

@Options({
  components: {
    ConnectionState
  }
})
export default class BattleshipPage extends Vue {
  socket?: Socket = undefined;
  generatedCode = '';

  created() {
    this.generatedCode = this.$route.query?.code as string;
    this.socket = io(`${SOCKET_URL}/battleship`, {
      reconnection: false
    });
    this.socket.connect();

    this.socket.on("connect", () => {
      console.log("connected");
      if (this.generatedCode) {
        this.socket.emit('joinRoom', this.generatedCode)
      }
    });
    this.socket.on("disconnect", () => {
      console.log("disconnected");
    });

    this.socket.on("newCode", (data: string) => {
      this.generatedCode = data
    })
    this.socket.on("gamePhaseUpdated", (data: number) => {
      if (data === 1) {
        this.socket.emit('randomizeFleet')
      }
    })
    this.socket.on("NewFleet", (data: any) => {
      const ships = data.map(s => {
        const cells = Array.from(atob(s.Cells)).map(v => v.charCodeAt(0))
        return cells;
      })
      console.log("NewFleet", ships);
    })
  }
	
  unmounted() {
    this.socket?.disconnect();
  }

  playWithFriend() {
    this.socket?.emit('playWithFriend');
  }

  get shareUrl() {
    return `${document.location.origin}${document.location.pathname}?code=${this.generatedCode}`
  }
}
</script>

<style lang=""></style>
