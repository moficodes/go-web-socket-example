var app = new Vue({
    el: '#app',
    data: {
        ws: null,
        serverUrl: "ws://" + location.host + "/ws",
        roomInput: null,
        rooms: [],
        user: {
            name: ""
        },
        users: []
    },

    methods: {
        connect() {
            this.connectToWebsocket()
        },
        connectToWebsocket() {
            this.ws = new WebSocket(this.serverUrl + "?name=" + this.user.name);
            this.ws.addEventListener('open', (event) => {
                this.onWebsocketOpen(event)
            });
            this.ws.addEventListener('message', (event) => {
                this.handleNewMessage(event)
            });
        },
        onWebsocketOpen() {
            console.log("connected to WS!");
        },
        handleNewMessage(event) {
            let data = event.data;
            data = data.split(/\r?\n/);
            for (let i = 0; i < data.length; i++) {
                let msg = JSON.parse(data[i]);
                switch (msg.action) {
                    case "send-message":
                        this.handleChatMessage(msg);
                        break;
                    case "user-join":
                        this.handleUserJoined(msg);
                        break;
                    case "user-left":
                        this.handleUserLeft(msg);
                        break;
                    case "room-joined":
                        this.handleRoomJoined(msg);
                        break;
                    default:
                        break;
                }
            }
        },
        handleRoomJoined(msg) {
            let room = msg.target;
            room.name = room.private ? msg.sender.name : room.name;
            room["messages"] = [];
            this.rooms.push(room);
        },
        handleUserLeft(msg) {
            for (let i = 0; i < this.users.length; i++) {
                if (this.users[i].id === msg.sender.id) {
                    this.users.splice(i, 1)
                }
            }
        },
        handleUserJoined(msg) {
            this.users.push(msg.sender);
        },
        handleChatMessage(msg) {
            const room = this.findRoom(msg.target.id);
            console.log(room);
            if (typeof room !== "undefined") {
                room.messages.push(msg);
            }
        },
        findRoom(roomID) {
            for (let i = 0; i < this.rooms.length; i++) {
                if (this.rooms[i].id === roomID) {
                    return this.rooms[i];
                }
            }
        },
        joinRoom() {
            this.ws.send(JSON.stringify({action: 'join-room', message: this.roomInput}));
            this.roomInput = "";
        },
        joinPrivateRoom(room) {
            this.ws.send(JSON.stringify({action: 'join-room-private', message: room.id}))
        },
        leaveRoom(room) {
            this.ws.send(JSON.stringify({action: 'leave-room', message: room.name}));
            for (let i = 0; i < this.rooms.length; i++) {
                if (this.rooms[i].name === room.name) {
                    this.rooms.splice(i, 1);
                    break;
                }
            }
        },
        sendMessage(room) {
            if (room.newMessage !== "") {
                this.ws.send(JSON.stringify({
                    action: 'send-message',
                    target: {
                        id: room.id,
                        name: room.name,
                    },
                    message: room.newMessage
                }));
                room.newMessage = "";
            }
        },

    }
})
