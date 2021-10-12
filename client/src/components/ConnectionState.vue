<template>
  <div class="grid grid-cols-2 w-max text-xs">
    <div class="font-bold text-uppercase">Status</div>
    <div class="flex justify-center items-center">
      <div class="w-3 h-3 rounded-full mr-2" :class="[connected ? 'bg-green-500' : 'bg-red-500']"></div>
      <div>{{ connected ? 'Connected' : 'Disconnected' }}</div>
    </div>
    <div class="font-bold text-uppercase">Host</div>
    <div>{{ opts?.hostname }}</div>
    <div class="font-bold text-uppercase">Port</div>
    <div>{{ opts?.port }}</div>
    <div class="font-bold text-uppercase">Secure</div>
    <div>{{ opts?.secure }}</div>
  </div>
</template>

<script lang="ts">
import { Socket } from 'socket.io-client';
import { Options, Vue } from "vue-class-component";

@Options({
  props: {
    socket: Socket,
  },
})
export default class ConnectionState extends Vue {
  socket?: Socket;
  connected = false;

  created() {
    this.socket?.on("connect", () => {
      this.connected = true;
    });
    this.socket?.on("disconnect", () => {
      this.connected = false;
    });
  }

  get opts() {
    return this.socket?.io?.opts;
  }
}
</script>

<style scoped>
</style>
