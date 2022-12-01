package controller

import (
	"github.com/gin-gonic/gin"
	"myself/logger"
	"myself/logic"
	"myself/module"
	"net/http"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	communityList := make([]module.Community, 0, 10)
	err := logic.GetCommunityList(&communityList)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	c.JSON(http.StatusOK, communityList)
	return
}
func CommunityIDHandler(c *gin.Context) {
	id := c.Param("id")
	ID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeInvalidParam)
	}
	data, err := logic.GetCommunityByID(ID)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
