package controller

import (
	"strconv"

	"github.com/wujeevan/douyinv0/service"
)

type FeedVideoResponse struct {
	code      int         `json:"status_code"`
	msg       string      `json:"status_msg"`
	nextTime  int64       `json:"next_time"`
	videoList interface{} `json:"video_list"`
}

func QueryFeedVideo(latestTimeStr, token string) *FeedVideoResponse {
	latestTime, err := strconv.ParseInt(latestTimeStr, 10, 64)
	if err != nil {
		return &FeedVideoResponse{
			code: -1,
			msg:  err.Error(),
		}
	}
	latestTime /= 1000
	feedVideo, err := service.QueryFeedVideo(latestTime, token)
	if err != nil {
		return &FeedVideoResponse{
			code: -1,
			msg:  err.Error(),
		}
	}
	return &FeedVideoResponse{
		code:      0,
		msg:       "success",
		nextTime:  0,
		videoList: feedVideo,
		// nextTime:  feedVideo.nextTime,
		// videoList: feedVideo.videoList,
	}
}
