// 自动生成模板SysOperationRecord
package model

import "newchat/global"

// 如果含有time.Time 请自行import time包
type Messages_list struct {
	global.GVA_MODEL
	Type        int    `json:"type"  gorm:"comment:1 私聊 ，2 群聊"`
	Friend_id   int    `json:"friend_id" gorm:"comment:好友id"`
	Group_id    int    `json:"group_id"  gorm:"comment:群ID"`
	Not_disturb int    `json:"not_disturb" gorm:"comment:是否开启免打扰"`
	Is_top      int    `json:"is_top"  gorm:"comment:是否置顶"`
	Unread_num  int    `json:"unread_num"  gorm:"comment:备用"`
	Name        string `json:"name"  gorm:"comment:用户名，或者群名"`
	Remark_name string `json:"remark_name"  gorm:"comment:好友备注名"`
	Msg_text    string `json:"msg_text"  gorm:"comment:消息内容"`
	Avatar      string `json:"avatar"  gorm:"comment:群或者用户头像"`
	Online      bool   `json:"online"  gorm:"comment:是否在线"`
}
