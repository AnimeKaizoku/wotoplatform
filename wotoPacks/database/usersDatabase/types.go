package usersDatabase

import (
	"sync"
	wv "wp-server/wotoPacks/core/wotoValues"
)

type UserFavorites struct {
	mut    *sync.Mutex
	values map[string]*wv.FavoriteInfo
}

type NewUserData struct {
	Username   string
	Password   string
	Email      string
	TelegramId int64
	FirstName  string
	LastName   string
	Birthday   string
	Permission wv.UserPermission
	By         wv.PublicUserId
}

type favoriteManager struct {
	mut    *sync.Mutex
	values map[wv.PublicUserId]*UserFavorites
}
