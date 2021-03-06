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
	"errors"
	"sync"
	"wp-server/wotoPacks/core/wotoValues"
	wa "wp-server/wotoPacks/wotoActions"
	"wp-server/wotoPacks/wotoActions/versioning"
	"wp-server/wotoPacks/wotoActions/wotoGroups"
	"wp-server/wotoPacks/wotoActions/wotoMedia"
	"wp-server/wotoPacks/wotoActions/wotoUsers"
)

//---------------------------------------------------------

var MainListener *wotoValues.WotoListener

//---------------------------------------------------------

var ErrEmptyRequest = errors.New("entryPoints: the incoming request is empty")
var ErrDataLengthInvalid = errors.New("entryPoints: the data length is not valid")
var ErrActionOrBatchInvalid = errors.New("entryPoints: action or batch is invalid")
var ErrConnectionUnavailable = errors.New("entryPoints: connection is no longer available")
var ErrConnectionNotRegistered = errors.New("entryPoints: the connection is not registered")
var ErrListenerNil = errors.New("entryPoints: listener cannot be nil")

//---------------------------------------------------------

var registrationMap = make(map[*wotoValues.WotoConnection]bool)
var registrationMutex = new(sync.Mutex)
var isCheckingRegistration = false

//---------------------------------------------------------

var (
	_handlersMap = map[wa.RequestAction]wotoValues.ReqHandler{
		wa.ActionVersion: versioning.HandleVersionAction,
		wa.ActionUser:    wotoUsers.HandleUserAction,
		wa.ActionMedia:   wotoMedia.HandleMediaAction,
		wa.ActionGroups:  wotoGroups.HandleGroupsAction,
	}

	_parsersMap = map[wa.RequestAction]wotoValues.ReqHandler{
		wa.ActionVersion: versioning.ParseBatchExecute,
		wa.ActionUser:    wotoUsers.ParseBatchExecute,
		wa.ActionMedia:   wotoMedia.ParseBatchExecute,
		wa.ActionGroups:  wotoGroups.ParseBatchExecute,
	}
)
