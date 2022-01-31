package wotoValidate

func IsCorrectPasswordFormat(password string) bool {
	return len(password) >= MinPasswordLength && len(password) <= MaxPasswordLength
}

func IsCorrectUsernameFormat(username string) bool {
	return len(username) >= 5 && len(username) <= MaxUsernameLength && isCorrectUsername(username)
}

func isCorrectUsername(username string) bool {
	for i, c := range username {
		if (c < 'a' && c > 'z') || (c < 'A' && c > 'Z') {
			if i == 0 || i == len(username)-1 {
				return false
			}

			if !allowedUsernameChars[c] {
				return false
			}
		}
	}

	return true
}

func IsKeyValid(key string) bool {
	return len(key) >= MinKeyLength && len(key) <= MaxKeyLength
}

func IsBioTooLong(bio string) bool {
	return len(bio) > MaxBioLength
}

func IsFirstNameTooLong(firstName string) bool {
	return len(firstName) > MaxFirstNameLength
}

func IsLastNameTooLong(lastName string) bool {
	return len(lastName) > MaxLastNameLength
}
