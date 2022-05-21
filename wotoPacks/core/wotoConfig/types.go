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

package wotoConfig

// config struct types defined in config.json
type Config struct {
	Name                   string      `json:"name"`
	Network                string      `json:"network"`
	Bind                   string      `json:"bind"`
	Port                   string      `json:"port"`
	Description            string      `json:"description"`
	Owners                 []WotoOwner `json:"owners"`
	ClientIDs              []string    `json:"client_ids"`
	UseCrypto              bool        `json:"use_crypto"`
	UseTLS                 bool        `json:"use_tls"`
	DatabaseUrl            string      `json:"db_url"`
	DatabaseName           string      `json:"db_name"`
	UseSQLite              bool        `json:"use_sqlite"`
	SkipDefaultTransaction bool        `json:"skip_default_transaction"`
	IsDebug                bool        `json:"is_debug"`
	IsDefault              bool        `json:"-"` // test purpose only
}

type WotoOwner struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
