package usersDatabase

import (
	"sync"
	"time"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/core/wotoValues/wotoValidate"

	ws "github.com/AnimeKaizoku/ssg/ssg"
)

// GetPasswordHash returns the hash256 associated with this password
// (will use the passed argument as the alternative password, if `NewUserData.Password` field is nil).
// This method has no usage anymore, it is only here for more compatibility
func (d *NewUserData) GetPasswordHash(rawPass string) string {
	if d.Password != nil {
		return d.Password.Hash256
	}

	return wotoValidate.GetPasswordHash([]byte(rawPass))
}

//---------------------------------------------------------

func (m *favoriteManager) Length(id wv.PublicUserId) int {
	m.mut.Lock()
	favorites := m.values[id]
	m.mut.Unlock()

	return favorites.Length()
}

func (m *favoriteManager) GetLikedListCount(id wv.PublicUserId, key string) int {
	return len(m.GetFavoritesAndLiked(id).GetLikedList(key))
}

func (m *favoriteManager) GetFavoritesAndLiked(id wv.PublicUserId) *UserFavoritesAndLiked {
	m.mut.Lock()
	f := m.values[id]
	m.mut.Unlock()

	return f
}

func (m *favoriteManager) GetUserFavorite(id wv.PublicUserId, key string) *wv.FavoriteInfo {
	return m.GetFavoritesAndLiked(id).GetFavoriteInfo(key)
}

func (m *favoriteManager) GetUserLikeList(id wv.PublicUserId, key string) []*wv.LikedListElement {
	return m.GetFavoritesAndLiked(id).GetLikedList(key)
}

func (m *favoriteManager) AddFavorite(f *wv.FavoriteInfo) {
	favorites := m.GetFavoritesAndLiked(f.UserId)
	if favorites == nil {
		favorites = &UserFavoritesAndLiked{
			mut:    &sync.Mutex{},
			values: make(map[string]*UserFavoriteAndLikedInfo),
		}

		m.mut.Lock()
		m.values[f.UserId] = favorites
		m.mut.Unlock()
	}

	favorites.AddFavorite(f)
}

func (m *favoriteManager) AddLiked(l *wv.LikedListElement) {
	favLikes := m.GetFavoritesAndLiked(l.OwnerId)
	if favLikes == nil {
		favLikes = &UserFavoritesAndLiked{
			mut:    &sync.Mutex{},
			values: make(map[string]*UserFavoriteAndLikedInfo),
		}

		m.mut.Lock()
		m.values[l.OwnerId] = favLikes
		m.mut.Unlock()
	}

	favLikes.AddLiked(l)
}

func (m *favoriteManager) NewFavorite(id wv.PublicUserId, key, value string) *wv.FavoriteInfo {
	info := &wv.FavoriteInfo{
		UserId:      id,
		FavoriteKey: key,
		TheValue:    value,
	}

	m.AddFavorite(info)
	return info
}

func (m *favoriteManager) DeleteFavorite(id wv.PublicUserId, key string) *wv.FavoriteInfo {
	info := &wv.FavoriteInfo{
		UserId:      id,
		FavoriteKey: key,
	}

	m.GetFavoritesAndLiked(id).Delete(key)
	return info
}

func (m *favoriteManager) FavoriteExists(id wv.PublicUserId, key string) bool {
	return m.GetFavoritesAndLiked(id).FavoriteExists(key)
}

func (m *favoriteManager) LikedListExists(id wv.PublicUserId, key string) bool {
	return m.GetFavoritesAndLiked(id).LikedListExists(key)
}

func (m *favoriteManager) LikedItemExists(id wv.PublicUserId, uniqueId string) bool {
	return m.GetFavoritesAndLiked(id).LikedItemExists(uniqueId)
}

func (m *favoriteManager) GetLikedItemByUniqueId(id wv.PublicUserId, uniqueId string) *wv.LikedListElement {
	return m.GetFavoritesAndLiked(id).GetLikedItemByUniqueId(uniqueId)
}

func (m *favoriteManager) DeleteLikedItemByUniqueId(id wv.PublicUserId, uniqueId string) *wv.LikedListElement {
	return m.GetFavoritesAndLiked(id).DeleteLikedItemByUniqueId(uniqueId)
}

func (m *favoriteManager) LoadAllFavorites(favorites []*wv.FavoriteInfo) {
	for _, current := range favorites {
		m.AddFavorite(current)
	}
}

func (m *favoriteManager) LoadAllLikedList(liked []*wv.LikedListElement) {
	for _, current := range liked {
		m.AddLiked(current)
	}
}

//---------------------------------------------------------

