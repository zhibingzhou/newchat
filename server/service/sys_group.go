package service

import (
	"errors"
	"newchat/global"
	"newchat/model"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/utils"
	"strconv"
	"strings"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func Group_List(id int) (err error, rep []response.ResponseGroup_list) {

	err = global.GVA_DB.Raw("SELECT group_list.id AS id , is_group_leader, not_disturb , group_profile ,group_name,avatar FROM group_member,group_list WHERE group_list.`id` = group_member.`group_id` AND group_member.`user_id` = ?", id).Scan(&rep).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func GroupDetail(group_id string, uid int) (err error, rep response.ResponseGroupDetail) {

	var notice response.ResponseGroupNotice
	err = global.GVA_DB.Raw("SELECT group_list.`id` as group_id,group_name,group_profile,group_list.avatar, group_list.created_at,group_member.`is_group_leader` AS is_manager , group_member.`visit_card` ,group_member.`not_disturb`,sys_user.`nickname` AS manager_nickname FROM group_list,group_member,sys_user WHERE group_list.`manager_id` = sys_user.`id` AND group_list.`id` = ? AND user_id = ? ", group_id, uid).Scan(&rep).Error
	if err != nil {
		return err, rep
	}
	err = global.GVA_DB.Table("group_notice").Select([]string{"title", "content"}).Where("group_id = ?", group_id).Order("created_at desc").Limit(1).Scan(&notice).Error
	rep.Notice = notice
	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func GroupMembers(group_id string) (err error, rep []response.ResponseGroupMember) {
	err = global.GVA_DB.Raw("SELECT group_member.`id`, group_member.`is_group_leader` AS is_manager , group_member.`visit_card` , group_member.`user_id`,sys_user.`avatar`,sys_user.`nickname`,sys_user.`gender`,sys_user.`motto` FROM group_member,sys_user WHERE group_member.`group_id` = ? AND sys_user.`id` = group_member.`user_id`", group_id).Scan(&rep).Error
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

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 邀请好友
//@param: id int
//@return: err error, user *model.SysUser

func InviteFriends(uid int, group_id string) (err error, rep []response.ResponseContactsFriend) {

	f_id := []model.Friend_id{}
	f_ids := []int{}

	if group_id == "0" {
		err = global.GVA_DB.Debug().Raw("SELECT sys_user.id AS id , sys_user.nickname AS nickname ,sys_user.gender AS gender, sys_user.motto AS motto, sys_user.avatar AS avatar , contacts.`friend_remark` AS friend_remark FROM  contacts LEFT JOIN  sys_user ON sys_user.`id` = contacts.`friend_id` WHERE contacts.`user_id` = ?", uid).Scan(&rep).Error
	} else {
		err = global.GVA_DB.Debug().Raw(" SELECT friend_id FROM contacts WHERE contacts.`friend_id` NOT IN ((SELECT user_id FROM group_member WHERE group_id = ?)) AND contacts.`user_id` = ?", group_id, uid).Scan(&f_id).Error
		for _, value := range f_id {
			f_ids = append(f_ids, value.Friend_id)
		}
		err = global.GVA_DB.Raw("SELECT sys_user.id AS id , sys_user.nickname AS nickname ,sys_user.gender as gender, sys_user.motto as motto, sys_user.avatar as avatar , contacts.`friend_remark` as friend_remark FROM  sys_user left join contacts on sys_user.`id` = contacts.`friend_id` where contacts.`friend_id` in (?) and contacts.user_id = ?", f_ids, uid).Scan(&rep).Error
	}

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 邀请好友
//@param: id int
//@return: err error, user *model.SysUser

func GroupCreate(uid int, rep request.RequestGroupCreate) (err error) {
	if rep.Uids == "" {
		return errors.New("群组成员不能为空！")
	}
	var groups model.Group_list
	groups.Group_name = rep.Group_name
	groups.Group_profile = rep.Group_profile
	groups.Avatar = rep.Group_avatar
	groups.Manager_id = uid
	uids := strings.Split(rep.Uids, ",")
	tx := global.GVA_DB.Begin()
	groups.ID, _ = strconv.Atoi(utils.Random("number", 4))

	err = tx.Table("group_list").Create(&groups).Error
	if err != nil {
		tx.Rollback()
	}

	group := model.Group_member{
		Group_id:      groups.ID,
		User_id:       uid,
		IsGroupLeader: 1,
	}

	err = tx.Table("group_member").Create(&group).Error
	if err != nil {
		tx.Rollback()
	}

	for _, value := range uids {
		userid, _ := strconv.Atoi(value)
		group := model.Group_member{
			Group_id:      groups.ID,
			User_id:       userid,
			IsGroupLeader: 0,
		}
		tx.Table("group_member").Create(&group)
		if err != nil {
			tx.Rollback()
		}

	}

	tx.Commit()
	return err
}

func GroupInvite(uid int, rep request.RequestGroupInvite) (err error) {
	if rep.Uids == "" {
		return errors.New("群组成员不能为空！")
	}
	var groups model.Group_list

	uids := strings.Split(rep.Uids, ",")
	tx := global.GVA_DB.Begin()

	err = tx.Table("group_list").Where("id = ?", rep.Group_id).Scan(&groups).Error
	if err != nil {
		tx.Rollback()
	}

	for _, value := range uids {
		userid, _ := strconv.Atoi(value)
		group := model.Group_member{
			Group_id:      groups.ID,
			User_id:       userid,
			IsGroupLeader: 0,
		}
		tx.Table("group_member").Create(&group)
		if err != nil {
			tx.Rollback()
		}

	}

	tx.Commit()
	return err
}

func GroupSecede(uid, group_secede int) (err error) {
	if group_secede == 0 {
		return errors.New("群id 不能为空！！")
	}
	err = global.GVA_DB.Unscoped().Delete(&model.Group_member{}, "group_id = ? and user_id = ?", group_secede, uid).Error
	return err
}

func SetGroupCard(uid, group_id int, card string) (err error) {
	if group_id == 0 || card == "" {
		return errors.New("信息不能为空！！")
	}
	err = global.GVA_DB.Table("group_member").Where("group_id = ? and user_id = ?", group_id, uid).Update(map[string]interface{}{"visit_card": card}).Error
	return err
}

func RemoveMembers(uid, group_id int, group_id_iist []int) (err error) {
	if group_id == 0 || len(group_id_iist) <= 0 {
		return errors.New("信息不能为空！！")
	}
	var g model.Group_list
	err = global.GVA_DB.Table("group_list").Where("id = ?", group_id).Scan(&g).Error
	if g.Manager_id != uid {
		return errors.New("不是管理员不能移除！！")
	}
	err = global.GVA_DB.Unscoped().Delete(&model.Group_member{}, "group_id = ? and user_id in (?)", group_id, group_id_iist).Error
	return err
}
