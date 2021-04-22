package servers

import (
	"bytes"
	"fmt"
	"log"
	"websocket/pkg/setting"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://zhou:123@127.0.0.1:5672/zhou" //amqp:// 账号 : 密码 @ 连接生产者、消费者的端口 /verhost name

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	//连接信息
	Mqurl string
}

var RabbitWebSocketService map[string]*RabbitMQ
var RabbitAdminService map[string]*RabbitMQ

func NewDirect(exchange, key string) *RabbitMQ {
	r := NewRabbitMQ("", exchange, key)
	return r
}

func InitRabbitService() {
	RabbitWebSocketService = make(map[string]*RabbitMQ)
	RabbitWebSocketService["even_talk"] = NewDirect("event_web", "even_talk")
	RabbitWebSocketService["login_event"] = NewDirect("event_web", "login_event")
	RabbitWebSocketService["event_keyboard"] = NewDirect("event_web_adm", "event_keyboard") //同个系统发送接受

	RabbitAdminService = make(map[string]*RabbitMQ)
	RabbitAdminService["even_talk"] = NewDirect("event_adm", "even_talk")
	RabbitAdminService["login_event"] = NewDirect("event_adm", "login_event")
	RabbitAdminService["join_group"] = NewDirect("event_adm", "join_group")
	RabbitAdminService["event_keyboard"] = NewDirect("event_web_adm", "event_keyboard")

	for _, value := range RabbitAdminService {
		go value.DirectConsume()
	}
}

func NewRabbitMQ(QueueName, Exchange, key string) *RabbitMQ {
	MQURL := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", setting.CommonRabbitMq.Admin, setting.CommonRabbitMq.Pwd, setting.CommonRabbitMq.Ip, setting.CommonRabbitMq.Port, setting.CommonRabbitMq.Verhost)
	fmt.Println(MQURL)
	rabbitmq := RabbitMQ{QueueName: QueueName, Exchange: Exchange, Key: key, Mqurl: MQURL}
	var err error
	//先连接
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接错误")
	//再获取通道
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel错误")
	return &rabbitmq
}

//中断连接
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		//错误日志打印
		log.Fatalf("%s:%s", message, err)
		//panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}