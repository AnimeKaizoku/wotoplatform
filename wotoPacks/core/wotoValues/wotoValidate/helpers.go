package wotoValidate

func IsCorrectPasswordFormat(password string) bool {
	return len(password) >= 8
}

func IsCorrectUsernameFormat(username string) bool {
	return len(username) >= 5 && isCorrectUsername(username)
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

	return false
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
