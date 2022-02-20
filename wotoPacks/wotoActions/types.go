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

package wotoActions

import (
	"wp-server/wotoPacks/serverErrors"

	wcr "github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

type RequestAction uint
type BatchExecution string

type ActionResp struct {
	Success bool                        `json:"success"`
	Error   *serverErrors.EndPointError `json:"error"`
	Result  interface{}                 `json:"result"`
	Keys    wcr.KeyCollection           `json:"keys"`
}
