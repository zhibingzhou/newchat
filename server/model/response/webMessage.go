package response

type ResponseEvenTalk struct {
	Send_user    string `json:"send_user"`
	Receive_user string `json:"receive_user"`
	Text_message string `json:"text_message"`
	Msg_type     string `json:"msg_type"`
	Source_type  string `json:"source_type"`
}
