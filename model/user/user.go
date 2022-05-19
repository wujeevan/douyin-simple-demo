package user

import "time"

type User struct {
	ID              int64
	Username        string
	Password        string
	Token           string
	TokenCreateTime time.Time
	FollowCount     int64     `gorm:"default:0"`
	FollowerCount   int64     `gorm:"default:0"`
	CreateTime      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status          bool      `gorm:"default:1"`
}
