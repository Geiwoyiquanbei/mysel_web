package mysql

import (
	"fmt"
	"myself/module"
	"strconv"
)

func CreatePost(pos *module.Post) error {
	sqlStr := "insert into post (post_id,title,content,author_id,community_id) values (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, pos.Post_id, pos.Title, pos.Content, pos.Author_id, pos.Community_id)
	if err != nil {
		return err
	}
	return nil
}
func GetPostDetail(id int64) (data *module.Post, err error) {
	data = new(module.Post)
	sqlStr := `select post_id,author_id,community_id,title,content,create_time from post where post_id=?`
	err = db.Get(data, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetPostsList(page, size int64) (data []*module.Post, err error) {
	sqlStr := `select post_id ,title ,content ,author_id ,community_id,create_time from post  ORDER BY create_time DESC limit ?,?`
	data = make([]*module.Post, 0, 2)
	db.Select(&data, sqlStr, (page-1)*size, size)
	return data, nil
}
func GetPostByOrders(ids []string) (data []*module.Post, err error) {

	l := len(ids)
	data = make([]*module.Post, 0, l)
	fmt.Println(ids)
	for i := 0; i < l; i++ {
		sqlStr := `select post_id ,title ,content ,author_id ,community_id,create_time from post where post_id = ? `
		post := module.Post{}
		parseInt, err := strconv.ParseInt(ids[i], 10, 64)
		err = db.Get(&post, sqlStr, parseInt)
		if err != nil {
			return nil, err
		}
		tmp := module.Post{}
		tmp = post
		data = append(data, &tmp)
	}
	fmt.Println(data)
	return data, nil
}
