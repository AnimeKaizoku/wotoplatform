package usersDatabase

import (
	"sync"
	wv "wp-server/wotoPacks/core/wotoValues"

	"github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

type UserFavoritesAndLiked struct {
	mut    *sync.Mutex
	values map[string]*UserFavoriteAndLikedInfo
}

type UserFavoriteAndLikedInfo struct {
	FavoriteInfo *wv.FavoriteInfo
	LikedList    []*wv.LikedListElement
}

type NewUserData struct {
	Username       string
	Password       *wotoCrypto.PasswordContainer256
	PasswordHash   string
	SaltedPassword string
	Email          string
	TelegramId     int64
	FirstName      string
	LastName       string
	Birthday       string
	Permission     wv.UserPermission
	By             wv.PublicUserId
}

type NewLikedListElementData struct {
	UserId       wv.PublicUserId
	MediaId      wv.MediaModelId
	Title        string
	LikedKey     string
	ReferenceUrl string
	SourceUrl    string
}

type favoriteManager struct {
	mut    *sync.Mutex
	values map[wv.PublicUserId]*UserFavoritesAndLiked
}
