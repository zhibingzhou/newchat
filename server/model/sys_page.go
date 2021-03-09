package model

type PageToal struct {
	Page       int `json:"page"`
	Page_size  int `json:"page_size"`
	Page_total int `json:"page_total"`
	Total      int `json:"total"`
}
