package wotoValues_test

import (
	"testing"
	wv "wp-server/wotoPacks/core/wotoValues"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func TestMediaGenreElementUniqueId(t *testing.T) {
	uniqueIdList := ssg.GetEmptyList[string]()
	for i := 0; i < 100; i++ {
		genreEl := &wv.MediaGenreElement{
			Genre: wv.GenreId(i),
		}
		genreEl.GenerateNewUniqueId()
		if uniqueIdList.Exists(genreEl.UniqueId) {
			t.Errorf("unique-id %s already exists in the list", genreEl.UniqueId)
			return
		}

		uniqueIdList.Add(genreEl.UniqueId)
	}

}
