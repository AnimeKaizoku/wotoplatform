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
	wa "wp-server/wotoPacks/wotoActions"
)

var (
	_batchHandlers = map[wa.BatchExecution]wv.ReqHandler{
		BATCH_REGISTER_MEDIA:     batchRegisterMedia,
		BATCH_GET_MEDIA_BY_ID:    batchGetMediaById,
		BATCH_CREATE_NEW_GENRE:   batchCreateNewGenre,
		BATCH_DELETE_GENRE:       batchDeleteGenre,
		BATCH_EDIT_GENRE_INFO:    batchEditGenreInfo,
		BATCH_ADD_MEDIA_GENRE:    batchAddMediaGenre,
		BATCH_REMOVE_MEDIA_GENRE: batchRemoveMediaGenre,
		BATCH_GET_MEDIA_GENRES:   batchGetMediaGenres,
		BATCH_SEARCH_MEDIA:       batchSearchGenre,
		BATCH_DELETE_MEDIA:       batchDeleteMedia,
	}
)
