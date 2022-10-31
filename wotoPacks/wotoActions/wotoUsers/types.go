package wotoUsers

import (
	wv "wp-server/wotoPacks/core/wotoValues"

	"github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

type RegisterUserData struct {
	Username   string                           `json:"username"`
	Password   *wotoCrypto.PasswordContainer256 `json:"password"`
	FirstName  string                           `json:"first_name"`
	LastName   string                           `json:"last_name"`
	Birthday   string                           `json:"birthday"`
	Permission wv.UserPermission                `json:"permission"`
}

type RegisterUserResult struct {
	UserId      wv.PublicUserId   `json:"user_id"`
	PrivateHash string            `json:"private_hash"`
	Email       string            `json:"email"`
	Website     string            `json:"website"`
	AuthKey     string            `json:"auth_key"`
	AccessHash  string            `json:"access_hash"`
	Permission  wv.UserPermission `json:"permission"`
	Bio         string            `json:"bio"`
	SourceUrl   string            `json:"source_url"`
	TelegramId  int64             `json:"telegram_id"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Username    string            `json:"username"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	IsVirtual   bool              `json:"is_virtual"`
	CreatedBy   wv.PublicUserId   `json:"created_by"`
}

type RegisterVirtualUserData struct {
	TelegramId int64             `json:"telegram_id"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	Birthday   string            `json:"birthday"`
	Permission wv.UserPermission `json:"permission"`
}

type RegisterVirtualUserResult struct {
	UserId      wv.PublicUserId   `json:"user_id"`
	PrivateHash string            `json:"private_hash"`
	Email       string            `json:"email"`
	Website     string            `json:"website"`
	AuthKey     string            `json:"auth_key"`
	AccessHash  string            `json:"access_hash"`
	Permission  wv.UserPermission `json:"permission"`
	Bio         string            `json:"bio"`
	SourceUrl   string            `json:"source_url"`
	TelegramId  int64             `json:"telegram_id"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	IsVirtual   bool              `json:"is_virtual"`
	CreatedBy   wv.PublicUserId   `json:"created_by"`
}

type LoginUserData struct {
	Username   string                           `json:"username"`
	Password   *wotoCrypto.PasswordContainer256 `json:"password"`
	AuthKey    string                           `json:"auth_key"`
	AccessHash string                           `json:"access_hash"`
}

type LoginUserResult struct {
	UserId      wv.PublicUserId   `json:"user_id"`
	PrivateHash string            `json:"private_hash"`
	Email       string            `json:"email"`
	Website     string            `json:"website"`
	Permission  wv.UserPermission `json:"permission"`
	Bio         string            `json:"bio"`
	SourceUrl   string            `json:"source_url"`
	TelegramId  int64             `json:"telegram_id"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Username    string            `json:"username"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	CreatedBy   wv.PublicUserId   `json:"created_by"`
}

type GetMeResult struct {
	UserId             wv.PublicUserId   `json:"user_id"`
	PrivateHash        string            `json:"private_hash"`
	Email              string            `json:"email"`
	Region             string            `json:"region"`
	Language           string            `json:"language"`
	Birthday           string            `json:"birthday"`
	AnilistUrl         string            `json:"anilist_url"`
	MyAnimelistUrl     string            `json:"my_animelist_url"`
	RedditUrl          string            `json:"reddit_url"`
	GithubUrl          string            `json:"github_url"`
	GitlabUrl          string            `json:"gitlab_url"`
	FavoriteColor      string            `json:"favorite_color"`
	FavoriteMusic      string            `json:"favorite_music"`
	FavoriteAnime      string            `json:"favorite_anime"`
	FavoriteMovie      string            `json:"favorite_movie"`
	FavoriteFood       string            `json:"favorite_food"`
	FavoriteSeries     string            `json:"favorite_series"`
	FavoriteLightNovel string            `json:"favorite_light_novel"`
	FavoriteNovel      string            `json:"favorite_novel"`
	Website            string            `json:"website"`
	Permission         wv.UserPermission `json:"permission"`
	Bio                string            `json:"bio"`
	SourceUrl          string            `json:"source_url"`
	TelegramId         int64             `json:"telegram_id"`
	FirstName          string            `json:"first_name"`
	LastName           string            `json:"last_name"`
	Username           string            `json:"username"`
	CreatedAt          string            `json:"created_at"`
	UpdatedAt          string            `json:"updated_at"`
}

type ChangeBioData struct {
	UserId wv.PublicUserId `json:"user_id"`
	Bio    string          `json:"bio"`
}

type ChangeNamesData struct {
	UserId    wv.PublicUserId `json:"user_id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
}

type GetUserInfoData struct {
	UserId   wv.PublicUserId `json:"user_id"`
	Username string          `json:"username"`
}

type GetUserInfoResult struct {
	UserId         wv.PublicUserId   `json:"user_id"`
	Email          string            `json:"email"`
	Region         string            `json:"region"`
	Language       string            `json:"language"`
	Birthday       string            `json:"birthday"`
	AnilistUrl     string            `json:"anilist_url"`
	MyAnimelistUrl string            `json:"my_animelist_url"`
	RedditUrl      string            `json:"reddit_url"`
	GithubUrl      string            `json:"github_url"`
	GitlabUrl      string            `json:"gitlab_url"`
	Website        string            `json:"website"`
	Permission     wv.UserPermission `json:"permission"`
	Bio            string            `json:"bio"`
	SourceUrl      string            `json:"source_url"`
	TelegramId     int64             `json:"telegram_id"`
	FirstName      string            `json:"first_name"`
	LastName       string            `json:"last_name"`
	Username       string            `json:"username"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	IsVirtual      bool              `json:"is_virtual"`
}

type ChangeUserPermissionData struct {
	UserId     wv.PublicUserId   `json:"user_id"`
	Permission wv.UserPermission `json:"permission"`
}

type ChangeUserPermissionResult struct {
	UserId             wv.PublicUserId   `json:"user_id"`
	PreviousPermission wv.UserPermission `json:"previous_permission"`
	CurrentPermission  wv.UserPermission `json:"current_permission"`
}

type GetUserByTelegramIdData struct {
	TelegramId int64 `json:"telegram_id"`
}

type GetUserByEmailData struct {
	Email string `json:"email"`
}

type GetUserFavoriteData struct {
	UserId      wv.PublicUserId `json:"user_id"`
	FavoriteKey string          `json:"favorite_key"`
}

type GetUserFavoriteResult struct {
	FavoriteValue string `json:"favorite_value"`
	UpdatedAt     string `json:"updated_at"`
}

type GetUserFavoriteCountData struct {
	UserId wv.PublicUserId `json:"user_id"`
}

type GetUserFavoriteCountResult struct {
	FavoritesCount int `json:"favorites_count"`
}

type SetUserFavoriteData struct {
	UserId        wv.PublicUserId `json:"user_id"`
	FavoriteKey   string          `json:"favorite_key"`
	FavoriteValue string          `json:"favorite_value"`
}

type DeleteUserFavoriteData struct {
	UserId      wv.PublicUserId `json:"user_id"`
	FavoriteKey string          `json:"favorite_key"`
}

type ResolveUsernameData struct {
	Username string `json:"username"`
}

type GetUserLikedListData struct {
	UserId   wv.PublicUserId `json:"user_id"`
	LikedKey string          `json:"liked_key"`
}

type GetUserLikedListResult struct {
	LikedList []*wv.LikedListElement `json:"liked_list"`
	UpdatedAt string                 `json:"updated_at"`
}

type GetUserLikedListCountData struct {
	UserId       wv.PublicUserId `json:"user_id"`
	LikedListKey string          `json:"liked_list_key"`
}

type GetUserLikedListCountResult struct {
	LikedListCount int `json:"liked_list_count"`
}

type AppendUserLikedListData struct {
	UserId       wv.PublicUserId `json:"user_id"`
	LikedListKey string          `json:"liked_list_key"`
	MediaId      wv.MediaModelId `json:"media_id"`
	Title        string          `json:"title"`
	ReferenceUrl string          `json:"reference_url"`
	SourceUrl    string          `json:"source_url"`
}

type AppendUserLikedListResult struct {
	UniqueId     string          `json:"unique_id"`
	UserId       wv.PublicUserId `json:"user_id"`
	LikedListKey string          `json:"liked_list_key"`
	MediaId      wv.MediaModelId `json:"media_id"`
	Title        string          `json:"title"`
	ReferenceUrl string          `json:"reference_url"`
	SourceUrl    string          `json:"source_url"`
}

type DeleteUserLikedListItemData struct {
	UserId   wv.PublicUserId `json:"user_id"`
	UniqueId string          `json:"unique_id"`
}
