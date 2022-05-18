package service

import (
	"errors"

	"github.com/wujeevan/douyinv0/repository"
)

func ProcessFollowUser(token string, action_type, followedUserId int64) error {
	user, err := QueryUserByToken(token)
	if err != nil {
		return err
	}
	if action_type == 1 {
		if err := repository.DoFollowUser(user.ID, followedUserId); err != nil {
			return err
		}
	} else if action_type == 2 {
		if err := repository.NotFollowUser(user.ID, followedUserId); err != nil {
			return err
		}
	} else {
		return errors.New("action_type is invalid")
	}
	return nil
}

func QueryFollowList(token string) ([]*repository.User, error) {
	user, err := QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	followList, err := repository.QueryFollowList(user.ID)
	if err != nil {
		return nil, err
	}
	return followList, nil
}

func QueryFollowerList(token string) ([]*repository.User, error) {
	user, err := QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	followerList, err := repository.QueryFollowerList(user.ID)
	if err != nil {
		return nil, err
	}
	return followerList, nil
}
