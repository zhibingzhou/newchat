// 自动生成模板SysOperationRecord
package model

import "newchat/global"

// 如果含有time.Time 请自行import time包
type Apply_records struct {
	global.GVA_MODEL
	Status    int    `json:"status"  gorm:"comment:0 审核中 ， 1 同意"`
	User_id   int    `json:"user_id" gorm:"comment:用户id"`
	Friend_id int    `json:"friend_id"  gorm:"comment:好友id"`
	Remarks   string `json:"remarks"  gorm:"comment:用户名，或者群名"`
	Nickname  string `json:"nickname"  gorm:"comment:用户昵称或者群组名称"`
	Mobile    string `json:"mobile"  gorm:"comment:用户手机号"`
	Avatar    string `json:"avatar"  gorm:"comment:群或者用户头像"`
}
