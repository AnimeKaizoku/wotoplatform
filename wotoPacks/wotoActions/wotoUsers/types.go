package wotoUsers

type RegisterUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
