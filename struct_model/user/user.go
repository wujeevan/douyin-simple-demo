
package user

type User struct{
    ID int64 
    Username string 
    Password string 
    Token string `gorm:"default:abcdefg"`
    FollowCount int64 `gorm:"default:0"`
    FollowerCount int64 `gorm:"default:0"`
    CreateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    UpdateTime time.Time `gorm:"default:CURRENT_TIMESTAMP"`
    Status bool `gorm:"default:1"`
}
