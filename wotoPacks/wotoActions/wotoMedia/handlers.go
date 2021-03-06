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
func batchDeleteGenreInfo(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginDeleteGenreInfo)
	}

	me := req.GetMe()
	meta := me.GetMeta()
	if meta != nil && !meta.GetBoolNoErr("can_delete_genre_info") {
		return we.SendPermissionDenied(req, OriginCreateNewGenre)
	}

	var entryData = new(DeleteGenreInfoData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	var info *wv.MediaGenreInfo

	if entryData.GenreId.IsInvalid() {
		if entryData.GenreTitle == "" {
			return we.SendInvalidGenreId(req, OriginDeleteGenreInfo)
		}

		info = mediaDatabase.GetGenreInfoByTitle(entryData.GenreTitle)
	} else {
		info = mediaDatabase.GetGenreInfoById(entryData.GenreId)
	}

	if info == nil {
		return we.SendGenreInfoNotFound(req, OriginDeleteGenreInfo)
	}

	mediaDatabase.DeleteGenreInfo(info)

	return req.SendResult(true)
}

// batchEditGenreInfo handler edits the specified genre-info.
func batchEditGenreInfo(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginEditGenreInfo)
	}

	me := req.GetMe()
	meta := me.GetMeta()
	if meta != nil && !meta.GetBoolNoErr("can_edit_genre_info") {
		return we.SendPermissionDenied(req, OriginCreateNewGenre)
	}

	var entryData = new(EditGenreInfoData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	entryData.GenreTitle = strings.TrimSpace(entryData.GenreTitle)

	var info *wv.MediaGenreInfo

	if entryData.GenreId.IsInvalid() {
		if entryData.GenreTitle == "" {
			return we.SendInvalidGenreId(req, OriginDeleteGenreInfo)
		}

		info = mediaDatabase.GetGenreInfoByTitle(entryData.GenreTitle)
	} else {
		info = mediaDatabase.GetGenreInfoById(entryData.GenreId)
		if info != nil && entryData.GenreTitle != "" {
			tmpInfo := mediaDatabase.GetGenreInfoByTitle(entryData.GenreTitle)
			if tmpInfo != nil && tmpInfo.GenreTitle != info.GenreTitle {
				// this title already exists in db, return error
				return we.SendGenreTitleAlreadyExists(req, OriginEditGenreInfo)
			}
		}
	}

	if info == nil {
		return we.SendGenreInfoNotFound(req, OriginDeleteGenreInfo)
	}

	entryData.UpdateGenreInfoFields(info, me.UserId)
	mediaDatabase.SaveNewGenreInfo(info)

	return req.SendResult(&EditGenreInfoResult{
		GenreInfo: info,
	})
}

// batchAddMediaGenre handler adds the specified genre-info to the
// target media-model.
func batchAddMediaGenre(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginAddMediaGenre)
	}

	me := req.GetMe()

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

	if !media.HasGenreId(entryData.MediaGenre) {
		mediaDatabase.AddMediaGenre(media, entryData.MediaGenre, me.UserId)
	}

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

	var entryData = new(RemoveMediaGenreData)
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

	element := media.RemoveGenreElement(entryData.MediaGenre)
	if element != nil {
		// there is a possibility that media.RemoveGenreElement method
		// returns a nil value; we won't return any error in that case,
		// this checker here is to make sure we don't waste any resources
		// in sending useless db queries.
		mediaDatabase.DeleteMediaGenreElement(element)
	}

	return req.SendResult(&RemoveAddMediaGenreResult{
		MediaId:     media.ModelId,
		MediaGenres: media.GetGenreIDs(),
	})
}

// batchGetMediaGenres handler returns all genre-ids that the target
// media-model contains.
func batchGetMediaGenres(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMediaGenres)
	}

	var entryData = new(GetMediaGenresData)
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

	return req.SendResult(&GetMediaGenresResult{
		MediaId:     media.ModelId,
		MediaGenres: media.GetGenreIDs(),
	})
}

// batchGetMediaGenresInfo handler returns all genre-infos that the target
// media-model contains.
func batchGetMediaGenresInfo(req interfaces.ReqBase) error {
	if !req.IsAuthorized() {
		return we.SendNotAuthorized(req, OriginGetMediaGenresInfo)
	}

	var entryData = new(GetMediaGenresInfoData)
	err := req.ParseJsonData(entryData)
	if err != nil {
		return err
	}

	if entryData.MediaId.IsInvalid() {
		return we.SendInvalidMediaId(req, OriginGetMediaGenresInfo)
	}

	media := mediaDatabase.GetMediaById(entryData.MediaId)
	if media == nil {
		return we.SendMediaNotFound(req, OriginGetMediaGenresInfo)
	}

	return req.SendResult(&GetMediaGenresInfoResult{
		MediaId:         media.ModelId,
		MediaGenresInfo: media.Genres,
	})
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
