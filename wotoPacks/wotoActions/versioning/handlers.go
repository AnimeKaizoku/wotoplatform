/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
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

package versioning

import (
	"strings"
	"wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/utils/logging"
	"wp-server/wotoPacks/utils/wotoTime"
	"wp-server/wotoPacks/wotoActions"
	"wp-server/wotoPacks/wotoConfig"
)

func HandleVersionAction(req interfaces.ReqBase) error {
	logging.Debug("received versioning action")
	//log.Println("received versioning action")
	b := req.GetBatchValues()
	var err error
	for _, ex := range b {
		switch ex {
		case BATCH_CHECK_VERSION:
			err = batchCheckVersion(req)
			if err != nil {
				logging.Debug("an error while executing batch execution: ", err)
				//log.Println("an error while executing batch execution: ", err)
				return err
			}

			continue
		default:
			logging.Warn("invalid batch:", ex)
			//log.Println("invalid batch")
			return wotoActions.ErrInvalidBatch
		}
	}

	req.LetExit()

	return nil
}

func batchCheckVersion(req interfaces.ReqBase) error {
	var entryData checkVersionEntry
	err := req.ParseJsonData(&entryData)
	if err != nil {
		logging.Error(err)
		return err
	}

	if !strings.EqualFold(entryData.UserAgent, userAgentValue) {
		logging.Error("user-agent wasn't correct")
		return wotoActions.ErrInvalidBatch
	}

	if !wotoConfig.IsClientIDValid(entryData.ClientID) {
		logging.Error("client id wasn't correct")
		return wotoActions.ErrInvalidBatch
	}

	logging.Debug("trying to send the json")
	//log.Println("trying to send the json")

	a := VersionAcceptable(entryData.VersionKey, entryData.VersionHashKey)
	if a {
		// if the client is verified and its version and version hash are
		// acceptable, register the connection, so it can execute another
		// batch executions as well
		req.RegisterConnection()
	}

	_, err = req.WriteResult(&VersionResults{
		IsAcceptable: a,
		ServerTime:   wotoTime.GenerateCurrentDateTime(),
	})

	if err != nil {
		logging.Error(err)
	}

	return err
}
