package v1

import (
	"newchat/global"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 查看好友列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Group_List(c *gin.Context) {
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	err, rep := service.Group_List(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 查看好友列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func GroupDetail(c *gin.Context) {
	uid := getUserID(c)
	group_id := c.Query("group_id")
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	err, rep := service.GroupDetail(group_id, uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 查看群成员
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func GroupMembers(c *gin.Context) {
	// uid := getUserID(c)
	group_id := c.Query("group_id")
	// if uid == 0 {
	// 	response.FailWithMessage("获取Uid失败", c)
	// }

	err, rep := service.GroupMembers(group_id)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 查看群内容
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func GroupNotices(c *gin.Context) {
	// uid := getUserID(c)
	group_id := c.Query("group_id")
	// if uid == 0 {
	// 	response.FailWithMessage("获取Uid失败", c)
	// }

	err, rep := service.GroupNotices(group_id)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 群编辑
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func GroupEdit(c *gin.Context) {
	uid := getUserID(c)
	var editjson request.RequestGroupEdit
	_ = c.ShouldBindJSON(&editjson)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	err := service.GroupEdit(uid, editjson)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.Ok(c)

}



// @Tags Base
// @Summary 群编辑
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func EditNotice(c *gin.Context) {
	uid := getUserID(c)
	var editjson request.RequestEditGroupEdit
	_ = c.ShouldBindJSON(&editjson)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	err := service.EditNotice(uid, editjson)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.Ok(c)

}