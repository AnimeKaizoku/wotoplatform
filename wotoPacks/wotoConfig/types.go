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

package wotoConfig

// config struct types defined in config.json
type Config struct {
	Name           string   `json:"name"`
	Network        string   `json:"network"`
	Bind           string   `json:"bind"`
	Port           string   `jsong:"port"`
	Description    string   `json:"description"`
	AdminPasswords []string `json:"admin_passwords"`
	AdminUsers     []string `json:"admin_users"`
	ClientIDs      []string `json:"client_ids"`
	DB_URL         string   `json:"db_url"`
	DB_Name        string   `json:"db_name"`
	UseSQLLite     bool     `json:"use_sql_lite"`
	IsDefault      bool     `json:"-"` // test purpose only
}
