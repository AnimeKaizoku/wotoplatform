/*
 * This file is part of wp-server project (https://github.com/AnimeKaizoku/wotoplatform).
 * Copyright (c) 2021 ALiwoto.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package wotoUsers

import (
	"wp-server/wotoPacks/core/utils/logging"
	"wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database/usersDatabase"
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

	if usersDatabase.UsernameExists(entryData.Username) {
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
