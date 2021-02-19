package request

import "newchat/model"

type {{.StructName}}Search struct{
    model.{{.StructName}}
    PageInfo
}