/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
 * Copyright (c) 2021 AmanoTeam.
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

import "wp-server/wotoPacks/wotoActions"

type ReqBase interface {
	GetAction() wotoActions.RequestAction
	GetBatchExecute() string
	GetData() string
	GetBatchValues() []wotoActions.BatchExecution
	CanWrite() bool
	SetBatchValues([]wotoActions.BatchExecution)
	WriteData(b []byte) (n int, err error)
	WriteJson(i interface{}) (n int, err error)
	WriteError(t int, msg string) (int, error)
	WriteResult(result interface{}) (int, error)
	WriteString(str string) (n int, err error)
	ParseJsonData(v interface{}) error
	ReadData() (n []byte, err error)
	ReadJson(i interface{}) error
	ReadString() (string, error)
	LetExit()
	ShouldExit() bool
	IsRegistered() bool
	RegisterConnection()
}
