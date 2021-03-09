package service

import (
	"newchat/global"
	"newchat/model"
	"newchat/model/response"
	"strconv"
)

func CreatTalk(send_id int, received_id, source, msg_type, msg string) (err error, rep model.Talk_list) {

	ireceived_id, _ := strconv.Atoi(received_id)
	imsg_type, _ := strconv.Atoi(msg_type)
	isource, _ := strconv.Atoi(source)
	rep = model.Talk_list{
		Msg_type:   imsg_type,
		Receive_id: ireceived_id,
		Source:     isource,
		User_id:    send_id,
		Content:    msg,
	}
	err = global.GVA_DB.Table("talk_list").Create(&rep).Error
	return err, rep
}

func GetTalk_listById(id int) (err error, rep response.ResponseTalkRecord) {

	err = global.GVA_DB.Debug().Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE  user_id = sys_user.`id`  AND talk_list.`id` = ? ORDER BY created_at DESC ", id).Scan(&rep).Error
	if err != nil {
		return err, rep
	}
	if rep.Msg_type == 2 {
		err, value := model.Redis_GetFileById(rep.Id)
		if err != nil {
			return err, rep
		}
		rep.File = value
	}
	return err, rep

}
