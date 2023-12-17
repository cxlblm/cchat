<script setup>

import {ref, defineEmits} from "vue";
let emit=defineEmits(['createWebsocketA'])
let nick = ref("")

function login() {
  emit('createWebsocketA', nick)
}

</script>

<template>
  昵称<input type="text" v-model="nick">
  <button @click="login">进入</button>
</template>

<style>
@media (min-width: 1024px) {
  .chart{
    min-height: 50vh;
    align-items: center;
  }
}
</style>