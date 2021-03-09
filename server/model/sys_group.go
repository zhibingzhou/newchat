package model

import "newchat/global"

type Group_id struct {
	Group_id    int `json:"group_id" gorm:"comment:用户id"`
}

type Group_notice struct {
	global.GVA_MODEL
	User_id  int    `json:"user_id"  gorm:"comment:用户id"`
	Group_id int    `json:"group_id"  gorm:"comment:群id"`
	Title    string `json:"title" gorm:"comment:公告标题"`
	Content  string `json:"content" gorm:"comment:公告内容"`
}

type Group_list struct {
	global.GVA_MODEL
	Manager_id    int    `json:"manager_id"  gorm:"comment:群主id"`
	Group_profile string `json:"group_profile"  gorm:"comment:群简介"`
	Avatar        string `json:"avatar"  gorm:"comment:群头像"`
	Group_name    string `json:"group_name"  gorm:"comment:群名称"`
}
