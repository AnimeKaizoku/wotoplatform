package wotoUsers

type RegisterUserData struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginUserData struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	AuthKey    string `json:"auth_key"`
	AccessHash string `json:"access_hash"`
}
