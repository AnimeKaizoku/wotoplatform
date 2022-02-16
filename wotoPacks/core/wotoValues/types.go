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
	"wp-server/wotoPacks/core/wotoCrypto"
	"wp-server/wotoPacks/core/wotoValues/wotoRaw"
	"wp-server/wotoPacks/interfaces"
)

type PublicUserId = wotoRaw.PublicUserId
type UserPermission = wotoRaw.UserPermission
type PublicGroupId = wotoRaw.PublicGroupId
type MediaModelId = wotoRaw.MediaModelId
type ProfilePictureModelId = wotoRaw.ProfilePictureModelId

type ReqHandler func(interfaces.ReqBase) error
type Registerer func(*WotoConnection)

type MediaModel struct {
	ModelId     MediaModelId  `json:"model_id" gorm:"primaryKey"`
	MediaType   string        `json:"media_type"`
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

type GroupInfo struct {
	GroupId          PublicGroupId `json:"group_id" gorm:"primaryKey"`
	GroupRegion      string        `json:"group_region"`
	GroupUsername    string        `json:"group_username"`
	TelegramId       int64         `json:"telegram_id"`
	TelegramUsername string        `json:"telegram_username"`
	CurrentPlaying   MediaModelId  `json:"current_playing"`
}

type UserMediaHistoryElement struct {
	UserId    PublicUserId  `json:"user_id" gorm:"primaryKey"`
	AtGroupId PublicGroupId `json:"at_group_id"`
	Media     MediaModelId  `json:"media_model_id"`
	PlayedBy  PublicUserId  `json:"played_by"`
}

type GroupMediaHistoryElement struct {
	GroupId PublicGroupId `json:"group_id" gorm:"primaryKey"`
	Media   MediaModelId  `json:"media_model_id"`
}

type MediaTag struct {
	ModelId   string    `json:"model_id" gorm:"primaryKey"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}

type UserInfo = wotoRaw.UserInfo
type FavoriteInfo = wotoRaw.FavoriteValue
type LikedListElement = wotoRaw.LikedListElement

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
	keys         *EntryKeys
	me           *UserInfo
}

type EntryKeys struct {
	PastKey    string `json:"past_key"`
	PresentKey string `json:"present_key"`
	FutureKey  string `json:"future_key"`

	_pastKey    wotoCrypto.WotoKey `json:"-"`
	_presentKey wotoCrypto.WotoKey `json:"-"`
	_futureKey  wotoCrypto.WotoKey `json:"-"`
}
