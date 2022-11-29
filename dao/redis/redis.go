package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

func RedisInit() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: fmt.Sprintf("%s", viper.GetString("redis.password")),
		DB:       0,
	})
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
func Close() {
	client.Close()
}