func (f *UserFavoritesAndLiked) FavoriteExists(key string) bool {
	if f == nil {
		return false
	}

	f.mut.Lock()
	value := f.values[key]
	f.mut.Unlock()

	return value != nil && value.FavoriteInfo != nil
}

func (f *UserFavoritesAndLiked) LikedListExists(key string) bool {
	if f == nil {
		return false
	}

	f.mut.Lock()
	value := f.values[key]
	f.mut.Unlock()

	return value != nil && len(value.LikedList) != 0
}

func (f *UserFavoritesAndLiked) LikedItemExists(uniqueId string) bool {
	if f == nil {
		return false
	}

	f.mut.Lock()
	for _, currentValue := range f.values {
		if currentValue == nil || len(currentValue.LikedList) == 0 {
			continue
		}

		for _, currentItem := range currentValue.LikedList {
			if currentItem.UniqueId == uniqueId {
				f.mut.Unlock()
				return true
			}
		}
	}
	f.mut.Unlock()

	return false
}

func (f *UserFavoritesAndLiked) GetLikedItemByUniqueId(uniqueId string) *wv.LikedListElement {
	if f == nil {
		return nil
	}

	f.mut.Lock()
	for _, currentValue := range f.values {
		if currentValue == nil || len(currentValue.LikedList) == 0 {
			continue
		}

		for _, currentItem := range currentValue.LikedList {
			if currentItem.UniqueId == uniqueId {
				f.mut.Unlock()
				return currentItem
			}
		}
	}
	f.mut.Unlock()

	return nil
}

func (f *UserFavoritesAndLiked) DeleteLikedItemByUniqueId(uniqueId string) *wv.LikedListElement {
	if f == nil {
		return nil
	}

	var item *wv.LikedListElement
	var newList []*wv.LikedListElement

	f.mut.Lock()
	for _, currentValue := range f.values {
		if currentValue == nil || len(currentValue.LikedList) == 0 {
			continue
		}

		newList = nil

		for _, currentItem := range currentValue.LikedList {
			if currentItem.UniqueId == uniqueId {
				item = currentItem
				continue
			}
			newList = append(newList, currentItem)
		}

		if item != nil {
			currentValue.LikedList = newList
			break
		}
	}
	f.mut.Unlock()

	return item
}

func (f *UserFavoritesAndLiked) Delete(key string) {
	if f == nil {
		return
	}

	f.mut.Lock()
	delete(f.values, key)
	f.mut.Unlock()
}

func (f *UserFavoritesAndLiked) AddFavorite(info *wv.FavoriteInfo) {
	if f == nil {
		return
	}

	f.mut.Lock()
	all := f.values[info.FavoriteKey]
	f.mut.Unlock()
	if all == nil {
		f.values[info.FavoriteKey] = &UserFavoriteAndLikedInfo{
			FavoriteInfo: info,
		}
		return
	}
	all.FavoriteInfo = info
}

func (f *UserFavoritesAndLiked) AddLiked(liked *wv.LikedListElement) {
	if f == nil {
		return
	}

	f.mut.Lock()
	all := f.values[liked.LikedKey]
	if all == nil {
		f.values[liked.LikedKey] = &UserFavoriteAndLikedInfo{
			LikedList: []*wv.LikedListElement{liked},
		}
		f.mut.Unlock()
		return
	}
	all.LikedList = append(all.LikedList, liked)
	f.mut.Unlock()
}

func (f *UserFavoritesAndLiked) GetFavoriteInfo(key string) *wv.FavoriteInfo {
	if f == nil {
		return nil
	}

	f.mut.Lock()
	v := f.values[key]
	f.mut.Unlock()

	if v == nil {
		return nil
	}

	return v.FavoriteInfo
}

func (f *UserFavoritesAndLiked) GetLikedList(key string) []*wv.LikedListElement {
	if f == nil {
		return nil
	}

	f.mut.Lock()
	v := f.values[key]
	f.mut.Unlock()

	if v == nil {
		return nil
	}

	return v.LikedList
}

func (f *UserFavoritesAndLiked) Length() int {
	if f == nil {
		return 0
	}

	f.mut.Lock()
	l := len(f.values)
	f.mut.Unlock()

	return l
}

//---------------------------------------------------------

func (d *NewLikedListElementData) ToLikedListElement() *wv.LikedListElement {
	return &wv.LikedListElement{
		UniqueId:  d.getUniqueId(),
		OwnerId:   d.UserId,
		MediaId:   d.MediaId,
		Title:     d.Title,
		LikedKey:  d.LikedKey,
		SourceUrl: d.SourceUrl,
	}
}

func (d *NewLikedListElementData) getUniqueId() string {
	return ws.ToBase30(time.Now().Unix()) + likedListUIDSep + d.UserId.ToBase32()
}
