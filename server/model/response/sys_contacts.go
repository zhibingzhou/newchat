package response

import "newchat/model"

type Apply_unread_num struct {
	Unread_num int `json:"unread_num"`
}

type ResponseApply struct {
	Rows       []model.Apply_records `json:"rows"`
	Page       int                   `json:"page"`
	Page_size  int                   `json:"page_size"`
	Page_total int                   `json:"page_total"`
	Total      int                   `json:"total"`
}

// 如果含有time.Time 请自行import time包
type Response_Contacts_list struct {
	ID            int    `json:"id" gorm:"primarykey"`
	Nickname      string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar        string `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Motto         string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Gender        int    `json:"gender" gorm:"default:-1;comment:性别"`
	Friend_remark string `json:"friend_remark" gorm:"default:'';comment:好友备注"`
	Online        bool   `json:"online" gorm:"default:false;comment:是否在线"`
}
