package response

type ReponseLoginEvent struct {
	User_id      int    `json:"user_id"`
	Status       int   `json:"status"`
	Receivedlist []int  `json:"received"`
	Event        string `json:"event"`
}
