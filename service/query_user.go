package service

import (
	"errors"

	"github.com/wujeevan/douyinv0/repository"
	"github.com/wujeevan/douyinv0/utils"
)

type UserInfo = repository.User

func QueryUserByToken(token string) (*UserInfo, error) {
	// invalidTime := 7 * 24 * time.Hour //token有效期
	if err := utils.CheckSqlInjection(token); err != nil {
		return nil, err
	}
	user, err := repository.QueryUserByToken_(token)
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("token not exists")
	}
	// if user.TokenCreateTime.Add(invalidTime).Before(time.Now()) {
	// 	return nil, errors.New("token is expired, please login in again")
	// }
	return user, nil
}
