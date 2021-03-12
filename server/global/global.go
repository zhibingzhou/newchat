package global

import (
	"go.uber.org/zap"

	"newchat/config"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG *zap.Logger
)

var (
	UserId    = "user_id"
	MsgNoRead = "no_read_msg"
	FileId    = "file_id"
)

var (
	UserStatus   = "user_Online" //websocekt
	UserIdSystem = "user_System" //websocket
)
