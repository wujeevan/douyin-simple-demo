package service

import (
	"github.com/wujeevan/douyinv0/repository"
)

type FeedVideo struct {
	NextTime  int64
	VideoList []*repository.Video
}

type UserVideo = []*repository.Video

func QueryFeedVideo(latestTime int64, token string) (*FeedVideo, error) {
	return NewQueryFeedVideoFlow(latestTime, token).Do()
}

func QueryUserVideo(token string) (UserVideo, error) {
	user, err := QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	videos, err := repository.QueryUserVideoById(user.ID)
	if err != nil {
		return nil, err
	}
	return videos, err
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
	userId     int64
	videoList  []*repository.Video
	feedVideo  *FeedVideo
}

func (f *QueryFeedVideoFlow) Do() (*FeedVideo, error) {
	if err := f.CheckParam(); err != nil {
		return nil, err
	}
	if err := f.PrepareFeedVideo(); err != nil {
		return nil, err
	}
	if err := f.PackFeedVideo(); err != nil {
		return nil, err
	}
	return f.feedVideo, nil
}

func (f *QueryFeedVideoFlow) CheckParam() error {
	user, err := QueryUserByToken(f.token)
	if f.token == "" {
		f.userId = 0
	} else if err != nil {
		return err
	} else {
		f.userId = user.ID
	}
	return nil
}

func (f *QueryFeedVideoFlow) PrepareFeedVideo() error {
	var err error
	f.videoList, err = repository.QueryFeedVideo(f.latestTime, f.userId)
	if err != nil {
		return err
	}
	if len(f.videoList) > 0 {
		f.nextTime = f.videoList[len(f.videoList)-1].CreateTime.UnixMilli()
	}
	return nil
}

func (f *QueryFeedVideoFlow) PackFeedVideo() error {
	f.feedVideo = &FeedVideo{
		NextTime:  f.nextTime,
		VideoList: f.videoList,
	}
	return nil
}
