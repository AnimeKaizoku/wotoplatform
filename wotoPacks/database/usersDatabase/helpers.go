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
	"strings"
	"sync"
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
)

func LoadUsersDatabase() error {
	var allUsers []*wv.UserInfo
	var allFavorites []*wv.FavoriteInfo
	var allLiked []*wv.LikedListElement

	lockDatabase()
	wv.SESSION.Find(&allUsers)
	wv.SESSION.Find(&allFavorites)
	wv.SESSION.Find(&allLiked)
	unlockDatabase()

	for _, user := range allUsers {
		userIdGenerator.SafeSet(user.UserId)

		usersMapById.Add(user.UserId, user)

		if user.HasUsername() {
			usersMapByUsername.Add(user.Username, user)
		}

		if user.HasTelegramId() {
			usersMapByTelegramId.Add(user.TelegramId, user)
		}

		if user.HasEmail() {
			usersMapByEmail.Add(strings.ToLower(user.Email), user)
		}
	}

	usersFavoriteManager.LoadAllFavorites(allFavorites)
	usersFavoriteManager.LoadAllLikedList(allLiked)

	migrateOwners()

	return nil
}

func UsernameExists(username string) bool {
	return usersMapByUsername.Exists(strings.ToLower(username))
}

func GetUserById(id wv.PublicUserId) *wv.UserInfo {
	return usersMapById.Get(id)
}

func GetUserByTelegramId(id int64) *wv.UserInfo {
	return usersMapByTelegramId.Get(id)
}

func GetUserByEmail(email string) *wv.UserInfo {
	return usersMapByEmail.Get(strings.ToLower(email))
}

func GetUserByUsername(username string) *wv.UserInfo {
	return usersMapByUsername.Get(strings.ToLower(username))
}

func GetUserFavorite(id wv.PublicUserId, key string) *wv.FavoriteInfo {
	return usersFavoriteManager.GetUserFavorite(id, key)
}

func GetUserLikedList(id wv.PublicUserId, key string) []*wv.LikedListElement {
	return usersFavoriteManager.GetUserLikeList(id, key)
}

func FavoriteValueExists(id wv.PublicUserId, key string) bool {
	return usersFavoriteManager.FavoriteExists(id, key)
}

func LikedListExists(id wv.PublicUserId, key string) bool {
	return usersFavoriteManager.LikedListExists(id, key)
}

func LikedItemExists(id wv.PublicUserId, uniqueId string) bool {
	return usersFavoriteManager.LikedItemExists(id, uniqueId)
}

func IsLikedItemUniqueIdValid(uniqueId string) bool {
	// a unique id should look like this:
	// "abcd=a1bcd"
	// we are sure that it should *always* contain more than
	// 6 characters (at least) and it should contain sep character ('=')
	return len(uniqueId) > minUniqueIdLen && strings.Contains(uniqueId, likedListUIDSep)
}

func GetUserFavoriteCount(id wv.PublicUserId) int {
	return usersFavoriteManager.Length(id)
}

func GetUserLikedListCount(id wv.PublicUserId, key string) int {
	return usersFavoriteManager.GetLikedListCount(id, key)
}

func SetUserFavorite(id wv.PublicUserId, key, value string) {
	info := usersFavoriteManager.NewFavorite(id, key, value)
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(info)
	tx.Commit()
	unlockDatabase()
}

func DeleteLikedListItem(id wv.PublicUserId, uniqueId string) {
	liked := usersFavoriteManager.DeleteLikedItemByUniqueId(id, uniqueId)
	lockDatabase()
	wv.SESSION.Delete(liked)
	unlockDatabase()
}

func AddUserLikedList(data *NewLikedListElementData) *wv.LikedListElement {
	liked := data.ToLikedListElement()
	usersFavoriteManager.AddLiked(liked)
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(liked)
	tx.Commit()
	unlockDatabase()

	return liked
}

func DeleteUserFavorite(id wv.PublicUserId, key string) {
	info := usersFavoriteManager.DeleteFavorite(id, key)
	lockDatabase()
	wv.SESSION.Delete(info)
	unlockDatabase()
}

func SaveUser(user *wv.UserInfo) {
	SaveUserNoCache(user)
	usersMapById.Add(user.UserId, user)

	if user.HasUsername() {
		usersMapByUsername.Add(strings.ToLower(user.Username), user)
	}

	if user.HasTelegramId() {
		usersMapByTelegramId.Add(user.TelegramId, user)
	}

	if user.HasEmail() {
		usersMapByEmail.Add(strings.ToLower(user.Email), user)
	}
}

func SaveUserNoCache(user *wv.UserInfo) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(user)
	tx.Commit()
	unlockDatabase()
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
		Email:      data.Email,
		CreatedBy:  data.By,
		Birthday:   data.Birthday,
		IsVirtual:  data.Username == "",
	}
	SaveUser(u)
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
		// save the user in the db, don't let it cache to save more time.
		SaveUserNoCache(currentUser)
	}
}

func generateUserId() wv.PublicUserId {
	return userIdGenerator.Next()
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

func _getFavoriteManager() *favoriteManager {
	return &favoriteManager{
		mut:    &sync.Mutex{},
		values: make(map[wv.PublicUserId]*UserFavoritesAndLiked),
	}
}
