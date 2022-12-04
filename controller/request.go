package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserID = "userid"

var ERRorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ERRorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ERRorUserNotLogin
		return
	}
	return
}
