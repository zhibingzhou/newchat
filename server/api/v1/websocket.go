package v1

import (
	"newchat/global"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//
func EvenTalk(c *gin.Context) {
	var requesteventalk request.RequestEvenTalk
	_ = c.ShouldBindJSON(&requesteventalk)
	send_id := requesteventalk.Send_user
	received_id := requesteventalk.Receive_user //
	source := requesteventalk.Source_type       //是否群聊
	msg := requesteventalk.Text_message
	msg_type := "1" //文字
	//创建记录
	err, rep := service.CreatTalk(send_id, received_id, source, msg_type, msg)
	if err != nil {
		global.GVA_LOG.Error("新增失败!", zap.Any("err", err))
		response.FailWithMessage("新增失败", c)
		return
	}
	err, res := service.GetTalk_listById(rep.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(res, c)
}
