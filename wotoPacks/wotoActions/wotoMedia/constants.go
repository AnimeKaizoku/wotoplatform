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

// batch execution values
const (
	BATCH_REGISTER_MEDIA     = "register_media"
	BATCH_GET_MEDIA_BY_ID    = "get_media_by_id"
	BATCH_CHANGE_MEDIA_GENRE = "change_media_genre"
	BATCH_SEARCH_MEDIA       = "search_media"
	BATCH_DELETE_MEDIA       = "delete_media"
)

// origin constants
const (
	OriginRegisterMedia    = "RegisterMedia"
	OriginGetMediaById     = "GetMediaById"
	OriginChangeMediaGenre = "ChangeMediaGenre"
)
