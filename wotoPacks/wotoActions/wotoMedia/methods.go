package wotoMedia

import mdb "wp-server/wotoPacks/database/mediaDatabase"

func (d *RegisterMediaData) ToNewMediaData() *mdb.NewMediaData {
	return &mdb.NewMediaData{
		Company:     d.Company,
		Author:      d.Author,
		Episode:     d.Episode,
		MediaType:   d.MediaType,
		Title:       d.Title,
		Duration:    d.Duration,
		Artist:      d.Artist,
		Album:       d.Album,
		Year:        d.Year,
		Cover:       d.Cover,
		File:        d.File,
		Thumbnail:   d.Thumbnail,
		Lyrics:      d.Lyrics,
		Lang:        d.Lang,
		LangCode:    d.LangCode,
		Region:      d.Region,
		SourceUrl:   d.SourceUrl,
		ExternalUrl: d.ExternalUrl,
		IsPrivate:   d.IsPrivate,
		Description: d.Description,
	}
}
