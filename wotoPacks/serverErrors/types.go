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

type ErrorType int

type EndPointError struct {
	// Type is type of the error.
	Type ErrorType `json:"type"`

	// Message field represents message of the error.
	Message string `json:"message"`

	// Origin field represents the method which has been caused the
	// error to be raised. (the batch execute)
	Origin string `json:"origin"`
}
