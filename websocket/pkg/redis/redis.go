package redis

import (
	"log"

	"websocket/pkg/setting"

	"github.com/go-redis/redis/v7"
)

var RedisDB *redis.Client

var UserIdClient = "user_Client"

var UserIdSystem = "user_System"

var UserStatus = "user_Online"

var UserCheck = "user_Check"

func InitRedis() *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     setting.CommonRedis.Host + ":" + setting.CommonRedis.Port,
		Password: setting.CommonRedis.Pwd,
		DB:       setting.CommonRedis.DBName,
	})
	err := client.Ping().Err()
	if err != nil {
		log.Fatalln(err)
	}
	RedisDB = client
	return client
}
