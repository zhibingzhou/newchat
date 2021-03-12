package servers

import (
	"time"

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
				MessageChannel.Request <- DRequest{Message: message}
			}

		}
	}()
}

func WebSendMessage(message []byte) {
	MessageChannel.Request <- DRequest{Message: message}
}
