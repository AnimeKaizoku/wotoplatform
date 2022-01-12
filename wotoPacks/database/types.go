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

package database

import "time"

type UserInfo struct {
	SoacialvoidUsername string `gorm:"column:sv_username" json:"sv_username"`
	PublicID            string `gorm:"primary_key;column:public_id" json:"public_id"`
	PrivateID           string `gorm:"primary_key;column:private_id" json:"private_id"`
	Password            string `gorm:"column:pass" json:"-"`
	UserName            string `json:"user_name"`
	UserLevel           uint16 `json:"level"`
	LastSeen            string `json:"last_seen"`
	UserIntro           string `json:"intro"`
	UserAvatar          string `json:"avatar"`
	UserAvatarFrame     string `json:"avatar_frame"`
	UserVIPLevel        uint8  `json:"vip_level"`
	// should be big.Int
	UserCurrentExp string `json:"current_exp"`
	UserMaxExp     string `json:"max_exp"`
	// should be big.Int
	UserTotalExp    string `json:"total_exp"`
	UserTotalVIPExp string `json:"total_vip_exp"`
	// should be big.Int
	UserCurrentVIPExp string    `json:"current_vip_exp"`
	UserMaxVIPExp     string    `json:"max_vip_exp"`
	UserCity          string    `json:"city"`
	CreatedAt         time.Time `json:"created_at"`
}

type SongInfo struct {
	SongName string `gorm:"primary_key" json:"name"`
	SongId   string `gorm:"primary_key" json:"id"`
	// the avilable direct links of this song
	// (separate it using space character)
	AvailableLinks string `json:"link_list"`
}
