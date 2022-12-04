package controller

import (
	"github.com/gin-gonic/gin"
	"myself/logger"
	"myself/logic"
	"myself/module"
	"strconv"
)

func VoteHandler(c *gin.Context) {
	p := new(module.ParamPostVoted)
	err := c.ShouldBindJSON(p)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeInvalidParam)
		return
	}
	id, err := GetCurrentUserID(c)
	itoa := strconv.Itoa(int(id))
	err = logic.PostVoted(itoa, p.PostID, float64(p.Vote))
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	ResponseWithSuccess(c, CodeSuccess)
}
