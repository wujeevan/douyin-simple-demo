package repository

import (
	"time"
)

type UserFavoriteVideo struct {
	ID         int64
	UserID     int64
	VideoID    int64
	Status     bool      `gorm:"default:1"`
	CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func QueryIsFavoriteVideo(userId, videoId int64) bool {
	var favorite UserFavoriteVideo
	err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(&favorite).Error
	if err != nil {
		return false
	}
	return favorite.Status
}

func DoFavoriteVideo(userId, videoId int64) error {
	favorite := &UserFavoriteVideo{
		UserID:  userId,
		VideoID: videoId,
	}
	if err := db.Where(favorite).FirstOrCreate(favorite, favorite).Error; err != nil {
		return err
	}
	if err := db.Model(favorite).Update("status", true).Error; err != nil {
		return err
	}
	if err := UpdateVideoFavoriteCount(videoId, 1); err != nil {
		return err
	}
	return nil
}

func NotFavoriteVideo(userId, videoId int64) error {
	favorite := &UserFavoriteVideo{
		UserID:  userId,
		VideoID: videoId,
	}
	if err := db.Model(favorite).Where(favorite).Update("status", false).Error; err != nil {
		return err
	}
	if err := UpdateVideoFavoriteCount(videoId, -1); err != nil {
		return err
	}
	return nil
}
