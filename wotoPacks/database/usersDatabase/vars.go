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
	usersMapById         = ssg.NewSafeMap[wv.PublicUserId, wv.UserInfo]()
	usersMapByUsername   = ssg.NewSafeMap[string, wv.UserInfo]()
	usersMapByEmail      = ssg.NewSafeMap[string, wv.UserInfo]()
	userIdGenerator      = ssg.NewNumIdGenerator[wv.PublicUserId]()
	usersFavoriteManager = _getFavoriteManager()
)
