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
		Group.GET("invite-friends", v1.InviteFriends)
		Group.POST("create", v1.GroupCreate)
		Group.POST("invite", v1.GroupInvite)
		Group.POST("secede", v1.GroupSecede)
		Group.POST("set-group-card", v1.SetGroupCard)
		Group.POST("remove-members", v1.RemoveMembers)

	}
}
