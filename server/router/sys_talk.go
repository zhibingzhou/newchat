package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("talk")
	{
		UserRouter.GET("list", v1.Talk_List)
		UserRouter.GET("records", v1.TalkRecords)
		UserRouter.POST("set-not-disturb", v1.NotDisturb)
		UserRouter.POST("create", v1.TalkCreate)
		UserRouter.POST("update-unread-num", v1.UpdateUnreadNum)
		UserRouter.GET("find-chat-records", v1.ChatRecords)
		UserRouter.POST("send-image", v1.SendImage)
		UserRouter.POST("send-emoticon", v1.SendEmoticon)
		UserRouter.POST("topping", v1.Topping)
		UserRouter.POST("delete", v1.TalkDelete)
       
	}
}
