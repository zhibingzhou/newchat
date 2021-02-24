package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitMessageRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("talk")
	{
		UserRouter.GET("list", v1.Talk_List)
	}
}
