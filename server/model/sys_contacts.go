// 自动生成模板SysOperationRecord
package model

import (
	"newchat/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type Apply_records struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	Status    int       `json:"status"  gorm:"comment:0 审核中 ， 1 同意"`
	User_id   int       `json:"user_id" gorm:"comment:用户id"`
	Friend_id int       `json:"friend_id"  gorm:"comment:好友id"`
	Remarks   string    `json:"remarks"  gorm:"comment:用户名，或者群名"`
	Nickname  string    `json:"nickname"  gorm:"comment:用户昵称或者群组名称"`
	Mobile    string    `json:"mobile"  gorm:"comment:用户手机号"`
	Avatar    string    `json:"avatar"  gorm:"comment:群或者用户头像"`
}

type Friend_id struct {
	Friend_id   int `json:"friend_id"  gorm:"comment:好友id"`
}

type Contacts struct {
	global.GVA_MODEL
	Friend_status int    `json:"friend_status"  gorm:"comment:0 审核中 ， 1 同意"`
	User_id       int    `json:"user_id" gorm:"comment:用户id"`
	Friend_id     int    `json:"friend_id"  gorm:"comment:好友id"`
	Friend_remark string `json:"friend_remark"  gorm:"comment:用户名，或者群名"`
	Not_disturb   int    `json:"not_disturb"  gorm:"comment:是否开启免打扰"`
	Is_top        int    `json:"is_top"  gorm:"comment:是否置顶"`
}

func RedisGetUserOnline(ids string) bool {
	result := false
	Online, _ := global.GVA_REDIS.HMGet(global.UserStatus, ids).Result()
	if Online[0] == "1" {
		result = true
	}
	return result
}
