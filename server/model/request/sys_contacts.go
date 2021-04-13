package request

type RequestContactsAdd struct {
	Friend_id int    `json:"friend_id"`
	Remarks   string `json:"remarks"`
}

type RequestDeleteContacts struct {
	Apply_id int `json:"apply_id"`
}

type RequestAcceptApply struct {
	Apply_id int    `json:"apply_id"`
	Remarks  string `json:"remarks"`
}

type RequestContactsDelete struct {
	Friend_id int `json:"friend_id"`
}

type RequestContactsEditRemarks struct {
	Friend_id int    `json:"friend_id"`
	Remarks   string `json:"remarks"`
}
