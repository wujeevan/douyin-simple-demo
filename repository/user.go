package repository

import "time"

type User struct {
	ID            int64
	Username      string
	Password      string
	Token         string `gorm:"default:abcdefg"`
	IsFollow      bool
	FollowCount   int64     `gorm:"default:0"`
	FollowerCount int64     `gorm:"default:0"`
	CreateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status        bool      `gorm:"default:1"`
}

func QueryUserById(id int64) (*User, error) {
	var user User
	err := db.Preload("Videos").Where("id = ?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
