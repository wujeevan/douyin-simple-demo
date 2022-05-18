package repository

import (
	"time"
)

type UserFollow struct {
	ID             int64
	UserID         int64
	User           User `gorm:""`
	FollowedUserID int64
	Status         bool      `gorm:"default:1"`
	CreateTime     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func QueryIsFollow(userId, followedUserId int64) bool {
	var follow UserFollow
	err := db.Where("user_id = ? and followed_user_id = ?", userId, followedUserId).First(&follow).Error
	if err != nil {
		return false
	}
	return follow.Status
}

func DoFollowUser(userId, followedUserId int64) error {
	follow := &UserFollow{
		UserID:         userId,
		FollowedUserID: followedUserId,
	}
	if err := db.Debug().Where(follow).FirstOrCreate(follow, follow).Error; err != nil {
		return err
	}
	if err := db.Model(follow).Update("status", true).Error; err != nil {
		return err
	}
	if err := UpdateUserFollowCount(userId, 1); err != nil {
		return err
	}
	if err := UpdateUserFollowerCount(followedUserId, 1); err != nil {
		return err
	}
	return nil
}

func NotFollowUser(userId, followedUserId int64) error {
	follow := &UserFollow{
		UserID:         userId,
		FollowedUserID: followedUserId,
	}
	if err := db.Model(follow).Where(follow).Update("status", false).Error; err != nil {
		return err
	}
	if err := UpdateUserFollowCount(userId, -1); err != nil {
		return err
	}
	if err := UpdateUserFollowerCount(followedUserId, -1); err != nil {
		return err
	}
	return nil
}

func QueryFollowList(userId int64) ([]*User, error) {
	var follows []*UserFollow
	if err := db.Where("user_id = ? and status = ?", userId, true).Find(&follows).Error; err != nil {
		return nil, err
	}
	var users []*User
	followedIds := make([]int64, len(follows))
	for i, follow := range follows {
		followedIds[i] = follow.FollowedUserID
	}
	if err := db.Find(&users, followedIds).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func QueryFollowerList(userId int64) ([]*User, error) {
	var follows []*UserFollow
	if err := db.Preload("User").Where("followed_user_id = ? and status = ?", userId, true).Find(&follows).Error; err != nil {
		return nil, err
	}
	users := make([]*User, len(follows))
	for i, follow := range follows {
		users[i] = &follow.User
	}
	return users, nil
}
