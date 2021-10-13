package serverErrors

import (
	"strconv"
)

func (e *EndPointError) ToString() string {
	return "type of " + e.GetType() + ": message of:" + e.Message
}

func (e *EndPointError) GetType() string {
	switch ErrorType(e.Type) {
	case ErrNoError:
		return StrNoError
	case ErrUnknown:
		return StrUnknownError
	case ErrServerUnavailable:
		return StrServerUnavailableError
	default:
		return strconv.Itoa(e.Type)

	}
}
