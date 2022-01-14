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
	BATCH_REGISTER_USER    = "register_user"
	BATCH_LOGIN_USER       = "login_user"
	BATCH_GET_ME           = "get_info"
	BATCH_GET_USER_INFO    = "get_user_info"
	BATCH_RESOLVE_USERNAME = "resolve_username"
)

// error types
const (
	ErrTypeUsernameExists = iota + 1
	ErrTypeUserPassInvalid
)

// error messages
const (
	ErrMsgUsernameExists  = "username is already registered in database"
	ErrMsgUserPassInvalid = "username or password are entered in a wrong format"
)

const (
	OriginRegisterUser = "RegisterUser"
	OriginLoginUser    = "LoginUser"
)
