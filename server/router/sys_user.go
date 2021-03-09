package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("users")
	{
		UserRouter.POST("change-password", v1.ChangePassword) // 修改密码
		UserRouter.GET("detail", v1.UserDetail)              // 用户信息
		UserRouter.GET("setting", v1.Setting)
		UserRouter.POST("search-user", v1.Search_user)
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
		UserRouter.POST("edit-user-detail", v1.EditUserDetail)            // 设置用户信息
		
	}
}
