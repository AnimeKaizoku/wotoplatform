package wotoUsers

import (
	wv "wp-server/wotoPacks/core/wotoValues"
)

type RegisterUserData struct {
	Username   string            `json:"username"`
	Password   string            `json:"password"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	Birthday   string            `json:"birthday"`
	Permission wv.UserPermission `json:"permission"`
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
	Username   string `json:"username"`
	Password   string `json:"password"`
	AuthKey    string `json:"auth_key"`
	AccessHash string `json:"access_hash"`
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

type GetUserByTelegramIdData struct {
	TelegramId int64 `json:"telegram_id"`
}

type GetUserByTelegramIdResult struct {
	UserId             wv.PublicUserId   `json:"user_id"`
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
	IsVirtual          bool              `json:"is_virtual"`
}

type GetUserByEmailData struct {
	Email string `json:"email"`
}

type GetUserByEmailResult struct {
	UserId             wv.PublicUserId   `json:"user_id"`
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
	IsVirtual          bool              `json:"is_virtual"`
}
