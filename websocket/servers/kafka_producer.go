package servers

import (
	"fmt"

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

var KafkaSendService map[string]TopicAndKey
var KafkaReceiveService map[string]TopicAndKey

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

	//发送者
	KafkaSendService = make(map[string]TopicAndKey)
	KafkaSendService["even_talk"] = NewTopicAndKey("event_web", "even_talk")
	KafkaSendService["login_event"] = NewTopicAndKey("event_web", "login_event")
	KafkaSendService["event_keyboard"] = NewTopicAndKey("event_web_adm", "event_keyboard") //同个系统发送接受

	//接收者  和rabbitmq 不一样，kafka 可以一个topic 下面有多个key ，只需要监控topic , rabbitmq 需要 通道 和 key 都监控
	KafkaReceiveService = make(map[string]TopicAndKey)
	KafkaReceiveService["event_adm"] = NewTopicAndKey("event_adm", "")
	KafkaReceiveService["event_web_adm"] = NewTopicAndKey("event_web_adm", "")

	for _, value := range KafkaReceiveService {
		go value.ListenKafka()
	}

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
