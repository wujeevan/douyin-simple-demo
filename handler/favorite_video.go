package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
)

func ProcessFavoriteVideo(ctx *gin.Context) {
	token := ctx.Query("token")
	videoIdStr := ctx.Query("video_id")
	actionTypeStr := ctx.Query("action_type")
	videoId, err1 := strconv.ParseInt(videoIdStr, 10, 64)
	actionType, err2 := strconv.ParseInt(actionTypeStr, 10, 64)
	if err1 != nil || err2 != nil {
		ctx.JSON(200, &SignInResponse{
			Code: -1,
			Msg:  err1.Error(),
		})
	} else {
		if err := service.ProcessFavoriteVideo(token, videoId, actionType); err != nil {
			ctx.JSON(200, &SignInResponse{
				Code: -1,
				Msg:  err1.Error(),
			})
		} else {
			ctx.JSON(200, &SignInResponse{
				Code: 0,
				Msg:  "success",
			})
		}
	}
}

func QueryFavoriteVideo(ctx *gin.Context) {
	token := ctx.Query("token")
	videoList, err := service.QueryFavoriteVideo(token)
	if err != nil {
		ctx.JSON(200, &SignInResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	for _, video := range videoList {
		video.PlayUrl = "http://" + ctx.Request.Host + video.PlayUrl
		video.CoverUrl = "http://" + ctx.Request.Host + video.CoverUrl
	}
	ctx.JSON(200, &UserVideoResponse{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}
