package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"newchat/global"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/google/uuid"
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

func Base64ToImage(file string) (error, string) {

	img := strings.Split(file, ",")
	// 上传文件至指定目录
	guid := uuid.New().String()

	filename := fmt.Sprintf("./uploads/file/img/%s.png", guid)
	if len(img) < 2 {
		return errors.New("图片格式错误"), ""
	}
	// 解压
	dist, _ := base64.StdEncoding.DecodeString(img[1])
	// 写入新文件
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	weburl := global.GVA_CONFIG.System.Url + ":" + strconv.Itoa(global.GVA_CONFIG.System.Addr)
	url := fmt.Sprintf("%s/uploads/file/img/%s.png", weburl, guid)

	return err, url
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
