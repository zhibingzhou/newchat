package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
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
