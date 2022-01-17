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

package wotoValues

// the base constant values.
const (
	BaseIndex                       = 0  // number 0
	BaseOneIndex                    = 1  // number 1
	BaseTwoIndex                    = 2  // number 2
	BaseThreeIndex                  = 3  // number 2
	Base4Bit                        = 4  // number 8
	Base8Bit                        = 8  // number 8
	Base16Bit                       = 16 // number 16
	Base32Bit                       = 32 // number 32
	Base64Bit                       = 64 // number 64
	BaseTimeOut                     = 40 // 40 seconds
	MaxBioLength                    = 200
	MaxFirstNameLength              = 50
	MaxLastNameLength               = 50
	BaseTen                         = 10 // 10 seconds
	BaseUserId         PublicUserId = 10000
	MAX_FIRST_BYTES                 = 8 // the max first bytes for sending the length
)

const (
	PermissionNormalUser UserPermission = iota
	PermissionSpecial
	PermissionAdmin
	PermissionDeveloper
	PermissionOwner
)

const (
	BaseIndexStr    = "0" // number 0
	BaseOneIndexStr = "1" // number 1
	DateTimeFormat  = "2006-01-02 15:04:05"
)

const (
	SpaceChar = ' ' // space: ' '
)
