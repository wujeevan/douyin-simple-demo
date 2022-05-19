package video

import "time"

type Video struct {
	ID            int64
	UserID        int64
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64     `gorm:"default:0"`
	CommentCount  int64     `gorm:"default:0"`
	CreateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdateTime    time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
	Status        bool      `gorm:"default:1"`
}
