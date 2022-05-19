package handler

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wujeevan/douyinv0/service"
	"github.com/wujeevan/douyinv0/utils"
)

type FeedVideoResponse struct {
	Code      int         `json:"status_code"`
	Msg       string      `json:"status_msg"`
	VideoList interface{} `json:"video_list"`
	NextTime  int64       `json:"next_time"`
}

type UserVideoResponse struct {
	Code      int         `json:"status_code"`
	Msg       string      `json:"status_msg"`
	VideoList interface{} `json:"video_list"`
}

type NULL struct{}

func QueryFeedVideo(ctx *gin.Context) {
	latest_time_ := ctx.Query("latest_time")
	token := ctx.Query("token")
	latestTime, err := strconv.ParseInt(latest_time_, 10, 64)
	if err != nil {
		latestTime = time.Now().UnixMicro()
	}
	feedVideo, err := service.QueryFeedVideo(latestTime, token)
	if err != nil {
		ctx.JSON(200, &FeedVideoResponse{
			Code:      0,
			Msg:       err.Error(),
			VideoList: []NULL{}, //防止客户端空指针异常，列表类型都应返回空列表，而不是空指针
		})
		return
	}
	for _, video := range feedVideo.VideoList {
		video.PlayUrl = utils.AddHostName(ctx.Request.Host, video.PlayUrl)
		video.CoverUrl = utils.AddHostName(ctx.Request.Host, video.CoverUrl)
	}
	ctx.JSON(200, &FeedVideoResponse{
		Code:      0,
		Msg:       "success",
		NextTime:  feedVideo.NextTime,
		VideoList: feedVideo.VideoList,
	})
}

func QueryUserVideo(ctx *gin.Context) {
	token := ctx.Query("token")
	userVideo, err := service.QueryUserVideo(token)
	if err != nil {
		ctx.JSON(200, &UserVideoResponse{
			Code:      -1,
			Msg:       err.Error(),
			VideoList: []NULL{},
		})
		return
	}
	for _, video := range userVideo {
		video.PlayUrl = utils.AddHostName(ctx.Request.Host, video.PlayUrl)
		video.CoverUrl = utils.AddHostName(ctx.Request.Host, video.CoverUrl)
	}
	ctx.JSON(200, &UserVideoResponse{
		Code:      0,
		Msg:       "success",
		VideoList: userVideo,
	})
}
