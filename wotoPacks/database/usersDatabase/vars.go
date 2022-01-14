package usersDatabase

import (
	"sync"
	wv "wp-server/wotoPacks/core/wotoValues"
)

// database models
var (
	ModelUserInfo *wv.UserInfo = &wv.UserInfo{}
)

// cache values and mutexes.
var (
	usersMapById              = make(map[wv.PublicUserId]*wv.UserInfo)
	usersMapByIdMutex         = &sync.Mutex{}
	usersMapByUsername        = make(map[string]*wv.UserInfo)
	usersMapByUsernameMutex   = &sync.Mutex{}
	usersMapByTelegramId      = make(map[int64]*wv.UserInfo)
	usersMapByTelegramIdMutex = &sync.Mutex{}
	lastUserId                = wv.BaseUserId
	userIdGeneratorMutex      = &sync.Mutex{}
)
