package service

import (
	"encoding/json"
	"fmt"
	"newchat/model/request"
	"newchat/model/response"

	"github.com/streadway/amqp"
)

func (r *RabbitMQ) DirectPulish(message []byte) {

	//name, kind string, durable, autoDelete, internal, noWait bool, args Table
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	//exchange, key string, mandatory, immediate bool, msg Publishing
	r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			Body:        message,
			ContentType: "text/pain",
		},
	)

}

func (r *RabbitMQ) DirectConsume() {
	//name, kind string, durable, autoDelete, internal, noWait bool, args Table
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	//name string, durable, autoDelete, exclusive, noWait bool, args Table
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		true,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
	//name, key, exchange string, noWait bool, args Table
	err = r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	//queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args Table
	result, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	forver := make(chan bool)
	for b := range result {
		switch r.Key {
		case "even_talk": //单聊，群聊事件
			talkReponse := Rabbit_even_talk(b.Body)
			result, _ := json.Marshal(talkReponse)
			RabbitAdminService["even_talk"].DirectPulish(result)
		case "login_event": //上线，下线事件
			talkReponse := Rabbit_login_event(b.Body)
			result, _ := json.Marshal(talkReponse)
			RabbitAdminService["login_event"].DirectPulish(result)
		}
	}
	<-forver
}

//处理完发消息给rabbit
func Rabbit_even_talk(body []byte) response.Response {
	var talkReponse response.Response
	requesteventalk := ReponseSyncPool["even_talk"].Get().(*request.RequestEvenTalk)
	_ = json.Unmarshal(body, &requesteventalk)
	ReponseSyncPool["even_talk"].Put(requesteventalk)

	send_id := requesteventalk.Send_user
	received_id := requesteventalk.Receive_user //
	source := requesteventalk.Source_type       //是否群聊
	msg := requesteventalk.Text_message
	msg_type := "1" //文字
	//创建记录
	err, rep := CreatTalk(send_id, received_id, source, msg_type, msg)
	if err != nil {
		talkReponse.Code = 100
		talkReponse.Msg = "新增失败"
		return talkReponse
	}
	err, res := GetTalk_listById(rep.ID)
	if err != nil {
		talkReponse.Code = 100
		talkReponse.Msg = "获取失败"
		return talkReponse
	}
	talkReponse.Code = 200
	talkReponse.Data = res
	talkReponse.Msg = "success"

	return talkReponse

}

func Rabbit_login_event(body []byte) response.Response {
	var talkReponse response.Response
	var requesteventalk request.UpdateUserStatus
	_ = json.Unmarshal(body, &requesteventalk)
	fmt.Println(requesteventalk)
	err, res := GetFriendIdList(requesteventalk.User_id, requesteventalk.Status, "login_event")
	if err != nil {
		talkReponse.Code = 100
		talkReponse.Msg = "获取失败"
		return talkReponse
	}
	talkReponse.Code = 200
	talkReponse.Data = res
	talkReponse.Msg = "success"
	return talkReponse
}
