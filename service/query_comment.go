package service

import (
	"errors"

	"github.com/wujeevan/douyinv0/repository"
)

func ProcessComment(videoId, action_type int64, token, content string) error {
	user, err := QueryUserByToken(token)
	if err != nil {
		return err
	}
	if action_type == 1 {
		if err := CreateComment(user.ID, videoId, content); err != nil {
			return err
		}
	} else if action_type == 2 {
		if err := DropComment(user.ID, videoId); err != nil {
			return err
		}
	} else {
		return errors.New("action type is invalid")
	}
	return nil
}

func CreateComment(userId, videoId int64, content string) error {
	comment := &repository.Comment{
		UserID:  userId,
		VideoID: videoId,
		Content: content,
	}
	if err := CheckComment(content); err != nil {
		return err
	}
	if err := repository.CreateComment(comment); err != nil {
		return err
	}
	return nil
}

func DropComment(userId, videoId int64) error {
	comment := &repository.Comment{
		UserID:  userId,
		VideoID: videoId,
	}
	if err := repository.DropComment(comment); err != nil {
		return err
	}
	return nil
}

func QueryCommentList(token string, videoId int64) ([]*repository.Comment, error) {
	user, err := QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	commentList, err := repository.QueryCommentList(user.ID, videoId)
	if err != nil {
		return nil, err
	}
	return commentList, nil
}

func CheckComment(comment string) error {
	if len(comment) == 0 || len(comment) > 500 {
		return errors.New("comment length is invalid")
	}
	return nil
}
