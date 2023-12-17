import {defineStore} from "pinia";
import {reactive} from "vue";

export const useWebsocketStore = defineStore('websocket', {
    state: () => ({
        ws: null,
        chats: reactive([]),
    }),
    actions: {
        init(name) {
            let conn = new WebSocket("ws://localhost:8088/ws?name=" + name);
            conn.onopen = function (evt) {
                conn.send("Hello my name is: " + name);
            };

            // 这里必须使用箭头函数引用外部的this
            conn.onmessage = (evt) => {
                console.log("Received Message: " + evt.data);
                this.chats.push(evt.data)
            };

            conn.onclose = function (evt) {
                console.log("Connection closed.");
            };

            this.ws = conn
        }
    }
})