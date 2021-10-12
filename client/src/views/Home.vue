<template>
  <div class="flex justify-center items-center w-full">
    <router-link to="/battleship">
      <button class="flex flex-col justify-center items-center p-6 w-80 h-80 bg-green-200 rounded-md transform transition-all scale-100 -rotate-0 hover:-rotate-3 hover:scale-105 focus:scale-100 focus:rotate-0">
          <h1 class="font-bold text-xl">Battleship</h1>
      </button>
    </router-link>
  </div>
</template>

<script lang="ts">
import { Options, Vue } from "vue-class-component";
// import HelloWorld from "@/components/HelloWorld.vue"; // @ is an alias to /src

@Options({
  components: {
    // HelloWorld,
  },
})
export default class Home extends Vue {}
</script>
