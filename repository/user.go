package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              int64     `json:"id"`
	Username        string    `json:"name"`
	Password        string    `json:"-"`
	Token           string    `json:"-"`
	TokenCreateTime time.Time `json:"-"`
	Avatar          string    `json:"avatar"`
	Signature       string    `json:"signature"`
	BackgroundImage string    `json:"background_image"`
	FollowCount     int64     `json:"follow_count" gorm:"default:0"`
	FollowerCount   int64     `json:"follower_count" gorm:"default:0"`
	IsFollow        bool      `json:"is_follow" gorm:"-"`
	CreateTime      time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime      time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	Status          bool      `json:"-" gorm:"default:1"`
}

func QueryUserById(id int64) (*User, error) {
	var user User
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByUsername(username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func QueryUserByToken(token string) (*User, error) {
	var user User
	err := db.Where("token = ?", token).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return &user, nil
}

func QueryUserByToken_(token string) (*User, error) {
	userIdStr, err := rdb.Get(context.TODO(), token).Result()
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return nil, err
	}
	return QueryUserById(userId)
}

func UpdateUserToken(userId int64, token string, tokenValidTime time.Duration) error {
	user := User{
		ID:              userId,
		Token:           token,
		TokenCreateTime: time.Now(),
	}
	err := db.Model(&user).Updates(user).Error
	if err != nil {
		return err
	}
	rdb.Set(context.TODO(), token, userId, tokenValidTime)
	return nil
}

func UpdateUserFollowCount(userId, delta int64) error {
	if err := db.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", delta)).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserFollowerCount(userId, delta int64) error {
	if err := db.Model(&User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count + ?", delta)).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
