package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
)

type FeedVideoResponse struct {
	Code      int         `json:"status_code"`
	Msg       string      `json:"status_msg"`
	VideoList interface{} `json:"video_list"`
	NextTime  int64       `json:"next_time"`
}

func QueryFeedVideo(ctx *gin.Context) {
	latest_time_ := ctx.Query("latest_time")
	token := ctx.Query("token")
	latestTime, err := strconv.ParseInt(latest_time_, 10, 64)
	if err != nil {
		ctx.JSON(200, &FeedVideoResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	feedVideo, err := service.QueryFeedVideo(latestTime, token)
	if err != nil {
		ctx.JSON(200, &FeedVideoResponse{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	for _, video := range feedVideo.VideoList {
		video.PlayUrl = "http://" + ctx.Request.Host + video.PlayUrl
		video.CoverUrl = "http://" + ctx.Request.Host + video.CoverUrl
	}
	ctx.JSON(200, &FeedVideoResponse{
		Code:      0,
		Msg:       "success",
		NextTime:  feedVideo.NextTime,
		VideoList: feedVideo.VideoList,
	})
}
