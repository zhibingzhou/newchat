package initialize

import (
	"net/http"
	_ "newchat/docs"
	"newchat/global"
	"newchat/middleware"
	"newchat/router"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}

	webSocketGroup := Router.Group("") //websocket服务 通讯部分
	webSocketGroup.Use(middleware.HostAuth())
	{
		router.InitWebSocketRouter(webSocketGroup)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		router.InitJwtRouter(PrivateGroup)                   // jwt相关路由
		router.InitUserRouter(PrivateGroup)                  // 注册用户路由
		router.InitMessageRouter(PrivateGroup)               //消息
		router.InitContactsRouter(PrivateGroup)              //联系人
		router.InitGroupRouter(PrivateGroup)                 //群
		router.InitEmoticonRouter(PrivateGroup)              //表情包
		router.InitArticleRouter(PrivateGroup)
		//router.InitWebSocketRouter(PrivateGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
