package router

import (
	"github.com/gin-gonic/gin"
	"myself/controller"
	"myself/middleware"
	"net/http"
	"time"
)

func Setup() (r *gin.Engine, err error) {
	r = gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong!")
	})
	r.Use(middleware.RareLimitMiddleware(time.Second, 10))
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LogInHandler)
	v := r.Group("api/talk")
	v.Use(middleware.JWTAuthMiddleware())
	{
		v.GET("/community", controller.CommunityHandler)
		v.GET("/community/:id", controller.CommunityIDHandler)
		v.POST("/post", controller.CreatePostHandler)
		v.GET("/post/:id", controller.GetPostDetailHandler)
		v.GET("/posts", controller.GetPostsHandler)
		v.POST("/vote", controller.VoteHandler)
		v.GET("/posts2", controller.GetPostListHandler)
	}
	return r, nil
}
