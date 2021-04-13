package model

import "newchat/global"

type Emoticon struct {
	global.GVA_MODEL
	User_id     int    `json:"user_id"  gorm:"comment:用户id"`
	Status      int    `json:"status"  gorm:"comment:状态"`
	Src         string `json:"src" gorm:"comment:链接"`
	Name        string `json:"name" gorm:"comment:名称"`
	Size        int    `json:"size" gorm:"comment:大小"`
	File_suffix string `json:"file_suffix" gorm:"comment:后缀"`
	Save_dir    string `json:"save_dir" gorm:"comment:路径"`
}
