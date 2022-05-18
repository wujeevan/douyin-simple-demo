package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"-"`
	User       User      `json:"user"`
	VideoID    int64     `json:"-"`
	Content    string    `json:"content"`
	Status     bool      `json:"-" gorm:"default:1"`
	CreateTime time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime time.Time `json:"create_date" gorm:"default:CURRENT_TIMESTAMP"`
}

func QueryIsComment(userId, videoId int64) (bool, error) {
	var comment Comment
	err := db.Where("user_id = ? and video_id = ?", userId, videoId).First(&comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return comment.Status, err
}

func CreateComment(comment *Comment) error {
	comment_ := &Comment{
		UserID:  comment.UserID,
		VideoID: comment.VideoID,
	}
	if err := db.Where(comment_).FirstOrCreate(comment_, comment).Error; err != nil {
		return err
	}
	if err := db.Model(comment_).Updates(comment).Error; err != nil {
		return err
	}
	if err := UpdateVideoCommentCount(comment.VideoID, 1); err != nil {
		return err
	}
	return nil
}

func DropComment(comment *Comment) error {
	if err := db.Model(comment).Where(comment).Update("status", 0).Error; err != nil {
		return err
	}
	if err := UpdateVideoCommentCount(comment.VideoID, -1); err != nil {
		return err
	}
	return nil
}

func QueryCommentList(user_id, video_id int64) ([]*Comment, error) {
	var comments []*Comment
	err := db.Preload("User").Where("video_id = ? and status = ?", video_id, true).Order("update_time desc").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	for _, comment := range comments {
		comment.User.IsFollow = QueryIsFollow(user_id, comment.UserID)
	}
	return comments, nil
}
