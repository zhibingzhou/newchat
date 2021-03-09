package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
	"github.com/woodylan/go-websocket/pkg/setting"
	"github.com/woodylan/go-websocket/tools/crypto"
)

//GenUUID 生成uuid
func GenUUID() string {
	uuidFunc := uuid.NewV4()
	uuidStr := uuidFunc.String()
	uuidStr = strings.Replace(uuidStr, "-", "", -1)
	uuidByt := []rune(uuidStr)
	return string(uuidByt[8:24])
}

//对称加密IP和端口，当做clientId
func GenClientId() string {
	raw := []byte(setting.GlobalSetting.LocalHost + ":" + setting.CommonSetting.RPCPort)
	str, err := crypto.Encrypt(raw, []byte(setting.CommonSetting.CryptoKey))
	if err != nil {
		panic(err)
	}

	return str
}

//解析redis的地址格式
func ParseRedisAddrValue(redisValue string) (host string, port string, err error) {
	if redisValue == "" {
		err = errors.New("解析地址错误")
		return
	}
	addr := strings.Split(redisValue, ":")
	if len(addr) != 2 {
		err = errors.New("解析地址错误")
		return
	}
	host, port = addr[0], addr[1]

	return
}

//判断地址是否为本机
func IsAddrLocal(host string, port string) bool {
	return host == setting.GlobalSetting.LocalHost && port == setting.CommonSetting.RPCPort
}

//是否集群
func IsCluster() bool {
	return setting.CommonSetting.Cluster
}

//获取client key地址信息
func GetAddrInfoAndIsLocal(clientId string) (addr string, host string, port string, isLocal bool, err error) {
	//解密ClientId
	addr, err = crypto.Decrypt(clientId, []byte(setting.CommonSetting.CryptoKey))
	if err != nil {
		return
	}

	host, port, err = ParseRedisAddrValue(addr)
	if err != nil {
		return
	}

	isLocal = IsAddrLocal(host, port)
	return
}

func GenGroupKey(systemId, groupName string) string {
	return systemId + ":" + groupName
}

func HttPRquest(url string, jsonStr []byte) (err error, rep []byte) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err, rep
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return err, body
}
