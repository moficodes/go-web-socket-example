package main

import (
	"encoding/json"
	"log"

	"github.com/moficodes/go-web-socket-example/models"
)

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const UserJoinedAction = "user-join"
const UserLeftAction = "user-left"
const JointRoomPrivateAction = "join-room-private"
const RoomJoinedAction = "room-joined"

type Message struct {
	Action  string      `json:"action"`
	Message string      `json:"message"`
	Target  *Room       `json:"target"`
	Sender  models.User `json:"sender"`
}

func (message *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	msg := &struct {
		Sender Client `json:"sender"`
		*Alias
	}{
		Alias: (*Alias)(message),
	}

	if err := json.Unmarshal(data, &msg); err != nil {
		return err
	}

	message.Sender = &msg.Sender
	return nil
}

func (message *Message) encode() []byte {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	return data
}
