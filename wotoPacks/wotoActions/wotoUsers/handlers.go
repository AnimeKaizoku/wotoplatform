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
	we "wp-server/wotoPacks/core/wotoErrors"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database/usersDatabase"
	"wp-server/wotoPacks/interfaces"
	wa "wp-server/wotoPacks/wotoActions"
)

func HandleUserAction(req interfaces.ReqBase) error {
	batchValues := req.GetBatchValues()
	var err error
	var handler wv.ReqHandler

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

	return req.LetExit()
}

func batchRegisterUser(req interfaces.ReqBase) error {
	var entryData = new(RegisterUserData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		logging.Error(err)
		return err
	}

	doer := req.GetMe()
	if doer != nil && !doer.CanCreateAccount() {
		return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	}

	if !wv.IsCorrectUsernameFormat(entryData.Username) {
		return we.SendInvalidUsernameFormat(req, OriginRegisterUser)
	}

	if !wv.IsCorrectPasswordFormat(entryData.Password) {
		return we.SendInvalidUsernameFormat(req, OriginRegisterUser)
	}

	if usersDatabase.UsernameExists(entryData.Username) {
		return we.SendUsernameExists(req, OriginRegisterUser)
	}

	var dbData = &usersDatabase.NewUserData{
		Username: entryData.Username,
		Password: entryData.Password,
	}

	if doer != nil {
		dbData.By = doer.UserId
		dbData.Permission = entryData.Permission
	}

	user := usersDatabase.CreateNewUser(dbData)
	if doer == nil {
		req.SetMe(user)
	}

	return req.SendResult(toRegisterUserResult(user))
}

func batchLoginUser(req interfaces.ReqBase) error {
	var entryData = new(LoginUserData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		logging.Error(err)
		return err
	}

	if req.IsAuthorized() {
		return we.SendAlreadyAuthorized(req, OriginLoginUser)
	}

	if !wv.IsCorrectUsernameFormat(entryData.Username) {
		return we.SendInvalidUsernameFormat(req, OriginLoginUser)
	}

	if !wv.IsCorrectPasswordFormat(entryData.Password) {
		return we.SendInvalidPasswordFormat(req, OriginLoginUser)
	}

	if !usersDatabase.UsernameExists(entryData.Username) {
		return we.SendWrongUsername(req, OriginLoginUser)
	}

	user := usersDatabase.GetUserByUsername(entryData.Username)

	if !user.IsPasswordCorrect(entryData.Password) {
		return we.SendWrongPassword(req, OriginLoginUser)
	}

	req.SetMe(user)

	return req.SendResult(toLoginUserResult(user))
}

func batchGetMe(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMe)
	}

	return req.SendResult(toGetMeResult(req.GetMe()))
}

func batchChangeUserBio(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginChangeUserBio)
	}

	var entryData = new(ChangeBioData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		logging.Error(err)
		return err
	}

	user := req.GetMe()
	if user.Bio == entryData.Bio {
		return we.SendNotModified(req, OriginChangeUserBio)
	}

	if wv.IsBioTooLong(entryData.Bio) {
		return we.SendBioTooLong(req, OriginChangeUserBio)
	}

	user.Bio = entryData.Bio
	user.SetCachedTime()
	usersDatabase.SaveUser(user, false)

	return req.SendResult(true)
}
