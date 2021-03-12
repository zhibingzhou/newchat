package model

import (
	"encoding/json"
	"fmt"
	"newchat/global"
	"time"

	"github.com/go-redis/redis"
)

type Group_id struct {
	Group_id int `json:"group_id" gorm:"comment:用户id"`
}

type Group_notice struct {
	global.GVA_MODEL
	User_id  int    `json:"user_id"  gorm:"comment:用户id"`
	Group_id int    `json:"group_id"  gorm:"comment:群id"`
	Title    string `json:"title" gorm:"comment:公告标题"`
	Content  string `json:"content" gorm:"comment:公告内容"`
}

type Group_list struct {
	global.GVA_MODEL
	Manager_id    int    `json:"manager_id"  gorm:"comment:群主id"`
	Group_profile string `json:"group_profile"  gorm:"comment:群简介"`
	Avatar        string `json:"avatar"  gorm:"comment:群头像"`
	Group_name    string `json:"group_name"  gorm:"comment:群名称"`
}

//群成员列表
func Redis_GetGroupbyId(group_id int) (err error, user_idlist []UserId) {

	redisKey := "newchat:group_id_member:" + fmt.Sprintf("%d", group_id)
	//优先查询redis 拿map
	filejson, err := global.GVA_REDIS.Get(redisKey).Result()

	if err == redis.Nil {
		// 查询数据库 得 map
		global.GVA_DB.Debug().Table("group_member").Select([]string{"user_id"}).Where("group_id = ?", group_id).Scan(&user_idlist)

		rejson, _ := json.Marshal(user_idlist)

		err = global.GVA_REDIS.Set(redisKey, string(rejson), time.Hour*1).Err()
		if err != nil {
			return err, user_idlist
		}

		//新增无序集合 所有的key头存在无序集合里面
		err = global.GVA_REDIS.SAdd(global.FileId, redisKey).Err()
		if err != nil {
			return err, user_idlist
		}

		filejson, err = global.GVA_REDIS.Get(redisKey).Result()
		if err != nil {
			return err, user_idlist
		}
	}

	if err != nil {
		return err, user_idlist
	}

	err = json.Unmarshal([]byte(filejson), &user_idlist)
	return err, user_idlist
}
