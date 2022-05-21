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
	BATCH_REGISTER_MEDIA        = "register_media"
	BATCH_GET_MEDIA_BY_ID       = "get_media_by_id"
	BATCH_CREATE_NEW_GENRE      = "create_new_genre"
	BATCH_DELETE_GENRE_INFO     = "delete_genre_info"
	BATCH_EDIT_GENRE_INFO       = "edit_genre_info"
	BATCH_ADD_MEDIA_GENRE       = "add_media_genre"
	BATCH_REMOVE_MEDIA_GENRE    = "remove_media_genre"
	BATCH_GET_MEDIA_GENRES      = "get_media_genres"
	BATCH_GET_MEDIA_GENRES_INFO = "get_media_genres_info"
	BATCH_SEARCH_MEDIA          = "search_media"
	BATCH_DELETE_MEDIA          = "delete_media"
)

// origin constants
const (
	OriginRegisterMedia      = "RegisterMedia"
	OriginGetMediaById       = "GetMediaById"
	OriginCreateNewGenre     = "CreateNewGenre"
	OriginDeleteGenreInfo    = "DeleteGenreInfo"
	OriginEditGenreInfo      = "EditGenreInfo"
	OriginAddMediaGenre      = "AddMediaGenre"
	OriginRemoveMediaGenre   = "RemoveMediaGenre"
	OriginGetMediaGenres     = "GetMediaGenres"
	OriginGetMediaGenresInfo = "GetMediaGenresInfo"
	OriginSearchGenre        = "SearchGenre"
	OriginDeleteMedia        = "DeleteMedia"
)
