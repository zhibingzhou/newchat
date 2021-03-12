package servers

import "time"

type RequestEvenTalk struct {
	Send_user    int    `json:"send_user"`
	Receive_user string `json:"receive_user"`
	Text_message string `json:"text_message"`
	Msg_type     string `json:"msg_type"`
	Source_type  string `json:"source_type"`
}

type ResponseEvenTalk struct {
	Source_type  int                 `json:"source_type"` //群或者组
	Receive_user int                 `json:"receive_user"`
	Send_user    int                 `json:"send_user"`
	Data         MessageInformaction `json:"data"`
}

type MessageInformaction struct {
	Id         int        `json:"id"`
	CreatedAt  time.Time  `json:"created_at" gorm:"comment:创建时间"`
	Is_revoke  int        `json:"is_revoke" gorm:"comment:是否是好友请求"`
	Msg_type   int        `json:"msg_type"  gorm:"comment:消息类型1 文字 5 代码 2 图片"`
	Receive_id int        `json:"receive_id" gorm:"comment:接收信息id"`
	Source     int        `json:"source"  gorm:"comment: //2群 还是 1私聊"`
	User_id    int        `json:"user_id"  gorm:"comment:用户id"`
	Content    string     `json:"content"  gorm:"comment:聊天内容"`
	Nickname   string     `json:"nickname"  gorm:"comment:用户名"`
	Avatar     string     `json:"avatar"  gorm:"comment:用户头像"`
	File       File       `json:"file"  gorm:"comment:文件名称"`
	Code_block Code_block `json:"code_block"  gorm:"comment:代码片段"`
	Forward    Forward    `json:"forward"  gorm:"comment:代码片段"`
	Invite     Invite     `json:"invite"  gorm:"comment:代码片段"`
}

type MessageFile struct {
	File_type int `json:"file_type"`
}

type File struct {
}

type Code_block struct {
}

type Forward struct {
}

type Invite struct {
}
