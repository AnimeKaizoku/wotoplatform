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

package entryPoints

import (
	"encoding/json"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/serverErrors"
	"wp-server/wotoPacks/wotoActions"

	wcr "github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

//---------------------------------------------------------

// GetAction returns the action associated with the current request.
func (e *RequestEntry) GetAction() wotoActions.RequestAction {
	return e.Action
}

// IsActionValid will check if the action value of the request entry
// is a valid action or not. it will return `true`, if it's valid,
// and will return `false`, if it's not.
// if the action is not a valid action, for security matters,
// you should close the connection. (because there is a possibility that
// the request sender, is not the official client itself, so you should
// close the connection.
// although there is the possibility that the client is official,
// but it has encounter a bug, or maybe it's an old version of the client.
// in anyway, you should close the connection.
func (e *RequestEntry) IsActionValid() bool {
	return wotoActions.IsActionValid(e.Action)
}

// IsRegistered function will check if the current connection
// is registered or not.
// if a connection is not registered, it needs to send its
// information to the server first, so it can be registered
// via batch execution of `versioning` package.
// but if a client is not registered and wants to execute
// a batch execute besides `check_version` of versioning
// package, its connection should be closed at once.
func (e *RequestEntry) IsRegistered() bool {
	if e.Connection == nil {
		return false
	}

	return e.Connection.IsRegistered()
}

// RegisterConnection function will mark the current connection
// that this request belongs to as a registered connection, so
// the connection can execute another batch executions besides
// `check_version` in versioning package.
// WARNING: this package should only be used in versioning
// package when we are checking for information of the
// client, such as client id, etc...
func (e *RequestEntry) RegisterConnection() {
	if e.Connection != nil {
		e.Connection.Register()
	}
}

// GetBatchExecute will return the batch execute of the
// request entry.
// you shouldn't check for batch execute being valid in request entry,
// since we don't know it belongs to which action, the action package
// itself should check and see if it's valid or not.
// but please do notice, that if the batch execute is not
// valid, you should close the connection.
func (e *RequestEntry) GetBatchExecute() string {
	return e.BatchExecute
}

// GetData will return the data sent by the client.
// it SHOULD be in json format, but we can't check it in request entry,
// because it's not our duty in this package.
// the action package itself should check if it's valid or not.
// please do notice, that if the data is not in a valid format,
// you should just close the connection.
// return an error from the handler, so we can close the connection.
func (e *RequestEntry) GetData() string {
	return e.Data
}

func (e *RequestEntry) WriteData(b []byte) (int, error) {
	if e.Connection == nil {
		return 0, ErrConnectionUnavailable
	}

	return e.Connection.WriteBytes(b)
}

func (e *RequestEntry) WriteJson(v wcr.KeysContainer) (int, error) {
	if e.Connection == nil {
		return 0, ErrConnectionUnavailable
	}

	return e.Connection.WriteJson(v)
}

func (e *RequestEntry) WriteError(errType int, message string) (int, error) {
	return e.WriteJson(&wotoActions.ActionResp{
		UniqueId: e.UniqueId,
		Success:  false,
		Error: &serverErrors.EndPointError{
			Code:    serverErrors.ErrorCode(errType),
			Message: message,
		},
	})
}

func (e *RequestEntry) SendError(err *serverErrors.EndPointError) (int, error) {
	return e.WriteJson(&wotoActions.ActionResp{
		UniqueId: e.UniqueId,
		Success:  false,
		Error:    err,
	})
}

func (e *RequestEntry) WriteResult(result interface{}) (int, error) {
	return e.WriteJson(&wotoActions.ActionResp{
		UniqueId: e.UniqueId,
		Success:  true,
		Result:   result,
	})
}

// SendResult is a wrapper method for `WriteResult` which returns
// only error value.
func (e *RequestEntry) SendResult(result interface{}) error {
	_, err := e.WriteResult(result)
	return err
}

func (e *RequestEntry) WriteResp(resp *wotoActions.ActionResp) (int, error) {
	return e.WriteJson(resp)
}

func (e *RequestEntry) WriteString(str string) (int, error) {
	if len(str) == wv.BaseIndex {
		return wv.BaseIndex, nil
	}

	return e.WriteData([]byte(str))
}

func (e *RequestEntry) ReadData() ([]byte, error) {
	if e.Connection == nil {
		return nil, ErrConnectionUnavailable
	}

	return e.Connection.ReadBytes()
}

func (e *RequestEntry) ReadJson(i interface{}) error {
	if e.Connection == nil {
		return ErrConnectionUnavailable
	}

	return e.Connection.ReadJson(i)
}

func (e *RequestEntry) ReadString() (string, error) {
	if e.Connection == nil {
		return "", nil
	}

	return e.Connection.ReadString()
}

func (e *RequestEntry) GetBatchValues() []wotoActions.BatchExecution {
	return e.batchValues
}

func (e *RequestEntry) SetBatchValues(v []wotoActions.BatchExecution) {
	e.batchValues = v
}

func (e *RequestEntry) ParseJsonData(v interface{}) error {
	if v == nil {
		return nil
	}

	return json.Unmarshal([]byte(e.Data), v)
}

func (e *RequestEntry) CanWrite() bool {
	return e.Connection != nil && e.Connection.CanReadAndWrite()
}

// ShouldExit returns true if and only if current request entry should
// exit.
func (e *RequestEntry) ShouldExit() bool {
	return e.exit
}

// LetExit method will mark the current batch execution request
// as exited; it will set the connection field to nil, so you
// won't be able to read and write through it anymore.
func (e *RequestEntry) LetExit() error {
	if !e.exit {
		e.exit = true
		e.batchValues = nil
		e.Connection = nil
	}
	return nil
}

func (e *RequestEntry) SetMe(user *wv.UserInfo) {
	e.Connection.SetMe(user)
}

func (e *RequestEntry) GetMe() *wv.UserInfo {
	return e.Connection.GetMe()
}

func (e *RequestEntry) IsAuthorized() bool {
	return e.Connection.GetMe() != nil
}

func (e *RequestEntry) GetUniqueId() string {
	return e.UniqueId
}

func (e *RequestEntry) SetAsUniqueId(value string) {
	e.UniqueId = value
}

//---------------------------------------------------------
