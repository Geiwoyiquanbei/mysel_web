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
