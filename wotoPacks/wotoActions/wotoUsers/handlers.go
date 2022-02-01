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
	we "wp-server/wotoPacks/core/wotoErrors"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"
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
		return err
	}

	doer := req.GetMe()
	if doer != nil && !doer.CanCreateAccount() {
		return we.SendAlreadyAuthorized(req, OriginRegisterUser)
	}

	if !wotoValidate.IsCorrectUsernameFormat(entryData.Username) {
		return we.SendInvalidUsernameFormat(req, OriginRegisterUser)
	}

	if !wotoValidate.IsCorrectPasswordFormat(entryData.Password) {
		return we.SendInvalidUsernameFormat(req, OriginRegisterUser)
	}

	if usersDatabase.UsernameExists(entryData.Username) {
		return we.SendUsernameExists(req, OriginRegisterUser)
	}

	var dbData = &usersDatabase.NewUserData{
		Username:  entryData.Username,
		Password:  entryData.Password,
		FirstName: entryData.FirstName,
		LastName:  entryData.LastName,
		Birthday:  entryData.Birthday,
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

func batchRegisterVirtualUser(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetUserInfo)
	}

	var entryData = new(RegisterVirtualUserData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	doer := req.GetMe()
	if doer != nil && !doer.CanCreateAccount() {
		return we.SendAlreadyAuthorized(req, OriginRegisterVirtualUser)
	}

	var dbData = &usersDatabase.NewUserData{
		TelegramId: entryData.TelegramId,
		FirstName:  entryData.FirstName,
		LastName:   entryData.LastName,
		Birthday:   entryData.Birthday,
	}

	if doer != nil {
		dbData.By = doer.UserId
		dbData.Permission = entryData.Permission
	}

	user := usersDatabase.CreateNewUser(dbData)

	return req.SendResult(toRegisterVirtualUserResult(user))
}

func batchLoginUser(req interfaces.ReqBase) error {
	var entryData = new(LoginUserData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if req.IsAuthorized() {
		return we.SendAlreadyAuthorized(req, OriginLoginUser)
	}

	if !wotoValidate.IsCorrectUsernameFormat(entryData.Username) {
		return we.SendInvalidUsernameFormat(req, OriginLoginUser)
	}

	if !wotoValidate.IsCorrectPasswordFormat(entryData.Password) {
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
		return we.SendNotAuthorized(req, OriginChangeNames)
	}

	var entryData = new(ChangeBioData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	user := req.GetMe()
	if user.IsAdmin() && !entryData.UserId.IsZero() {
		user = usersDatabase.GetUserById(entryData.UserId)
		if user.IsInvalid() {
			return we.SendUserNotFound(req, OriginChangeNames)
		}
	}

	if entryData.HasNotModified(user) {
		return we.SendNotModified(req, OriginChangeNames)
	}

	if entryData.IsBioTooLong() {
		return we.SendBioTooLong(req, OriginChangeNames)
	}

	user.Bio = entryData.Bio
	user.SetCachedTime()
	usersDatabase.SaveUser(user, false)

	return req.SendResult(true)
}

func batchChangeNames(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginChangeNames)
	}

	var entryData = new(ChangeNamesData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	user := req.GetMe()
	if user.IsAdmin() && !entryData.UserId.IsZero() {
		user = usersDatabase.GetUserById(entryData.UserId)
		if user.IsInvalid() {
			return we.SendUserNotFound(req, OriginChangeNames)
		}
	}

	if entryData.HasNotModified(user) {
		return we.SendNotModified(req, OriginChangeNames)
	}

	if entryData.IsFirstNameTooLong() {
		return we.SendFirstNameTooLong(req, OriginChangeNames)
	}

	if entryData.IsLastNameTooLong() {
		return we.SendLastNameTooLong(req, OriginChangeNames)
	}

	user.FirstName = entryData.FirstName
	user.LastName = entryData.LastName
	user.SetCachedTime()
	usersDatabase.SaveUser(user, false)

	return req.SendResult(true)
}

func batchGetUserInfo(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetUserInfo)
	}

	var entryData = new(GetUserInfoData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.IsInvalid() {
		return we.SendInvalidUsernameAndUserId(req, OriginGetUserInfo)
	}

	var user *wv.UserInfo
	if !entryData.UserId.IsZero() {
		user = usersDatabase.GetUserById(entryData.UserId)
	} else {
		user = usersDatabase.GetUserByUsername(entryData.Username)
	}

	if user.IsInvalid() {
		return we.SendUserNotFound(req, OriginGetUserInfo)
	}

	return req.SendResult(toGetUserInfoResult(user))
}

func batchGetUserByTelegramID(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetUserByTelegramID)
	}

	var entryData = new(GetUserByTelegramIdData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.TelegramId == 0 {
		return we.SendInvalidTelegramId(req, OriginGetUserByTelegramID)
	}

	user := usersDatabase.GetUserByTelegramId(entryData.TelegramId)
	if user.IsInvalid() {
		return we.SendUserNotFound(req, OriginGetUserByTelegramID)
	}

	return req.SendResult(toGetUserInfoResult(user))
}

