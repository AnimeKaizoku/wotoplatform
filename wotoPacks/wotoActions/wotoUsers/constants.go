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

package wotoUsers

// batch execution values
const (
	BATCH_REGISTER_USER               = "register_user"
	BATCH_REGISTER_VIRTUAL_USER       = "register_virtual_user"
	BATCH_LOGIN_USER                  = "login_user"
	BATCH_GET_ME                      = "get_me"
	BATCH_CHANGE_USER_BIO             = "change_user_bio"
	BATCH_CHANGE_NAMES                = "change_names"
	BATCH_GET_USER_INFO               = "get_user_info"
	BATCH_GET_USER_BY_TELEGRAM_ID     = "get_user_by_telegram_id"
	BATCH_GET_USER_BY_EMAIL           = "get_user_by_email"
	BATCH_CHANGE_USER_PERMISSION      = "change_user_permission"
	BATCH_RESOLVE_USERNAME            = "resolve_username"
	BATCH_GET_USER_FAVORITE           = "get_user_favorite"
	BATCH_GET_USER_FAVORITE_COUNT     = "get_user_favorite_count"
	BATCH_SET_USER_FAVORITE           = "set_user_favorite"
	BATCH_DELETE_USER_FAVORITE        = "delete_user_favorite"
	BATCH_GET_USER_LIKED_LIST         = "get_user_liked_list"
	BATCH_GET_USER_LIKED_LIST_COUNT   = "get_user_liked_list_count"
	BATCH_APPEND_USER_LIKED_LIST      = "append_user_liked_list"
	BATCH_DELETE_USER_LIKED_LIST_ITEM = "delete_user_liked_list_item"
)

// origin constants
const (
	OriginRegisterUser            = "RegisterUser"
	OriginLoginUser               = "LoginUser"
	OriginGetMe                   = "GetMe"
	OriginChangeUserBio           = "ChangeUserBio"
	OriginChangeNames             = "ChangeNames"
	OriginGetUserInfo             = "GetUserInfo"
	OriginRegisterVirtualUser     = "RegisterVirtualUser"
	OriginGetUserByTelegramID     = "GetUserByTelegramID"
	OriginGetUserByEmail          = "GetUserByEmail"
	OriginResolveUsername         = "ResolveUsername"
	OriginChangeUserPermission    = "ChangeUserPermission"
	OriginGetUserFavorite         = "GetUserFavorite"
	OriginGetUserFavoriteCount    = "GetUserFavoriteCount"
	OriginSetUserFavorite         = "SetUserFavorite"
	OriginDeleteUserFavorite      = "DeleteUserFavorite"
	OriginGetUserLikedList        = "GetUserLikedList"
	OriginGetUserLikedListCount   = "GetUserLikedListCount"
	OriginAppendUserLikedList     = "AppendUserLikedList"
	OriginDeleteUserLikedListItem = "DeleteUserLikedListItem"
)
