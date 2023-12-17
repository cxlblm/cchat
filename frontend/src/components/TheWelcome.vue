<script setup>

import {ref} from "vue";
const props = defineProps(['wsSocket', 'chats'])
let count = ref(0)
let input = ref("")
let chats = props.chats
function submit() {
  if (input.value !== "") {
    count.value++
  }
  props.wsSocket.send(input.value)
  console.log(input.value)
}
</script>

<template>
  <div>
    回复：{{ count }} 次
  </div>
  <div class="chart">
    <p>聊天内容</p>
    <ul>
      <li :key="i" v-for="(v, i) in chats">{{v}}</li>
    </ul>
  </div>
  <div>
    <input type="text" v-model="input">
    <button @click="submit">提交</button>
  </div>
</template>
<style>
@media (min-width: 1024px) {
  .chart{
    min-height: 50vh;
    align-items: center;
  }
}
</style>