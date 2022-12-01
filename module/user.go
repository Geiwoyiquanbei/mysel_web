package module

type User struct {
	Username string `json:"username" db:"username" bind:"required"`
	Password string `json:"password" db:"password" bind:"required" `
	User_id  int64  `json:"user_id" db:"user_id"`
}
