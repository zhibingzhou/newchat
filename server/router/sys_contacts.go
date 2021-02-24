package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitContactsRouter(Router *gin.RouterGroup) {
	Contacts := Router.Group("contacts")
	{
		Contacts.GET("apply-unread-num", v1.Apply_unread_num)
		Contacts.GET("apply-records", v1.Apply_records)
	}
}
