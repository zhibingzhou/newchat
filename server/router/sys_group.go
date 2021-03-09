package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitGroupRouter(Router *gin.RouterGroup) {
	Group := Router.Group("group")
	{
		Group.GET("list", v1.Group_List)
		Group.GET("detail", v1.GroupDetail)
		Group.GET("members", v1.GroupMembers)
		Group.GET("notices", v1.GroupNotices)
		Group.POST("edit", v1.GroupEdit)
		Group.POST("edit-notice", v1.EditNotice)
	}
}