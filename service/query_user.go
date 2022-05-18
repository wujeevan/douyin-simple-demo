package service

import (
	"errors"

	"github.com/wujeevan/douyinv0/repository"
	"github.com/wujeevan/douyinv0/utils"
)

type UserInfo = repository.User

func QueryUserByToken(token string) (*UserInfo, error) {
	if err := utils.CheckSqlInjection(token); err != nil {
		return nil, err
	}
	user, err := repository.QueryUserByToken(token)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("token not exists")
	}
	return user, nil
}
