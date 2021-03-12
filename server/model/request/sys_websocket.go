package request

type Wregister struct {
	SystemId string `json:"systemId"`
}

type RequestEvenTalk struct {
	Send_user    int    `json:"send_user"`
	Receive_user string `json:"receive_user"`
	Text_message string `json:"text_message"`
	Msg_type     string `json:"msg_type"`    //事件类型
	Source_type  string `json:"source_type"` //群聊或者私聊
}

