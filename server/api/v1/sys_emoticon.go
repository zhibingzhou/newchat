package v1

import (
	"fmt"
	"newchat/global"
	"newchat/model"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/service"
	"newchat/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 查看好友列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func UserEmoticon(c *gin.Context) {
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	err, rep := service.UserEmoticon(uid)
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
func SystemEmoticon(c *gin.Context) {
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	err, rep := service.SystemEmoticon()
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
func SetUserEmoticon(c *gin.Context) {
	uid := getUserID(c)

	if uid != 1 {
		response.FailWithMessage("权限不足", c)
	}

	var editjson request.RequestSetUserEmoticon
	_ = c.ShouldBindJSON(&editjson)

	err := service.SetUserEmoticon(editjson.Emoticon_id, editjson.Type)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.Ok(c)

}

// @Tags Base
// @Summary 查看好友列表
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func UploadEmoticon(c *gin.Context) {
	uid := getUserID(c)

	if uid == 0 {
		response.FailWithMessage("无此用户", c)
	}

	width := c.PostForm("width")
	height := c.PostForm("height")

	files, err := c.FormFile("emoticon")
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Any("err", err))
		response.FailWithMessage("上传失败", c)
		return
	}
	// 上传文件至指定目录
	guid := uuid.New().String()

	widthandheight := "_" + width + "x" + height
	singleFile := "uploads/file/img/" + guid + widthandheight + utils.GetExt(files.Filename)
	_ = c.SaveUploadedFile(files, singleFile)

	weburl := global.GVA_CONFIG.System.Url
	url := fmt.Sprintf("%s/%s", weburl, singleFile)

	file := model.Emoticon{
		User_id:     uid,
		Size:        int(files.Size),
		Name:        files.Filename,
		File_suffix: strings.Replace(utils.GetExt(files.Filename), ".", "", 1),
		Save_dir:    singleFile,
		Src:         url,
		Status:      1,
	}

	err, rep := service.UploadEmoticon(file)
	if err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Any("err", err))
		response.FailWithMessage("上传失败", c)
		return
	}

	response.OkWithData(response.ResponseEmoticon{
		Src:      rep.Src,
		Media_id: rep.ID,
	}, c)

}
