<script setup>
import {ref} from "vue";
import {useRouter} from "vue-router";
import {useWebsocketStore} from "@/stores/websocket.js";

const router = useRouter()
const wsStore = useWebsocketStore()
let nick = ref("")

function createWebsocket() {
  let name = nick.value
  wsStore.init(name)

  router.push("/chat")
}

</script>

<template>
  <main>
    <div class="nick">
      <input type="text" v-model="nick">
    </div>
    <div>
      <a class="confirm" @click.prevent="createWebsocket">进入</a>
    </div>
  </main>
</template>

<style scoped>
main {
  display: flex;
  flex-direction: column;
  align-content: center;
  align-items: center;
}

main div {
  margin-top: 10px;
}

input {
  box-sizing: border-box;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
  outline: none;
  transition: border-color 0.3s;
}

input:focus {
  border-color: dodgerblue;
}

.confirm {
  display: inline-block;
  background: dodgerblue;
  color: white;
  border-radius: 10px;
  padding: 10px 20px;
}
</style>
