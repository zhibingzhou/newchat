package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitWebSocketRouter(Router *gin.RouterGroup) {
	webrouter := Router.Group("")
	//websocket
	webrouter.POST("/even_talk", v1.EvenTalk)
	webrouter.POST("/login_event", v1.Login_Event)
	
}
