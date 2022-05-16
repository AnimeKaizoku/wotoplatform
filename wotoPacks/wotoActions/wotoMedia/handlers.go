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

package wotoMedia

import (
	we "wp-server/wotoPacks/core/wotoErrors"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database/mediaDatabase"
	"wp-server/wotoPacks/interfaces"
	wa "wp-server/wotoPacks/wotoActions"
)

// HandleMediaAction handles the media-action.
func HandleMediaAction(req interfaces.ReqBase) error {
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

//
func batchRegisterMedia(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginRegisterMedia)
	}

	var entryData = new(RegisterMediaData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	media := mediaDatabase.GetMediaByTitle(entryData.Title)
	if media != nil {
		return we.SendMediaTitleAlreadyExists(req, OriginRegisterMedia)
	}

	// TODO: add more checkers here to check for already-existing values.
	media = mediaDatabase.SaveNewMedia(entryData.ToNewMediaData())

	return req.SendResult(&RegisterMediaResult{
		MediaId: media.ModelId,
	})
}

//
func batchGetMediaById(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMediaById)
	}

	var entryData = new(GetMediaByIdData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.MediaId.IsInvalid() {
		return we.SendInvalidMediaId(req, OriginGetMediaById)
	}

	media := mediaDatabase.GetMediaById(entryData.MediaId)
	if media == nil {
		return we.SendMediaNotFound(req, OriginGetMediaById)
	}

	return req.SendResult(media)
}

//
func batchChangeMediaGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMediaById)
	}

	var entryData = new(ChangeMediaGenreData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.MediaId.IsInvalid() {
		return we.SendInvalidMediaId(req, OriginGetMediaById)
	}

	media := mediaDatabase.GetMediaById(entryData.MediaId)
	if media == nil {
		return we.SendMediaNotFound(req, OriginGetMediaById)
	}

	return req.SendResult(media)
}

//
func batchCreateNewGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginCreateNewGenre)
	}

	return we.SendMethodNotImplemented(req, OriginCreateNewGenre)
}

//
func batchDeleteGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteGenre)
	}

	return we.SendMethodNotImplemented(req, OriginDeleteGenre)
}

//
func batchAddMediaGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginAddMediaGenre)
	}

	return we.SendMethodNotImplemented(req, OriginAddMediaGenre)
}

//
func batchRemoveMediaGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginRemoveMediaGenre)
	}

	return we.SendMethodNotImplemented(req, OriginRemoveMediaGenre)
}

//
func batchGetMediaGenres(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMediaGenres)
	}

	return we.SendMethodNotImplemented(req, OriginGetMediaGenres)
}

//
func batchSearchGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginSearchGenre)
	}

	return we.SendMethodNotImplemented(req, OriginSearchGenre)
}

//
func batchDeleteMedia(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteMedia)
	}

	return we.SendMethodNotImplemented(req, OriginDeleteMedia)
}
