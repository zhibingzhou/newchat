package service

import (
	"newchat/global"
	"newchat/model"
	"newchat/model/response"

	"gorm.io/gorm"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func UserEmoticon(id int) (err error, rep response.ResponseUserEmoticon) {

	result := global.GVA_DB.Raw("SELECT emoticon.id AS media_id , emoticon.`src` FROM emoticon WHERE user_id = ? AND STATUS = 1 ", id).Scan(&rep.Collect_emoticon)

	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return result.Error, rep
		}
	}

	err = global.GVA_DB.Raw("SELECT emoticon.id AS media_id , emoticon.`src` FROM emoticon WHERE user_id = 1 AND STATUS = 1 ").Scan(&rep.Sys_emoticon).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func SystemEmoticon() (err error, rep []response.ResponseSysEmoticon) {

	err = global.GVA_DB.Raw(" SELECT emoticon.id , emoticon.`src` AS url, emoticon.`name` ,emoticon.`status`  FROM emoticon WHERE user_id = 1 AND STATUS = 1 ").Scan(&rep).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func SetUserEmoticon(id, status int) (err error) {

	err = global.GVA_DB.Table("emoticon").Where("id = ?", id).Update("status", status).Error

	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func UploadEmoticon(id int, url string) (err error, emoticon model.Emoticon) {

	emoticon = model.Emoticon{
		User_id: id,
		Src:     url,
		Status:  1,
	}
	err = global.GVA_DB.Table("emoticon").Create(&emoticon).Error

	return err, emoticon
}
