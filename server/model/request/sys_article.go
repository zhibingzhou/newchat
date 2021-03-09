package request

type RequestEditArticleClass struct {
	Class_id   int    `json:"class_id"`
	Class_name string `json:"class_name"`
}


type RequestDelArticleClass struct {
	Class_id   int    `json:"class_id"`
}
