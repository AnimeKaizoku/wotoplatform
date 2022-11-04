package usersDatabase

import (
	wv "wp-server/wotoPacks/core/wotoValues"

	"github.com/AnimeKaizoku/ssg/ssg"
)

// database models
var (
	ModelUserInfo         *wv.UserInfo         = &wv.UserInfo{}
	userInfoExists        *wv.UserInfo         = &wv.UserInfo{}
	userInfoNotFound      *wv.UserInfo         = &wv.UserInfo{}
	ModelUserFavorite     *wv.FavoriteInfo     = &wv.FavoriteInfo{}
	ModelLikedListElement *wv.LikedListElement = &wv.LikedListElement{}
)

// cache values and mutexes.
var (
	usersMapById         = _getExpiringMap[wv.PublicUserId, wv.UserInfo]()
	usersMapByUsername   = _getExpiringMap[string, wv.UserInfo]()
	usersMapByEmail      = _getExpiringMap[string, wv.UserInfo]()
	userIdGenerator      = ssg.NewNumIdGenerator[wv.PublicUserId]()
	usersFavoriteManager = _getFavoriteManager()
)
