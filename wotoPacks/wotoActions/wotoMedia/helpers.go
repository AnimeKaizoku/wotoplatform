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
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/interfaces"
	wa "wp-server/wotoPacks/wotoActions"
)

func ParseBatchExecute(b interfaces.ReqBase) error {
	arr, err := wa.ParseBatchExecute(b.GetBatchExecute())
	if err != nil {
		return err
	}

	if len(arr) == 0 || !batchValuesValid(arr) {
		return wa.ErrBatchParseFailed
	}

	b.SetBatchValues(arr)

	return nil
}

func batchValuesValid(e []wa.BatchExecution) bool {
	for _, b := range e {
		if _batchHandlers[b] == nil {
			return false
		}
	}

	return true
}

func toGetMediaByIdResult(media *wv.MediaModel) *GetMediaByIdResult {
	return &GetMediaByIdResult{
		ModelId:     media.ModelId,
		Genres:      media.Genres,
		Company:     media.Company,
		Author:      media.Author,
		Episode:     media.Episode,
		MediaType:   media.MediaType,
		Title:       media.Title,
		Duration:    media.Duration,
		Artist:      media.Artist,
		Album:       media.Album,
		Year:        media.Year,
		Cover:       media.Cover,
		File:        media.File,
		Thumbnail:   media.Thumbnail,
		Lyrics:      media.Lyrics,
		Lang:        media.Lang,
		LangCode:    media.LangCode,
		Region:      media.Region,
		SourceUrl:   media.SourceUrl,
		ExternalUrl: media.ExternalUrl,
		IsPrivate:   media.IsPrivate,
		Description: media.Description,
		CreatedAt:   media.CreatedAt,
		CreatedBy:   media.CreatedBy,
		UpdatedBy:   media.UpdatedBy,
	}
}
