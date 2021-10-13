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

//func EnsureBotInDb(b *gotgbot.Bot) {
// Insert bot user only if it doesn't exist already
//	self, err := b.GetMe()
//	if err != nil {
//		log.Fatal(err.Error())
//	}
//	botUser := &User{UserId: self.Id, UserName: self.Username}
//	SESSION.Save(botUser)
//}

/*
func UpdateUser(userId int64, username string, chatId string, chatName string) {
	username = strings.ToLower(username)
	tx := SESSION.Begin()

	// upsert user
	user := &PlayerInfo{UserId: userId, UserName: username}
	tx.Save(user)

	if chatId == "nil" || chatName == "nil" {
		tx.Commit()
		return
	}

	// upsert chat
	chat := &Chat{ChatId: chatId, ChatName: chatName}
	tx.Save(chat)
	tx.Commit()
}

func GetUserIDByUsername(username string) (user *PlayerInfo, err error) {
	username = strings.ToLower(username)
	if SESSION == nil {
		return nil, errors.New("cannot access to SESSION " +
			"of db, because it's nil")
	}

	p := User{}
	SESSION.Where("user_name = ?", username).Take(&p)
	return &p, nil
}

func GetUserId(username string) int64 {
	if len(username) <= 5 {
		return 0
	}
	if username[0] == '@' {
		username = username[1:]
	}
	users, err := GetUserIDByUsername(username)
	if err != nil {
		return 0
	}
	if users == nil {
		return 0
	}

	return (*users).UserId
}

*/
