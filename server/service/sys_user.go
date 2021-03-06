package service

import (
	"errors"
	"newchat/global"
	"newchat/model"
	"newchat/model/request"
	"newchat/model/response"
	"newchat/utils"
	"strconv"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(global.GVA_DB.Where("mobile = ?", u.Mobile).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("手机号已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	u.ID, _ = strconv.Atoi(utils.Random("number", 4))
	err = global.GVA_DB.Create(&u).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Debug().Where("mobile = ? AND password = ?", u.Mobile, u.Password).First(&user).Error
	return err, &user
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户密码
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = global.GVA_DB.Where("id = ? AND password = ?", u.ID, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ChangePassword
//@description: 修改用户信息
//@param: u *model.SysUser, newPassword string
//@return: err error, userInter *model.SysUser

func EditUserDetail(id int, rep map[string]interface{}) (err error) {
	err = global.GVA_DB.Table("sys_user").Where("id = ?", id).Update(rep).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUserInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
	return err, userList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUser
//@description: 删除用户
//@param: id float64
//@return: err error

func DeleteUser(id int) (err error) {
	var user model.SysUser
	err = global.GVA_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserInfo
//@description: 设置用户信息
//@param: reqUser model.SysUser
//@return: err error, user model.SysUser

func SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	err = global.GVA_DB.Updates(&reqUser).Error
	return err, reqUser
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = global.GVA_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserByUuid
//@description: 通过uuid获取用户信息
//@param: uuid string
//@return: err error, user *model.SysUser

func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = global.GVA_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: FindUserById
//@description: 通过id获取用户信息
//@param: id int
//@return: err error, user *model.SysUser

func SearchUserById(user_id, friend_id int) (err error, user *response.ResponseSearchUser) {
	u := response.ResponseSearchUser{}
	var c model.Contacts
	var a model.ApplyRecords
	err = global.GVA_DB.Debug().Table("sys_user").Where("id = ?", friend_id).Scan(&u).Error
	if u.ID <= 0 {
		return errors.New("无此用户"), &u
	}
	//是好友
	if !errors.Is(global.GVA_DB.Debug().Table("contacts").Where("user_id = ? and friend_id = ? ", user_id, friend_id).Scan(&c).Error, gorm.ErrRecordNotFound) {
		u.Friend_status = 2
		u.Nickname_remark = c.Friend_remark
	} else {
		//不是好友
		u.Friend_status = 1
		//查找是否有已经审请中的记录 , 有审请记录 apply=1
		if !errors.Is(global.GVA_DB.Debug().Table("apply_records").Where("user_id = ? and friend_id = ? and status = 0 ", user_id, friend_id).Scan(&a).Error, gorm.ErrRecordNotFound) {
			u.Friend_apply = 1
		}
		if user_id == friend_id {
			u.Friend_status = 0
			u.Friend_apply = 0
		}
	}

	return err, &u
}
