package wotoValidate

const (
	MaxBioLength       = 200
	MaxFirstNameLength = 50
	MaxLastNameLength  = 50
	MinUsernameLength  = 5
	MaxUsernameLength  = 24
	MinPasswordLength  = 8
	MaxPasswordLength  = 32
	MinKeyLength       = 3
	MaxKeyLength       = 20
)

const (
	EmailRegex = `/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/`
)
