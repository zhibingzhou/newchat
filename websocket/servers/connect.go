package servers

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/woodylan/go-websocket/api"
	"github.com/woodylan/go-websocket/define/retcode"
	"github.com/woodylan/go-websocket/pkg/redis"
	"github.com/woodylan/go-websocket/tools/util"
)

const (
	// 最大的消息大小
	maxMessageSize = 8192
)

type Controller struct {
}

type renderData struct {
	ClientId string `json:"clientId"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)

	if err != nil {
		log.Errorf("upgrade error: %v", err)
		http.NotFound(w, r)
		return
	}

	//设置读取消息大小上线
	conn.SetReadLimit(maxMessageSize)

	//解析参数
	systemId := r.FormValue("systemId")
	if len(systemId) == 0 {
		_ = Render(conn, "", "", retcode.SYSTEM_ID_ERROR, "系统ID不能为空", []string{})
		_ = conn.Close()
		return
	}

	uid := r.FormValue("uid")
	if len(uid) == 0 {
		_ = Render(conn, "", "", retcode.SYSTEM_ID_ERROR, "用户ID不能为空", []string{})
		_ = conn.Close()
		return
	}

	clientId := util.GenClientId()

	clientSocket := NewClient(clientId, systemId, uid, conn)

	Manager.AddClient2SystemClient(systemId, clientSocket)

	//读取客户端消息
	clientSocket.Read()

	//保存user信息
	redis.RedisDB.HSet(redis.UserIdClient, uid, clientId)
	redis.RedisDB.HSet(redis.UserIdSystem, uid, systemId)
	redis.RedisDB.HSet(redis.UserStatus, uid, "1")

	if err = api.ConnRender(conn, renderData{ClientId: clientId}); err != nil {
		_ = conn.Close()
		return
	}

	// 用户连接事件
	Manager.Connect <- clientSocket
}
