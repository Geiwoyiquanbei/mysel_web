package mysql

import (
	"errors"
	"myself/logger"
	"myself/module"
)

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUSerNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

func QueryByName(username string) (err error) {
	sqlStr := "select count(user_id)  from user where username = ?"
	count := 0
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	if count > 0 {
		logger.Log.Info("username has existed")
		return ErrorUserExist
	}
	return nil
}
func InsertUser(u *module.User) (err error) {
	sqlStr := "insert into user (user_id,username,password) values (?,?,?)"
	_, err = db.Exec(sqlStr, u.User_id, u.Username, u.Password)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	return nil
}
func GetUser(u *module.User) (err error) {
	sqlStr := "select count(username) from user where username = ?"
	var count = 0
	err = db.Get(&count, sqlStr, u.Username)
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrorUSerNotExist
	}
	sqlStr = "select user_id , password from user where username = ?"
	err = db.Get(u, sqlStr, u.Username)
	if err != nil {
		return err
	}
	return nil
}
func GetUserByID(uid int64) *module.UserInfo {
	sqlStr := `select username,user_id from user where user_id = ?`
	uinfo := new(module.UserInfo)
	db.Get(uinfo, sqlStr, uid)
	return uinfo
}
