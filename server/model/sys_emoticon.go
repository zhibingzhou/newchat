package model

import "newchat/global"

type Emoticon struct {
	global.GVA_MODEL
	User_id int    `json:"user_id"  gorm:"comment:用户id"`
	Status  int    `json:"status"  gorm:"comment:状态"`
	Src     string `json:"src" gorm:"comment:链接"`
	Name    string `json:"name" gorm:"comment:名称"`
}
