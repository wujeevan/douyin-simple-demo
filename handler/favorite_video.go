package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
	"github.com/wujeevan/douyinv0/utils"
)

func ProcessFavoriteVideo(ctx *gin.Context) {
	token := ctx.Query("token")
	videoIdStr := ctx.Query("video_id")
	actionTypeStr := ctx.Query("action_type")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		SendFailResponse(ctx, err)
		return
	}
	if err := service.ProcessFavoriteVideo(token, videoId, actionType); err != nil {
		SendFailResponse(ctx, err)
	} else {
		SendSuccessResponse(ctx)
	}
}

func QueryFavoriteVideo(ctx *gin.Context) {
	token := ctx.Query("token")
	videoList, err := service.QueryFavoriteVideo(token)
	if err != nil {
		ctx.JSON(200, &UserVideoResponse{
			Code:      -1,
			Msg:       err.Error(),
			VideoList: []NULL{},
		})
		return
	}
	for _, video := range videoList {
		video.PlayUrl = utils.AddHostName(ctx.Request.Host, video.PlayUrl)
		video.CoverUrl = utils.AddHostName(ctx.Request.Host, video.CoverUrl)
	}
	ctx.JSON(200, &UserVideoResponse{
		Code:      0,
		Msg:       "success",
		VideoList: videoList,
	})
}
