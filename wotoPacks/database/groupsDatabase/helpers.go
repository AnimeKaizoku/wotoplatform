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

package groupsDatabase

import (
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
)

func LoadGroupsDatabase() error {
	var allGroups []*wv.GroupInfo

	lockDatabase()
	wv.SESSION.Find(&allGroups)
	unlockDatabase()

	for _, group := range allGroups {
		groupsInfo.Add(group.GroupId, group)

		if group.HasUsername() {
			groupsInfoByUsername.Add(group.GroupUsername, group)
		}

		if group.HasTelegramId() {
			groupsInfoByTelegramId.Add(group.TelegramId, group)
		}
	}

	return nil
}

// CreateNewGroup function creates a new group and saves it to db
// and caches the value (it generates and new unique group-id).
// this function doesn't validate any value, it's completely up to
// the caller to check and see if the group-username, telegram-id, etc
// already exists in the db or not.
func CreateNewGroup(data *CreateNewGroupData) *wv.GroupInfo {
	group := &wv.GroupInfo{
		GroupId:          generateGroupId(),
		GroupRegion:      data.GroupRegion,
		GroupUsername:    data.GroupUsername,
		TelegramId:       data.TelegramId,
		TelegramUsername: data.TelegramUsername,
	}

	SaveGroup(group)
	return group
}

func GetGroupInfoByUsername(username string) *wv.GroupInfo {
	return groupsInfoByUsername.Get(username)
}

func GetGroupInfoByTelegramUsername(username string) *wv.GroupInfo {
	return groupsInfoByTelegramUsername.Get(username)
}

func GetGroupInfoByTelegramId(id int64) *wv.GroupInfo {
	return groupsInfoByTelegramId.Get(id)
}

func SaveGroup(group *wv.GroupInfo) {
	SaveGroupNoCache(group)
	groupsInfo.Add(group.GroupId, group)

	if group.HasUsername() {
		groupsInfoByUsername.Add(group.GroupUsername, group)
	}

	if group.HasTelegramId() {
		groupsInfoByTelegramId.Add(group.TelegramId, group)
	}
}

func SaveGroupNoCache(group *wv.GroupInfo) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(group)
	tx.Commit()
	unlockDatabase()
}

func GetGroupInfo(id wv.PublicGroupId) *wv.GroupInfo {
	return groupsInfo.Get(id)
}

func GetGroupQueue(id wv.PublicGroupId) ([]wv.MediaModelId, error) {
	if !groupsQueue.Exists(id) {
		return nil, ErrGroupCallNotFound
	}

	queue := groupsQueue.Get(id)
	if queue == nil {
		return nil, ErrGroupHasNoQueue
	}

	return queue.mediaList, nil
}

func generateGroupId() wv.PublicGroupId {
	return ""
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
