package module

type ParamSignUp struct {
	UserName   string `json:"username"  binding:"required"`
	Password   string `json:"password"  binding:"required"`
	RePassword string `json:"repassword"  binding:"required,eqfield=Password"`
}

type ParamLogIn struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username" bind:"required"`
	Password string `json:"password" bind:"required"`
	Token    string `json:"token"`
	Rtoken   string `json:"rtoken"`
}
