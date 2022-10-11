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
	"sync"
	"time"
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"

	"github.com/AnimeKaizoku/ssg/ssg"
)

// LoadGroupsDatabase function loads all the variables related to groupsDatabase
// package.
func LoadGroupsDatabase() error {
	var allGroups []*wv.GroupInfo

	lockDatabase()
	wv.SESSION.Find(&allGroups)
	unlockDatabase()

	for _, group := range allGroups {
		cacheGroupInfo(ssg.Clone(group))
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

// cacheGroupInfo function caches the passed group-info into the memory.
func cacheGroupInfo(group *wv.GroupInfo) {
	groupsInfo.Add(group.GroupId, group)

	if group.HasUsername() {
		groupsInfoByUsername.Add(group.GroupUsername, group)
	}

	if group.HasTelegramId() {
		groupsInfoByTelegramId.Add(group.TelegramId, group)
	}
}

// GetGroupInfoByUsername returns the group-info by the specified group username.
func GetGroupInfoByUsername(username string) *wv.GroupInfo {
	return groupsInfoByUsername.Get(username)
}

// GetGroupInfoByTelegramUsername returns the group-info specified by its telegram
// username. this method is not really a good way of getting information
// about a group.
func GetGroupInfoByTelegramUsername(username string) *wv.GroupInfo {
	return groupsInfoByTelegramUsername.Get(username)
}

// GetGroupInfoByTelegramId returns the group-info specified by its telegram-id.
func GetGroupInfoByTelegramId(id int64) *wv.GroupInfo {
	return groupsInfoByTelegramId.Get(id)
}

// SaveGroup function saves the specified group-info into the database
// and caches it into memory.
func SaveGroup(group *wv.GroupInfo) {
	SaveGroupNoCache(group)
	cacheGroupInfo(ssg.Clone(group))
}

// SaveGroupNoCache saves the specified group-info into the db, but it does not
// caches the value into memory.
func SaveGroupNoCache(group *wv.GroupInfo) {
	lockDatabase()
	tx := wv.SESSION.Begin()
	tx.Save(group)
	tx.Commit()
	unlockDatabase()
}

// GetGroupInfo function returns the group-info using its specified id.
func GetGroupInfo(id wv.PublicGroupId) *wv.GroupInfo {
	return groupsInfo.Get(id)
}

// GetGroupQueue returns an array of
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

// CreateGroupCall creates a new group call. it does not validate the
// passed info var as argument. it's up to the caller to check for that.
func CreateGroupCall(info *wv.GroupInfo, statedBy wv.PublicUserId) {
	manager := &groupQueueManager{
		groupInfo: info,
		callInfo:  generateCallInfo(info.GroupId, statedBy),
		mut:       &sync.Mutex{},
	}

	groupsQueue.Add(info.GroupId, manager)
}

// generateCallInfo generates a new group-call info struct variable
// and return it.
func generateCallInfo(id wv.PublicGroupId, startedBy wv.PublicUserId) *wv.GroupCallInfo {
	return &wv.GroupCallInfo{
		GroupId:   id,
		StartedBy: startedBy,
		StartedAt: time.Now(),
	}
}

// generateGroupId generates a new unique group-id. this function should be used when and
// only when we want to create a new group-info into db.
func generateGroupId() wv.PublicGroupId {
	id := ssg.ToBase10(groupIdGenerator.Next())
	tStr := ssg.ToBase32(time.Now().Unix())
	return wv.PublicGroupId(GroupIdPrefix + id + tStr)
}

// GetNewGroupId function just simply calls generateGroupId function. it's exported
// only and only to be used in test packages.
func GetNewGroupId() wv.PublicGroupId {
	return generateGroupId()
}

// lockDatabase will put a lock on db mutex.
func lockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Lock()
	}
}

// unlockDatabase will unlock db mutex.
func unlockDatabase() {
	if wotoConfig.UseSqlite() {
		wv.SessionMutex.Unlock()
	}
}
