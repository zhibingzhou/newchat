package request

import "newchat/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
