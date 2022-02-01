package usersDatabase

import (
	"sync"
	wv "wp-server/wotoPacks/core/wotoValues"
)

//---------------------------------------------------------

func (m *favoriteManager) Length(id wv.PublicUserId) int {
	m.mut.Lock()
	favorites := m.values[id]
	m.mut.Unlock()

	return favorites.Length()
}

func (m *favoriteManager) GetFavorites(id wv.PublicUserId) *UserFavorites {
	m.mut.Lock()
	f := m.values[id]
	m.mut.Unlock()

	return f
}

func (m *favoriteManager) GetUserFavorite(id wv.PublicUserId, key string) *wv.FavoriteInfo {
	return m.GetFavorites(id).GetValue(key)
}

func (m *favoriteManager) AddFavorite(f *wv.FavoriteInfo) {
	favorites := m.GetFavorites(f.UserId)
	if favorites == nil {
		favorites = &UserFavorites{
			mut:    &sync.Mutex{},
			values: make(map[string]*wv.FavoriteInfo),
		}

		m.mut.Lock()
		m.values[f.UserId] = favorites
		m.mut.Unlock()
	}

	favorites.Add(f)
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

	m.GetFavorites(id).Delete(key)
	return info
}

func (m *favoriteManager) Exists(id wv.PublicUserId, key string) bool {
	return m.GetFavorites(id).Exists(key)
}

func (m *favoriteManager) LoadAll(favorites []*wv.FavoriteInfo) {
	for _, current := range favorites {
		m.AddFavorite(current)
	}
}

//---------------------------------------------------------

func (f *UserFavorites) Exists(key string) bool {
	if f == nil {
		return false
	}

	f.mut.Lock()
	b := f.values[key] != nil
	f.mut.Unlock()

	return b
}

func (f *UserFavorites) Delete(key string) {
	if f == nil {
		return
	}

	f.mut.Lock()
	delete(f.values, key)
	f.mut.Unlock()
}

func (f *UserFavorites) Add(info *wv.FavoriteInfo) {
	if f == nil {
		return
	}

	f.mut.Lock()
	f.values[info.FavoriteKey] = info
	f.mut.Unlock()
}

func (f *UserFavorites) GetValue(key string) *wv.FavoriteInfo {
	if f == nil {
		return nil
	}

	f.mut.Lock()
	v := f.values[key]
	f.mut.Unlock()

	return v
}

func (f *UserFavorites) Length() int {
	if f == nil {
		return 0
	}

	f.mut.Lock()
	l := len(f.values)
	f.mut.Unlock()

	return l
}
