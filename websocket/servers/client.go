package servers

import (
	"time"
	"websocket/pkg/redis"
	"websocket/pkg/setting"

	"github.com/gorilla/websocket"
)

type Client struct {
	ClientId    string          // 标识ID
	SystemId    string          // 系统ID
	Socket      *websocket.Conn // 用户连接
	ConnectTime uint64          // 首次连接时间
	IsDeleted   bool            // 是否删除或下线
	UserId      string          // 业务端标识用户ID
	Extend      string          // 扩展字段，用户可以自定义
	GroupList   []string
}

type SendData struct {
	Code int
	Msg  string
	Data *interface{}
}

func NewClient(clientId string, systemId, userId string, socket *websocket.Conn) *Client {
	return &Client{
		ClientId:    clientId,
		SystemId:    systemId,
		UserId:      userId,
		Socket:      socket,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
	}
}

func (c *Client) Read() {
	go func() {
		for {
			messageType, message, err := c.Socket.ReadMessage()
			if err != nil {
				if messageType == -1 && websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
					Manager.DisConnect <- c
					return
				} else if messageType != websocket.PingMessage {
					return
				}
			}
			if string(message) != "PING" {
				//chan  方法
				// MessageChannel.Request <- DRequest{Message: message}
				//rabitmq 方法
				switch setting.CommonTool.ToolName {
				case "channel":
					MessageChannel.Request <- DRequest{Message: message}
				case "rabbitmq":
					r := RabbitMqSend{Message: message}
					ChannelAll.ChannelReceiveMessage <- r
				case "kafka":
					r := KafkaSend{Message: message}
					ChannelAll.ChannelReceiveMessage <- r
				}

			} else {
				userstatus, _ := redis.RedisDB.HGet(redis.UserStatus, c.UserId).Result()
				//更新在线状态，上线通知
				if userstatus != "1" {
					redis.RedisDB.HSet(redis.UserStatus, c.UserId, "1")
					SendUserStatus(c.UserId, 1)
					// check, _ := redis.RedisDB.HGet(redis.UserCheck, c.UserId).Result()
					// if check != "1" {
					// 	redis.RedisDB.HSet(redis.UserCheck, c.UserId, "1")
					// }
				}
			}
		}
	}()
}

func WebSendMessage(message []byte) {
	MessageChannel.Request <- DRequest{Message: message}
}
