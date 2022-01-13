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

import (
	"strconv"
)

func (e *EndPointError) ToString() string {
	return "type of " + e.GetType() + ": message of:" + e.Message
}

func (e *EndPointError) GetType() string {
	str := ErrTypeStrMap[e.Type]
	if str != "" {
		return str
	}

	return strconv.Itoa(int(e.Type))
}
