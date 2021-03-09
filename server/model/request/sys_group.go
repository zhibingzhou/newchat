package request

type RequestGroupEdit struct {
	Avatar        string `json:"avatar"`
	Group_id      string `json:"group_id"`
	Group_name    string `json:"group_name"`
	Group_profile string `json:"group_profile"`
}

type RequestEditGroupEdit struct {
	Group_id  int `json:"group_id"`
	Content   string `json:"content"`
	Notice_id int    `json:"notice_id"`
	Title     string `json:"title"`
}
