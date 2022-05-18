package repository

import (
	"time"

	"gorm.io/gorm"
)

type Video struct {
	ID            int64     `json:"id"`
	User          User      `json:"author"`
	UserID        int64     `json:"-"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count" gorm:"default:0"`
	IsFavorite    bool      `json:"is_favorite" gorm:"-"`
	CommentCount  int64     `json:"comment_count" gorm:"default:0"`
	CreateTime    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdateTime    time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP(3)"`
	Status        bool      `json:"-" gorm:"default:1"`
}

func QueryFeedVideo(latestTime, userId int64) ([]*Video, error) {
	// it means not logging in when userId == 0
	var videos []*Video
	err := db.Preload("User").Where("replace(unix_timestamp(create_time),'.','') <= ?", latestTime).Order("create_time desc").Limit(30).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	if userId == 0 {
		return videos, nil
	}
	for _, video := range videos {
		video.IsFavorite = QueryIsFavoriteVideo(userId, video.ID)
		video.User.IsFollow = QueryIsFollow(userId, video.User.ID)
	}
	return videos, nil
}

func QueryUserVideoById(userId int64) ([]*Video, error) {
	var videos []*Video
	err := db.Preload("User").Where("user_id = ?", userId).Order("create_time desc").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		video.IsFavorite = QueryIsFavoriteVideo(userId, video.ID)
		video.User.IsFollow = QueryIsFollow(userId, video.User.ID)
	}
	return videos, nil
}

func CreateVideo(video *Video) error {
	if err := db.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideoFavoriteCount(videoId, delta int64) error {
	if err := db.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", delta)).Error; err != nil {
		return err
	}
	return nil
}

func UpdateVideoCommentCount(videoId, delta int64) error {
	if err := db.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", delta)).Error; err != nil {
		return err
	}
	return nil
}

func QueryFavoriteVideo(userId int64) ([]*Video, error) {
	var videos []*Video
	err := db.Preload("User").Joins("inner join user_favorite_video on user_favorite_video.user_id = ? and video.id = user_favorite_video.video_id and user_favorite_video.status = 1", userId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	for _, video := range videos {
		video.IsFavorite = true
		video.User.IsFollow = QueryIsFollow(userId, video.User.ID)
	}
	return videos, nil
}
