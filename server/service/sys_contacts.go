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

	err = global.GVA_DB.Debug().Raw("SELECT count(*) as unread_num FROM apply_records WHERE  friend_id = ?  and status = 0", id).Scan(&count).Error

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

	err = global.GVA_DB.Table("apply_records").Where(" friend_id = ? ", id).Limit(pageSize).Offset(page - 1).Scan(&rep).Count(&total).Error
	page_total = len(rep)

	return err, rep, total, page_total
}
