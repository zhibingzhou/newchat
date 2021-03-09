package response

type ResponseUserEmoticon struct {
	Sys_emoticon     []ResponseEmoticon `json:"sys_emoticon"`
	Collect_emoticon []ResponseEmoticon `json:"collect_emoticon"`
}

type ResponseSysEmoticon struct {
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Url    string `json:"url"`
	Name   string `json:"name"`
}

type ResponseEmoticon struct {
	Media_id int    `json:"media_id"`
	Src      string `json:"src"`
}

