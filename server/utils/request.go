package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
)

func HttpPostjson(url string, js []byte) (int, string) {
	status := 100
	msg := "fail"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(js))
	if err != nil {
		return status, err.Error()
	}
	request.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	response, err := client.Do(request)
	if err != nil {

		return status, err.Error()
	}
	defer response.Body.Close()

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return status, err.Error()
	}

	status = 200
	msg = string(all)
	return status, msg
}

/**
*  处理分页
 */
func ThreadPage(page, page_size string) (int, int) {
	page_int, _ := strconv.Atoi(page)
	if page_int < 1 {
		page_int = 1
	}
	size_int, _ := strconv.Atoi(page_size)
	if size_int < 1 {
		size_int = 20
	} else if size_int > 100 {
		size_int = 100
	}
	return page_int, size_int
}
