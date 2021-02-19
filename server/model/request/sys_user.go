package request

import uuid "github.com/satori/go.uuid"

// User register structure
type Register struct {
	Mobile   string `json:"mobile"`
	Password string `json:"passWord"`
	NickName string `json:"nickname" gorm:"default:''"`
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
	Mobile    string `json:"mobile"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// Modify  user's auth structure
type SetUserAuth struct {
	UUID        uuid.UUID `json:"uuid"`
	AuthorityId string    `json:"authorityId"`
}

