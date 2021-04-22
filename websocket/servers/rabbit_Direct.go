package servers

import (
	"encoding/json"
	"fmt"

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
			var reponseweb ReponseFromWeb
			_ = json.Unmarshal(b.Body, &reponseweb)
			if reponseweb.Code == 200 {
				var result DResultTalkEvent
				result.DataResponse.Send_user = reponseweb.Data.Messagedata.Send_user
				result.DataResponse.Source_type = reponseweb.Data.Messagedata.Source_type
				result.Event = reponseweb.Data.Event
				result.DataResponse.Receive_user = reponseweb.Data.Messagedata.Receive_user
				result.Status = 200
				result.DataResponse.Data = reponseweb.Data.Messagedata.Data
				result.Receivedlist = reponseweb.Data.Receive_list
				result.EventDo()
			}
		case "login_event": //上线，下线事件
			var reponselogin ReponseLoginEvents
			_ = json.Unmarshal(b.Body, &reponselogin)
			reponselogin.EventDo()
		case "event_keyboard":
			var evenKeyboard KeyBoard
			err = json.Unmarshal(b.Body, &evenKeyboard)
			var reponsekeyboard KeyBoardEvent
			reponsekeyboard.Send_user = evenKeyboard.Data.Send_user
			reponsekeyboard.Event = evenKeyboard.Event
			reponsekeyboard.Receive_user = evenKeyboard.Data.Receive_user
			reponsekeyboard.EventDo()
		case "join_group":
			var join_group GroupList
			err = json.Unmarshal(b.Body, &join_group)
			join_group.EventDo()
		}
	}
	<-forver
}
