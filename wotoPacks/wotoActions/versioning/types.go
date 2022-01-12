/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
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

package versioning

type Version struct {
	Num1 uint8
	Num2 uint8
	Num3 uint8
	Num4 uint8
}

type VersionResults struct {
	IsAcceptable bool   `json:"is_acceptable"`
	ServerTime   string `json:"server_time"`
}

type checkVersionEntry struct {
	UserAgent      string `json:"user_agent"`
	VersionKey     string `json:"version_key"`
	VersionHashKey string `json:"version_hash"`
	ClientID       string `json:"client_id"`
}
