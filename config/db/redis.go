package dbs

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

//ConnectRedis returns a connection to posgress instance
func ConnectRedis() *redis.Client {
	env := viper.GetString("env")
	client := redis.NewClient(&redis.Options{
		Addr:        viper.GetString(env + ".db.redis.host"),
		Password:    viper.GetString(env + ".db.redis.password"),
		DialTimeout: time.Second * 20,
		DB:          0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}
