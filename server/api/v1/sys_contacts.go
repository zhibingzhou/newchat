package v1

import (
	"newchat/global"
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
	}

	page_size := c.Request.FormValue("page_size")

	page := c.Request.FormValue("page")

	pageIndex, pageSize = utils.ThreadPage(page, page_size)

	err, rep, total, page_total := service.Apply_records(uid, pageIndex, pageSize)
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
	}

	err, rep := service.Contacts_List(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}