func batchGetUserByEmail(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetUserByEmail)
	}

	var entryData = new(GetUserByEmailData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.Email == "" || !wotoValidate.IsEmailValid(entryData.Email) {
		return we.SendInvalidEmail(req, OriginGetUserByTelegramID)
	}

	user := usersDatabase.GetUserByEmail(entryData.Email)
	if user.IsInvalid() {
		return we.SendUserNotFound(req, OriginGetUserByTelegramID)
	}

	return req.SendResult(toGetUserInfoResult(user))
}

func batchResolveUsername(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginResolveUsername)
	}

	var entryData = new(ResolveUsernameData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if !wotoValidate.IsCorrectUsernameFormat(entryData.Username) {
		return we.SendInvalidUsernameFormat(req, OriginResolveUsername)
	}

	user := usersDatabase.GetUserByUsername(entryData.Username)
	if user.IsInvalid() {
		return we.SendUserNotFound(req, OriginResolveUsername)
	}

	return req.SendResult(toGetUserInfoResult(user))
}

func batchChangeUserPermission(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginChangeUserPermission)
	}

	me := req.GetMe()
	if !me.CanChangePermission() {
		return we.SendPermissionDenied(req, OriginChangeUserPermission)
	}

	var entryData = new(ChangeUserPermissionData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	target := usersDatabase.GetUserById(entryData.UserId)
	if target.IsInvalid() {
		return we.SendUserNotFound(req, OriginChangeUserPermission)
	}

	if target.Permission == entryData.Permission {
		return we.SendNotModified(req, OriginChangeUserPermission)
	}

	last := target.Permission
	target.Permission = entryData.Permission
	usersDatabase.SaveUser(target, false)

	return req.SendResult(&ChangeUserPermissionResult{
		UserId:             target.UserId,
		PreviousPermission: last,
		CurrentPermission:  target.Permission,
	})
}

func batchGetUserFavorite(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetUserFavorite)
	}

	var entryData = new(GetUserFavoriteData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	user := req.GetMe()
	if user.IsAdmin() && !entryData.UserId.IsZero() {
		user = usersDatabase.GetUserById(entryData.UserId)
		if user.IsInvalid() {
			return we.SendUserNotFound(req, OriginGetUserFavorite)
		}
	}

	entryData.FavoriteKey = wotoValidate.PurifyKey(entryData.FavoriteKey)
	if entryData.FavoriteKey == "" {
		return we.SendInvalidKey(req, OriginSetUserFavorite)
	}

	fav := usersDatabase.GetUserFavorite(user.UserId, entryData.FavoriteKey)
	if fav.IsInvalid() {
		return we.SendKeyNotFound(req, OriginGetUserFavorite)
	}

	return req.SendResult(toGetUserFavoriteResult(fav))
}

func batchGetUserFavoriteCount(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetUserFavoriteCount)
	}

	var entryData = new(GetUserFavoriteCountData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	var user *wv.UserInfo
	if entryData.UserId.IsZero() {
		user = req.GetMe()
	} else {
		user = usersDatabase.GetUserById(entryData.UserId)
		if user == nil {
			return we.SendUserNotFound(req, OriginGetUserFavoriteCount)
		}
	}

	return req.SendResult(&GetUserFavoriteCountResult{
		FavoritesCount: usersDatabase.GetUserFavoriteCount(user.UserId),
	})
}

func batchSetUserFavorite(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginSetUserFavorite)
	}

	var entryData = new(SetUserFavoriteData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	user := req.GetMe()
	if user.IsAdmin() && !entryData.UserId.IsZero() {
		user = usersDatabase.GetUserById(entryData.UserId)
		if user.IsInvalid() {
			return we.SendUserNotFound(req, OriginSetUserFavorite)
		}
	}

	entryData.FavoriteKey = wotoValidate.PurifyKey(entryData.FavoriteKey)
	if entryData.FavoriteKey == "" {
		return we.SendInvalidKey(req, OriginSetUserFavorite)
	}

	fav := usersDatabase.GetUserFavorite(user.UserId, entryData.FavoriteKey)
	if !fav.IsInvalid() && fav.TheValue == entryData.FavoriteValue {
		return we.SendNotModified(req, OriginSetUserFavorite)
	}

	usersDatabase.SetUserFavorite(
		user.UserId,
		entryData.FavoriteKey,
		entryData.FavoriteValue,
	)

	return req.SendResult(true)
}

func batchDeleteUserFavorite(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteUserFavorite)
	}

	var entryData = new(DeleteUserFavoriteData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	user := req.GetMe()
	if user.IsAdmin() && !entryData.UserId.IsZero() {
		user = usersDatabase.GetUserById(entryData.UserId)
		if user.IsInvalid() {
			return we.SendUserNotFound(req, OriginDeleteUserFavorite)
		}
	}

	entryData.FavoriteKey = wotoValidate.PurifyKey(entryData.FavoriteKey)
	if entryData.FavoriteKey == "" {
		return we.SendInvalidKey(req, OriginSetUserFavorite)
	}

	if !usersDatabase.FavoriteValueExists(user.UserId, entryData.FavoriteKey) {
		return we.SendKeyNotFound(req, OriginDeleteUserFavorite)
	}

	usersDatabase.DeleteUserFavorite(user.UserId, entryData.FavoriteKey)

	return req.SendResult(true)
}
