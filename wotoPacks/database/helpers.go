/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
 * Copyright (c) 2021 AmanoTeam.
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

import (
	"errors"
	"fmt"
	"strings"
	"wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/utils/logging"
	"wp-server/wotoPacks/wotoConfig"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartDatabase() (err error) {
	if wotoConfig.WConfig == nil {
		return errors.New("config cannot be nil, make sure to config the server")
	}

	if wotoConfig.WConfig.UseSQLLite {
		SESSION, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", wotoConfig.WConfig.DB_Name)), &gorm.Config{})
	} else {
		if len(wotoConfig.WConfig.DB_URL) < 4 {
			return errors.New("invalid db url provided")
		}

		SESSION, err = gorm.Open(postgres.Open(wotoConfig.WConfig.DB_URL), &gorm.Config{})
	}

	if err != nil {
		logging.SUGARED.Error("failed to connect database", err)
		return
	}

	logging.SUGARED.Debug("Database connected")
	//log.Println("Database connected")

	// Create tables if they don't exist
	err = SESSION.AutoMigrate(&UserInfo{}, &SongInfo{})
	if err != nil {
		logging.SUGARED.Error(err)
		return
	}

	logging.SUGARED.Debug("Auto-migrated database schema")
	//log.Println("Auto-migrated database schema")

	return nil
}

func CreateNewUser(raw interfaces.RawUser) {
	tx := SESSION.Begin()
	tx.Save(toUserInfo(raw))
	tx.Commit()
}

func UpdateUser(raw interfaces.RawUser) {
	tx := SESSION.Begin()
	tx.Save(toUserInfo(raw))
	tx.Commit()
}

func CheckUserPassword(raw interfaces.RawUser) (bool, error) {
	username := strings.ToLower(raw.GetName())
	if SESSION == nil {
		return false, errors.New("db doesn't exist")
	}

	p := &UserInfo{}
	SESSION.Where("user_name = ?", username).Take(p)

	if p.IsEmpty() {
		return false, errors.New("user doesn't exist")
	}

	return p.Password == raw.GetPassword(), nil
}

// GetPlayerInfo will return a player's info in the database
func GetUserInfoPublic(publicID string) (interfaces.RawUser, error) {
	p := UserInfo{}
	SESSION.Where("public_id = ?", publicID).Take(&p)

	if p.IsEmpty() {
		return nil, errors.New("user doesn't exist")
	}

	return &p, nil
}

func GetUserInfoPrivate(privateID string) (interfaces.RawUser, error) {
	p := UserInfo{}
	SESSION.Where("private_id = ?", privateID).Take(&p)

	if p.IsEmpty() {
		return nil, errors.New("user doesn't exist")
	}

	return &p, nil
}

func GetUserBySv(username string) (raw []interfaces.RawUser, err error) {
	p := []UserInfo{}

	SESSION.Where("sv_username = ?", username).Find(p)

	if len(p) == 0 {
		return nil, errors.New("doesn't exist")
	}

	return convertToRawArray(p...), nil
}

func GetUserProfilePreview(username string) (interfaces.RawUser, error) {
	p := UserInfo{}
	SESSION.Where("user_name = ?", username).Take(&p)

	if p.IsEmpty() {
		return nil, errors.New("user doesn't exist")
	}

	return &p, nil
}

func UsernameExists(username string) bool {
	p := UserInfo{}
	SESSION.Where("user_name = ?", username).Take(&p)

	return !p.IsEmpty()
}

func convertToRawArray(u ...UserInfo) []interfaces.RawUser {
	if len(u) == 0 {
		return nil
	} else if len(u) == 1 {
		return []interfaces.RawUser{&u[0]}
	}
	var r []interfaces.RawUser
	for _, i := range u {
		r = append(r, &i)
	}

	return r
}

func toUserInfo(raw interfaces.RawUser) *UserInfo {
	// prevent from allocating memory, if the type is already `*UserInfo`,
	// simply return it.
	user, ok := raw.(*UserInfo)
	if ok {
		return user
	}

	return &UserInfo{
		SoacialvoidUsername: raw.GetSocialvoidUsername(),
		PublicID:            raw.GetPublicID(),
		PrivateID:           raw.GetPrivateID(),
		UserName:            raw.GetName(),
		Password:            raw.GetPassword(),
		UserLevel:           raw.GetUserLever(),
		LastSeen:            raw.GetLastSeen(),
		UserIntro:           raw.GetUserIntro(),
		UserAvatar:          raw.GetAvatar(),
		UserAvatarFrame:     raw.GetAvatarFrame(),
		UserVIPLevel:        raw.GetUserVIPLevel(),
		UserCurrentExp:      raw.GetCurrentExp(),
		UserCurrentVIPExp:   raw.GetCurrentVIPExp(),
		UserTotalExp:        raw.GetTotalExp(),
		UserTotalVIPExp:     raw.GetTotalVIPExp(),
		UserMaxExp:          raw.GetMaxExp(),
		UserMaxVIPExp:       raw.GetMaxVIPExp(),
		UserCity:            raw.GetCity(),
	}
}
