package repository

import "time"

type Video struct {
	User          User
	ID            int64
	UserID        int64 `json:"-"`
	PlayUrl       string
	CoverUrl      string
	IsFavorite    bool
	FavoriteCount int64     `gorm:"default:0"`
	CommentCount  int64     `gorm:"default:0"`
	CreateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Status        bool      `gorm:"default:1"`
}

func QueryFeedVideo(latestTime int64, token string) ([]*Video, error) {
	var videos []*Video
	err := db.Preload("User").Where("unix_timestamp(update_time) <= ?", latestTime).Order("create_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		video.IsFavorite = false
		video.User.IsFollow = false
	}
	return videos, nil
}
