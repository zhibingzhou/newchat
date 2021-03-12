package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"newchat/global"
	"newchat/model"
	"newchat/model/response"
	"newchat/utils"
	"strconv"

	"go.uber.org/zap"
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


func GetTalk_listById(id int) (err error, rep response.WebsocketMessage) {
	var repdata response.ResponseTalkRecord
	err = global.GVA_DB.Raw(" SELECT talk_list.`id`, source,msg_type,user_id,receive_id,is_revoke,content,sys_user.`avatar`,sys_user.`nickname`,talk_list.`created_at` FROM talk_list,sys_user WHERE  user_id = sys_user.`id`  AND talk_list.`id` = ? ORDER BY created_at DESC ", id).Scan(&repdata).Error
	if err != nil {
		return err, rep
	}
	if repdata.Msg_type == 2 {
		err, value := model.Redis_GetFileById(repdata.Id)
		if err != nil {
			return err, rep
		}
		repdata.File = value
	}
	if repdata.Source == 2 {
		var user_idlist []model.UserId
		var received_list []int

		err, user_idlist := model.Redis_GetGroupbyId(repdata.Receive_id)

		if err != nil {
			return err, rep
		}

		for key, _ := range user_idlist {
			received_list = append(received_list, user_idlist[key].User_id)
		}
		rep.Receive_list = received_list
	}

	rep.Messagedata.Data = repdata

	return err, rep

}

func SendToClient(event string, request response.WebsocketMessage) {

	switch event {
	case "event_img":
		client, _ := global.GVA_REDIS.HGet(global.UserIdSystem, fmt.Sprintf("%d", request.Messagedata.Send_user)).Result()
		onlineStatus := model.RedisGetUserOnline(fmt.Sprintf("%d", request.Messagedata.Send_user))
		if client != "" && onlineStatus {
			requestbyte, _ := json.Marshal(request)
			url := global.GVA_CONFIG.Websocket.Url + "/send_message"
			status, msg := utils.HttpPostjson(url, requestbyte, map[string]string{
				"SystemId": client,
			})
			if status != 200 {
				global.GVA_LOG.Error(msg, zap.Any("err", errors.New(msg)))
			}
		}

		break
	}
	//	HttpPostjson()
}
