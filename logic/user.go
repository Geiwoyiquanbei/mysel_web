package logic

import (
	"golang.org/x/crypto/bcrypt"
	"myself/dao/mysql"
	"myself/logger"
	"myself/module"
	"myself/pkg/JWt"
	"myself/pkg/snowflake"
)

func SignUp(p *module.ParamSignUp) (err error) {
	if mysql.QueryByName(p.UserName) != nil {
		return err
	}
	p.Password, err = HashAndSalt([]byte(p.Password))
	if err != nil {
		logger.Log.Error("密码加密失败")
		return err
	}
	var u = module.User{
		User_id:  snowflake.GetID(),
		Username: p.UserName,
		Password: p.Password,
	}
	err = mysql.InsertUser(&u)
	if err != nil {
		logger.Log.Error("新用户插入失败")
		return err
	}
	return nil
}

func LogIn(p *module.ParamLogIn) error {
	//获取数据库的该用户的密码
	var u = &module.User{
		Username: p.Username,
		Password: p.Password,
	}
	err := mysql.GetUser(u)
	if err != nil {
		logger.Log.Error(err)
		return err
	}
	if ValidatePasswords(u.Password, []byte(p.Password)) == false {
		return mysql.ErrorInvalidPassword
	}
	p.UserID = u.User_id
	//返回token
	token, rToken, err := JWt.GenToken2(u.User_id, u.Username, u.Password)
	p.Token = token
	p.Rtoken = rToken
	return nil
}
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//验证密码
func ValidatePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}
	return true
}
