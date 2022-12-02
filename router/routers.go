package router

import (
	"github.com/gin-gonic/gin"
	"myself/controller"
	"myself/middleware"
	"net/http"
)

func Setup() (r *gin.Engine, err error) {
	r = gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong!")
	})
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
	}
	return r, nil
}
