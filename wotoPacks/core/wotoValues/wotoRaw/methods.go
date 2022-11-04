package wotoRaw

import (
	"strings"
	"time"
	"wp-server/wotoPacks/core/utils/logging"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"

	ws "github.com/AnimeKaizoku/ssg/ssg"
	"github.com/TheGolangHub/wotoCrypto/wotoCrypto"
	"golang.org/x/crypto/bcrypt"
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

func (u *UserInfo) IsPasswordCorrect(password *wotoCrypto.PasswordContainer256) bool {
	return u.IsRawPasswordCorrect(wotoValidate.GetPassAsBytes(password))
}

func (u *UserInfo) HasInvalidPasswordHash() bool {
	// TODO: Add more checks here.
	return u.PasswordHash == ""
}

func (u *UserInfo) IsRawPasswordCorrect(password []byte) bool {
	if u.HasInvalidPasswordHash() {
		if u.RegenerateSaltedPassword != nil {
			u.RegenerateSaltedPassword(u)
			// no return here, because CompareHashAndPassword should run again
			// at the end.
		} else {
			// DEPRECATED: if we have reached here, it means
			// something had went wrong in database package, and
			// RegenerateSaltedPassword field isn't set, DO CHECK.
			// this is very unsafe to do... (this is for old versions of wotoplatform)
			// see also: https://github.com/AnimeKaizoku/wotoplatform/pull/61
			logging.Warn("DEPRECATED: IsRawPasswordCorrect: PasswordHash is invalid and RegenerateSaltedPassword is nil ")
			return u.Password == string(password)
		}
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), password) == nil
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

func (u *UserInfo) SetAsMeta(meta ws.MetaDataProvider) {
	u.metaProvider = meta
}

func (u *UserInfo) GetMeta() ws.MetaDataProvider {
	return u.metaProvider
}

//---------------------------------------------------------

func (i PublicGroupId) IsInvalid() bool {
	return !i.IsValid()
}

func (i PublicGroupId) IsValid() bool {
	return len(i) > len(GroupIdPrefix)+2 && strings.HasPrefix(string(i), GroupIdPrefix)
}

//---------------------------------------------------------

func (m *MediaModel) SetAsMeta(meta ws.MetaDataProvider) {
	m.mediaMetaData = meta
}

func (m *MediaModel) GetMeta() ws.MetaDataProvider {
	return m.mediaMetaData
}

func (m *MediaModel) GetGenreIDs() []GenreId {
	if len(m.Genres) == 0 {
		return nil
	}

	var result []GenreId
	for _, current := range m.Genres {
		result = append(result, current.GenreId)
	}

	return result
}

func (m *MediaModel) HasGenreId(id GenreId) bool {
	for _, current := range m.Genres {
		if current != nil && current.GenreId == id {
			return true
		}
	}

	return false
}

// RemoveGenreElement removes the target genre-info and element from
// the media-model and returns the removed element (it will return nil if not found).
func (m *MediaModel) RemoveGenreElement(id GenreId) *MediaGenreElement {
	var newArray []*MediaGenreElement
	var target *MediaGenreElement
	for _, current := range m.GenreElements {
		if current != nil && current.Genre == id {
			target = current
			continue
		}

		newArray = append(newArray, current)
	}

	if target == nil {
		// not found
		return nil
	}

	m.GenreElements = newArray

	var newGenres []*MediaGenreInfo
	for _, current := range m.Genres {
		if current != nil && current.GenreId == id {
			continue
		}

		newGenres = append(newGenres, current)
	}

	m.Genres = newGenres
	return target
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
	e.UniqueId = MediaGenreElementPrefix + ws.ToBase32(time.Now().Unix())
}

func (e *MediaGenreElement) GenerateNewUniqueId() {
	if e.UniqueId != "" {
		return
	}
	e.UniqueId = MediaGenreElementPrefix + e.Genre.ToString() +
		UniqueIdInnerSeparator + ws.ToBase32(time.Now().Unix())
}

//---------------------------------------------------------

func (e GenreId) ToString() string {
	return ws.ToBase10(int64(e))
}

func (e GenreId) IsInvalid() bool {
	return e == 0
}

//---------------------------------------------------------
