package service

import (
	"path"

	"github.com/wujeevan/douyinv0/repository"
	"github.com/wujeevan/douyinv0/utils"
)

type UploadVideoFlow struct {
	Token    string
	Filepath string
	video    repository.Video
}

func UploadVideo(token, filename string) error {
	return NewUploadVideoFlow(token, filename).Do()
}

func NewUploadVideoFlow(token, filepath string) *UploadVideoFlow {
	return &UploadVideoFlow{
		Token:    token,
		Filepath: filepath,
	}
}

func (f *UploadVideoFlow) Do() error {
	if err := f.CheckParam(); err != nil {
		return err
	}
	if err := f.PrepareVideo(); err != nil {
		return err
	}
	return nil
}

func (f *UploadVideoFlow) CheckParam() error {
	user, err := QueryUserByToken(f.Token)
	if err != nil {
		return err
	}
	f.video.UserID = user.ID
	return nil
}

func (f *UploadVideoFlow) PrepareVideo() error {
	f.video.PlayUrl = path.Join("/upload", f.Filepath)
	coverUrl, err := utils.GenerateVideoCover(f.video.PlayUrl)
	if err != nil {
		return err
	}
	f.video.CoverUrl = coverUrl
	if err := repository.CreateVideo(&f.video); err != nil {
		return err
	}
	return nil
}
