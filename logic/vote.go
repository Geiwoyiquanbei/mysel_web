package logic

import (
	"myself/dao/redis"
)

func PostVoted(user_id, post_id string, value float64) (err error) {
	if err != nil {
		return err
	}
	err = redis.SavePostVoted(user_id, post_id, value)
	if err != nil {
		return err
	}
	return nil
}
