
package comment

type Comment struct{
    ID int64 
    UserID int64 
    VideoID int64 
    Content string 
    Status bool `gorm:"default:1"`
    CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
