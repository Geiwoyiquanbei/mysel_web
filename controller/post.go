package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myself/logger"
	"myself/logic"
	"myself/module"
	"net/http"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	p := new(module.ParamPost)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeInvalidParam)
		return
	}
	value, _ := GetCurrentUserID(c)
	fmt.Println(value)
	pos := new(module.Post)
	pos.Author_id = value
	pos.Title = p.Title
	pos.Content = p.Content
	pos.Community_id = p.Community_id
	err = logic.CreatePost(pos)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	ResponseWithSuccess(c, CodeSuccess)
	return
}
func GetPostDetailHandler(c *gin.Context) {
	value := c.Param("id")
	fmt.Println(value)
	id, err2 := strconv.ParseInt(value, 10, 64)
	if err2 != nil {
		logger.Log.Error(err2)
		ResponseWithError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetail(id)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeInvalidParam)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
func GetPostsHandler(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	p, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	s, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	data, err := logic.GetPosts(p, s)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeError)
		return
	}
	c.JSON(http.StatusOK, data)
	return
}
func GetPostListHandler(c *gin.Context) {
	p := &module.ParamPostList{
		Page:  1,
		Size:  10,
		Order: "time",
	}
	err := c.ShouldBindQuery(p)
	if err != nil {
		logger.Log.Error(err)
		ResponseWithError(c, CodeInvalidParam)
		return
	}
	if p.CommunityID == 0 {
		data, err := logic.GetPostList(p.Page, p.Size, p.Order)
		if err != nil {
			logger.Log.Error(err)
			ResponseWithError(c, CodeError)
			return
		}
		ResponseWithMsg(c, CodeSuccess, data)
		return
	} else {
		data, err := logic.GetPostListByCommunity(p.CommunityID, p.Page, p.Size, p.Order)
		if err != nil {
			logger.Log.Error(err)
			ResponseWithError(c, CodeError)
			return
		}
		ResponseWithMsg(c, CodeSuccess, data)
		return
	}
}
