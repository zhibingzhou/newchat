package response

type ResponseEvenTalk struct {
	Send_user    string `json:"send_user"`
	Receive_user string `json:"receive_user"`
	Text_message string `json:"text_message"`
	Msg_type     string `json:"msg_type"`
	Source_type  string `json:"source_type"`
}

//返回参数
type WebsocketMessage struct {
	Messagedata  WebsocketImage `json:"messagedata"`
	Receive_list []int          `json:"receive_list"`
	Event        string         `json:"event"`
}

type WebsocketImage struct {
	Source_type  int                `json:"source_type"` //群或者组
	Receive_user int                `json:"receive_user"`
	Send_user    int                `json:"send_user"`
	Data         ResponseTalkRecord `json:"data"`
}

type GroupListJoin struct {
	Send_user    int    `json:"send_user"`
	Event        string `json:"event"`
	Receivedlist []int  `json:"received"`
}
