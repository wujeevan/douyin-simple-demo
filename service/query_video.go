package service

import (
	"errors"
	"strings"

	"github.com/wujeevan/douyinv0/repository"
)

type FeedVideo struct {
	nextTime  int64
	videoList []*repository.Video
}

func QueryFeedVideo(latestTime int64, token string) (*FeedVideo, error) {
	return NewQueryFeedVideoFlow(latestTime, token).Do()
}

func NewQueryFeedVideoFlow(latestTime int64, token string) *QueryFeedVideoFlow {
	return &QueryFeedVideoFlow{
		latestTime: latestTime,
		token:      token,
		nextTime:   0,
	}
}

type QueryFeedVideoFlow struct {
	latestTime int64
	token      string
	nextTime   int64
	videoList  []*repository.Video
	feedVideo  *FeedVideo
}

func (f *QueryFeedVideoFlow) Do() (*FeedVideo, error) {
	if err := f.CheckParam(); err != nil {
		return nil, nil
	}
	if err := f.PrepareFeedVideo(); err != nil {
		return nil, nil
	}
	if err := f.PackFeedVideo(); err != nil {
		return nil, nil
	}
	return f.feedVideo, nil
}

func (f *QueryFeedVideoFlow) CheckParam() error {
	//TODO: 检查时间是否正常，token是否带有特殊字符防注入
	if strings.ContainsAny(f.token, "'<>&*") {
		return errors.New("the token is invalid")
	}
	return nil
}

func (f *QueryFeedVideoFlow) PrepareFeedVideo() error {
	var err error
	f.videoList, err = repository.QueryFeedVideo(f.latestTime, f.token)
	if err != nil {
		return err
	}
	if len(f.videoList) > 0 {
		f.nextTime = f.videoList[len(f.videoList)-1].CreateTime.UnixMilli()
	}
	return err
}

func (f *QueryFeedVideoFlow) PackFeedVideo() error {
	f.feedVideo = &FeedVideo{
		nextTime:  f.nextTime,
		videoList: f.videoList,
	}
	return nil
}
