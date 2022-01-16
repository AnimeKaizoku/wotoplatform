package wotoUsers

import (
	wv "wp-server/wotoPacks/core/wotoValues"
)

type RegisterUserData struct {
	UserId      wv.PublicUserId   `json:"user_id"`
	PrivateHash string            `json:"private_hash"`
	Username    string            `json:"username"`
	Password    string            `json:"password"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Email       string            `json:"email"`
	Permission  wv.UserPermission `json:"permission"`
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
}
