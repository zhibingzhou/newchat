package servers

import (
	"encoding/json"
)

var ChannelAll *ChannelMessage

type ChannelMessage struct {
	//接受消息
	ChannelReceiveMessage chan []byte
}

func NewChannelMessage() *ChannelMessage {
	return &ChannelMessage{
		ChannelReceiveMessage: make(chan []byte),
	}
}

func ChannelListenMessage() {
	go func() {
		for {
			select {
			case r := <-ChannelAll.ChannelReceiveMessage:
				SendToRabbitMQ(r)
			}
		}
	}()
}

func SendToRabbitMQ(Message []byte) {

	var eventalk RequestEvenTalk
	err := json.Unmarshal(Message, &eventalk)
	if eventalk.Msg_type != "" { //判断类型
		if err != nil || eventalk.Msg_type == "" {
			return
		}
		//发送到队列
		RabbitWebSocketService["even_talk"].DirectPulish(Message)
		return
	}

	var evenlogin UpdateUserStatus
	err = json.Unmarshal(Message, &evenlogin)
	if evenlogin.Event != "" && evenlogin.User_id != 0 {
		if err != nil || evenlogin.Event == "" {
			return
		}
		//发送到队列
		RabbitWebSocketService["login_event"].DirectPulish(Message)
		return
	}
	var evenKeyboard KeyBoard

	err = json.Unmarshal(Message, &evenKeyboard)
	if evenKeyboard.Event != "" {
		//发送到队列
		RabbitWebSocketService["event_keyboard"].DirectPulish(Message)
		return
	}
}
