package servers

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"websocket/pkg/redis"
	"websocket/pkg/setting"
	"websocket/tools/util"
)

var MessageChannel *Channel_Pool

type Channel_Pool struct {
	Count    int
	Request  chan DRequest
	Requests chan chan DRequest
	Result   chan DResult
}

type DRequest struct {
	Message []byte
}

type DResultTalkEvent struct {
	Status       int              `json:"status"`
	Receivedlist []int            `json:"received"`
	Event        string           `json:"event"`
	DataResponse ResponseEvenTalk `json:"data"`
}

type DResult struct {
	Result interface{} `json:"result"`
}

//接收参数
type ReponseFromWeb struct {
	Code int                  `json:"code"`
	Data ReponseFromTalkEvent `json:"data"`
	Msg  string               `json:"msg"`
}

type ReponseFromTalkEvent struct {
	Messagedata  ResponseEvenTalk `json:"messagedata"`
	Receive_list []int            `json:"receive_list"`
	Event        string           `json:"event"`
}

type ReponseLoginEvents struct {
	Code int               `json:"code"`
	Data ReponseLoginEvent `json:"data"`
	Msg  string            `json:"msg"`
}

type ReponseLoginEvent struct {
	User_id      int    `json:"user_id"`
	Status       int    `json:"status"`
	Receivedlist []int  `json:"received"`
	Event        string `json:"event"`
}

type KeyBoardEvent struct {
	Send_user    int    `json:"send_user"`
	Event        string `json:"event"`
	Receive_user string `json:"receive_user"`
}

type ReponseGroupJoin struct {
	Code int       `json:"code"`
	Data GroupList `json:"data"`
	Msg  string    `json:"msg"`
}

type GroupList struct {
	Send_user    int    `json:"send_user"`
	Event        string `json:"event"`
	Receivedlist []int  `json:"received"`
}

func NewChannel_Pool(count int) *Channel_Pool {
	return &Channel_Pool{
		Count:    count,
		Request:  make(chan DRequest),
		Requests: make(chan chan DRequest),
		Result:   make(chan DResult)}
}

func (c *Channel_Pool) NewChannel_PoolGo() {

	for i := 0; i < c.Count; i++ { //工作
		w := Newwoker()
		w.Run(c.Requests, c.Result)
	}

	// go func() {
	// 	for {
	// 		select {
	// 		case woker := <-c.Request:
	// 			wokers := <-c.Requests
	// 			wokers <- woker
	// 		}
	// 	}
	// }()

	go func() {

		var DRs []chan DRequest
		var DR []DRequest

		for {

			var goRs chan DRequest
			var goDR DRequest

			if len(DRs) > 0 && len(DR) > 0 {
				goRs = DRs[0]
				goDR = DR[0]
			}

			select {

			case requests := <-c.Requests:
				DRs = append(DRs, requests)
			case request := <-c.Request:
				DR = append(DR, request)
			case goRs <- goDR:
				DRs = DRs[1:]
				DR = DR[1:]
			}

		}

	}()

}

//工作 发送管理，接收任务
//管理器，接受管理，分配任务

type Woker struct {
	Wrequest chan DRequest
}

func Newwoker() Woker {
	return Woker{Wrequest: make(chan DRequest)}
}

//工作
func (w Woker) Run(crequest chan chan DRequest, cresult chan DResult) {

	go func() {

		for {
			crequest <- w.Wrequest
			select {
			case r := <-w.Wrequest:
				result := r.Run()
				cresult <- result
			}
		}

	}()
}

