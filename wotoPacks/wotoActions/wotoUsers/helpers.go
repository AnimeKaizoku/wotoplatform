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
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/wotoActions"
)

func ParseBatchExecute(b interfaces.ReqBase) error {
	arr, err := wotoActions.ParseBatchExecute(b.GetBatchExecute())
	if err != nil {
		return err
	}

	if len(arr) == 0 || !batchValuesValid(arr) {
		return wotoActions.ErrBatchParseFailed
	}

	b.SetBatchValues(arr)

	return nil
}

func batchValuesValid(e []wotoActions.BatchExecution) bool {
	for _, b := range e {
		if _batchHandlers[b] == nil {
			return false
		}
	}

	return true
}

func toRegisterUserResult(user *wv.UserInfo) *RegisterUserResult {
	return &RegisterUserResult{
		UserId:      user.UserId,
		PrivateHash: user.PrivateHash,
		Email:       user.Email,
		Website:     user.Website,
		AuthKey:     user.AuthKey,
		AccessHash:  user.AccessHash,
		Permission:  user.Permission,
		Bio:         user.Bio,
		SourceUrl:   user.SourceUrl,
		TelegramId:  user.TelegramId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		CreatedAt:   user.CreatedAt.Format(wv.DateTimeFormat),
		UpdatedAt:   user.UpdatedAt.Format(wv.DateTimeFormat),
		IsVirtual:   user.IsVirtual,
		CreatedBy:   user.CreatedBy,
	}
}

func toLoginUserResult(user *wv.UserInfo) *LoginUserResult {
	return &LoginUserResult{
		UserId:      user.UserId,
		PrivateHash: user.PrivateHash,
		Email:       user.Email,
		Website:     user.Website,
		Permission:  user.Permission,
		Bio:         user.Bio,
		SourceUrl:   user.SourceUrl,
		TelegramId:  user.TelegramId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		CreatedAt:   user.CreatedAt.Format(wv.DateTimeFormat),
		UpdatedAt:   user.UpdatedAt.Format(wv.DateTimeFormat),
		CreatedBy:   user.CreatedBy,
	}
}

func toGetMeResult(user *wv.UserInfo) *GetMeResult {
	return &GetMeResult{
		UserId:      user.UserId,
		PrivateHash: user.PrivateHash,
		Email:       user.Email,
		Website:     user.Website,
		Permission:  user.Permission,
		Bio:         user.Bio,
		SourceUrl:   user.SourceUrl,
		TelegramId:  user.TelegramId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Username:    user.Username,
		CreatedAt:   user.CreatedAt.Format(wv.DateTimeFormat),
		UpdatedAt:   user.UpdatedAt.Format(wv.DateTimeFormat),
	}
}
