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

package mediaDatabase

import (
	"time"
	wv "wp-server/wotoPacks/core/wotoValues"
)

type NewMediaData struct {
	Company     wv.CompanyId
	Author      wv.AuthorId
	Episode     int
	MediaType   string
	Title       string
	Duration    time.Duration
	Artist      string
	Album       string
	Year        int
	Cover       string
	File        string
	Thumbnail   string
	Lyrics      string
	Lang        string
	LangCode    string
	Region      string
	SourceUrl   string
	ExternalUrl string
	IsPrivate   bool
	Description string
}
