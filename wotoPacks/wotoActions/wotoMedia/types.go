package wotoMedia

import (
	"time"
	wv "wp-server/wotoPacks/core/wotoValues"
)

type RegisterMediaData struct {
	Genre       wv.GenreId    `json:"genre"`
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
