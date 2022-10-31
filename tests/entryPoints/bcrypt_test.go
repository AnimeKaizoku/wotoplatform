package entryPoints_test

import (
	"testing"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"

	"golang.org/x/crypto/bcrypt"
)

func TestRawBCrypt(t *testing.T) {
	err := bcrypt.CompareHashAndPassword([]byte("hello"), []byte("hello"))
	if err != nil {
		// crypto/bcrypt: hashedSecret too short to be a bcrypted password
		t.Log("The error when pass is not really hashed:", err)
	} else if err == nil {
		t.Error("err variable should not be nil when it's not actually hashed.")
		return
	}

	b, err := bcrypt.GenerateFromPassword([]byte("hello"), wotoValidate.PasswordSaltCost)
	if err != nil {
		t.Error("error when trying to generate hashed password:", err)
		return
	}

	t.Log(b)

	err = bcrypt.CompareHashAndPassword(b, []byte("hello ."))
	if err != nil {
		//crypto/bcrypt: hashedPassword is not the hash of the given password
		t.Log("The error when wrong pass is passed, but the first arg is really salted.")
	} else if err == nil {
		t.Error("err variable should not be nil when the password isn't correct at all.")
	}

	err = bcrypt.CompareHashAndPassword(b, []byte("hello"))
	if err != nil {
		t.Error("error when trying to compare correct passwords:", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte("hello there, how are you doing, this is a long string. 1234567891011123456"), []byte("hello"))
	if err != nil {
		//error(golang.org/x/crypto/bcrypt.InvalidHashPrefixError) 104
		t.Log("The error when pass is not really hashed:", err)
	} else if err == nil {
		t.Error("err variable should not be nil when it's not actually hashed.")
		return
	}
}
