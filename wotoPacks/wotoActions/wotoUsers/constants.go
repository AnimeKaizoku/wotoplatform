package wotoUsers

// batch execution values
const (
	BATCH_REGISTER_USER = "register_user"
	BATCH_LOGIN_USER    = "login_user"
)

// error types
const (
	ErrTypeUsernameExists = iota + 1
	ErrTypeUserPassInvalid
)

// error messages
const (
	ErrMsgUsernameExists  = "username is already registered in database"
	ErrMsgUserPassInvalid = "username or password are entered in a wrong format"
)
