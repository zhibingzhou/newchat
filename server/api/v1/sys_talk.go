package v1

import (
	"newchat/global"
	"newchat/model/response"
	"newchat/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Talk_List(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}
	err, rep := service.FindTalk_List(uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return 
	}
	response.OkWithDetailed(rep, "success", c)

}
