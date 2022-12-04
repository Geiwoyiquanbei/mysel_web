package logic

import (
	"fmt"
	"myself/dao/mysql"
	"myself/dao/redis"
	"myself/logger"
	"myself/module"
	"myself/pkg/snowflake"
)

func CreatePost(pos *module.Post) error {
	pos.Post_id = snowflake.GetID()

	err := mysql.CreatePost(pos)
	err = redis.SavePostTime(pos.Post_id, pos.Community_id)
	if err != nil {
		return err
	}
	return nil
}
func GetPostDetail(id int64) (data *module.Post, err error) {
	data, err = mysql.GetPostDetail(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetPosts(page, size int64) (data []*module.ApiPostDetail, err error) {
	list, err := mysql.GetPostsList(page, size)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(list); i++ {
		user_id := list[i].Author_id
		user_info := mysql.GetUserByID(user_id)
		community_id := list[i].Community_id
		community_info, err := mysql.GetCommunityByID(community_id)
		if err != nil {
			return nil, err
		}
		detail := new(module.ApiPostDetail)
		detail = &module.ApiPostDetail{
			Post:      list[i],
			Community: community_info,
			UserInfo:  user_info,
		}
		data = append(data, detail)
	}
	return data, nil
}
func GetPostList(page, size int64, order string) (data []*module.ApiPostDetail, err error) {
	ids, err := redis.GetPostListByOrder(page, size, order)
	if err != nil {
		return nil, err
	}
	posts, err := mysql.GetPostByOrders(ids)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	votes, err := redis.GetVoteData(ids)
	fmt.Println(votes)
	for index, post := range posts {
		user := mysql.GetUserByID(post.Author_id)
		community, err := mysql.GetCommunityByID(post.Community_id)
		if err != nil {
			return nil, err
		}
		detail := &module.ApiPostDetail{
			UserInfo:  user,
			Votes:     votes[index],
			Post:      post,
			Community: community,
		}
		data = append(data, detail)
	}
	return data, nil
}
func GetPostListByCommunity(community_id, page, size int64, order string) (data []*module.ApiPostDetail, err error) {
	ids, err := redis.GetPostListByCommunity(community_id, page, size, order)
	if len(ids) == 0 {
		return
	}
	if err != nil {
		return nil, err
	}
	posts, err := mysql.GetPostByOrders(ids)
	if err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	votes, err := redis.GetVoteData(ids)
	fmt.Println(votes)
	for index, post := range posts {
		user := mysql.GetUserByID(post.Author_id)
		community, err := mysql.GetCommunityByID(post.Community_id)
		if err != nil {
			return nil, err
		}
		detail := &module.ApiPostDetail{
			UserInfo:  user,
			Votes:     votes[index],
			Post:      post,
			Community: community,
		}
		data = append(data, detail)
	}
	return data, nil
}
