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
	"wp-server/wotoPacks/core/utils/logging"
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"

	"github.com/AnimeKaizoku/ssg/ssg"
	"github.com/TheGolangHub/wotoCrypto/wotoCrypto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoadUsersDatabase() error {
	var allFavorites []*wv.FavoriteInfo
	var allLiked []*wv.LikedListElement

	lockDatabase()
	wv.SESSION.Find(&allFavorites)
	wv.SESSION.Find(&allLiked)
	unlockDatabase()

	userIdGenerator.SafeSet(getLastUserId())

	usersFavoriteManager.LoadAllFavorites(allFavorites)
	usersFavoriteManager.LoadAllLikedList(allLiked)

	migrateOwners()

	return nil
}

func UsernameExists(username string) bool {
	user := usersMapByUsername.Get(username)
	if user == userInfoExists {
		return true
	} else if user == userInfoNotFound {
		return false
	} else if user != nil && user.Username == username {
		// user info exists, and is a valid user info.
		return true
	}

	exists := false
	wv.SESSION.Raw(
		"SELECT true AS RESULT WHERE EXISTS (SELECT * FROM user_infos WHERE UPPER(username) = ?)",
		strings.ToUpper(username),
	).Scan(&exists)

	if exists {
		usersMapByUsername.Add(username, userInfoExists)
	}

	return exists
}

func GetUserById(id wv.PublicUserId) *wv.UserInfo {
	return getUserByField(usersMapById, id, "user_id")
}

func getUserByField[T comparable](theMap *ssg.SafeMap[T, wv.UserInfo], key T, columnName string) *wv.UserInfo {
	myStr, isString := interface{}(key).(string)
	if isString && !strings.HasPrefix(columnName, "UPPER(") {
		columnName = "UPPER(" + columnName + ")"
		key = interface{}(strings.ToUpper(myStr)).(T)
	}

	user := theMap.Get(key)
	if user == userInfoNotFound {
		return nil
	} else if user != nil && user.UserId != 0 {
		// this makes sure that the user is actually a valid user.
		return user
	}

	user = new(wv.UserInfo)
	err := wv.SESSION.Table("user_infos").Take(user, columnName+" = ?", key).Error
	if err == gorm.ErrRecordNotFound {
		theMap.Add(key, userInfoNotFound)
		return nil
	} else if err != nil {
		logging.Debugf("GetUserById: returned error for %s %v: %v", columnName, key, err)
		return nil
	}

	theMap.Add(key, user)

	return user
}

// getLastUserId returns the last user id saved inside of db.
func getLastUserId() wv.PublicUserId {
	var theId wv.PublicUserId
	// or can do another raw query:
	// SELECT * FROM user_infos ORDER BY user_id DESC LIMIT 0, 1
	err := wv.SESSION.Raw("SELECT MAX(user_id) FROM user_infos").Scan(&theId).Error
	if err != nil || theId == 0 {
		return wv.BaseUserId
	}

	return theId
}

// GetUserByTelegramId returns the user which had connected a telegram account with the
// given id to their wotoplatform account.
// WARNING: this function has to be rewritten, it does nothing by now, should be
// implemented later (TODO).
// skipcq
func GetUserByTelegramId(id int64) *wv.UserInfo {
	return userInfoExists
}

func GetUserByEmail(email string) *wv.UserInfo {
	return getUserByField(usersMapByEmail, email, "email")
}

func GetUserByUsername(username string) *wv.UserInfo {
	return getUserByField(usersMapByUsername, username, "username")
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
	if data.SaltedPassword == "" {
		data.SaltedPassword = getSaltedPasswordAsStr(data.Password)
		data.PasswordHash = data.Password.Hash256
	}
	u := &wv.UserInfo{
		UserId:       generateUserId(),
		Username:     data.Username,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		TelegramId:   data.TelegramId,
		Password:     data.SaltedPassword,
		PasswordHash: data.PasswordHash,
		Permission:   data.Permission,
		Email:        data.Email,
		CreatedBy:    data.By,
		Birthday:     data.Birthday,
		IsVirtual:    data.Username == "",

		RegenerateSaltedPassword: regenerateSaltedPassword,
	}
	SaveUser(u)
	return u
}

func getSaltedPasswordAsStr(password *wotoCrypto.PasswordContainer256) string {
	return getSaltedPasswordFromBytes(wotoValidate.GetPassAsBytes(password))
}

func getSaltedPasswordFromBytes(password []byte) string {
	b, err := bcrypt.GenerateFromPassword(password, wotoValidate.PasswordSaltCost)
	if err != nil {
		logging.Error(err)
		return ""
	}

	return string(b)
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
				Username:       current.Username,
				SaltedPassword: getSaltedPasswordFromBytes([]byte(current.Password)),
				PasswordHash:   wotoValidate.GetPasswordHash([]byte(current.Password)),
				Permission:     wv.PermissionOwner,
			})
			continue
		}

		if currentUser.IsOwner() && currentUser.IsRawPasswordCorrect([]byte(current.Password)) {
			continue
		}

		currentUser.Permission = wv.PermissionOwner
		currentUser.Password = getSaltedPasswordFromBytes([]byte(current.Password))
		currentUser.PasswordHash = wotoValidate.GetPasswordHash([]byte(current.Password))
		// save the user in the db, don't let it cache to save more time.
		SaveUserNoCache(currentUser)
	}
}

// regenerateSaltedPassword function will regenerate the salted password
// of a user.
// a tricky solution to prevent from breaking changes.
// WARNING: Beware of deadlocks, this function might try to lock internal db mute,
// if it's already locked, it will reach a deadlock.
func regenerateSaltedPassword(u *wv.UserInfo) {
	u.PasswordHash = wotoValidate.GetPasswordHash([]byte(u.Password))
	u.Password = getSaltedPasswordFromBytes([]byte(u.Password))
	SaveUserNoCache(u)
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
