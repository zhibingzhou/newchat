package model

import "newchat/global"

type Group_member struct {
	global.GVA_MODEL
	Not_disturb   int    `json:"not_disturb"  gorm:"comment:是否免打扰"`
	IsGroupLeader int    `json:"is_group_leader"  gorm:"comment:是否是群主"`
	Group_id      int    `json:"group_id"  gorm:"comment:群id"`
	User_id       int    `json:"user_id"  gorm:"comment:用户id"`
	Is_top        int    `json:"is_top"  gorm:"comment:是否置顶"`
	Talk_remove   int    `json:"talk_remove"  gorm:"comment:是否移除"`
	Visit_card    string `json:"visit_card"  gorm:"comment:群备注名"`
}
