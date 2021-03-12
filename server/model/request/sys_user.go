package request

import uuid "github.com/satori/go.uuid"

// User register structure
type Register struct {
	Mobile   string `json:"mobile"`
	Password string `json:"passWord"`
	Nickname string `json:"nickname" gorm:"default:''"`
	Avatar   string `json:"avatar" gorm:"default:'http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg'"`
}

// User login structure
type Login struct {
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	Captcha   string `json:"captcha"`
	CaptchaId string `json:"captchaId"`
}

// Modify password structure
type ChangePasswordStruct struct {
	Old_password string `json:"old_password"`
	New_password string `json:"new_password"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}


type RequestUserEdit struct {
	Nickname string `json:"nickname" gorm:"default:系统用户;comment:用户昵称" `
	Avatar   string `json:"avatar" gorm:"default:http://img.wxcha.com/m00/b0/2b/65252be5c6e7e8ace4458e517cb5ad08.jpg;comment:用户头像"`
	Motto    string `json:"motto" gorm:"default:这个人很懒什么都没留下;comment:修改签名"`
	Gender   int    `json:"gender" gorm:"default:-1;comment:性别"`
}
