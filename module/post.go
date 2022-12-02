package module

import "time"

type Post struct {
	Status       int32     `json:"status" db:"status" `
	Post_id      int64     `json:"id" db:"post_id"`
	Author_id    int64     `json:"author_id" db:"author_id"`
	Community_id int64     `json:"community_id" db:"community_id"`
	Title        string    `json:"title" db:"title" binding:"required"`
	Content      string    `json:"content" db:"content" binding:"required"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	Votes      int64 `json:"votes"`
	*Post      `json:"post"`
	*UserInfo  `json:"user_info"`
	*Community `json:"community_detail"`
}
