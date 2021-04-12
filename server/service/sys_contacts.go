package service

import (
	"newchat/global"
	"newchat/model"
	"newchat/model/response"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func Apply_unread_num(id int) (err error, count response.Apply_unread_num) {

	err = global.GVA_DB.Raw("SELECT count(*) as unread_num FROM apply_records WHERE  friend_id = ?  and status = 0 and `apply_records`.`deleted_at` IS NULL", id).Scan(&count).Error

	return err, count
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func Apply_records(id, page, pageSize int) (err error, rep []model.Apply_records, total, page_total int) {

	// var num model.Totaln

	// err = global.GVA_DB.Table("apply_records").Select("count(*) as num ").Where(" friend_id = ? ", id).Scan(&num).Error
	// if err != nil {
	// 	return err, rep, total, page_total
	// }

	err = global.GVA_DB.Debug().Raw("SELECT apply_records.id,status,user_id,remarks,friend_id,apply_records.`created_at`,sys_user.`nickname`,sys_user.`mobile`,sys_user.`avatar` FROM apply_records , sys_user WHERE apply_records.`friend_id` = ? AND apply_records.`user_id` = sys_user.`id` and `apply_records`.`deleted_at` IS NULL ", id).Order("created_at desc").Limit(pageSize).Offset(page - 1).Scan(&rep).Error
	page_total = len(rep)
	total = page_total
	return err, rep, total, page_total
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func Contacts_List(id int) (err error, rep []response.Response_Contacts_list) {

	// f_id := []model.Friend_id{}
	// f_ids := []int{}

	// err = global.GVA_DB.Raw("SELECT friend_id from contacts where user_id = ? ", id).Scan(&f_id).Error
	// for _, value := range f_id {
	// 	f_ids = append(f_ids, value.Friend_id)
	// }

	err = global.GVA_DB.Raw("SELECT sys_user.id AS id , sys_user.nickname AS nickname ,sys_user.gender as gender, sys_user.motto as motto, sys_user.avatar as avatar , contacts.`friend_remark` as friend_remark FROM  sys_user,contacts WHERE sys_user.`id` = contacts.`friend_id` and contacts.user_id = ?", id).Scan(&rep).Error

	for key, _ := range rep {
		rep[key].Online = model.RedisGetUserOnline(rep[key].ID)
	}

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func Contacts_Search(mobile string) (err error, rep response.ResponseSearchMoblie) {

	err = global.GVA_DB.Debug().Table("sys_user").Select([]string{"id", "mobile", "nickname", "avatar", "gender"}).Where("mobile = ?", mobile).Scan(&rep).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 发送加好友请求
//@param: id int
//@return: err error, user *model.SysUser

func Contacts_Add(uid, friend_id int, remarks string) (err error) {
	var apply model.ApplyRecords
	apply = model.ApplyRecords{
		User_id:   uid,
		Friend_id: friend_id,
		Remarks:   remarks,
	}
	err = global.GVA_DB.Table("apply_records").Create(&apply).Error
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 发送加好友请求
//@param: id int
//@return: err error, user *model.SysUser

func Contacts_Del(apply_id int) (err error) {
	err = global.GVA_DB.Unscoped().Delete(model.ApplyRecords{}, "id = ?", apply_id).Error
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 发送加好友请求
//@param: id int
//@return: err error, user *model.SysUser

func AcceptInvitation(apply_id int, remarks, nickname string) (err error) {
	var contacts_one model.Contacts
	var contacts_two model.Contacts
	var apply_records model.ApplyRecords

	err = global.GVA_DB.Table("apply_records").Where("id = ?", apply_id).Scan(&apply_records).Error

	if err != nil {
		return err
	}

	tx := global.GVA_DB.Begin()

	contacts_one = model.Contacts{
		User_id:       apply_records.User_id,
		Friend_id:     apply_records.Friend_id,
		Friend_remark: nickname,
	}

	contacts_two = model.Contacts{
		Friend_id:     apply_records.User_id,
		User_id:       apply_records.Friend_id,
		Friend_remark: remarks,
	}

	err = tx.Table("contacts").Create(&contacts_one).Error
	if err != nil {
		tx.Rollback()
	}

	err = tx.Table("contacts").Create(&contacts_two).Error
	if err != nil {
		tx.Rollback()
	}

	err = tx.Table("apply_records").Where("id = ?", apply_id).Update("status", 1).Error

	if err != nil {
		tx.Rollback()
	}

	tx.Commit()

	return err
}

//删除记录
//db.Unscoped().Delete(&order)

func Contacts_Delete(user_id, friend_id int) (err error) {

	tx := global.GVA_DB.Begin()
	err = tx.Table("contacts").Unscoped().Delete(&model.Contacts{}, "user_id = ? and friend_id = ?", user_id, friend_id).Error
	if err != nil {
		tx.Rollback()
	}
	err = tx.Table("contacts").Unscoped().Delete(&model.Contacts{}, "user_id = ? and friend_id = ?", friend_id, user_id).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	return err
}

//编辑备注
func Topping(user_id, friend_id, t_ype, is_top int) (err error) {

	if t_ype == 1 {
		err = global.GVA_DB.Table("contacts").Where("user_id = ? and friend_id = ?", user_id, friend_id).Update(map[string]interface{}{"is_top": is_top}).Error
	} else {
		err = global.GVA_DB.Table("group_member").Where("user_id = ? and group_id = ?", user_id, friend_id).Update(map[string]interface{}{"is_top": is_top}).Error
	}

	return err
}

//编辑备注
func Edit_Remark(user_id, friend_id int, remarks string) (err error) {

	err = global.GVA_DB.Table("contacts").Where("user_id = ? and friend_id = ?", user_id, friend_id).Update(map[string]interface{}{"friend_remark": remarks}).Error

	return err
}

//删除聊天列表
func TalkDelete(id int) (err error) {
	var mes model.Messages_list
	tx := global.GVA_DB.Begin()
	err = tx.Table("messages_list").Where("id = ?", id).Scan(&mes).Error
	if err != nil {
		tx.Rollback()
	}
	if mes.Type == 1 {
		err = tx.Debug().Table("messages_list").Unscoped().Delete(&model.Messages_list{}, "id = ?", id).Error
		if err != nil {
			tx.Rollback()
		}
	}
	if mes.Type == 2 {
		err = tx.Table("group_member").Where("group_id = ? and user_id = ?", mes.Group_id, id).Update(map[string]interface{}{"talk_remove": 1}).Error
		if err != nil {
			tx.Rollback()
		}
	}
	tx.Commit()
	return err
}
