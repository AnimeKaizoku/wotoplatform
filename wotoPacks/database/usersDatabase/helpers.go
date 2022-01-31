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

package usersDatabase

import (
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
)

func LoadUsersDatabase() error {
	var allUsers []*wv.UserInfo
	var allFavorites []*wv.FavoriteInfo

	lockDatabase()
	wv.SESSION.Find(&allUsers)
	wv.SESSION.Find(&allFavorites)
	unlockDatabase()

	usersMapByIdMutex.Lock()
	usersMapByUsernameMutex.Lock()
	usersMapByTelegramIdMutex.Lock()
	for _, user := range allUsers {
		if user.UserId > lastUserId {
			lastUserId = user.UserId
		}

		usersMapById[user.UserId] = user

		if user.HasUsername() {
			usersMapByUsername[user.Username] = user
		}

		if user.HasTelegramId() {
			usersMapByTelegramId[user.TelegramId] = user
		}
	}
	usersMapByIdMutex.Unlock()
	usersMapByUsernameMutex.Unlock()
	usersMapByTelegramIdMutex.Unlock()

	usersFavoriteManager.LoadAll(allFavorites)

	migrateOwners()

	return nil
}

func UsernameExists(username string) bool {
	usersMapByUsernameMutex.Lock()
	b := usersMapByUsername[username] != nil
	usersMapByUsernameMutex.Unlock()

	return b
}

func GetUserById(id wv.PublicUserId) *wv.UserInfo {
	usersMapByIdMutex.Lock()
	user := usersMapById[id]
	usersMapByIdMutex.Unlock()

	return user
}

func GetUserByTelegramId(id int64) *wv.UserInfo {
	usersMapByTelegramIdMutex.Lock()
	user := usersMapByTelegramId[id]
	usersMapByTelegramIdMutex.Unlock()

	return user
}

func GetUserByUsername(username string) *wv.UserInfo {
	usersMapByUsernameMutex.Lock()
	user := usersMapByUsername[username]
	usersMapByUsernameMutex.Unlock()

	return user
}

func GetUserFavorite(id wv.PublicUserId, key string) string {
	return usersFavoriteManager.GetUserFavorite(id, key)
}

func GetUserFavoriteCount(id wv.PublicUserId) int {
	return usersFavoriteManager.Length(id)
}

func SetUserFavorite(id wv.PublicUserId, key, value string) {
	info := usersFavoriteManager.NewFavorite(id, key, value)
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(info)
	tx.Commit()
	unlockDatabase()
}

func SaveUser(user *wv.UserInfo, cache bool) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockDatabase()

	if cache {
		usersMapByIdMutex.Lock()
		usersMapById[user.UserId] = user
		usersMapByIdMutex.Unlock()

		if user.HasUsername() {
			usersMapByUsernameMutex.Lock()
			usersMapByUsername[user.Username] = user
			usersMapByUsernameMutex.Unlock()
		}

		if user.HasTelegramId() {
			usersMapByTelegramIdMutex.Lock()
			usersMapByTelegramId[user.TelegramId] = user
			usersMapByTelegramIdMutex.Unlock()
		}
	}
}

// CreateNewUser creates a new user and saves it to the database.
// This function doesn't check for existing username.
// It doesn't validate username or password. User parameters need
// to be validated before this function is called.
func CreateNewUser(data *NewUserData) *wv.UserInfo {
	u := &wv.UserInfo{
		UserId:     generateUserId(),
		Username:   data.Username,
		FirstName:  data.FirstName,
		LastName:   data.LastName,
		TelegramId: data.TelegramId,
		Password:   data.Password,
		Permission: data.Permission,
		CreatedBy:  data.By,
		Birthday:   data.Birthday,
		IsVirtual:  data.Username == "",
	}
	SaveUser(u, true)
	return u
}

func migrateOwners() {
	owners := wotoConfig.GetOwners()
	if len(owners) == 0 {
		return
	}

	var currentUser *wv.UserInfo

	for _, current := range owners {
		currentUser = GetUserByUsername(current.Username)
		if currentUser == nil {
			CreateNewUser(&NewUserData{
				Username:   current.Username,
				Password:   current.Password,
				Permission: wv.PermissionOwner,
			})
			continue
		}

		if currentUser.IsOwner() && currentUser.IsPasswordCorrect(current.Password) {
			continue
		}

		currentUser.Permission = wv.PermissionOwner
		currentUser.Password = current.Password
		SaveUser(currentUser, false)
	}
}

func generateUserId() wv.PublicUserId {
	userIdGeneratorMutex.Lock()
	lastUserId++
	userIdGeneratorMutex.Unlock()

	return lastUserId
}

func lockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Lock()
	}
}

func unlockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Unlock()
	}
}
