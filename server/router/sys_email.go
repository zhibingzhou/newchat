package router

import (
	"github.com/gin-gonic/gin"
	"newchat/api/v1"
	"newchat/middleware"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.OperationRecord())
	{
		UserRouter.POST("emailTest", v1.EmailTest) // 发送测试邮件
	}
}
