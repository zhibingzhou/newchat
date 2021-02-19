package model

import (
	"time"

	"gorm.io/gorm"
)

type SysUser struct {
	ID        int `json:"uid" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Mobile    string         `json:"mobile" gorm:"comment:用户手机号"`
	Password  string         `json:"password"  gorm:"comment:用户登录密码"`
	NickName  string         `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar    string         `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Motto     string         `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Email     string         `json:"email" gorm:"default:888;comment:用户邮箱"`
}
