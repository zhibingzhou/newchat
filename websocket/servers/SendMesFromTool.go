package servers

import (
	"encoding/json"
)

var ChannelAll *ChannelMessage

type ChannelMessage struct {
	//接受消息
	ChannelReceiveMessage chan SendFrom
}

type SendFrom interface {
	Do()
}

func NewChannelMessage() *ChannelMessage {
	return &ChannelMessage{
		ChannelReceiveMessage: make(chan SendFrom),
	}
}

func ChannelListenMessage() {
	go func() {
		for {
			select {
			case r := <-ChannelAll.ChannelReceiveMessage:
				r.Do()
			}
		}
	}()
}

type RabbitMqSend struct {
	Message []byte
}

//rabbitmq
func (r RabbitMqSend) Do() {

	var eventalk RequestEvenTalk
	err := json.Unmarshal(r.Message, &eventalk)
	if eventalk.Msg_type != "" { //判断类型
		if err != nil || eventalk.Msg_type == "" {
			return
		}
		//发送到队列
		RabbitWebSocketService["even_talk"].DirectPulish(r.Message)
		return
	}

	var evenlogin UpdateUserStatus
	err = json.Unmarshal(r.Message, &evenlogin)
	if evenlogin.Event != "" && evenlogin.User_id != 0 {
		if err != nil || evenlogin.Event == "" {
			return
		}
		//发送到队列
		RabbitWebSocketService["login_event"].DirectPulish(r.Message)
		return
	}
	var evenKeyboard KeyBoard

	err = json.Unmarshal(r.Message, &evenKeyboard)
	if evenKeyboard.Event != "" {
		//发送到队列
		RabbitWebSocketService["event_keyboard"].DirectPulish(r.Message)
		return
	}
}

type KafkaSend struct {
	Message []byte
}

func (k KafkaSend) Do() {

	var eventalk RequestEvenTalk
	err := json.Unmarshal(k.Message, &eventalk)
	if eventalk.Msg_type != "" { //判断类型
		if err != nil || eventalk.Msg_type == "" {
			return
		}
		KafkaSendService["even_talk"].KafkaSendMessage(k.Message)
		//发送到队列
		return
	}

	var evenlogin UpdateUserStatus
	err = json.Unmarshal(k.Message, &evenlogin)
	if evenlogin.Event != "" && evenlogin.User_id != 0 {
		if err != nil || evenlogin.Event == "" {
			return
		}
		KafkaSendService["login_event"].KafkaSendMessage(k.Message)
		return
	}
	var evenKeyboard KeyBoard

	err = json.Unmarshal(k.Message, &evenKeyboard)
	if evenKeyboard.Event != "" {
		//发送到队列
		KafkaSendService["event_keyboard"].KafkaSendMessage(k.Message)
		return
	}

}
