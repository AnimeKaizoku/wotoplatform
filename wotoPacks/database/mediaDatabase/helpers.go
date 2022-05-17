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

// LoadMediaDatabase
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

// SaveNewMedia
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

// GetMediaByTitle
func GetMediaByTitle(title string) *wv.MediaModel {
	return mediaModelsByTitle.Get(title)
}

// GetMediaById
func GetMediaById(id wv.MediaModelId) *wv.MediaModel {
	return mediaModels.Get(id)
}

// GetGenreInfoById
func GetGenreInfoById(id wv.GenreId) *wv.MediaGenreInfo {
	return mediaGenreInfos.Get(id)
}

// GetGenreInfoByTitle
func GetGenreInfoByTitle(title string) *wv.MediaGenreInfo {
	return mediaGenreInfosByTitle.Get(title)
}

// SaveNewGenreInfo saves the info into db and caches it in
// the memory.
func SaveNewGenreInfo(info *wv.MediaGenreInfo) {
	SaveNewGenreInfoNoCache(info)

	mediaGenreInfos.Add(info.GenreId, info)
	mediaGenreInfosByTitle.Add(info.GenreTitle, info)
}

// UpdateGenreInfo updates the info in db and caches it in
// the memory.
func UpdateGenreInfo(info *wv.MediaGenreInfo) {
	DeleteGenreInfoFromCache(info)
	SaveNewGenreInfoNoCache(info)

	mediaGenreInfos.Add(info.GenreId, info)
	mediaGenreInfosByTitle.Add(info.GenreTitle, info)
}

// SaveNewGenreInfoNoCache saves the info into db only, this
// function won't touch cache.
func SaveNewGenreInfoNoCache(info *wv.MediaGenreInfo) {
	if info.GenreId.IsInvalid() {
		info.GenreId = genreInfoIdGenerator.Next()
	}

	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(info)
	tx.Commit()
	unlockDatabase()
}

// AddMediaGenre adds a specific genre (using its genre-id) to a
// media-model.
// this function does not validate the genre-id value passed to it;
// the caller has to make sure both media-model and genre-id are valid
// before using this function.
func AddMediaGenre(media *wv.MediaModel, id wv.GenreId) {
	element := &wv.MediaGenreElement{
		MediaId: media.ModelId,
		Genre:   id,
	}

	AddMediaGenreElement(media, element)
}

// AddMediaGenreElement adds the target element to the database and caches
// it in memory.
func AddMediaGenreElement(media *wv.MediaModel, element *wv.MediaGenreElement) {
	SaveNewMediaGenreElementNoCache(element)

	genreInfo := mediaGenreInfos.Get(element.Genre)
	elements := mediaGenreElements.GetValue(element.MediaId)
	elements = append(elements, element)
	mediaGenreElements.Set(element.MediaId, elements)

	elements = mediaGenreElementsByGenreId.GetValue(element.Genre)
	elements = append(elements, element)
	mediaGenreElementsByGenreId.Set(element.Genre, elements)

	media.Genres = append(media.Genres, genreInfo)
}

// SaveNewMediaGenreElementNoCache tries to generate a new unique-id for
// the passed element and save it to the database.
// this function doesn't cache the element.
func SaveNewMediaGenreElementNoCache(element *wv.MediaGenreElement) {
	element.GenerateNewUniqueId()
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(element)
	tx.Commit()
	unlockDatabase()
}

// SaveMediaModel
func SaveMediaModel(media *wv.MediaModel) {
	SaveMediaModelNoCache(media)

	mediaModels.Add(media.ModelId, media)
	mediaModelsByTitle.Add(media.Title, media)
}

// SaveMediaModelNoCache saves the target media-model in the database.
// this function doesn't cache the media-model into memory.
func SaveMediaModelNoCache(media *wv.MediaModel) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(media)
	tx.Commit()
	unlockDatabase()
}

// DeleteMediaGenreElement
func DeleteMediaGenreElement(element *wv.MediaGenreElement) {
	lockDatabase()
	deleteMediaGenreElement(element)
	unlockDatabase()

	mediaGenreElements.Delete(element.MediaId)
	mediaGenreElementsByGenreId.Delete(element.Genre)
}

// deleteMediaGenreElement
func deleteMediaGenreElement(element *wv.MediaGenreElement) {
	wv.SESSION.Delete(element)
}

// DeleteGenreInfo deletes the specified genre-info.
// this function doesn't check for existence of the passed info,
// it's up to caller to validate the info.
func DeleteGenreInfo(info *wv.MediaGenreInfo) {
	lockDatabase()
	deleteGenreInfo(info)
	unlockDatabase()

	DeleteGenreInfoFromCache(info)
}

// DeleteGenreInfoFromCache will deletes the genre info only from memory cache.
func DeleteGenreInfoFromCache(info *wv.MediaGenreInfo) {
	mediaGenreInfos.Delete(info.GenreId)
	mediaGenreInfosByTitle.Delete(info.GenreTitle)
}

func deleteGenreInfo(info *wv.MediaGenreInfo) {
	wv.SESSION.Delete(info)
}

// lockDatabase
func lockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Lock()
	}
}

// unlockDatabase
func unlockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Unlock()
	}
}
