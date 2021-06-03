package servers

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
					var reponseweb ReponseFromWeb
					_ = json.Unmarshal(msg.Value, &reponseweb)
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
					_ = json.Unmarshal(msg.Value, &reponselogin)
					reponselogin.EventDo()
				case "event_keyboard":
					var evenKeyboard KeyBoard
					err = json.Unmarshal(msg.Value, &evenKeyboard)
					var reponsekeyboard KeyBoardEvent
					reponsekeyboard.Send_user = evenKeyboard.Data.Send_user
					reponsekeyboard.Event = evenKeyboard.Event
					reponsekeyboard.Receive_user = evenKeyboard.Data.Receive_user
					reponsekeyboard.EventDo()
				case "join_group":
					var join_group GroupList
					err = json.Unmarshal(msg.Value, &join_group)
					join_group.EventDo()
				case "event_img":
					var inputData ReponseFromTalkEvent
					var rep DResultTalkEvent
					err = json.Unmarshal(msg.Value, &inputData)
					rep.Event = inputData.Event
					rep.Status = 200
					rep.Receivedlist = inputData.Receive_list
					rep.DataResponse.Data = inputData.Messagedata.Data
					rep.DataResponse.Data.File = inputData.Messagedata.Data.File
					rep.DataResponse.Receive_user = inputData.Messagedata.Receive_user
					rep.DataResponse.Send_user = inputData.Messagedata.Send_user
					rep.DataResponse.Source_type = inputData.Messagedata.Source_type
					rep.EventDo()
				}
				fmt.Printf("topic:%s Key:%v Value:%v", msg.Topic, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}

	select {}
}
