<script setup>
import {useWebsocketStore} from "@/stores/websocket.js";
import {nextTick, onMounted, ref} from "vue";
import {useRouter} from "vue-router";

const wsStore = useWebsocketStore()
const router = useRouter()

let content = ref("")
const messageList = ref(null)

onMounted(() => {
  if (!wsStore.ws) {
    router.push("/")
  }
})

function submit() {
  wsStore.ws.send(content.value)
}

wsStore.$subscribe((mutation, state) => {
  if (mutation.storeId === 'websocket') {
    console.log(mutation, state)
    nextTick(() => {
      messageList.value.scrollTop = messageList.value.scrollHeight;
    })
  }
})


</script>

<template>
  <div class="chat">
    <div class="chat-window">
      <ul class="message-list" ref="messageList">
        <li :key="i" v-for="(v, i) in wsStore.chats" class="user-message">
          {{v}}
        </li>
      </ul>
    </div>
    <div class="input-area">
      <input v-model="content" @keyup.enter="submit" placeholder="Type your message...">
      <button @click="submit">Send</button>
    </div>

  </div>
</template>

<style>
.chat-window {
  width: 800px;
  border: 1px solid #ccc;
  border-radius: 8px;
  overflow: hidden;
}

.message-list {
  list-style-type: none;
  padding: 10px;
  margin: 0;
  overflow-y: auto;
  height: 500px;
}

.user-message {
  background-color: #64b5f6;
  color: white;
  border-radius: 5px;
  padding: 8px;
  margin-bottom: 8px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.other-message {
  background-color: #e0e0e0;
  color: #333;
  border-radius: 5px;
  padding: 8px;
  margin-bottom: 8px;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.input-area {
  display: flex;
  align-items: center;
  padding: 8px;
}

.input-area input {
  flex: 1;
  padding: 8px;
  margin-right: 8px;
}

.input-area button {
  padding: 8px;
  cursor: pointer;
}
</style>
