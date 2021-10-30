<template>
  <div class="grid grid-cols-2 w-max text-xs">
    <div class="font-bold text-uppercase">Status</div>
    <div class="flex justify-center items-center">
      <div
        class="w-3 h-3 rounded-full mr-2"
        :class="[connected ? 'bg-green-500' : 'bg-red-500']"
      ></div>
      <div class="flex">
        <div class="mr-2">{{ connected ? "Connected" : "Disconnected" }}</div>
        <button
          class="w-3 h-3 rounded-sm"
          :class="[connected ? 'bg-red-500' : 'bg-green-500']"
          @click="toggleConnect"
        ></button>
      </div>
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
import { ref, computed } from "vue";
import { Socket } from "socket.io-client";

export default {
  props: {
    socket: Socket,
  },
  setup(props) {
    const connected = ref(false);
    props.socket?.on("connect", () => {
      connected.value = true;
    });
    props.socket?.on("disconnect", () => {
      connected.value = false;
    });

    function toggleConnect() {
      !connected.value ? props.socket?.connect() : props.socket?.disconnect();
    }
    const opts = computed(() => props.socket?.io?.opts);

    return {
      toggleConnect,
      opts,
      connected,
    };
  },
};
</script>

<style scoped></style>
