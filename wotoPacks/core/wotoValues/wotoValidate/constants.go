package wotoValidate

import "golang.org/x/crypto/bcrypt"

const (
	MaxBioLength         = 200
	MaxFirstNameLength   = 50
	MaxLastNameLength    = 50
	MinUsernameLength    = 5
	MaxUsernameLength    = 24
	MinPasswordLength    = 8
	MaxPasswordLength    = 32
	PassHeadersLen       = 0x03
	MinKeyLength         = 3
	MaxKeyLength         = 20
	MaxUserFavoriteCount = 32
	MinTitleLength       = 2
	MaxTitleLength       = 64
	PasswordSaltCost     = bcrypt.MinCost + 0x05
)

const (
	EmailRegex = `/^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/`
)
