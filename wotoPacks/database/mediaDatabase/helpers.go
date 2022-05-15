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

package mediaDatabase

import (
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
)

func LoadMediaDatabase() error {
	var allMedias []*wv.MediaModel
	var allGenreElements []*wv.MediaGenreElement
	var allGenreInfo []*wv.MediaGenreInfo

	lockDatabase()
	wv.SESSION.Find(&allMedias)
	wv.SESSION.Find(&allGenreElements)
	wv.SESSION.Find(&allGenreInfo)
	unlockDatabase()

	for _, media := range allMedias {
		mediaModels.Add(media.ModelId, media)
	}

	for _, info := range allGenreInfo {
		mediaGenreInfos.Add(info.GenreId, info)
		mediaGenreInfosByTitle.Add(info.GenreTitle, info)
	}

	for _, element := range allGenreElements {
		media := mediaModels.Get(element.MediaId)
		genreInfo := mediaGenreInfos.Get(element.Genre)
		if media == nil || genreInfo == nil {
			// the genre-element belongs to a media-model that's already
			// deleted from the database OR the genre info which this genre-element
			// is reffering to is invalid (perhaps has been removed
			// from the database), remove the genre-element.
			deleteMediaGenreElement(element)
			continue
		}

		elements := mediaGenreElements.GetValue(element.MediaId)
		elements = append(elements, element)
		mediaGenreElements.Set(element.MediaId, elements)

		elements = mediaGenreElementsByGenreId.GetValue(element.Genre)
		elements = append(elements, element)
		mediaGenreElementsByGenreId.Set(element.Genre, elements)

		media.Genres = append(media.Genres, genreInfo)
	}

	return nil
}

func SaveNewMedia(m *NewMediaData) *wv.MediaModel {
	model := &wv.MediaModel{
		Company:     m.Company,
		Author:      m.Author,
		Episode:     m.Episode,
		MediaType:   m.MediaType,
		Title:       m.Title,
		Duration:    m.Duration,
		Artist:      m.Artist,
		Album:       m.Album,
		Year:        m.Year,
		Cover:       m.Cover,
		File:        m.File,
		Thumbnail:   m.Thumbnail,
		Lyrics:      m.Lyrics,
		Lang:        m.Lang,
		LangCode:    m.LangCode,
		Region:      m.Region,
		SourceUrl:   m.SourceUrl,
		ExternalUrl: m.ExternalUrl,
		IsPrivate:   m.IsPrivate,
		Description: m.Description,
	}

	SaveMediaModel(model)
	return model
}

func GetMediaByTitle(title string) *wv.MediaModel {
	return mediaModelsByTitle.Get(title)
}

func GetMediaById(id wv.MediaModelId) *wv.MediaModel {
	return mediaModels.Get(id)
}

func GetGenreInfoById(id wv.GenreId) *wv.MediaGenreInfo {
	return mediaGenreInfos.Get(id)
}

func GetGenreInfoByTitle(title string) *wv.MediaGenreInfo {
	return mediaGenreInfosByTitle.Get(title)
}

func SaveMediaModel(media *wv.MediaModel) {
	SaveMediaModelNoCache(media)

	mediaModels.Add(media.ModelId, media)
	mediaModelsByTitle.Add(media.Title, media)
}

func SaveMediaModelNoCache(media *wv.MediaModel) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(media)
	tx.Commit()
	unlockDatabase()
}

func DeleteMediaGenreElement(element *wv.MediaGenreElement) {
	lockDatabase()
	deleteMediaGenreElement(element)
	unlockDatabase()
}

func deleteMediaGenreElement(element *wv.MediaGenreElement) {
	wv.SESSION.Delete(element)
}

func lockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Lock()
	}
}

func unlockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Unlock()
	}
}
