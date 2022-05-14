
package userfollow

type UserFollow struct{
    ID int64 
    UserID int64 
    FollowedUserID int64 
    Status bool `gorm:"default:1"`
    CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
