package service

import (
	"fmt"
	"newchat/model/request"
	"sync"

	"github.com/Shopify/sarama"
)

var (
	ServerCon sarama.SyncProducer
)

type TopicAndKey struct {
	Topic string
	Key   string
}

//中断连接
func Destory() {
	ServerCon.Close()
}

var (
	ReponseSyncPool     map[string]*sync.Pool
	KafkaSendService    map[string]TopicAndKey
	KafkaReceiveService map[string]TopicAndKey
)

//复用池子
func InitPoolOfReponse() {
	ReponseSyncPool = make(map[string]*sync.Pool)
	ReponseSyncPool["even_talk"] = &sync.Pool{New: func() interface{} { return new(request.RequestEvenTalk) }}
}

func InitkafkaService() error {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// 连接kafka
	Server, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		panic(err)
		return err
	}
	ServerCon = Server

	//后台要发送信息过去
	KafkaSendService = make(map[string]TopicAndKey)
	KafkaSendService["even_talk"] = NewTopicAndKey("event_adm", "even_talk")
	KafkaSendService["login_event"] = NewTopicAndKey("event_adm", "login_event")
	KafkaSendService["join_group"] = NewTopicAndKey("event_adm", "join_group")
	KafkaSendService["event_img"] = NewTopicAndKey("event_adm", "event_img")

	//监听者
	KafkaReceiveService = make(map[string]TopicAndKey)
	KafkaReceiveService["event_web"] = NewTopicAndKey("event_web", "")

	for _, value := range KafkaReceiveService {
		go value.ListenKafka()
	}

	//sync.Pool
	InitPoolOfReponse()

	return err

}

func NewTopicAndKey(topic, key string) TopicAndKey {
	return TopicAndKey{Topic: topic, Key: key}
}

func (t TopicAndKey) KafkaSendMessage(Message []byte) {

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = t.Topic
	msg.Value = sarama.ByteEncoder(Message)
	msg.Key = sarama.StringEncoder(t.Key)

	// 发送消息
	pid, offset, err := ServerCon.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
