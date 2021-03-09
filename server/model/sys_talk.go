// 自动生成模板SysOperationRecord
package model

import (
	"encoding/json"
	"fmt"
	"newchat/global"
	"strconv"

	"github.com/go-redis/redis"
)

// 如果含有time.Time 请自行import time包
type Messages_list struct {
	global.GVA_MODEL
	Type       int    `json:"type"  gorm:"comment:1 私聊 ，2 群聊"`
	Friend_id  int    `json:"friend_id" gorm:"comment:好友id"`
	Group_id   int    `json:"group_id"  gorm:"comment:群ID"`
	User_id    int    `json:"user_id"  gorm:"comment:用户id"`
	Unread_num int    `json:"unread_num"  gorm:"comment:备用"`
	Msg_text   string `json:"msg_text"  gorm:"comment:消息内容"`
	Online     bool   `json:"online"  gorm:"comment:是否在线"`
}

type Talk_list struct {
	global.GVA_MODEL
	Is_revoke  int          `json:"is_revoke" gorm:"comment:是否是好友请求"`
	Msg_type   int          `json:"msg_type"  gorm:"comment:消息类型1 文字 5 代码 2 图片"`
	Receive_id int          `json:"receive_id" gorm:"comment:接收信息id"`
	Source     int          `json:"source"  gorm:"comment: //2群 还是 1私聊"`
	User_id    int          `json:"user_id"  gorm:"comment:用户id"`
	Content    string       `json:"content"  gorm:"comment:聊天内容"`
	File       File       `json:"file"  gorm:"comment:文件名称"`
	Code_block []Code_block `json:"code_block"  gorm:"comment:代码片段"`
	Forward    []Forward    `json:"forward"  gorm:"comment:代码片段"`
	Invite     []Invite     `json:"invite"  gorm:"comment:代码片段"`
}

type File struct {
	Id            int    `json:"id"`
	Save_type     int    `json:"save_type"  gorm:"comment:保存类型"`
	Record_id     int    `json:"record_id"  gorm:"comment:消息id"`
	User_id       int    `json:"user_id"  gorm:"comment:用户id"`
	File_source   int    `json:"file_source"  gorm:"comment:群聊，还是私聊"`
	File_type     int    `json:"file_type"  gorm:"comment:消息id"`
	File_size     int    `json:"file_size"  gorm:"comment:文件大小"`
	Original_name string `json:"original_name"  gorm:"comment:文件名称"`
	File_suffix   string `json:"file_suffix"  gorm:"comment:文件扩展名"`
	Save_dir      string `json:"save_dir"  gorm:"comment:保存路径"`
	File_url      string `json:"file_url"  gorm:"comment:文件url"`
}

type Code_block struct {
}

type Forward struct {
}

type Invite struct {
}

func Redis_GetMsgNoRead(msg_type, id, friend_id, group_id int) (int, error) {
	result := 0

	msgKey := fmt.Sprintf("%v_%v_%v_%v_%v", global.MsgNoRead, id, msg_type, friend_id, group_id)
	re, err := global.GVA_REDIS.Get(msgKey).Result()
	result, _ = strconv.Atoi(re)
	return result, err
}

func Redis_SetMsgNoRead(status, msg_type, id, friend_id, group_id int) (err error) {

	msgKey := fmt.Sprintf("%v_%v_%v_%v_%v", global.MsgNoRead, id, msg_type, friend_id, group_id)

	if status == 1 {
		err = global.GVA_REDIS.Incr(msgKey).Err()
		if err != nil {
			return err
		}
	} else {
		err = global.GVA_REDIS.Set(msgKey, "0", -1).Err()
	}

	return nil
}

func Redis_GetFileById(id int) (err error, file File) {

	redisKey := "newchat:file_id:" + fmt.Sprintf("%d", id)
	//优先查询redis 拿map
	filejson, err := global.GVA_REDIS.Get(redisKey).Result()
	if err == redis.Nil {
		// 查询数据库 得 map
		global.GVA_DB.Table("file").Where("record_id = ?", id).Scan(&file)

		rejson, _ := json.Marshal(file)

		err = global.GVA_REDIS.Set(redisKey, string(rejson), -1).Err()
		if err != nil {
			return err, file
		}

		//新增无序集合 所有的key头存在无序集合里面
		err = global.GVA_REDIS.SAdd(global.FileId, redisKey).Err()
		if err != nil {
			return err, file
		}

		filejson, err = global.GVA_REDIS.Get(redisKey).Result()
		if err != nil {
			return err, file
		}
	}

	if err != nil {
		return err, file
	}

	err = json.Unmarshal([]byte(filejson), &file)
	return err, file
}

//清理缓存
func Delcash(head string) error {

	//拿到key头在集合中的数量
	num, err := global.GVA_REDIS.SCard(head).Result()
	if err != nil {
		return err
	}
	var i int64
	for i = 0; i < num; i++ {

		//删除一条数据返回被删除的元素，逐个删除，但这个会返回对应元素
		red_key, err := global.GVA_REDIS.SPop(head).Result()

		if err != nil {
			return err
		}

		if global.GVA_REDIS.Del(red_key).Err() != nil {
			return err
		}

	}

	return err
}
