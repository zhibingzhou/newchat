package response

import (
	"newchat/model"
	"time"
)

type ResponseMessages_list struct {
	Id          int       `json:"id" gorm:"primarykey"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"comment:更新时间"`
	Type        int       `json:"type"  gorm:"comment:1 私聊 ，2 群聊"`
	Friend_id   int       `json:"friend_id" gorm:"comment:好友id"`
	Group_id    int       `json:"group_id"  gorm:"comment:群ID"`
	Not_disturb int       `json:"not_disturb" gorm:"comment:是否开启免打扰"`
	Is_top      int       `json:"is_top"  gorm:"comment:是否置顶"`
	Unread_num  int       `json:"unread_num"  gorm:"comment:备用"`
	Name        string    `json:"name"  gorm:"comment:用户名，或者群名"`
	Remark_name string    `json:"remark_name"  gorm:"comment:好友备注名"`
	Msg_text    string    `json:"msg_text"  gorm:"comment:消息内容"`
	Avatar      string    `json:"avatar"  gorm:"comment:群或者用户头像"`
	Online      bool      `json:"online"  gorm:"comment:是否在线"`
}

type ResponseCreateTalk struct {
	TalkItem ResponseMessages_list `json:"talkItem"`
}

type ResponseTalkRecord struct {
	Id         int                `json:"id"`
	CreatedAt  time.Time          `json:"created_at" gorm:"comment:创建时间"`
	Is_revoke  int                `json:"is_revoke" gorm:"comment:是否是好友请求"`
	Msg_type   int                `json:"msg_type"  gorm:"comment:消息类型1 文字 5 代码 2 图片"`
	Receive_id int                `json:"receive_id" gorm:"comment:接收信息id"`
	Source     int                `json:"source"  gorm:"comment: //2群 还是 1私聊"`
	User_id    int                `json:"user_id"  gorm:"comment:用户id"`
	Content    string             `json:"content"  gorm:"comment:聊天内容"`
	Nickname   string             `json:"nickname"  gorm:"comment:用户名"`
	Avatar     string             `json:"avatar"  gorm:"comment:用户头像"`
	File       model.File       `json:"file"  gorm:"comment:文件名称"`
	Code_block model.Code_block `json:"code_block"  gorm:"comment:代码片段"`
	Forward    model.Forward    `json:"forward"  gorm:"comment:代码片段"`
	Invite     model.Invite     `json:"invite"  gorm:"comment:代码片段"`
}

type ResponseTalkRecords struct {
	Rows      []ResponseTalkRecord `json:"rows"`
	Record_id int                  `json:"record_id"`
	Limit     int                  `json:"limit"`
}
