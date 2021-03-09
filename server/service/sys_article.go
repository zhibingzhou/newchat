package service

import (
	"newchat/global"
	"newchat/model"
	"newchat/model/response"
	"newchat/utils"
)

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func EditArticleClass(uid int, name string) (err error, article model.ArticleClass) {

	article = model.ArticleClass{
		Name:    name,
		User_id: uid,
	}
	var sort model.ArticleSort
	err = global.GVA_DB.Raw("SELECT MAX(sort) as sort FROM article_class where user_id = ?", uid).Scan(&sort).Error
	if err != nil {
	} else {
		article.Sort = sort.Sort + 1
	}

	err = global.GVA_DB.Create(&article).Error

	return err, article
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func DelArticleClass(id int) (err error) {

	err = global.GVA_DB.Delete(model.ArticleClass{}, "id = ?", id).Error

	return err
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func ArticleClass(id int) (err error, rep response.ResponseArticleClassList) {

	err = global.GVA_DB.Debug().Raw("SELECT  article_class.`id` , article_class.`name` AS class_name, article_class.`default` AS is_default ,( SELECT COUNT(*) FROM article_list WHERE  article_list.`class_id` = article_class.`id` AND article_class.`user_id` = ? AND article_list.`status` = 1 ) AS count  FROM article_class WHERE article_class.`user_id` = ? AND article_class.`deleted_at` IS NULL ", id, id).Scan(&rep.Rows).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func ArticleTags(id int) (err error, rep response.ResponseTagsList) {

	err = global.GVA_DB.Debug().Raw("SELECT article_tags.`id`,article_tags.`name` AS tag_name , (SELECT COUNT(*) FROM article_tag_list WHERE  article_tags.`id` = article_tag_list.`tags_id` AND article_tags.user_id = ?   ) AS count FROM article_tags WHERE article_tags.user_id = ?  ", id, id).Scan(&rep.Tags).Error

	return err, rep
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 好友列表
//@param: id int
//@return: err error, user *model.SysUser

func ArticleList(id int, pages, keyword, find_type, cid string) (err error, rep response.ResponseArticleList) {

	page, pagesize := utils.ThreadPage(pages, "0")
	rep.Rows = make([]response.ResponseArticle, 0)
	total := 0
	rep.Page = page
	rep.Page_size = pagesize
	rep.Page_total = 0
	rep.Total = 0

	switch find_type {

	case "1":

		err = global.GVA_DB.Table("article_list").Where("user_id = ?", id).Count(&total).Error
		if err != nil || total == 0 {
			return err, rep
		}
		//近期笔记
		err = global.GVA_DB.Raw("SELECT article_list.`id`,article_list.`class_id`,article_list.`title`,article_list.`image`,article_list.`abstract`,article_list.`updated_at`,article_list.`status`,article_class.`name` FROM article_list,article_class WHERE article_list.`class_id` = article_class.`id` AND article_class.`user_id` = ? AND  article_class.`deleted_at` IS NULL  ORDER BY article_list.`updated_at` DESC  LIMIT ?  OFFSET ? ", id, pagesize, page-1).Scan(&rep.Rows).Error

		break
	case "2":
		err = global.GVA_DB.Table("article_list").Where("user_id = ? and heart = 1", id).Count(&total).Error

		if err != nil || total == 0 {
			return err, rep
		}
		//星标笔记
		err = global.GVA_DB.Debug().Raw("SELECT article_list.`id`,article_list.`class_id`,article_list.`title`,article_list.`image`,article_list.`abstract`,article_list.`updated_at`,article_list.`status`,article_class.`name` FROM article_list,article_class WHERE article_list.`class_id` = article_class.`id` AND article_class.`user_id` = ? AND article_list.`heart` = 1 AND  article_class.`deleted_at` IS NULL  ORDER BY article_list.`updated_at` DESC  LIMIT ?  OFFSET ? ", pagesize, page-1).Scan(&rep.Rows).Error
		break
	case "3":

		err = global.GVA_DB.Table("article_list").Where("user_id = ? and class_id = ?", id, cid).Count(&total).Error

		if err != nil || total == 0 {
			return err, rep
		}

		//分类笔记
		err = global.GVA_DB.Raw("SELECT article_list.`id`,article_list.`class_id`,article_list.`title`,article_list.`image`,article_list.`abstract`,article_list.`updated_at`,article_list.`status`,article_class.`name` FROM article_list,article_class WHERE article_list.`class_id` = article_class.`id` AND article_class.`user_id` = ? AND article_list.`class_id` = ? AND  article_class.`deleted_at` IS NULL  ORDER BY article_list.`updated_at` DESC  LIMIT ?  OFFSET ? ", id, cid, pagesize, page-1).Scan(&rep.Rows).Error
		break
	case "4":
		var article_id []model.ArticleId
		var idlist []int

		err = global.GVA_DB.Table("article_tag_list").Select([]string{"article_id"}).Where("tags_id = ?", cid).Scan(&article_id).Error

		for key, _ := range article_id {
			idlist = append(idlist, article_id[key].Article_id)
		}
		total = len(idlist)

		if total == 0 {
			return err, rep
		}

		err = global.GVA_DB.Debug().Raw("SELECT article_list.`id`,article_list.`class_id`,article_list.`title`,article_list.`image`,article_list.`abstract`,article_list.`updated_at`,article_list.`status`,article_class.`name` FROM article_list,article_class WHERE article_list.`class_id` = article_class.`id` AND article_class.`user_id` = ? AND article_list.`id` IN (?) AND  article_class.`deleted_at` IS NULL  ORDER BY article_list.`updated_at` DESC  LIMIT ?  OFFSET ? ", id, idlist, pagesize, page-1).Scan(&rep.Rows).Error
		//标签
		break
	case "5":
		err = global.GVA_DB.Table("article_list").Where("user_id = ? and status = 2", id).Count(&total).Error
		if err != nil || total == 0 {
			return err, rep
		}

		err = global.GVA_DB.Raw("SELECT article_list.`id`,article_list.`class_id`,article_list.`title`,article_list.`image`,article_list.`abstract`,article_list.`updated_at`,article_list.`status`,article_class.`name` FROM article_list,article_class WHERE article_list.`class_id` = article_class.`id` AND article_class.`user_id` = ? and article_list.status = 2  AND  article_class.`deleted_at` IS NULL  ORDER BY article_list.`updated_at` DESC  LIMIT ?  OFFSET ? ", id, pagesize, page-1).Scan(&rep.Rows).Error
		//回收站
		break
	}
	rep.Page = page
	rep.Page_size = pagesize
	rep.Page_total = len(rep.Rows)
	rep.Total = total
	return err, rep
}
