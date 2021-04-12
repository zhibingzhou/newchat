package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	BaseRouter := Router.Group("auth")
	{
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("register_websocket", v1.RegisterWebsocket)
		BaseRouter.POST("register", v1.UserRegister)
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.POST("logout", v1.JsonInBlacklist) // jwt加入黑名单
	}
	return BaseRouter
}
