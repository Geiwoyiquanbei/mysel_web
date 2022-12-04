package redis

import (
	"github.com/go-redis/redis"
	"myself/logger"
	"strconv"
	"time"
)

func GetPostListByOrder(page, size int64, order string) (ids []string, err error) {
	key := GetKeys(KeyPostScoreZSet)
	if order == "time" {
		key = GetKeys(KeyPostTimeZSet)
	}
	ids, err = client.ZRevRange(key, (page-1)*size, (page-1)*size+size-1).Result()
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return ids, nil
}
func GetPostListByOrder2(page, size int64, key string) (ids []string, err error) {
	ids, err = client.ZRevRange(key, (page-1)*size, (page-1)*size+size-1).Result()
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return ids, nil
}
func GetPostListByCommunity(community_id, page, size int64, order string) (ids []string, err error) {
	orderKey := GetKeys(KeyPostTimeZSet)
	if order == "score" {
		orderKey = GetKeys(KeyPostScoreZSet)
	}
	ckey := GetKeys(KeyCommunitySetPF + strconv.Itoa(int(community_id)))
	key := orderKey + strconv.Itoa(int(community_id))
	if client.Exists(key).Val() < 1 {
		client.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		client.Expire(key, 60*time.Second) //设置超时时间
	}
	return GetPostListByOrder2(page, size, key)
}
