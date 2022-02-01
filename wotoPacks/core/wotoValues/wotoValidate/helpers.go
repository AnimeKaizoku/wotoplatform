package wotoValidate

import (
	"regexp"
	"strings"
)

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
	return len(key) >= MinKeyLength && len(key) <= MaxKeyLength && isCorrectKey(key)
}

func isCorrectKey(key string) bool {
	for i, c := range key {
		if !IsEnglish(c) {
			if i == 0 || i == len(key)-1 {
				return false
			}

			if !allowedKeyChars[c] {
				return false
			}
		}
	}

	return true
}

func IsEnglish(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func IsNonEnglish(r rune) bool {
	return !IsEnglish(r)
}

func PurifyKey(key string) string {
	key = strings.ToLower(strings.TrimSpace(key))
	result := ""
	lastSpecial := false
	firstPassed := false

	for i, c := range key {
		if !firstPassed {
			if !IsEnglish(c) {
				continue
			}
			firstPassed = true
		}

		if !IsEnglish(c) || (c == 32) {
			if i == 0 || i == len(key)-1 {
				lastSpecial = false
				continue
			}

			if lastSpecial {
				continue
			}

			lastSpecial = true
			result += replaceKeyChars[c]
			continue
		}

		if replaceKeyChars[c] != "" {
			if lastSpecial {
				continue
			}

			lastSpecial = true
			result += replaceKeyChars[c]
			continue
		}

		result += string(c)
	}

	return strings.TrimFunc(result, IsNonEnglish)
}

func IsEmailValid(email string) bool {
	b, err := regexp.MatchString(EmailRegex, email)
	return err == nil && b
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
