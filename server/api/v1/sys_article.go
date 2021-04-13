package v1

import (
	"newchat/global"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// @Tags Base
// @Summary 新增类型
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func EditArticleClass(c *gin.Context) {

	var article request.RequestEditArticleClass

	_ = c.ShouldBindJSON(&article)

	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}

	err, rep := service.EditArticleClass(uid, article.Class_name)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"id": rep.ID,
	}, "success", c)

}

// @Tags Base
// @Summary 新增类型
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func DelArticleClass(c *gin.Context) {

	var article request.RequestDelArticleClass

	_ = c.ShouldBindJSON(&article)

	err := service.DelArticleClass(article.Class_id)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.Ok(c)
}

// @Tags Base
// @Summary 新增类型
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func ArticleClass(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}
	err, rep := service.ArticleClass(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rep, c)

}

// @Tags Base
// @Summary 新增类型
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func ArticleTags(c *gin.Context) {
	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}
	err, rep := service.ArticleTags(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rep, c)

}

// @Tags Base
// @Summary 新增类型
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func ArticleList(c *gin.Context) {

	page := c.Query("page")
	keyword := c.Query("keyword")
	find_type := c.Query("find_type")
	cid := c.Query("cid")

	uid := getUserID(c)
	if uid == 0 {
		response.FailWithMessage("获取Uid失败", c)
		return
	}
	err, rep := service.ArticleList(uid, page, keyword, find_type, cid)
	if err != nil && err != gorm.ErrRecordNotFound {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(rep, c)

}
