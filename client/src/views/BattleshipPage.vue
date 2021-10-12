<template lang="">
  <div>
    <ConnectionState :socket="socket"/>
    <div ref="chat">
      <div :key="i" v-for="(msg, i) in messages">{{ msg }}</div>
    </div>
    <input class="border" v-model="text"/>
    <button @click="send">Send</button>
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
  messages: string[] = [];
  text = '';

  created() {
    this.socket = io(`${SOCKET_URL}`);
    this.socket.connect();

    this.socket.on("connect", () => {
      // ...
    });
    this.socket.on("disconnect", () => {
      // ...
    });
    this.socket.on("ping", () => {
      console.log('ping');
    });
    this.socket.on("reply", (msg) => {
      this.messages.push(msg);
    });
  }

  send() {
    this.socket?.emit('notice', this.text)
    this.text = '';
  }
}
</script>

<style lang=""></style>
