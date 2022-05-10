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

func SaveNewMedia(m *NewMediaData) {

}

func SaveMediaModel(media *wv.MediaModel, cache bool) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(media)
	tx.Commit()
	unlockDatabase()

	if cache {
		mediaModels.Add(media.ModelId, media)
	}
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
