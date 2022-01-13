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

import (
	"net"
	"time"
	"wp-server/wotoPacks/interfaces"
)

type PublicUserId int64
type UserPermission int
type PublicGroupCallId string
type SongModelId string
type ProfilePictureModelId string

type ReqHandler func(interfaces.ReqBase) error
type Registerer func(*WotoConnection)

type SongModel struct {
	ModelId     SongModelId   `json:"model_id" gorm:"primaryKey"`
	Title       string        `json:"title"`
	Duration    time.Duration `json:"duration"`
	Artist      string        `json:"artist"`
	Album       string        `json:"album"`
	Genre       string        `json:"genre"`
	Year        int           `json:"year"`
	Cover       string        `json:"cover"`
	File        string        `json:"file"`
	Thumbnail   string        `json:"thumbnail"`
	Lyrics      string        `json:"lyrics"`
	Lang        string        `json:"lang"`
	LangCode    string        `json:"lang_code"`
	Region      string        `json:"region"`
	SourceUrl   string        `json:"source_url"`
	TelegramUrl string        `json:"telegram_url"`
	CreatedAt   time.Time     `json:"created_at"`
	CreatedBy   PublicUserId  `json:"created_by"`
	UpdatedBy   PublicUserId  `json:"updated_by"`
}

type GroupCallInfo struct {
	GroupCallId       PublicGroupCallId `json:"group_call_id" gorm:"primaryKey"`
	GroupCallUsername string            `json:"group_call_username"`
	TelegramId        int64             `json:"telegram_id"`
	TelegramUsername  string            `json:"telegram_username"`
	CurrentPlaying    SongModelId       `json:"current_playing"`
}

type UserSongHistoryElement struct {
	UserId        PublicUserId      `json:"user_id" gorm:"primaryKey"`
	AtGroupCallId PublicGroupCallId `json:"at_group_call_id"`
	Song          SongModelId       `json:"song_model_id"`
	PlayedBy      PublicUserId      `json:"played_by"`
}

type GroupCallSongHistoryElement struct {
	GroupCallId PublicGroupCallId `json:"group_call_id" gorm:"primaryKey"`
}

type SongTag struct {
	ModelId   string    `json:"model_id" gorm:"primaryKey"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}

type UserInfo struct {
	UserId         PublicUserId          `json:"user_id" gorm:"primaryKey"`
	PrivateHash    string                `json:"private_hash"`
	Password       string                `json:"password"`
	Permission     UserPermission        `json:"-"`
	Bio            string                `json:"bio"`
	ProfilePicture ProfilePictureModelId `json:"profile_picture"`
	SourceUrl      string                `json:"source_url"`
	TelegramId     int64                 `json:"telegram_id"`
	FirstName      string                `json:"first_name"`
	LastName       string                `json:"last_name"`
	Username       string                `json:"username"`
	CreatedAt      time.Time             `json:"created_at"`
	UpdatedAt      time.Time             `json:"updated_at"`
	IsVirtual      bool                  `json:"is_virtual"`
	CreatedBy      PublicUserId          `json:"created_by"`
}

type WotoListener struct {
	listener net.Listener
	isClosed bool
}

type WotoConnection struct {
	conn         net.Conn
	origin       *WotoListener
	isClosed     bool
	isRegistered bool
	registerer   Registerer
}
