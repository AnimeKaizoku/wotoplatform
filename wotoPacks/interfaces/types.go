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

package interfaces

import (
	"time"
	"wp-server/wotoPacks/wotoActions"
)

type RawUser interface {
	GetName() string
	GetSvUsername() string
	GetAvatar() string
	GetPublicID() string
	GetPassword() string
	GetUserLever() uint16
	GetLastSeen() string
	GetUserIntro() string
	GetPrivateID() string
	GetAvatarFrame() string
	GetUserVIPLevel() uint8
	GetCurrentExp() string
	GetCurrentVIPExp() string
	GetMaxExp() string
	GetMaxVIPExp() string
	GetTotalExp() string
	GetCity() string
	GetTotalVIPExp() string
	GetSocialvoidUsername() string
	GetCreatedAt() time.Time
}

type ReqBase interface {
	GetAction() wotoActions.RequestAction
	GetBatchExecute() string
	GetData() string
	GetBatchValues() []wotoActions.BatchExecution
	CanWrite() bool
	SetBatchValues([]wotoActions.BatchExecution)
	WriteData(b []byte) (n int, err error)
	WriteJson(i interface{}) (n int, err error)
	WriteError(errCode int, errMessage string) (int, error)
	WriteResult(result interface{}) (int, error)
	WriteString(str string) (n int, err error)
	ParseJsonData(v interface{}) error
	ReadData() (n []byte, err error)
	ReadJson(value interface{}) error
	ReadString() (string, error)
	LetExit()
	ShouldExit() bool
	IsRegistered() bool
	RegisterConnection()
}
