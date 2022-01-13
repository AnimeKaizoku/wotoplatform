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
	"wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/wotoActions"
)

type RequestEntry struct {
	// Action is the request entry's request action.
	Action wotoActions.RequestAction `json:"action"`

	// BatchExecute is the batch execution values that client
	// wants us to execute. it should be formatted in special format.
	BatchExecute string `json:"batch_execute"`

	// Data is the specific data that client sends alongside of
	// its request. the data will be in different formats for
	// the different packages and different batch executions.
	Data string `json:"data"`

	// Connection field is the connection between the client and
	// server; use this field to communicate with the client.
	// this field should be ignored by json decoder.
	// when this request entry is in exit mode, this value
	// will be nil.
	Connection *wotoValues.WotoConnection `json:"-"`

	// exit tells us if should we exit from the batch execution
	// and let the client to execute another batch execution?
	// this field should be ignored by json decoder.
	exit bool `json:"-"`

	// batchValues is an array of batch execution values.
	// when the client wants to execute a batch execution, it
	// can send more that one batch execution, they will be
	// executed and the `reqBase` will be provided to the
	// specific handler in the specific package; that handler
	// can block its connection's goroutine as much as it wants
	// and it can communicate with the client there.
	// when it wants to end its commination (like if battle is over),
	// it needs to call the method.
	// when this request entry is in exit mode, this value
	// will be nil.
	batchValues []wotoActions.BatchExecution `json:"-"`
}
