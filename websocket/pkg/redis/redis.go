package redis

import (
	"log"

	"github.com/woodylan/go-websocket/pkg/setting"

	"github.com/go-redis/redis/v7"
)

var RedisDB *redis.Client

var UserId = "user_id" 

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
