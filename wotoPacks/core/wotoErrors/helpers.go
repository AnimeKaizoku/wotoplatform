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

func SendPermissionDenied(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrPermissionDenied,
		Message: MessagePermissionDenied,
		Origin:  origin,
	})
	return err
}

func SendKeyNotFound(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrKeyNotFound,
		Message: MessageKeyNotFound,
		Origin:  origin,
	})
	return err
}

func SendInvalidTelegramId(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidTelegramId,
		Message: MessageInvalidTelegramId,
		Origin:  origin,
	})
	return err
}

func SendInvalidEmail(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidEmail,
		Message: MessageInvalidEmail,
		Origin:  origin,
	})
	return err
}

func SendInvalidKey(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidKey,
		Message: MessageInvalidKey,
		Origin:  origin,
	})
	return err
}

func SendTooManyFavorites(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrTooManyFavorites,
		Message: MessageTooManyFavorites,
		Origin:  origin,
	})
	return err
}

func SendLikedListElementAlreadyExists(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrLikedListElementAlreadyExists,
		Message: MessageLikedListElementAlreadyExists,
		Origin:  origin,
	})
	return err
}

func SendTooManyLikedList(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrTooManyLikedList,
		Message: MessageTooManyLikedList,
		Origin:  origin,
	})
	return err
}

func SendInvalidUniqueId(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidUniqueId,
		Message: MessageInvalidUniqueId,
		Origin:  origin,
	})
	return err
}

func SendMediaNotFound(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrMediaNotFound,
		Message: MessageMediaNotFound,
		Origin:  origin,
	})
	return err
}

func SendInvalidMediaId(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidMediaId,
		Message: MessageInvalidMediaId,
		Origin:  origin,
	})
	return err
}

func SendMediaTitleAlreadyExists(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrMediaTitleAlreadyExists,
		Message: MessageMediaTitleAlreadyExists,
		Origin:  origin,
	})
	return err
}

func SendInvalidGenreId(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrInvalidGenreId,
		Message: MessageInvalidGenreId,
		Origin:  origin,
	})
	return err
}

func SendGenreInfoNotFound(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrGenreInfoNotFound,
		Message: MessageGenreInfoNotFound,
		Origin:  origin,
	})
	return err
}

func SendGenreTitleAlreadyExists(req inf.ReqBase, origin string) error {
	_, err := req.SendError(&serverErrors.EndPointError{
		Code:    serverErrors.ErrGenreTitleAlreadyExists,
		Message: MessageGenreTitleAlreadyExists,
		Origin:  origin,
	})
	return err
}
