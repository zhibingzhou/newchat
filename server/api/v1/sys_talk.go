package v1

import (
	"fmt"
	"newchat/global"
	"newchat/model"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/service"
	"newchat/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// @Tags Base
// @Summary 信息详情
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func TalkRecords(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	record_id := c.Query("record_id")
	receive_id := c.Query("receive_id")
	source := c.Query("source")

	err, rep := service.TalkRecords(record_id, source, receive_id, uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func NotDisturb(c *gin.Context) {
	uid := getUserID(c)
	var dis request.RequestNotDistraub
	_ = c.ShouldBindJSON(&dis)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}
	err := service.NotDisturb(uid, dis)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.Ok(c)

}

// @Tags Base
// @Summary 创建消息列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func TalkCreate(c *gin.Context) {
	uid := getUserID(c)
	var dis request.RequestCreate
	_ = c.ShouldBindJSON(&dis)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}
	err, rep := service.TalkCreate(uid, dis)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithData(response.ResponseCreateTalk{
		TalkItem: rep,
	}, c)

}

// @Tags Base
// @Summary 创建消息列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func UpdateUnreadNum(c *gin.Context) {
	uid := getUserID(c)
	var dis request.RequestUpdateNoread
	_ = c.ShouldBindJSON(&dis)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}
	msg_type, err := strconv.Atoi(dis.Type)
	if err != nil {
		response.FailWithMessage("Type 请求格式错误", c)
	}
	received, err := strconv.Atoi(dis.Receive)
	if err != nil {
		response.FailWithMessage("Receive 请求格式错误", c)
	}
	err = service.UpdateUnreadNum(uid, msg_type, received)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.Ok(c)

}

// @Tags Base
// @Summary 信息详情
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func SendImage(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	//record_id := c.Query("record_id")
	receive_id := c.PostForm("receive_id")
	source := c.PostForm("source")
	width := c.PostForm("width")
	height := c.PostForm("height")

	files, err := c.FormFile("img")
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Any("err", err))
		response.FailWithMessage("上传失败", c)
		return
	}
	// 上传文件至指定目录
	guid := uuid.New().String()
	// _690x690.jpg
	widthandheight := "_" + width + "x" + height
	singleFile := "uploads/file/img/" + guid + widthandheight + utils.GetExt(files.Filename)
	_ = c.SaveUploadedFile(files, singleFile)

	weburl := global.GVA_CONFIG.System.Url
	url := fmt.Sprintf("%s/%s", weburl, singleFile)

	err, res := service.CreatTalk(uid, receive_id, source, "2", "")
	file := model.File{
		Record_id:     res.ID,
		User_id:       res.User_id,
		File_source:   res.Source,
		File_size:     int(files.Size),
		Original_name: files.Filename,
		File_suffix:   strings.Replace(utils.GetExt(files.Filename), ".", "", 1),
		Save_dir:      singleFile,
		File_url:      url,
		File_type:     1,
	}

	err = service.SendImage(file)

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}

	err, data := service.GetTalk_listById(res.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	}

	go service.SendToClient("event_img", data)

	response.Ok(c)

}

// @Tags Base
// @Summary 信息详情
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func ChatRecords(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
	}

	//record_id := c.Query("record_id")
	receive_id := c.Query("receive_id")
	source := c.Query("source")
	msg_type := c.Query("msg_type")

	err, rep := service.ChatRecords(source, msg_type, receive_id, uid)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(rep, "success", c)

}
