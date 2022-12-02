package logic

import (
	"myself/dao/mysql"
	"myself/module"
	"myself/pkg/snowflake"
)

func CreatePost(pos *module.Post) error {
	pos.Post_id = snowflake.GetID()
	err := mysql.CreatePost(pos)
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
func GetPostsHandler(page, size int64) (data []*module.ApiPostDetail, err error) {
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
