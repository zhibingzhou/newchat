package request

type RequestGroupEdit struct {
	Avatar        string `json:"avatar"`
	Group_id      string `json:"group_id"`
	Group_name    string `json:"group_name"`
	Group_profile string `json:"group_profile"`
}

type RequestEditGroupEdit struct {
	Group_id  int    `json:"group_id"`
	Content   string `json:"content"`
	Notice_id int    `json:"notice_id"`
	Title     string `json:"title"`
}

type RequestGroupCreate struct {
	Group_avatar  string `json:"group_avatar"`
	Group_name    string `json:"group_name"`
	Group_profile string `json:"group_profile"`
	Uids          string `json:"uids"`
}

type RequestGroupInvite struct {
	Group_id int    `json:"group_id"`
	Uids     string `json:"uids"`
}

type RequestGroupSecede struct {
	Group_id int `json:"group_id"`
}

type RequestGroupSetCard struct {
	Group_id   int    `json:"group_id"`
	Visit_card string `json:"visit_card"`
}

type RequestGroupRemove struct {
	Group_id    int   `json:"group_id"`
	Members_ids []int `json:"members_ids"`
}
