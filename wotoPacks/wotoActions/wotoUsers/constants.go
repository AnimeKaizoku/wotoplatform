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
	BATCH_REGISTER_USER           = "register_user"
	BATCH_LOGIN_USER              = "login_user"
	BATCH_GET_ME                  = "get_me"
	BATCH_CHANGE_USER_BIO         = "change_user_bio"
	BATCH_CHANGE_NAMES            = "change_names"
	BATCH_GET_USER_INFO           = "get_user_info"
	BATCH_GET_USER_BY_TELEGRAM_ID = "get_user_by_telegram_id"
	BATCH_GET_USER_BY_EMAIL       = "get_user_by_email"
	BATCH_RESOLVE_USERNAME        = "resolve_username"
)
const (
	OriginRegisterUser  = "RegisterUser"
	OriginLoginUser     = "LoginUser"
	OriginGetMe         = "GetMe"
	OriginChangeUserBio = "ChangeUserBio"
	OriginChangeNames   = "ChangeNames"
	OriginGetUserInfo   = "GetUserInfo"
)
