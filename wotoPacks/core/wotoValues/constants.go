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

// the file formats used in this project.
const (
	Jpg  = ".jpg"  // photo : jpg
	Jpeg = ".jpeg" // photo : jpeg
	Gif  = ".gif"  // gif (framed photos?) : gif
	Json = ".json"
	Png  = ".png"  // photo : png
	Wav  = ".wav"  // oh yeah music : wav
	Wave = ".wave" // don't if it's really music: wave
	Mp3  = ".mp3"  // I love mp3 more: mp3
	Mp4  = ".mp4"  // mp4 ...
	M4a  = ".m4a"  // I hate musics with this format, my Nokia can't play them
)

const (
	Yes    = "yes"
	CapYes = "Yes"
	No     = "no"
	CapNo  = "No"
)

// the base constant values.
const (
	BaseIndex       = 0  // number 0
	BaseOneIndex    = 1  // number 1
	BaseTwoIndex    = 2  // number 2
	BaseThreeIndex  = 3  // number 2
	Base4Bit        = 4  // number 8
	Base8Bit        = 8  // number 8
	Base16Bit       = 16 // number 16
	Base32Bit       = 32 // number 32
	Base64Bit       = 64 // number 64
	BaseTimeOut     = 40 // 40 seconds
	BaseTen         = 10 // 10 seconds
	MAX_FIRST_BYTES = 8  // the max first bytes for sending the length
)

const (
	BaseIndexStr    = "0"  // number 0
	BaseOneIndexStr = "1"  // number 1
	DotStr          = "."  // dot : .
	LineStr         = "-"  // line : -
	EMPTY           = ""   //an empty string.
	UNDER           = "_"  // an underscope : _
	STR_SIGN        = "\"" // the string sign : "
	CHAR_STR        = '"'  // the string sign : '"'
	OrRegexp        = "|"  // the or string sign: "|"
)

const (
	LineChar   = '-' // line : '-'
	EqualChar  = '=' // equal: '='
	SpaceChar  = ' ' // space: ' '
	DPointChar = ':' // double point: ':'
)
