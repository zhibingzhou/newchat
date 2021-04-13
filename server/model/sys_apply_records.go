package model

import "newchat/global"

type ApplyRecords struct {
	global.GVA_MODEL
	status    int    `json:"status" gorm:"comment:状态,0为审请中，1为好友"`  
	User_id   int    `json:"user_id" gorm:"comment:用户id"`
	Friend_id int    `json:"friend_id" gorm:"comment:朋友id"`
	Remarks   string `json:"remarks" gorm:"comment:备注"`
}
