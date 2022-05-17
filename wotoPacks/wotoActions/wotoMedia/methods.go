package wotoMedia

import (
	wv "wp-server/wotoPacks/core/wotoValues"
	mdb "wp-server/wotoPacks/database/mediaDatabase"
)

//---------------------------------------------------------

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

//---------------------------------------------------------

func (g *CreateNewGenreData) ToMediaGenreInfo(by wv.PublicUserId) *wv.MediaGenreInfo {
	return &wv.MediaGenreInfo{
		GenreTitle:       g.GenreTitle,
		GenreDescription: g.GenreDescription,
		AgeRange:         g.AgeRange,
		CreatedBy:        by,
		UpdatedBy:        by,
	}
}

//---------------------------------------------------------

func (e *EditGenreInfoData) UpdateGenreInfoFields(info *wv.MediaGenreInfo, by wv.PublicUserId) {
	info.GenreTitle = e.GenreTitle
	info.GenreDescription = e.GenreDescription
	info.AgeRange = e.AgeRange
	info.UpdatedBy = by
}

//---------------------------------------------------------
