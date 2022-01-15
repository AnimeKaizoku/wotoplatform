package wotoUsers

import wv "wp-server/wotoPacks/core/wotoValues"

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
	Username   string            `json:"username"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	Permission wv.UserPermission `json:"permission"`
	Bio        string            `json:"bio"`
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
	Username    string            `json:"username"`
	FirstName   string            `json:"first_name"`
	LastName    string            `json:"last_name"`
	Bio         string            `json:"bio"`
	Website     string            `json:"website"`
	Email       string            `json:"email"`
	TelegramId  int64             `json:"telegram_id"`
	AuthKey     *string           `json:"auth_key"`
	AccessHash  *string           `json:"access_hash"`
	Permission  wv.UserPermission `json:"permission"`
}
