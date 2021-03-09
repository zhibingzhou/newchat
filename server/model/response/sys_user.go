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
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar   string `json:"avatar" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Motto    string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Gender   int    `json:"gender" gorm:"default:-1;comment:性别"`
}

type ResponseSearchUser struct {
	ID              int    `json:"id" gorm:"primarykey"`
	Mobile          string `json:"mobile" gorm:"comment:用户手机号"`
	Nickname        string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar          string `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Motto           string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Gender          int    `json:"gender" gorm:"default:-1;comment:性别"`
	Friend_status   int    `json:"friend_status" gorm:"default:1;comment:好友状态"`
	Nickname_remark string `json:"nickname_remark" gorm:"default:-1;comment:好友备注"`
	// Friend_apply    int    `json:"friend_apply" gorm:"default:0;comment:好友请求"`
}

type ResponseUserDetail struct {
	Mobile   string `json:"mobile" gorm:"comment:用户手机号"`
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar   string `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Motto    string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Email    string `json:"email" gorm:"default:888;comment:用户邮箱"`
	Gender   int    `json:"gender" gorm:"default:-1;comment:性别"`
}
