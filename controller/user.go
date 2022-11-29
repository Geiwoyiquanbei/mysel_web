package controller

import (
	"github.com/gin-gonic/gin"
	"myself/logger"
	"myself/logic"
	"myself/module"
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
		ResponseWithError(c, CodeError)
		return
	}
	ResponseWithSuccess(c, CodeSuccess)
	return
}
