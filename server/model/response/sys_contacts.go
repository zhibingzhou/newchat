package response

import "newchat/model"

type Apply_unread_num struct {
	Unread_num int `json:"unread_num"`
}

type ResponseApply struct {
	Rows       []model.Apply_records `json:"rows"`
	Page       int `json:"page"`
	Page_size  int `json:"page_size"`
	Page_total int `json:"page_total"`
	Total      int `json:"total"`
}
