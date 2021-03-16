package service

import (
	"errors"
	"newchat/global"
	"newchat/model"
	"newchat/model/request"
	"newchat/model/response"
	"sort"
	"strconv"

	"github.com/jinzhu/gorm"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindTalk_List(id int) (err error, messages []response.ResponseMessages_list) {

	var groupmessage []response.ResponseMessages_list
	err = global.GVA_DB.Raw("SELECT messages_list.id , messages_list.`updated_at`,messages_list.`type`,messages_list.`online`,messages_list.`msg_text`,contacts.`is_top`,contacts.`not_disturb`,contacts.`friend_remark` AS remark_name ,sys_user.`avatar`,sys_user.`nickname` AS name ,sys_user.`id` AS friend_id FROM messages_list,sys_user,contacts WHERE messages_list.user_id = ? AND messages_list.`group_id` = 0 AND contacts.`user_id` = messages_list.user_id AND sys_user.`id` = contacts.`friend_id`  AND contacts.`user_id` = ? and status = 1 ORDER BY messages_list.`created_at` DESC ", id, id).Scan(&messages).Error

	err = global.GVA_DB.Raw("SELECT messages_list.id , messages_list.`updated_at`,messages_list.`type`,messages_list.`online`,messages_list.`msg_text`,group_member.`is_top`,group_member.`not_disturb`,group_list.`avatar`,group_list.`group_name` AS name,messages_list.group_id FROM messages_list,group_list,group_member WHERE   group_list.id = messages_list.`group_id` AND group_member.`user_id` = ? and messages_list.`friend_id` = 0 and status = 1 ORDER BY messages_list.`created_at` DESC ", id).Scan(&groupmessage).Error

	messages = append(messages, groupmessage...)

	for key, _ := range messages {
		messages[key].Unread_num, _ = model.Redis_GetMsgNoRead(messages[key].Type, id, messages[key].Friend_id, messages[key].Group_id)
		messages[key].Online = model.RedisGetUserOnline(strconv.Itoa(messages[key].Friend_id))
	}
	messages = MessageSortbyUpdate(messages)
	return err, messages
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func TalkRecords(source, receive_id string, uid int) (err error, records response.ResponseTalkRecords) {

	switch source {
	case "1":
		err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE (receive_id = ? OR receive_id = ?)  AND user_id = sys_user.`id`  AND talk_list.`source` = 1 ORDER BY created_at DESC ", uid, receive_id).Scan(&records.Rows).Error
		break
	case "2":
		err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE receive_id = ? AND user_id = sys_user.`id` AND talk_list.`source` = 2 ORDER BY created_at DESC  ", receive_id).Scan(&records.Rows).Error
		break
	}
	if err != nil {
		return err, records
	}
	records.Limit = 30
	if len(records.Rows) < 1 {
		return err, records
	}
	for key, _ := range records.Rows {
		if records.Rows[key].Msg_type == 2 {
			err, value := model.Redis_GetFileById(records.Rows[key].Id)
			if err != nil {
				return err, records
			}
			records.Rows[key].File = value
		}
	}
	records.Record_id = records.Rows[len(records.Rows)-1].Id
	return err, records
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 设置免打扰
//@param: id int
//@return: err error, user *model.SysUser

func NotDisturb(id int, rep request.RequestNotDistraub) (err error) {

	switch rep.Type {
	case 1:
		err = global.GVA_DB.Table("contacts").Where("friend_id =  ? and user_id = ?", rep.Receive_id, id).Updates(map[string]interface{}{
			"not_disturb": rep.Not_disturb,
		}).Error

		if err != nil {
			return err
		}
		break
	case 2:
		err = global.GVA_DB.Table("group_member").Where("group_id =  ? and user_id = ?", rep.Receive_id, id).Updates(map[string]interface{}{
			"not_disturb": rep.Not_disturb,
		}).Error

		if err != nil {
			return err
		}

		break
	default:
		return errors.New("无此类型")
	}

	return err

}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 设置免打扰
//@param: id int
//@return: err error, user *model.SysUser

func TalkCreate(id int, rep request.RequestCreate) (err error, mes response.ResponseMessages_list) {

	var con model.Contacts
	var talk model.Talk_list
	switch rep.Type {
	case 1:
		err = global.GVA_DB.Table("contacts").Where("friend_id =  ? and user_id = ?", rep.Receive_id, id).First(&con).Error
		if con.ID < 1 {
			return errors.New("不是好友关系，无法聊天！！"), mes
		}
		result := global.GVA_DB.Table("messages_list").Where("friend_id =  ?", rep.Receive_id).First(&mes)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				err = global.GVA_DB.Raw("SELECT * FROM talk_list WHERE source = ? AND ((user_id = ? AND receive_id = ?) OR (user_id = ? AND receive_id = ?)) ORDER BY created_at DESC LIMIT 1 ", rep.Type, id, rep.Receive_id, rep.Receive_id, id).Scan(&talk).Error
				if err != nil {
					return err, mes
				}

				friend_id, _ := strconv.Atoi(rep.Receive_id)
				_, user := FindUserById(friend_id)

				mes = response.ResponseMessages_list{
					Type:        rep.Type,
					Friend_id:   friend_id,
					Group_id:    0,
					Not_disturb: con.Not_disturb,
					Is_top:      con.Is_top,
					Remark_name: con.Friend_remark,
					Online:      model.RedisGetUserOnline(rep.Receive_id),
					Name:        user.Nickname,
					Avatar:      user.Avatar,
					Msg_text:    talk.Content,
				}

				meg := model.Messages_list{
					Type:      rep.Type,
					Friend_id: friend_id,
					Group_id:  0,
					User_id:   id,
					Online:    model.RedisGetUserOnline(rep.Receive_id),
					Msg_text:  talk.Content,
				}

				err = global.GVA_DB.Table("messages_list").Create(&meg).Error
			}
			return err, mes
		}

		break
	case 2:
		err = global.GVA_DB.Table("group_member").Where("group_id =  ? and user_id = ?", rep.Receive_id, id).First(&con).Error
		if con.ID < 1 {
			return errors.New("不在该群里面，无法聊天！！"), mes
		}
		result := global.GVA_DB.Table("messages_list").Where("group_id =  ?", rep.Receive_id).First(&mes)
		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				err = global.GVA_DB.Raw("SELECT * FROM talk_list WHERE source = ? AND receive_id = ? ORDER BY created_at DESC LIMIT 1 ", rep.Type, rep.Receive_id).Scan(&talk).Error
				if err != nil {
					return err, mes
				}

				group_id, _ := strconv.Atoi(rep.Receive_id)
				_, group := FindGroupById(group_id)

				mes = response.ResponseMessages_list{
					Type:        rep.Type,
					Friend_id:   0,
					Group_id:    group_id,
					Not_disturb: con.Not_disturb,
					Is_top:      con.Is_top,
					Remark_name: con.Friend_remark,
					Online:      model.RedisGetUserOnline(rep.Receive_id),
					Name:        group.Group_name,
					Avatar:      group.Avatar,
					Msg_text:    talk.Content,
				}

				meg := model.Messages_list{
					Type:      rep.Type,
					Friend_id: 0,
					Group_id:  group_id,
					User_id:   id,
					Online:    model.RedisGetUserOnline(rep.Receive_id),
					Msg_text:  talk.Content,
				}

				err = global.GVA_DB.Table("messages_list").Create(&meg).Error
			}
			return err, mes
		}
		break
	default:
		return errors.New("无此类型"), mes
	}

	return err, mes

}

