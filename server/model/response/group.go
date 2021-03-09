package response

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type ResponseGroup_list struct {
	ID            int    `json:"id" gorm:"primarykey"`
	IsGroupLeader bool   `json:"isGroupLeader" gorm:"default:系统用户;comment:是否管理员" `
	Avatar        string `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Group_profile string `json:"group_profile" gorm:"default:这个人很懒什么都没留下;comment:群说明"`
	Not_disturb   int    `json:"not_disturb" gorm:"default:-1;comment:是否免打扰"`
	Group_name    string `json:"group_name" gorm:"default:'';comment:群名称"`
}

// 如果含有time.Time 请自行import time包
type ResponseGroupDetail struct {
	Group_id         int                 `json:"group_id" gorm:"primarykey"`
	CreatedAt        time.Time           `json:"created_at" gorm:"comment:创建时间"`
	Avatar           string              `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Group_name       string              `json:"group_name" gorm:"default:'';comment:群名称"`
	Group_profile    string              `json:"group_profile" gorm:"default:这个人很懒什么都没留下;comment:群说明"`
	Not_disturb      int                 `json:"not_disturb" gorm:"default:-1;comment:是否免打扰"`
	Is_manager       int                 `json:"is_manager" gorm:"default:-1;comment:是否管理员"`
	Manager_nickname string              `json:"manager_nickname" gorm:"default:-1;comment:是否免打扰"`
	Visit_card       string              `json:"visit_card" gorm:"default:-1;comment:群名称"`
	Notice           ResponseGroupNotice `json:"notice"`
}

type ResponseGroupMember struct {
	ID         int    `json:"id" gorm:"primarykey"`
	User_id    int    `json:"user_id"`
	Is_manager int    `json:"is_manager" gorm:"default:-1;comment:是否管理员"`
	Visit_card string `json:"visit_card" gorm:"default:-1;comment:群名称"`
	Nickname   string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar     string `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Motto      string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Gender     int    `json:"gender" gorm:"default:-1;comment:性别"`
}

type ResponseGroupNotice struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ResponseGroupNotices struct {
	ID        int       `json:"id" gorm:"primarykey"`
	User_id   int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
	Avatar    string    `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Nickname  string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Is_show   bool    `json:"isShow" gorm:"default:系统用户;comment:是否显示" `
}
