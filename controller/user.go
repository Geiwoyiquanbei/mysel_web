package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"myself/dao/mysql"
	"myself/logger"
	"myself/logic"
	"myself/module"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	var p module.ParamSignUp
	err := c.ShouldBindJSON(&p)
	if err != nil {
		logger.Log.Info(err)
		ResponseWithError(c, CodeError)
		return
	}
	err = logic.SignUp(&p)
	if err != nil {
		logger.Log.Info(err)
		ResponseWithError(c, CodeError)
		return
	}
	ResponseWithSuccess(c, CodeSuccess)
	return
}
func LogInHandler(c *gin.Context) {
	var p module.ParamLogIn
	err := c.ShouldBindJSON(&p)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	err = logic.LogIn(&p)
	if errors.Is(err, mysql.ErrorInvalidPassword) {
		logger.Log.Error(err)
		ResponseWithMsg(c, CodeError, "用户名或密码不正确")
		return
	}
	if errors.Is(err, mysql.ErrorUSerNotExist) {
		logger.Log.Error(err)
		ResponseWithMsg(c, CodeError, "用户不存在")
		return
	}
	if err != nil {
		ResponseWithError(c, CodeError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":  p.UserID,
		"username": p.Username,
		"token":    p.Token,
		"rtoken":   p.Rtoken,
	})
	return
}
