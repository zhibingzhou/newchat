package service

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

func (t TopicAndKey) ListenKafka() error {
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(t.Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println("start Listen -> ", t.Topic, partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(t.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {

				switch string(msg.Key) {
				case "even_talk": //单聊，群聊事件
					talkReponse := Rabbit_even_talk(msg.Value)
					result, _ := json.Marshal(talkReponse)

					KafkaSendService["even_talk"].KafkaSendMessage(result)

				case "login_event": //上线，下线事件
					talkReponse := Rabbit_login_event(msg.Value)
					result, _ := json.Marshal(talkReponse)

					KafkaSendService["login_event"].KafkaSendMessage(result)
				}
				fmt.Printf("topic:%s Key:%v Value:%v", msg.Topic, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}

	select {}
}
