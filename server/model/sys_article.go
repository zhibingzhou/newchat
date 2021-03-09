package model

import "newchat/global"

type ArticleClass struct {
	global.GVA_MODEL
	Default int    `json:"default" gorm:"comment:备用"`
	User_id int    `json:"user_id" gorm:"comment:用户id"`
	Sort    int    `json:"sort" gorm:"comment:排序"`
	Name    string `json:"name" gorm:"comment:类型名称"`
}

type ArticleSort struct {
	Sort int `json:"sort" gorm:"comment:排序"`
}

type ArticleId struct {
	Article_id int `json:"article_id" gorm:"comment:排序"`
}

type ArticleTags struct {
	global.GVA_MODEL
	Name    int `json:"name" gorm:"comment:备用"`
	User_id int `json:"user_id" gorm:"comment:用户id"`
}
