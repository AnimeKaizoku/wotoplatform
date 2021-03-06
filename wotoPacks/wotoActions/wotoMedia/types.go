package wotoMedia

import (
	"time"
	wv "wp-server/wotoPacks/core/wotoValues"
)

type RegisterMediaData struct {
	Company     wv.CompanyId  `json:"company"`
	Author      wv.AuthorId   `json:"author"`
	Episode     int           `json:"episode"`
	MediaType   string        `json:"media_type"`
	Title       string        `json:"title"`
	Duration    time.Duration `json:"duration"`
	Artist      string        `json:"artist"`
	Album       string        `json:"album"`
	Year        int           `json:"year"`
	Cover       string        `json:"cover"`
	File        string        `json:"file"`
	Thumbnail   string        `json:"thumbnail"`
	Lyrics      string        `json:"lyrics"`
	Lang        string        `json:"lang"`
	LangCode    string        `json:"lang_code"`
	Region      string        `json:"region"`
	SourceUrl   string        `json:"source_url"`
	ExternalUrl string        `json:"external_url"`
	IsPrivate   bool          `json:"is_private"`
	Description string        `json:"description"`
}

type RegisterMediaResult struct {
	MediaId wv.MediaModelId `json:"model_id"`
}

type GetMediaByIdData struct {
	MediaId wv.MediaModelId `json:"model_id"`
}

type GetMediaByIdResult struct {
	// ModelId is the unique id of this media-model. this field is used
	// to identify this media-model to another media-models. clients have
	// to use this id to interact with a media-model.
	ModelId wv.MediaModelId `json:"model_id"`

	// Genres field is an array of MediaGenreInfo type used to specify the
	// genres of this media-model. they shouldn't be repeated. this field is
	// ignored by sql and gorm, database packages have to use MediaGenreElement
	// to get this array.
	Genres      []*wv.MediaGenreInfo `json:"genres"`
	Company     wv.CompanyId         `json:"company"`
	Author      wv.AuthorId          `json:"author"`
	Episode     int                  `json:"episode"`
	MediaType   string               `json:"media_type"`
	Title       string               `json:"title"`
	Duration    time.Duration        `json:"duration"`
	Artist      string               `json:"artist"`
	Album       string               `json:"album"`
	Year        int                  `json:"year"`
	Cover       string               `json:"cover"`
	File        string               `json:"file"`
	Thumbnail   string               `json:"thumbnail"`
	Lyrics      string               `json:"lyrics"`
	Lang        string               `json:"lang"`
	LangCode    string               `json:"lang_code"`
	Region      string               `json:"region"`
	SourceUrl   string               `json:"source_url"`
	ExternalUrl string               `json:"external_url"`
	IsPrivate   bool                 `json:"is_private"`
	Description string               `json:"description"`
	CreatedAt   time.Time            `json:"created_at"`
	CreatedBy   wv.PublicUserId      `json:"created_by"`
	UpdatedBy   wv.PublicUserId      `json:"updated_by"`
}

type CreateNewGenreData struct {
	GenreTitle       string         `json:"genre_title"`
	GenreDescription string         `json:"genre_description"`
	AgeRange         wv.StringRange `json:"age_range"`
}

type CreateNewGenreResult struct {
	GenreId wv.GenreId `json:"genre_id"`
}

type DeleteGenreInfoData struct {
	GenreId    wv.GenreId `json:"genre_id"`
	GenreTitle string     `json:"genre_title"`
}

type EditGenreInfoData struct {
	GenreId          wv.GenreId     `json:"genre_id"`
	GenreTitle       string         `json:"genre_title"`
	GenreDescription string         `json:"genre_description"`
	AgeRange         wv.StringRange `json:"age_range"`
}

type EditGenreInfoResult struct {
	GenreInfo *wv.MediaGenreInfo `json:"genre_info"`
}

type AddMediaGenreData struct {
	MediaId    wv.MediaModelId `json:"model_id"`
	MediaGenre wv.GenreId      `json:"genre_id"`
}

type AddMediaGenreResult struct {
	MediaId     wv.MediaModelId `json:"model_id"`
	MediaGenres []wv.GenreId    `json:"media_genres"`
}

type RemoveMediaGenreData struct {
	MediaId    wv.MediaModelId `json:"model_id"`
	MediaGenre wv.GenreId      `json:"genre_id"`
}

type RemoveAddMediaGenreResult struct {
	MediaId     wv.MediaModelId `json:"model_id"`
	MediaGenres []wv.GenreId    `json:"media_genres"`
}

type GetMediaGenresData struct {
	MediaId wv.MediaModelId `json:"model_id"`
}

type GetMediaGenresResult struct {
	MediaId     wv.MediaModelId `json:"model_id"`
	MediaGenres []wv.GenreId    `json:"media_genres"`
}

type GetMediaGenresInfoData struct {
	MediaId wv.MediaModelId `json:"model_id"`
}

type GetMediaGenresInfoResult struct {
	MediaId         wv.MediaModelId      `json:"model_id"`
	MediaGenresInfo []*wv.MediaGenreInfo `json:"media_genres_info"`
}
type SearchGenreData struct {
}

type SearchGenreResult struct {
}

type DeleteMediaData struct {
}
