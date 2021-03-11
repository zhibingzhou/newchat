package service

import (
	"newchat/global"
	"newchat/model"
	"newchat/model/response"
	"strconv"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func Apply_unread_num(id int) (err error, count response.Apply_unread_num) {

	err = global.GVA_DB.Raw("SELECT count(*) as unread_num FROM apply_records WHERE  friend_id = ?  and status = 0", id).Scan(&count).Error

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

	err = global.GVA_DB.Raw("SELECT apply_records.id,status,user_id,remarks,friend_id,apply_records.`created_at`,sys_user.`nickname`,sys_user.`mobile`,sys_user.`avatar` FROM apply_records , sys_user WHERE apply_records.`friend_id` = 12 AND apply_records.`user_id` = sys_user.`id` ").Limit(pageSize).Offset(page - 1).Scan(&rep).Error
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

	f_id := []model.Friend_id{}
	f_ids := []int{}

	err = global.GVA_DB.Raw("SELECT friend_id from contacts where user_id = ? ", id).Scan(&f_id).Error
	for _, value := range f_id {
		f_ids = append(f_ids, value.Friend_id)
	}

	err = global.GVA_DB.Raw("SELECT sys_user.id AS id , sys_user.nickname AS nickname ,sys_user.gender as gender, sys_user.motto as motto, sys_user.avatar as avatar , contacts.`friend_remark` as friend_remark FROM  sys_user,contacts WHERE sys_user.`id` = contacts.`friend_id` AND contacts.`friend_id` in (?)", f_ids).Scan(&rep).Error

	for key, _ := range rep {
		ids := strconv.Itoa(rep[key].ID)
		rep[key].Online = model.RedisGetUserOnline(ids)
	}

	return err, rep
}
