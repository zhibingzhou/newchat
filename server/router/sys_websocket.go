package router

import (
	"fmt"
	"newchat/global"

	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

func InitWebSocketRouter(Router *gin.RouterGroup) {
	webrouter := Router.Group("")
	//websocket
	server, err := socketio.NewServer(nil)
	if err != nil {
		global.GVA_LOG.Info("router register WebSocket Fail" + err.Error())
	}
	InitWebSocketEvent(server)
	webrouter.GET("/socket.io", gin.WrapH(server))
	webrouter.POST("/socket.io", gin.WrapH(server))
}

func InitWebSocketEvent(server *socketio.Server) {

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		s.Emit("reply", "have "+msg)
	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})
    
	go server.Serve()
}



