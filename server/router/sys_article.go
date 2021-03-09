package router

import (
	v1 "newchat/api/v1"

	"github.com/gin-gonic/gin"
)

func InitArticleRouter(Router *gin.RouterGroup) {
	Article := Router.Group("article")
	{
		Article.POST("edit-article-class", v1.EditArticleClass)
		Article.POST("del-article-class", v1.DelArticleClass)
		Article.GET("article-class", v1.ArticleClass)
		Article.GET("article-list", v1.ArticleList)
		Article.GET("article-tags", v1.ArticleTags)

	}
}
