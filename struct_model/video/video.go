
package video

type Video struct{
    ID int64 
    UserID int64 
    PlayUrl string 
    CoverUrl string 
    FavoriteCount int64 `gorm:"default:0"`
    CommentCount int64 `gorm:"default:0"`
    CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    Status bool `gorm:"default:1"`
}
