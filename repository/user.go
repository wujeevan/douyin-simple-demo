package repository

import "time"

type User struct {
	ID            int64     `json:"id"`
	Username      string    `json:"name"`
	Password      string    `json:"-"`
	Token         string    `json:"-" gorm:"default:abcdefg"`
	FollowCount   int64     `json:"follow_count" gorm:"default:0"`
	FollowerCount int64     `json:"follower_count" gorm:"default:0"`
	IsFollow      bool      `json:"is_follow"`
	CreateTime    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	Status        bool      `json:"-" gorm:"default:1"`
}

func QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Preload("Videos").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
