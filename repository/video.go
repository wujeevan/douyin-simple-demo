package repository

import "time"

type Video struct {
	ID            int64     `json:"id"`
	User          User      `json:"author"`
	UserID        int64     `json:"-"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count" gorm:"default:0"`
	IsFavorite    bool      `json:"is_favorite"`
	CommentCount  int64     `json:"comment_count" gorm:"default:0"`
	CreateTime    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	Status        bool      `json:"-" gorm:"default:1"`
}

func QueryFeedVideo(latestTime int64, token string) ([]*Video, error) {
	var videos []*Video
	err := db.Preload("User").Where("replace(unix_timestamp(create_time),'.','') <= ?", latestTime).Order("create_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		video.IsFavorite = false
		video.User.IsFollow = false
	}
	return videos, nil
}