//请求后台，拿数据，插入聊天记录
func (d DRequest) Run() DResult {
	var dresult DResult
	var eventalk RequestEvenTalk
	err := json.Unmarshal(d.Message, &eventalk)
	fmt.Println(string(d.Message))
	if eventalk.Msg_type != "" {
		var result DResultTalkEvent
		if err != nil || eventalk.Msg_type == "" {
			result.Status = 0
			dresult.Result = result
			return dresult
		}

		url := setting.CommonSetting.Weburl + "/even_talk"
		err, rep := util.HttPRquest(url, d.Message)
		if err != nil {
			result.Status = 100
			dresult.Result = result
			return dresult
		}
		result.Status = 100
		result.DataResponse.Send_user = eventalk.Send_user
		result.DataResponse.Source_type, _ = strconv.Atoi(eventalk.Source_type)
		result.Event = eventalk.Msg_type
		result.DataResponse.Receive_user, _ = strconv.Atoi(eventalk.Receive_user)
		var reponseweb ReponseFromWeb
		_ = json.Unmarshal(rep, &reponseweb)
		if reponseweb.Code == 200 {
			result.Status = 200
			result.DataResponse.Data = reponseweb.Data.Messagedata.Data
			result.Receivedlist = reponseweb.Data.Receive_list
		}
		dresult.Result = result
		return dresult
	}

	var evenlogin UpdateUserStatus
	err = json.Unmarshal(d.Message, &evenlogin)
	if evenlogin.Event != "" && evenlogin.User_id != 0 {
		var reponselogin ReponseLoginEvents
		if err != nil || evenlogin.Event == "" {
			reponselogin.Code = 100
			dresult.Result = reponselogin
			return dresult
		}
		url := setting.CommonSetting.Weburl + "/login_event"
		err, rep := util.HttPRquest(url, d.Message)
		if err != nil {
			reponselogin.Code = 100
			dresult.Result = reponselogin
			return dresult
		}
		_ = json.Unmarshal(rep, &reponselogin)
		reponselogin.Data.Status = evenlogin.Status
		reponselogin.Data.Event = "login_event"
		reponselogin.Data.User_id = evenlogin.User_id
		dresult.Result = reponselogin
		return dresult

	}

	var evenKeyboard KeyBoard

	err = json.Unmarshal(d.Message, &evenKeyboard)
	if evenKeyboard.Event != "" {
		var reponsekeyboard KeyBoardEvent
		reponsekeyboard.Send_user = evenKeyboard.Data.Send_user
		reponsekeyboard.Event = evenKeyboard.Event
		reponsekeyboard.Receive_user = evenKeyboard.Data.Receive_user
		dresult.Result = reponsekeyboard
		return dresult
	}

	return dresult
}

//发送给各个用户
func MessageGetResult() {

	go func() {
		for {

			select {
			case rep := <-MessageChannel.Result:

				//反射
				getValue := reflect.ValueOf(rep.Result)
				methodvalue := getValue.MethodByName("EventDo")
				methodvalue.Call(make([]reflect.Value, 0))

				//类型断言
				// switch eventalk := rep.Result.(type) {
				// case DResultTalkEvent:
				// 	if eventalk.Status == 200 {
				// 		SendMessage(eventalk)
				// 	}
				// case ReponseLoginEvents:
				// 	if eventalk.Code == 200 {
				// 		SendStatus(eventalk.Data.Receivedlist, eventalk.Data.Status, eventalk.Data.User_id, eventalk.Data.Event)
				// 	}
				// case KeyBoardEvent:
				// 	SendKeyboardEvent(eventalk)
				// }
			}

		}
	}()

}

func SendMessage(rep DResultTalkEvent) {
	switch rep.DataResponse.Source_type {
	case 1:
		userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", rep.DataResponse.Receive_user)).Result()
		if userstatus == "1" {
			client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", rep.DataResponse.Receive_user)).Result()
			if client != "" {
				evenTalk, _ := json.Marshal(rep.DataResponse)
				result := string(evenTalk)
				SendMessage2Client(client, "0", 200, rep.Event, &result)
			}
		}

		client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", rep.DataResponse.Send_user)).Result()
		if client != "" {
			evenTalk, _ := json.Marshal(rep.DataResponse)
			result := string(evenTalk)
			SendMessage2Client(client, "0", 200, rep.Event, &result)
		}

	case 2:
		if len(rep.Receivedlist) > 0 { //组
			for _, value := range rep.Receivedlist {
				userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", value)).Result()
				if userstatus == "1" {
					client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", value)).Result()
					if client != "" {
						evenTalk, _ := json.Marshal(rep.DataResponse)
						result := string(evenTalk)
						SendMessage2Client(client, "0", 200, rep.Event, &result)
					}
				}
			}
		}
	}

}

