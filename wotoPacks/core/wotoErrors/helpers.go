package wotoErrors

import (
	inf "wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/serverErrors"
)

func SendInvalidUsernameFormat(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidUsernameFormat,
		Message: MessageInvalidUsernameFormat,
		Origin:  origin,
	})
	return err
}

func SendInvalidPasswordFormat(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidPasswordFormat,
		Message: MessageInvalidPasswordFormat,
		Origin:  origin,
	})
	return err
}

func SendUsernameExists(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrUsernameExists,
		Message: MessageUsernameExists,
		Origin:  origin,
	})
	return err
}

func SendWrongUsername(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrWrongUsername,
		Message: MessageWrongUsername,
		Origin:  origin,
	})
	return err
}

func SendWrongPassword(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrWrongPassword,
		Message: MessageWrongPassword,
		Origin:  origin,
	})
	return err
}

func SendInvalidAuthKeyFormat(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidAuthKeyFormat,
		Message: MessageInvalidAuthKeyFormat,
		Origin:  origin,
	})
	return err
}

func SendInvalidAccessHashFormat(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidAccessHashFormat,
		Message: MessageInvalidAccessHashFormat,
		Origin:  origin,
	})
	return err
}

func SendWrongAuthKey(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrWrongAuthKey,
		Message: MessageWrongAuthKey,
		Origin:  origin,
	})
	return err
}

func SendLoginAccessHashExpired(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrLoginAccessHashExpired,
		Message: MessageLoginAccessHashExpired,
		Origin:  origin,
	})
	return err
}

func SendInvalidFirstName(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidFirstName,
		Message: MessageInvalidFirstName,
		Origin:  origin,
	})
	return err
}

func SendInvalidLastName(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidLastName,
		Message: MessageInvalidLastName,
		Origin:  origin,
	})
	return err
}

func SendInvalidTitle(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidTitle,
		Message: MessageInvalidTitle,
		Origin:  origin,
	})
	return err
}

func SendAlreadyAuthorized(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrAlreadyAuthorized,
		Message: MessageAlreadyAuthorized,
		Origin:  origin,
	})
	return err
}

func SendNotAuthorized(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrNotAuthorized,
		Message: MessageNotAuthorized,
		Origin:  origin,
	})
	return err
}

func SendNotModified(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrNotModified,
		Message: MessageNotModified,
		Origin:  origin,
	})
	return err
}

func SendBioTooLong(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrBioTooLong,
		Message: MessageBioTooLong,
		Origin:  origin,
	})
	return err
}

func SendUserNotFound(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrUserNotFound,
		Message: MessageUserNotFound,
		Origin:  origin,
	})
	return err
}

func SendFirstNameTooLong(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrFirstNameTooLong,
		Message: MessageFirstNameTooLong,
		Origin:  origin,
	})
	return err
}

func SendLastNameTooLong(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrLastNameTooLong,
		Message: MessageLastNameTooLong,
		Origin:  origin,
	})
	return err
}

func SendInvalidUsernameAndUserId(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidUsernameAndUserId,
		Message: MessageInvalidUsernameAndUserId,
		Origin:  origin,
	})
	return err
}

func SendMethodNotImplemented(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrMethodNotImplemented,
		Message: MessageMethodNotImplemented,
		Origin:  origin,
	})
	return err
}
