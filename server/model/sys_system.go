package model

import (
	"newchat/config"
)

// 配置文件结构体
type System struct {
	Config config.Server
}

type Totaln struct {
	Num int `json:"num"`
}
