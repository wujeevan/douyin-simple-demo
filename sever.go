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
	r.GET("/douyin/feed/", handler.QueryFeedVideo)
	r.GET("/douyin/feed", handler.QueryFeedVideo)
	r.POST("/douyin/user/register/", handler.SignUp)
	r.POST("/douyin/user/login/", handler.SignIn)
	r.GET("/douyin/user/", handler.QueryUserInfo)
	r.GET("douyin/publish/list/", handler.QueryUserVideo)
	r.POST("/douyin/publish/action/", handler.UploadVideo)
	r.POST("/douyin/favorite/action/", handler.ProcessFavoriteVideo)
	r.GET("/douyin/favorite/list/", handler.QueryFavoriteVideo)
	r.POST("/douyin/comment/action/", handler.ProcessComment)
	r.GET("/douyin/comment/list/", handler.QueryCommentList)
	r.POST("/douyin/relation/action/", handler.ProcessFollowUser)
	r.GET("/douyin/relation/follow/list/", handler.QueryFollowList)
	r.GET("/douyin/relation/follower/list/", handler.QueryFollowerList)
	r.GET("/upload/:filename", handler.FileSever)
	err := r.Run(":8888")
	if err != nil {
		return
	}
}
