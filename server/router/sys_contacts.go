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
		Contacts.GET("list", v1.Contacts_List)
		Contacts.GET("search", v1.Contacts_Search)
		Contacts.POST("add", v1.Contacts_Add)
		Contacts.POST("delete-apply", v1.DeleteApply)
		Contacts.POST("accept-invitation", v1.AcceptInvitation)
		Contacts.POST("delete", v1.Contacts_Delete)
		Contacts.POST("edit-remark", v1.Edit_Remark)
	

	}
}
