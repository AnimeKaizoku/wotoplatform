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

	lockDatabase()
	wv.SESSION.Find(&allMedias)
	unlockDatabase()

	for _, media := range allMedias {
		mediaModels.Add(media.ModelId, media)
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
