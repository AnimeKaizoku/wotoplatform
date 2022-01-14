package wotoUsers

import wv "wp-server/wotoPacks/core/wotoValues"

type RegisterUserData struct {
	Username   string            `json:"username"`
	Password   string            `json:"password"`
	FirstName  string            `json:"first_name"`
	LastName   string            `json:"last_name"`
	Permission wv.UserPermission `json:"permission"`
}

type LoginUserData struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	AuthKey    string `json:"auth_key"`
	AccessHash string `json:"access_hash"`
}
