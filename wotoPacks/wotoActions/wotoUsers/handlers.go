package wotoUsers

import (
	"wp-server/wotoPacks/database"
	"wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/utils/logging"
	"wp-server/wotoPacks/wotoActions"
)

func HandleUserAction(req interfaces.ReqBase) error {
	logging.SUGARED.Debug("received user action")

	b := req.GetBatchValues()
	var err error
	for _, ex := range b {
		var handler func(req interfaces.ReqBase) error

		switch ex {
		case BATCH_REGISTER_USER:
			handler = BatchRegisterUser
			continue
		case BATCH_LOGIN_USER:
			handler = BatchLoginUser
		default:
			logging.SUGARED.Warn("invalid batch:", ex)
			return wotoActions.ErrInvalidBatch
		}

		err = handler(req)
		if err != nil {
			logging.SUGARED.Debug("an error while executing batch execution: ", err)
			return err
		}
	}

	req.LetExit()

	return nil
}

func BatchRegisterUser(req interfaces.ReqBase) error {
	var entryData RegisterUserData
	err := req.ParseJsonData(&entryData)
	if err != nil {
		logging.SUGARED.Error(err)
		return err
	}

	if len(entryData.Password) < 8 || len(entryData.Username) < 4 {
		_, err = req.WriteError(ErrTypeUserPassInvalid, ErrMsgUserPassInvalid)
		if err != nil {
			logging.SUGARED.Debug(err)
			return err
		}
	}

	if database.UsernameExists(entryData.Username) {
		_, err = req.WriteError(ErrTypeUsernameExists, ErrMsgUsernameExists)
		if err != nil {
			logging.SUGARED.Debug(err)
			return err
		}
	}

	return nil
}

func BatchLoginUser(req interfaces.ReqBase) error {

	return nil
}
