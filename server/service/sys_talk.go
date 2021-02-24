package service

import (
	"newchat/global"
	"newchat/model"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindTalk_List(id int) (err error, messages []model.Messages_list) {

	var g_id []int
	var f_id []int

	if err = global.GVA_DB.Table("contacts").Select([]string{"friend_id"}).Where("user_id=?", id).Scan(&g_id).Error; err != nil {
		return err, messages
	}

	if err = global.GVA_DB.Table("group_member").Select([]string{"group_id"}).Where("user_id=?", id).Scan(&f_id).Error; err != nil {
		return err, messages
	}

	err = global.GVA_DB.Debug().Raw("SELECT * FROM messages_list WHERE group_id IN (?) OR friend_id IN (?) ORDER BY updated_at ", f_id, g_id).Scan(&messages).Error

	return err, messages
}
