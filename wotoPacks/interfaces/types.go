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
	"wp-server/wotoPacks/core/wotoValues/wotoRaw"
	"wp-server/wotoPacks/serverErrors"
	"wp-server/wotoPacks/wotoActions"

	wcr "github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

type UserInfo = wotoRaw.UserInfo

type ReqBase interface {
	GetAction() wotoActions.RequestAction
	GetBatchExecute() string
	GetData() string
	GetBatchValues() []wotoActions.BatchExecution
	CanWrite() bool
	SetBatchValues([]wotoActions.BatchExecution)
	WriteData(b []byte) (n int, err error)
	WriteJson(value wcr.KeysContainer) (n int, err error)
	WriteError(errCode int, errMessage string) (int, error)
	SendError(err *serverErrors.EndPointError) (int, error)
	WriteResult(result interface{}) (int, error)

	// SendResult is a wrapper method for `WriteResult` which returns
	// only error value.
	SendResult(result interface{}) error

	WriteString(str string) (n int, err error)
	ParseJsonData(v interface{}) error
	ReadData() (n []byte, err error)
	ReadJson(value interface{}) error

	// ReadString will read the incoming bytes from the tcp
	// connection and will return it as a string value.
	// You should always use this method to read
	// all receiving data from the client.
	ReadString() (string, error)

	// LetExit method will mark the current batch execution request
	// as exited; it will set the connection field to nil, so you
	// won't be able to read and write through it anymore.
	// This method SHOULD always return nil.
	LetExit() error

	// ShouldExit returns true if and only if current request entry should
	// exit.
	ShouldExit() bool

	// IsRegistered function will check if the current connection
	// is registered or not.
	// if a connection is not registered, it needs to send its
	// information to the server first, so it can be registered
	// via batch execution of `versioning` package.
	// but if a client is not registered and wants to execute
	// a batch execute besides `check_version` of versioning
	// package, its connection should be closed at once.
	IsRegistered() bool

	// RegisterConnection function will mark the current connection
	// that this request belongs to as a registered connection, so
	// the connection can execute another batch executions besides
	// `check_version` in versioning package.
	// WARNING: this package should only be used in versioning
	// package when we are checking for information of the
	// client, such as client id, etc...
	RegisterConnection()

	// SetMe method will set the user information of the current
	// connection.
	SetMe(user *UserInfo)

	// IsAuthorized returns true if and only if the current connection
	// is authorized as a valid user.
	IsAuthorized() bool

	// GetMe method will return the user information of the current
	// connection.
	GetMe() *UserInfo
}
