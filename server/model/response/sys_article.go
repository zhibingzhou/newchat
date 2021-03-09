package response

import (
	"newchat/model"
	"time"
)

type ResponseArticleClass struct {
	Id         int    `json:"id"`
	Class_name string `json:"class_name"`
	Is_default int    `json:"is_default"`
	Count      int    `json:"count"`
}

type ResponseArticleClassList struct {
	Rows []ResponseArticleClass `json:"rows"`
}

type ResponseArticleList struct {
	Rows []ResponseArticle `json:"rows"`
	model.PageToal
}

type ResponseArticle struct {
	Id         int       `json:"id"`
	Class_id   int       `json:"class_id"`
	Title      string    `json:"title"`
	Image      string    `json:"image"`
	Abstract   string    `json:"abstract"`
	Updated_at time.Time `json:"updated_at"`
	Class_name string    `json:"class_name"`
	Status     int       `json:"status"`
}

type ResponseTags struct {
	Id       int    `json:"id"`
	Tag_name string `json:"tag_name"`
	Count    int    `json:"count"`
}

type ResponseTagsList struct {
	Tags []ResponseTags `json:"tags"`
}
