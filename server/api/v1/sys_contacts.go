package v1

import (
	"newchat/global"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/service"
	"newchat/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 查看好友请求列表数量
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Apply_unread_num(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}
	err, rep := service.Apply_unread_num(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 查看好友请求列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Apply_records(c *gin.Context) {
	uid := getUserID(c)
	var pageSize = 10
	var pageIndex = 1
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	page_size := c.Request.FormValue("page_size")

	page := c.Request.FormValue("page")

	pageIndex, pageSize = utils.ThreadPage(page, page_size)
	offset := pageSize * (pageIndex - 1)
	err, rep, total, page_total := service.Apply_records(uid, offset, pageSize)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.ResponseApply{
		Rows:       rep,
		Page:       pageIndex,
		Page_size:  pageSize,
		Page_total: page_total,
		Total:      total,
	}, "success", c)

}

// @Tags Base
// @Summary 查看好友列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Contacts_List(c *gin.Context) {
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	err, rep := service.Contacts_List(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 手机号查看用户
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Contacts_Search(c *gin.Context) {
	uid := getUserID(c)
	mobile := c.Query("mobile")
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	err, rep := service.Contacts_Search(mobile)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 手机号查看用户
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Contacts_Add(c *gin.Context) {
	uid := getUserID(c)
	var contacts_add request.RequestContactsAdd
	_ = c.ShouldBindJSON(&contacts_add)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	err := service.Contacts_Add(uid, contacts_add.Friend_id, contacts_add.Remarks)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithMessage("加好友审请成功", c)

}

// @Tags Base
// @Summary 手机号查看用户
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func DeleteApply(c *gin.Context) {

	var contacts_del request.RequestDeleteContacts
	_ = c.ShouldBindJSON(&contacts_del)

	err := service.Contacts_Del(contacts_del.Apply_id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)

}

// @Tags Base
// @Summary 接收请求
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func AcceptInvitation(c *gin.Context) {
	var apply_accept request.RequestAcceptApply
	_ = c.ShouldBindJSON(&apply_accept)
	userInfo := getUserInfobytoken(c)
	err := service.AcceptInvitation(apply_accept.Apply_id, apply_accept.Remarks, userInfo.Nickname)
	if err != nil {
		global.GVA_LOG.Error("加好友失败!", zap.Any("err", err))
		response.FailWithMessage("加好友失败", c)
		return
	}
	response.OkWithMessage("加好友成功", c)

}

// @Tags Base
// @Summary 接收请求
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Contacts_Delete(c *gin.Context) {
	var concacts_delete request.RequestContactsDelete
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}
	_ = c.ShouldBindJSON(&concacts_delete)

	err := service.Contacts_Delete(uid, concacts_delete.Friend_id)
	if err != nil {
		global.GVA_LOG.Error("删除好友失败!", zap.Any("err", err))
		response.FailWithMessage("删除好友失败", c)
		return
	}
	response.OkWithMessage("删除好友成功", c)

}

// @Tags Base
// @Summary 编辑备注
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Edit_Remark(c *gin.Context) {
	var concacts_edit request.RequestContactsEditRemarks
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}
	_ = c.ShouldBindJSON(&concacts_edit)

	err := service.Edit_Remark(uid, concacts_edit.Friend_id, concacts_edit.Remarks)
	if err != nil {
		global.GVA_LOG.Error("编辑备注失败!", zap.Any("err", err))
		response.FailWithMessage("编辑备注失败", c)
		return
	}
	response.OkWithMessage("编辑备注成功", c)

}
