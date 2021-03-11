package service

import (
	"errors"
	"newchat/global"
	"newchat/model"
	"newchat/model/request"
	"newchat/model/response"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func Group_List(id int) (err error, rep []response.ResponseGroup_list) {

	err = global.GVA_DB.Raw("SELECT group_list.id AS id , isGroupLeader, not_disturb , group_profile ,group_name,avatar FROM group_member,group_list WHERE group_list.`id` = group_member.`group_id` AND group_member.`user_id` = ?", id).Scan(&rep).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func GroupDetail(group_id string, uid int) (err error, rep response.ResponseGroupDetail) {

	var notice response.ResponseGroupNotice
	err = global.GVA_DB.Raw("SELECT group_list.`id` as group_id,group_name,group_profile,group_list.avatar, group_list.created_at,group_member.`isGroupLeader` AS is_manager , group_member.`visit_card` ,group_member.`not_disturb`,sys_user.`nickname` AS manager_nickname FROM group_list,group_member,sys_user WHERE group_list.`manager_id` = sys_user.`id` AND group_list.`id` = ? AND user_id = ? ", group_id, uid).Scan(&rep).Error
	if err != nil {
		return err, rep
	}
	err = global.GVA_DB.Table("group_notice").Select([]string{"title", "content"}).Where("group_id = ?", group_id).Order("created_at desc").Limit(1).Scan(&notice).Error
	if err != nil {
		return err, rep
	}
	rep.Notice = notice
	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func GroupMembers(group_id string) (err error, rep []response.ResponseGroupMember) {
	err = global.GVA_DB.Raw("SELECT group_member.`id`, group_member.`isGroupLeader` AS is_manager , group_member.`visit_card` , group_member.`user_id`,sys_user.`avatar`,sys_user.`nickname`,sys_user.`gender`,sys_user.`motto` FROM group_member,sys_user WHERE group_member.`group_id` = ? AND sys_user.`id` = group_member.`user_id`", group_id).Scan(&rep).Error
	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func GroupNotices(group_id string) (err error, rep []response.ResponseGroupNotices) {
	err = global.GVA_DB.Raw("SELECT group_notice.`id`,group_notice.`user_id`,title,content,group_notice.`created_at`,group_notice.`updated_at`,avatar,nickname FROM group_notice,sys_user WHERE group_notice.`group_id` = ? AND group_notice.`user_id` = sys_user.`id` order by created_at desc", group_id).Scan(&rep).Error
	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 群信息编辑
//@param: id int
//@return: err error, user *model.SysUser

func GroupEdit(uid int, groups request.RequestGroupEdit) (err error) {

	var grouplist model.Group_list
	db := global.GVA_DB.Where("id = ? and manager_id = ?", groups.Group_id, uid).First(&grouplist)

	if grouplist.ID <= 1 {
		return errors.New("不是群主无法修改！！")
	}

	groupmap := map[string]interface{}{
		"avatar":        groups.Avatar,
		"group_name":    groups.Group_name,
		"group_profile": groups.Group_profile,
	}

	err = db.Updates(groupmap).Error
	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindGroupById(id int) (err error, group *model.Group_list) {
	var u model.Group_list
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 群信息编辑
//@param: id int
//@return: err error, user *model.SysUser

func EditNotice(uid int, groups request.RequestEditGroupEdit) (err error) {

	var groupnotice model.Group_notice

	if groups.Group_id == 0 {
		return errors.New("错误的群组id")
	}

	groupnotice = model.Group_notice{
		User_id:  uid,
		Group_id: groups.Group_id,
		Title:    groups.Title,
		Content:  groups.Content,
	}

	err = global.GVA_DB.Table("group_notice").Create(&groupnotice).Error

	return err
}
