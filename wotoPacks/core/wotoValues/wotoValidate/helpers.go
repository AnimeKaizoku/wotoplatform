package wotoValidate

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

func GetPasswordHash(b []byte) string {
	hash := sha256.Sum256(b)
	return hex.EncodeToString(hash[:])
}

func IsCorrectPasswordFormat(password *wotoCrypto.PasswordContainer256) bool {
	if password.Hash256 == "" || !password.HasSignature() {
		return false
	}

	headers := ssg.Split(password.Header, _passHeaderStrs...)
	signatures := ssg.Split(password.Signature, _passSignatureStrs...)

	if len(headers) != PassHeadersLen {
		return false
	}

	charsLen := int(ssg.ToInt32(headers[0x00]))
	sigPayloadLen := int(ssg.ToInt8(headers[0x01]))
	if len(signatures)-sigPayloadLen != charsLen || (charsLen < MinPasswordLength && charsLen > MaxPasswordLength) {
		return false
	}

	hashCheckStr := ""
	for i, current := range signatures {
		if i < sigPayloadLen {
			// TODO: handle payload data
			continue
		}

		if i >= charsLen+sigPayloadLen {
			break
		}

		cInt := ssg.ToInt32(current)
		if cInt == 0 || cInt < 0x061 {
			return false
		}

		hashCheckStr += string(rune(cInt))
	}

	if len(hashCheckStr) != charsLen {
		return false
	} else if GetPasswordHash([]byte(hashCheckStr)) != password.Hash256 {
		return false
	}

	return true
}

func GetPassAsBytes(password *wotoCrypto.PasswordContainer256) []byte {
	headers := ssg.Split(password.Header, _passHeaderStrs...)
	signatures := ssg.Split(password.Signature, _passSignatureStrs...)

	charsLen := int(ssg.ToInt32(headers[0x00]))
	p := headers[0x02]
	for i, current := range signatures {
		if i >= charsLen {
			return []byte(p)
		}

		p += string(rune(ssg.ToInt32(current)))
	}

	return []byte(p)
}

func IsCorrectUsernameFormat(username string) bool {
	return len(username) >= 5 && len(username) <= MaxUsernameLength && isCorrectUsername(username)
}

func isCorrectUsername(username string) bool {
	for i, c := range username {
		if IsEnglish(c) {
			// a valid english letter, let it pass
			continue
		}

		if ssg.IsRuneNumber(c) && i != 0 {
			// a valid number, let it pass
			continue
		} else if i == 0 {
			// we only allow a-z at the beginning, but if we are
			// at the end of the username, we allow numbers as well.
			return false
		}

		if !allowedUsernameChars[c] {
			return false
		}
	}

	return true
}

func IsKeyValid(key string) bool {
	return len(key) >= MinKeyLength && len(key) <= MaxKeyLength && isCorrectKey(key)
}

func IsTitleValid(title string) bool {
	return len(title) >= MinTitleLength && len(title) <= MaxTitleLength
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

func IsUserFavoriteTooMany(count int) bool {
	return count > MaxUserFavoriteCount
}
