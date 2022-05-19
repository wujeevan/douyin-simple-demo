package service

import (
	"errors"

	"github.com/wujeevan/douyinv0/repository"
)

func ProcessFavoriteVideo(token string, videoId, action_type int64) error {
	user, err := QueryUserByToken(token)
	if err != nil {
		return err
	}
	if action_type == 1 {
		if repository.QueryIsFavoriteVideo(user.ID, videoId) {
			return nil
		}
		if err := repository.DoFavoriteVideo(user.ID, videoId); err != nil {
			return err
		}
	} else if action_type == 2 {
		if !repository.QueryIsFavoriteVideo(user.ID, videoId) {
			return nil
		}
		if err := repository.NotFavoriteVideo(user.ID, videoId); err != nil {
			return err
		}
	} else {
		return errors.New("action type is invalid")
	}
	return nil
}

func QueryFavoriteVideo(token string) ([]*repository.Video, error) {
	user, err := QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	videoList, err := repository.QueryFavoriteVideo(user.ID)
	if err != nil {
		return nil, err
	}
	return videoList, nil
}