func UpdateUnreadNum(uid, msg_type, receive_id int) (err error) {

	if msg_type == 1 {
		err = model.Redis_SetMsgNoRead(0, msg_type, uid, receive_id, 0)
	} else {
		err = model.Redis_SetMsgNoRead(0, msg_type, uid, 0, receive_id)
	}

	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func ChatRecords(source, msg_type, receive_id string, uid int) (err error, records response.ResponseTalkRecords) {

	switch source {
	case "1":
		if msg_type != "0" {
			err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE (receive_id = ? OR receive_id = ?)  AND user_id = sys_user.`id`  AND talk_list.`source` = 1 and msg_type = ? ORDER BY created_at DESC ", uid, receive_id, msg_type).Scan(&records.Rows).Error
		} else {
			err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE (receive_id = ? OR receive_id = ?)  AND user_id = sys_user.`id`  AND talk_list.`source` = 1  ORDER BY created_at DESC ", uid, receive_id).Scan(&records.Rows).Error
		}

		break
	case "2":
		if msg_type != "0" {
			err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE receive_id = ? AND user_id = sys_user.`id` AND talk_list.`source` = 2 and msg_type = ?  ORDER BY created_at DESC  ", receive_id, msg_type).Scan(&records.Rows).Error
		} else {
			err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE receive_id = ? AND user_id = sys_user.`id` AND talk_list.`source` = 2 ORDER BY created_at DESC  ", receive_id).Scan(&records.Rows).Error
		}

		break
	}
	if err != nil {
		return err, records
	}
	if len(records.Rows) < 1 {
		return err, records
	}
	records.Limit = 30
	records.Record_id = records.Rows[len(records.Rows)-1].Id
	return err, records
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func SendImage(file model.File) (err error) {
	err = global.GVA_DB.Table("file").Create(&file).Error
	return err
}

func MessageSortbyUpdate(rep []response.ResponseMessages_list) (result []response.ResponseMessages_list) {

	sort.Sort(MessagesSlice(rep))

	//sort.Sort(sort.Reverse(PersonSlice(people))) // 按照 Age 的升序排序
	return rep
}

func UpdateMessageList() {

}

// 按照 Person.Age 从大到小排序
type MessagesSlice []response.ResponseMessages_list

func (a MessagesSlice) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a MessagesSlice) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a MessagesSlice) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].UpdatedAt.Before(a[i].UpdatedAt)
}
