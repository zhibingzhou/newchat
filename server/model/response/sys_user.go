package response

import (
	"newchat/model"
)

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}

type Setting struct {
	ID       int    `json:"uid" gorm:"primarykey"`
	NickName string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar   string `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Motto    string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Gender   int    `json:"gender" gorm:"default:-1;comment:性别"`
}
