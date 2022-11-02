package wotoValidate

var (
	allowedUsernameChars = map[rune]bool{
		'_': true,
		'-': true,
	}
	allowedKeyChars = map[rune]bool{
		' ': true,
		'_': true,
		'-': true,
	}
	replaceKeyChars = map[rune]string{
		' ': "_",
		'-': "_",
		'_': "_",
		'~': "_",
		'.': "_",
		'+': "_",
		'/': "_",
		'|': "_",
		'*': "_",
		'^': "_",
	}
	_passHeaderStrs = []string{
		"*",
		"!",
		"/",
	}
	_passSignatureStrs = []string{
		"<",
		"@",
		"^",
		"-",
		"~",
		"]",
		"=",
		")",
	}
)
