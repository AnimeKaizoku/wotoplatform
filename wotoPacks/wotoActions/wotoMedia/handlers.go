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
	"strings"
	we "wp-server/wotoPacks/core/wotoErrors"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"
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

// batchRegisterMedia handler registers a new media-model.
func batchRegisterMedia(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginRegisterMedia)
	}

	meta := req.GetMe().GetMeta()
	if meta != nil && !meta.GetBoolNoErr("can_create_media_model") {
		return we.SendPermissionDenied(req, OriginRegisterMedia)
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

// batchGetMediaById handler returns a GetMediaByIdResult variable to
// the client.
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

	return req.SendResult(toGetMediaByIdResult(media))
}

// batchCreateNewGenre handler creates a new genre.
func batchCreateNewGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginCreateNewGenre)
	}

	me := req.GetMe()
	meta := me.GetMeta()
	if meta != nil && !meta.GetBoolNoErr("can_create_genre_info") {
		return we.SendPermissionDenied(req, OriginCreateNewGenre)
	}

	var entryData = new(CreateNewGenreData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	entryData.GenreTitle = strings.TrimSpace(entryData.GenreTitle)
	if !wotoValidate.IsTitleValid(entryData.GenreTitle) {
		return we.SendInvalidTitle(req, OriginCreateNewGenre)
	}

	info := entryData.ToMediaGenreInfo(me.UserId)
	mediaDatabase.SaveNewGenreInfo(info)

	return req.SendResult(&CreateNewGenreResult{
		GenreId: info.GenreId,
	})
}

// batchDeleteGenre handler deletes the specified genre-info
// (client has to pass the genre-id) from db.
func batchDeleteGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteGenre)
	}

	return we.SendMethodNotImplemented(req, OriginDeleteGenre)
} //

// batchEditGenreInfo handler edits the specified genre-info.
func batchEditGenreInfo(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteGenre)
	}

	return we.SendMethodNotImplemented(req, OriginDeleteGenre)
}

// batchAddMediaGenre handler adds the specified genre-info to the
// target media-model.
func batchAddMediaGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginAddMediaGenre)
	}

	var entryData = new(AddMediaGenreData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.MediaId.IsInvalid() {
		return we.SendInvalidMediaId(req, OriginAddMediaGenre)
	}

	media := mediaDatabase.GetMediaById(entryData.MediaId)
	if media == nil {
		return we.SendMediaNotFound(req, OriginAddMediaGenre)
	}

	mediaDatabase.AddMediaGenre(media, entryData.MediaGenre)

	return req.SendResult(&AddMediaGenreResult{
		MediaId:     media.ModelId,
		MediaGenres: media.GetGenreIDs(),
	})
}

// batchRemoveMediaGenre handler removes the specified genre-info
// from the target media-model.
func batchRemoveMediaGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginRemoveMediaGenre)
	}

	return we.SendMethodNotImplemented(req, OriginRemoveMediaGenre)
}

// batchGetMediaGenres handler returns all genre-infos that the target
// media-model contains.
func batchGetMediaGenres(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMediaGenres)
	}

	return we.SendMethodNotImplemented(req, OriginGetMediaGenres)
}

// batchSearchGenre handler searches the given title in database
// and finds most suitable results to return to the client.
func batchSearchGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginSearchGenre)
	}

	return we.SendMethodNotImplemented(req, OriginSearchGenre)
}

// batchDeleteMedia handler will delete the target media-model
// from the database. this handler will return error if the user is not
// the creator of the media-model (administrators can delete any media, even if
// they are not the creator of the media-model).
func batchDeleteMedia(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteMedia)
	}

	return we.SendMethodNotImplemented(req, OriginDeleteMedia)
}
