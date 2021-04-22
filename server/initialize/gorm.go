package initialize

import (
	"newchat/global"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
)

//@author: SliverHorn
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

// MysqlTables
//@author: SliverHorn
//@function: MysqlTables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func MysqlTables(db *gorm.DB) {
	// err := db.AutoMigrate(
	// 	//model.SysUser{},
	// 	// model.SysAuthority{},
	// 	// model.SysApi{},
	// 	// model.SysBaseMenu{},
	// 	// model.SysBaseMenuParameter{},
	// 	model.JwtBlacklist{},
	// 	// model.SysDictionary{},
	// 	// model.SysDictionaryDetail{},
	// 	// model.ExaFileUploadAndDownload{},
	// 	// model.ExaFile{},
	// 	// model.ExaFileChunk{},
	// 	// model.ExaSimpleUploader{},
	// 	// model.ExaCustomer{},
	// 	// model.SysOperationRecord{},
	// 	// model.WorkflowProcess{},
	// 	// model.WorkflowNode{},
	// 	// model.WorkflowEdge{},
	// 	// model.WorkflowStartPoint{},
	// 	// model.WorkflowEndPoint{},
	// 	// model.WorkflowMove{},
	// 	// model.ExaWfLeave{},

	// )
	// if err != nil {
	// 	global.GVA_LOG.Error("register table failed", zap.Any("err", err))
	// 	os.Exit(0)
	// }
	global.GVA_LOG.Info("register table success")
}

//
//@author: SliverHorn
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	db, err := gorm.Open("mysql", dsn)
	db.SingularTable(true)
	if err != nil {
		global.GVA_LOG.Error("MySQL启动异常 "+dsn, zap.Any("err", err))
		os.Exit(0)
	} else {
		global.GVA_DB = db
		// global.GVA_DB.DB().SetMaxIdleConns(admin.MaxIdleConns)
		// global.GVA_DB.DB().SetMaxOpenConns(admin.MaxOpenConns)
		global.GVA_DB.LogMode(false)
	}
	return db
}

//@author: SliverHorn
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

// func gormConfig(mod bool) *gorm.Config {
// 	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
// 	switch global.GVA_CONFIG.Mysql.LogZap {
// 	case "silent", "Silent":
// 		config.Logger = internal.Default.LogMode(logger.Silent)
// 	case "error", "Error":
// 		config.Logger = internal.Default.LogMode(logger.Error)
// 	case "warn", "Warn":
// 		config.Logger = internal.Default.LogMode(logger.Warn)
// 	case "info", "Info":
// 		config.Logger = internal.Default.LogMode(logger.Info)
// 	case "zap", "Zap":
// 		config.Logger = internal.Default.LogMode(logger.Info)
// 	default:
// 		if mod {
// 			config.Logger = internal.Default.LogMode(logger.Info)
// 			break
// 		}
// 		config.Logger = internal.Default.LogMode(logger.Silent)
// 	}
// 	return config
// }
