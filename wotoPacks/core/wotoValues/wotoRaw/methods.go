package wotoRaw

import (
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
	ws "github.com/AnimeKaizoku/ssg/ssg"
)

//---------------------------------------------------------

func (u *UserInfo) HasUsername() bool {
	return u.Username != ""
}

func (u *UserInfo) HasTelegramId() bool {
	return u.TelegramId != 0
}

func (u *UserInfo) HasEmail() bool {
	return u.Email != ""
}

func (u *UserInfo) SetCachedTime() {
	u.cachedTime = time.Now()
}

func (u *UserInfo) IsCacheExpired(d time.Duration) bool {
	return time.Since(u.cachedTime) > d
}

func (u *UserInfo) IsPasswordCorrect(password string) bool {
	// TODO: encrypt the password
	return u.Password == password
}

func (u *UserInfo) GetPublicId() PublicUserId {
	return u.UserId
}

func (u *UserInfo) CanCreateAccount() bool {
	return u.Permission >= PermissionAdmin
}

func (u *UserInfo) CanChangePermission() bool {
	return u.Permission >= PermissionDeveloper
}

func (u *UserInfo) IsAdmin() bool {
	return u.Permission >= PermissionAdmin
}

func (u *UserInfo) IsOwner() bool {
	return u.Permission >= PermissionOwner
}

func (u *UserInfo) IsValid() bool {
	return !u.UserId.IsZero()
}

func (u *UserInfo) IsInvalid() bool {
	return u == nil || u.UserId.IsZero()
}

func (u *UserInfo) SetAsMeta(meta MetaDataProvider) {
	u.metaProvider = meta
}

func (u *UserInfo) GetMeta() MetaDataProvider {
	return u.metaProvider
}

//---------------------------------------------------------

func (m *MediaModel) SetAsMeta(meta MetaDataProvider) {
	m.mediaMetaData = meta
}

func (m *MediaModel) GetMeta() MetaDataProvider {
	return m.mediaMetaData
}

//---------------------------------------------------------

func (g *GroupInfo) HasUsername() bool {
	return g.GroupUsername != ""
}

func (g *GroupInfo) HasTelegramId() bool {
	return g.TelegramId != 0
}

func (g *GroupInfo) HasTelegramUsername() bool {
	return g.TelegramUsername != ""
}

//---------------------------------------------------------

func (i PublicUserId) IsZero() bool {
	return i == 0
}

func (i PublicUserId) ToBase32() string {
	return ws.ToBase32(int64(i))
}

func (i PublicUserId) ToBase16() string {
	return ws.ToBase16(int64(i))
}

func (i PublicUserId) ToBase18() string {
	return ws.ToBase18(int64(i))
}

func (i PublicUserId) ToBase20() string {
	return ws.ToBase20(int64(i))
}

func (i PublicUserId) ToBase28() string {
	return ws.ToBase28(int64(i))
}

func (i PublicUserId) ToBase30() string {
	return ws.ToBase30(int64(i))
}

//---------------------------------------------------------

func (f *FavoriteValue) IsInvalid() bool {
	return f == nil || f.UserId.IsZero() || f.FavoriteKey == ""
}

//---------------------------------------------------------

func (i MediaModelId) IsInvalid() bool {
	return i == ""
}

//---------------------------------------------------------

func (e *LikedListElement) IsInvalid() bool {
	return e == nil || e.OwnerId.IsZero() || e.LikedKey == ""
}

func (e *LikedListElement) CompareWith(title string, media MediaModelId) bool {
	return e.Title == title || e.MediaId == media
}

//---------------------------------------------------------

func (e *MediaGenreElement) GenerateUniqueId() {
	e.UniqueId = MediaGenreElementPrefix + ssg.ToBase32(time.Now().Unix())
}

func (e *MediaGenreElement) GenerateNewUniqueId() {
	if e.UniqueId != "" {
		return
	}
	e.UniqueId = MediaGenreElementPrefix + e.Genre.ToString() +
		UniqueIdInnerSeparator + ssg.ToBase32(time.Now().Unix())
}

//---------------------------------------------------------

func (e GenreId) ToString() string {
	return ssg.ToBase10(int64(e))
}

//---------------------------------------------------------
