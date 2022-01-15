package wotoRaw

import "time"

type PublicUserId int64
type UserPermission int
type PublicGroupCallId string
type MediaModelId string
type ProfilePictureModelId string

// UserInfo struct in wotoRaw is a low level struct.
// It shouldn't be used directly in any package.
// Instead, use the UserInfo struct in `wotoValues` package.
type UserInfo struct {
	UserId      PublicUserId   `json:"user_id" gorm:"primaryKey"`
	PrivateHash string         `json:"private_hash"`
	Email       string         `json:"email"`
	Website     string         `json:"website"`
	AuthKey     string         `json:"auth_key"`
	AccessHash  string         `json:"access_hash"`
	Password    string         `json:"password"`
	Permission  UserPermission `json:"permission"`
	Bio         string         `json:"bio"`
	SourceUrl   string         `json:"source_url"`
	TelegramId  int64          `json:"telegram_id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Username    string         `json:"username"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	IsVirtual   bool           `json:"is_virtual"`
	CreatedBy   PublicUserId   `json:"created_by"`
	cachedTime  time.Time      `json:"-" gorm:"-" sql:"-"`
}
