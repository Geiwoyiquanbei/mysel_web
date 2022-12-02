package mysql

import "myself/module"

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
