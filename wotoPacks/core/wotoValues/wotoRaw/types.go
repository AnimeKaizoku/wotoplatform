package wotoRaw

import "time"

type (
	PublicUserId          int64
	UserPermission        int
	PublicGroupId         string
	MediaModelId          string
	GenreId               uint32
	CompanyId             uint32
	AuthorId              uint32
	ProfilePictureModelId string
)

// TODO: move this type to ssg package
type MetaDataProvider interface {
	Get(key string) (string, error)
	GetInt(key string) (int, error)
	GetInt8(key string) (int8, error)
	GetInt16(key string) (int16, error)
	GetInt32(key string) (int32, error)
	GetInt64(key string) (int64, error)
	GetUInt(key string) (uint, error)
	GetUInt8(key string) (uint8, error)
	GetUInt16(key string) (uint16, error)
	GetUInt32(key string) (uint32, error)
	GetUInt64(key string) (uint64, error)
	GetBool(key string) (bool, error)
}

// UserInfo struct in wotoRaw is a low level struct.
// It shouldn't be used directly in any package.
// Instead, use the UserInfo struct in `wotoValues` package.
type UserInfo struct {
	UserId         PublicUserId   `json:"user_id" gorm:"primaryKey"`
	PrivateHash    string         `json:"private_hash"`
	Email          string         `json:"email"`
	Region         string         `json:"region"`
	Language       string         `json:"language"`
	Website        string         `json:"website"`
	AuthKey        string         `json:"auth_key"`
	AccessHash     string         `json:"access_hash"`
	Password       string         `json:"password"`
	Permission     UserPermission `json:"permission"`
	Bio            string         `json:"bio"`
	SourceUrl      string         `json:"source_url"`
	Birthday       string         `json:"birthday"`
	AnilistUrl     string         `json:"anilist_url"`
	MyAnimelistUrl string         `json:"my_animelist_url"`
	RedditUrl      string         `json:"reddit_url"`
	GithubUrl      string         `json:"github_url"`
	GitlabUrl      string         `json:"gitlab_url"`
	TelegramId     int64          `json:"telegram_id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Username       string         `json:"username"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	IsVirtual      bool           `json:"is_virtual"`
	CreatedBy      PublicUserId   `json:"created_by"`
	cachedTime     time.Time      `json:"-" gorm:"-" sql:"-"`

	metaProvider MetaDataProvider `json:"-" gorm:"-" sql:"-"`
}

type GroupInfo struct {
	GroupId          PublicGroupId `json:"group_id" gorm:"primaryKey"`
	GroupRegion      string        `json:"group_region"`
	GroupUsername    string        `json:"group_username"`
	TelegramId       int64         `json:"telegram_id"`
	TelegramUsername string        `json:"telegram_username"`
	CurrentPlaying   MediaModelId  `json:"current_playing"`
}

type FavoriteValue struct {
	UserId      PublicUserId `json:"user_id" gorm:"primaryKey"`
	FavoriteKey string       `json:"favorite_key"`
	TheValue    string       `json:"the_value"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type LikedListElement struct {
	UniqueId     string       `json:"unique_id" gorm:"primaryKey"`
	OwnerId      PublicUserId `json:"owner_id"`
	MediaId      MediaModelId `json:"media_id"`
	Title        string       `json:"title"`
	LikedKey     string       `json:"liked_key"`
	SourceUrl    string       `json:"source_url"`
	ReferenceUrl string       `json:"reference_url"`
	UpdatedAt    time.Time    `json:"-"`
}
