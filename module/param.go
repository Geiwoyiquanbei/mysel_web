package module

type ParamSignUp struct {
	UserName   string `json:"username"  binding:"required"`
	Password   string `json:"password"  binding:"required"`
	RePassword string `json:"repassword"  binding:"required,eqfield=Password"`
}

type ParamLogIn struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Token    string `json:"token"`
	Rtoken   string `json:"rtoken"`
}
type ParamPost struct {
	Title        string `json:"title" binding:"required"`
	Content      string `json:"content" binding:"required"`
	Community_id int64  `json:"community_id"`
}
type ParamPostVoted struct {
	PostID string `json:"post_id"  binding:"required"`
	Vote   int64  `json:"vote" binding:"oneof=1 0 -1 "`
}
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
