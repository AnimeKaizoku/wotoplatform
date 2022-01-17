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
	UserId             PublicUserId   `json:"user_id" gorm:"primaryKey"`
	PrivateHash        string         `json:"private_hash"`
	Email              string         `json:"email"`
	Region             string         `json:"region"`
	Language           string         `json:"language"`
	Website            string         `json:"website"`
	AuthKey            string         `json:"auth_key"`
	AccessHash         string         `json:"access_hash"`
	Password           string         `json:"password"`
	Permission         UserPermission `json:"permission"`
	Bio                string         `json:"bio"`
	SourceUrl          string         `json:"source_url"`
	Birthday           string         `json:"birthday"`
	AnilistUrl         string         `json:"anilist_url"`
	MyAnimelistUrl     string         `json:"my_animelist_url"`
	RedditUrl          string         `json:"reddit_url"`
	GithubUrl          string         `json:"github_url"`
	GitlabUrl          string         `json:"gitlab_url"`
	FavoriteColor      string         `json:"favorite_color"`
	FavoriteMusic      string         `json:"favorite_music"`
	FavoriteAnime      string         `json:"favorite_anime"`
	FavoriteMovie      string         `json:"favorite_movie"`
	FavoriteFood       string         `json:"favorite_food"`
	FavoriteSeries     string         `json:"favorite_series"`
	FavoriteLightNovel string         `json:"favorite_light_novel"`
	FavoriteNovel      string         `json:"favorite_novel"`
	TelegramId         int64          `json:"telegram_id"`
	FirstName          string         `json:"first_name"`
	LastName           string         `json:"last_name"`
	Username           string         `json:"username"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	IsVirtual          bool           `json:"is_virtual"`
	CreatedBy          PublicUserId   `json:"created_by"`
	cachedTime         time.Time      `json:"-" gorm:"-" sql:"-"`
}
