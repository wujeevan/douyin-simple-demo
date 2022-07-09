package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/handler"
	"github.com/wujeevan/douyinv0/repository"
)

func main() {
	if err := repository.Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.TrustedPlatform = gin.PlatformGoogleAppEngine
	r.Static("/upload", "./upload")
	rg := r.Group("/douyin")
	rg.GET("/feed/", handler.QueryFeedVideo)

	rg.POST("/user/register/", handler.SignUp)
	rg.POST("/user/login/", handler.SignIn)
	rg.GET("/user/", handler.QueryUserInfo)

	rg.GET("/publish/list/", handler.QueryUserVideo)
	rg.POST("/publish/action/", handler.UploadVideo)

	rg.POST("/favorite/action/", handler.ProcessFavoriteVideo)
	rg.GET("/favorite/list/", handler.QueryFavoriteVideo)

	rg.POST("/comment/action/", handler.ProcessComment)
	rg.GET("/comment/list/", handler.QueryCommentList)

	rg.POST("/relation/action/", handler.ProcessFollowUser)
	rg.GET("/relation/follow/list/", handler.QueryFollowList)
	rg.GET("/relation/follower/list/", handler.QueryFollowerList)

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
