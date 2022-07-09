package service

import (
	"errors"
	"time"

	"github.com/wujeevan/douyinv0/repository"
	"github.com/wujeevan/douyinv0/utils"
)

type SignInfo struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type SignFlow struct {
	UserId   int64
	Username string
	Password string
	Token    string
	User     *repository.User
	signInfo *SignInfo
}

func SignUp(username, password string) (*SignInfo, error) {
	return NewSignUpFlow(username, password).DoSignUp()
}

func SignIn(username, password string) (*SignInfo, error) {
	return NewSignUpFlow(username, password).DoSignIn()
}

func NewSignUpFlow(username, password string) *SignFlow {
	return &SignFlow{
		Username: username,
		Password: password,
	}
}

func NewSignInFlow(username, password string) *SignFlow {
	return &SignFlow{
		Username: username,
		Password: password,
	}
}

func (f *SignFlow) DoSignUp() (*SignInfo, error) {
	if err := f.CheckSignUp(); err != nil {
		return nil, err
	}
	if err := f.PrepareSignUp(); err != nil {
		return nil, err
	}
	if err := f.PackSignInfo(); err != nil {
		return nil, err
	}
	return f.signInfo, nil
}

func (f *SignFlow) DoSignIn() (*SignInfo, error) {
	if err := f.CheckSignIn(); err != nil {
		return nil, err
	}
	if err := f.PrepareSignIn(); err != nil {
		return nil, err
	}
	if err := f.PackSignInfo(); err != nil {
		return nil, err
	}
	return f.signInfo, nil
}

func (f *SignFlow) CheckSignUp() error {
	if err := utils.CheckSqlInjection(f.Username); err != nil {
		return errors.New("username contains invalid char")
	}
	if len(f.Username) <= 1 || len(f.Username) > 20 {
		return errors.New("username length is invalid")
	}
	if len(f.Password) <= 5 || len(f.Password) > 30 {
		return errors.New("password length is invalid")
	}
	user, err := repository.QueryUserByUsername(f.Username)
	if err != nil {
		return err
	}
	if user.ID != 0 {
		return errors.New("username alreadly exists")
	}
	return nil
}

func (f *SignFlow) PrepareSignUp() error {
	f.Token = utils.GenerateToken()
	user := &repository.User{
		Username:        f.Username,
		Password:        f.Password,
		Token:           f.Token,
		TokenCreateTime: time.Now(),
	}
	err := repository.CreateUser(user)
	if err != nil {
		return err
	}
	f.UserId = user.ID
	f.Token = utils.GenerateToken()
	if err := repository.UpdateUserToken(f.UserId, f.Token, utils.TokenValidTime); err != nil {
		return err
	}
	return nil
}

func (f *SignFlow) CheckSignIn() error {
	if err := utils.CheckSqlInjection(f.Username); err != nil {
		return errors.New("username contains invalid char")
	}
	user, err := repository.QueryUserByUsername(f.Username)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return errors.New("username not exists")
	}
	if user.Password != f.Password {
		return errors.New("password is wrong")
	}
	f.User = user
	return nil
}

func (f *SignFlow) PrepareSignIn() error {
	f.UserId = f.User.ID
	f.Token = utils.GenerateToken()
	if err := repository.UpdateUserToken(f.UserId, f.Token, utils.TokenValidTime); err != nil {
		return err
	}
	return nil
}

func (f *SignFlow) PackSignInfo() error {
	f.signInfo = &SignInfo{
		UserId: f.UserId,
		Token:  f.Token,
	}
	return nil
}
