package wotoErrors

import (
	inf "wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/serverErrors"
)

func SendInvalidUsernameFormat(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Type:    serverErrors.ErrInvalidUsernameFormat,
		Message: MessageInvalidUsernameFormat,
		Origin:  origin,
	})
	return err
}

func SendInvalidPasswordFormat(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Type:    serverErrors.ErrInvalidPasswordFormat,
		Message: MessageInvalidPasswordFormat,
		Origin:  origin,
	})
	return err
}

func SendUsernameExists(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Type:    serverErrors.ErrUsernameExists,
		Message: MessageUsernameExists,
		Origin:  origin,
	})
	return err
}

func SendWrongUsername(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Type:    serverErrors.ErrWrongUsername,
		Message: MessageWrongUsername,
		Origin:  origin,
	})
	return err
}
