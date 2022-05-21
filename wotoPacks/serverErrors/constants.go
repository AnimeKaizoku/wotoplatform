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

package serverErrors

const (
	ErrNoError ErrorCode = iota
	ErrUnknown
	ErrServerUnavailable
	ErrInvalidUsernameFormat
	ErrInvalidPasswordFormat
	ErrUsernameExists
	ErrWrongUsername
	ErrWrongPassword
	ErrInvalidAuthKeyFormat
	ErrInvalidAccessHashFormat
	ErrWrongAuthKey
	ErrLoginAccessHashExpired
	ErrInvalidFirstName
	ErrInvalidLastName
	ErrInvalidTitle
	ErrAlreadyAuthorized
	ErrNotAuthorized
	ErrNotModified
	ErrBioTooLong
	ErrUserNotFound
	ErrFirstNameTooLong
	ErrLastNameTooLong
	ErrInvalidUsernameAndUserId
	ErrMethodNotImplemented
	ErrPermissionDenied
	ErrKeyNotFound
	ErrInvalidTelegramId
	ErrInvalidEmail
	ErrInvalidKey
	ErrTooManyFavorites
	ErrLikedListElementAlreadyExists
	ErrTooManyLikedList
	ErrInvalidUniqueId
	ErrMediaNotFound
	ErrInvalidMediaId
	ErrMediaTitleAlreadyExists
	ErrInvalidGenreId
	ErrGenreInfoNotFound
	ErrGenreTitleAlreadyExists
)

const (
	StrNoError                = "NoError"
	StrUnknownError           = "UnknownError"
	StrServerUnavailableError = "ServerUnavailableError"
)
