package model

import (
	"time"
)

type SysUser struct {
	ID        int `json:"uid" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `gorm:"default:NULL"`
	Mobile    string    `json:"mobile" gorm:"comment:用户手机号"`
	Password  string    `json:"password"  gorm:"comment:用户登录密码"`
	Nickname  string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar    string    `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Motto     string    `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Email     string    `json:"email" gorm:"default:888;comment:用户邮箱"`
	Gender    int       `json:"gender" gorm:"default:-1;comment:性别"`
}

type UserId struct {
	User_id int `json:"user_id"`
}
