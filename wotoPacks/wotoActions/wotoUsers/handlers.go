package wotoUsers

import (
	"wp-server/wotoPacks/core/utils/logging"
	"wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database"
	"wp-server/wotoPacks/interfaces"
	wa "wp-server/wotoPacks/wotoActions"
)

func HandleUserAction(req interfaces.ReqBase) error {
	batchValues := req.GetBatchValues()
	var err error
	var handler wotoValues.ReqHandler

	for _, currentBatch := range batchValues {
		handler = _batchHandlers[currentBatch]
		if handler == nil {
			return wa.ErrInvalidBatch
		}

		err = handler(req)
		if err != nil {
			return err
		}
	}

	req.LetExit()

	return nil
}

func batchRegisterUser(req interfaces.ReqBase) error {
	var entryData = new(RegisterUserData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		logging.Error(err)
		return err
	}

	if len(entryData.Password) < 8 || len(entryData.Username) < 4 {
		_, err = req.WriteError(ErrTypeUserPassInvalid, ErrMsgUserPassInvalid)
		if err != nil {
			logging.Debug(err)
			return err
		}
	}

	if database.UsernameExists(entryData.Username) {
		_, err = req.WriteError(ErrTypeUsernameExists, ErrMsgUsernameExists)
		if err != nil {
			logging.Debug(err)
			return err
		}
	}

	return nil
}

func batchLoginUser(req interfaces.ReqBase) error {
	return nil
}