func (rep DResultTalkEvent) EventDo() {
	if rep.Status != 200 {
		return
	}
	//消息事件
	switch rep.DataResponse.Source_type {
	case 1:
		userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", rep.DataResponse.Receive_user)).Result()
		if userstatus == "1" {
			client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", rep.DataResponse.Receive_user)).Result()
			if client != "" {
				evenTalk, _ := json.Marshal(rep.DataResponse)
				result := string(evenTalk)
				SendMessage2Client(client, "0", 200, rep.Event, &result)
			}
		}

		client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", rep.DataResponse.Send_user)).Result()
		if client != "" {
			evenTalk, _ := json.Marshal(rep.DataResponse)
			result := string(evenTalk)
			SendMessage2Client(client, "0", 200, rep.Event, &result)
		}

	case 2:
		if len(rep.Receivedlist) > 0 { //组
			for _, value := range rep.Receivedlist {
				userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", value)).Result()
				if userstatus == "1" {
					client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", value)).Result()
					if client != "" {
						evenTalk, _ := json.Marshal(rep.DataResponse)
						result := string(evenTalk)
						SendMessage2Client(client, "0", 200, rep.Event, &result)
					}
				}
			}
		}
	}

}

func SendStatus(user_list []int, status, user_id int, event string) {

	if len(user_list) > 0 { //组
		for _, value := range user_list {
			userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", value)).Result()
			if userstatus == "1" {
				client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", value)).Result()
				if client != "" {
					data := UserStatus{
						Status:  status,
						User_id: user_id,
					}
					evenTalk, _ := json.Marshal(data)
					result := string(evenTalk)
					SendMessage2Client(client, "0", 200, event, &result)
				}
			}
		}
	}

}

func (rep ReponseLoginEvents) EventDo() {
	if rep.Code != 200 {
		return
	}
	// 修改状态事件
	if len(rep.Data.Receivedlist) > 0 { //组
		for _, value := range rep.Data.Receivedlist {
			userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", value)).Result()
			if userstatus == "1" {
				client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", value)).Result()
				if client != "" {
					data := UserStatus{
						Status:  rep.Data.Status,
						User_id: rep.Data.User_id,
					}
					evenTalk, _ := json.Marshal(data)
					result := string(evenTalk)
					SendMessage2Client(client, "0", 200, rep.Data.Event, &result)
				}
			}
		}
	}

}

func SendKeyboardEvent(v KeyBoardEvent) {
	userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, v.Receive_user).Result()
	fmt.Println(userstatus, v)
	if userstatus == "1" {
		client, _ := redis.RedisDB.HGet(redis.UserIdClient, v.Receive_user).Result()
		if client != "" {
			evenTalk, _ := json.Marshal(v)
			result := string(evenTalk)
			SendMessage2Client(client, "0", 200, v.Event, &result)
		}
	}
}

func (v KeyBoardEvent) EventDo() {
	//键盘敲击事件
	userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, v.Receive_user).Result()
	fmt.Println(userstatus, v)
	if userstatus == "1" {
		client, _ := redis.RedisDB.HGet(redis.UserIdClient, v.Receive_user).Result()
		if client != "" {
			evenTalk, _ := json.Marshal(v)
			result := string(evenTalk)
			SendMessage2Client(client, "0", 200, v.Event, &result)
		}
	}
}

func (rep GroupList) EventDo() {
	if len(rep.Receivedlist) > 0 { //组
		for _, value := range rep.Receivedlist {
			userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, fmt.Sprintf("%d", value)).Result()
			if userstatus == "1" {
				client, _ := redis.RedisDB.HGet(redis.UserIdClient, fmt.Sprintf("%d", value)).Result()
				if client != "" {
					result := string(`{"msg":"join_group"}`)
					SendMessage2Client(client, "0", 200, rep.Event, &result)
				}
			}
		}
	}
}
