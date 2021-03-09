package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitEmoticonRouter(Router *gin.RouterGroup) {
	Group := Router.Group("emoticon")
	{
		Group.GET("user-emoticon", v1.UserEmoticon)
		Group.GET("system-emoticon", v1.SystemEmoticon)
		Group.POST("set-user-emoticon", v1.SetUserEmoticon)
		Group.POST("upload-emoticon", v1.UploadEmoticon)
	}
}
