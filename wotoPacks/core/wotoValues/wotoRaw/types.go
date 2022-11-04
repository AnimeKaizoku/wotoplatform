package wotoRaw

import (
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
)

type (
	PublicUserId          int64
	UserPermission        int
	PublicGroupId         string
	MediaModelId          string
	StringRange           string
	GenreId               uint32
	CompanyId             uint32
	AuthorId              uint32
	ProfilePictureModelId string
)

// UserInfo struct in wotoRaw is a low level struct.
// It shouldn't be used directly in any package.
// Instead, use the UserInfo struct in `wotoValues` package.
type UserInfo struct {
	UserId         PublicUserId   `json:"user_id" gorm:"primaryKey"`
	PrivateHash    string         `json:"private_hash"`
	Email          string         `json:"email"`
	Region         string         `json:"region"`
	Language       string         `json:"language"`
	Website        string         `json:"website"`
	AuthKey        string         `json:"auth_key"`
	AccessHash     string         `json:"access_hash"`
	Password       string         `json:"password"`
	PasswordHash   string         `json:"password_hash"`
	Permission     UserPermission `json:"permission"`
	Bio            string         `json:"bio"`
	SourceUrl      string         `json:"source_url"`
	Birthday       string         `json:"birthday"`
	AnilistUrl     string         `json:"anilist_url"`
	MyAnimelistUrl string         `json:"my_animelist_url"`
	RedditUrl      string         `json:"reddit_url"`
	GithubUrl      string         `json:"github_url"`
	GitlabUrl      string         `json:"gitlab_url"`
	TelegramId     int64          `json:"telegram_id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Username       string         `json:"username"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	IsVirtual      bool           `json:"is_virtual"`
	CreatedBy      PublicUserId   `json:"created_by"`

	RegenerateSaltedPassword func(*UserInfo)      `json:"-" gorm:"-" sql:"-"`
	metaProvider             ssg.MetaDataProvider `json:"-" gorm:"-" sql:"-"`
}

type GroupInfo struct {
	GroupId          PublicGroupId `json:"group_id" gorm:"primaryKey"`
	GroupRegion      string        `json:"group_region"`
	GroupUsername    string        `json:"group_username"`
	TelegramId       int64         `json:"telegram_id"`
	TelegramUsername string        `json:"telegram_username"`
	CurrentPlaying   MediaModelId  `json:"current_playing" gorm:"-" sql:"-"`
}

type GroupCallInfo struct {
	GroupId   PublicGroupId
	StartedBy PublicUserId
	StartedAt time.Time
}

type FavoriteValue struct {
	UserId      PublicUserId `json:"user_id" gorm:"primaryKey"`
	FavoriteKey string       `json:"favorite_key"`
	TheValue    string       `json:"the_value"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

type LikedListElement struct {
	UniqueId     string       `json:"unique_id" gorm:"primaryKey"`
	OwnerId      PublicUserId `json:"owner_id"`
	MediaId      MediaModelId `json:"media_id"`
	Title        string       `json:"title"`
	LikedKey     string       `json:"liked_key"`
	SourceUrl    string       `json:"source_url"`
	ReferenceUrl string       `json:"reference_url"`
	UpdatedAt    time.Time    `json:"-"`
}

type MediaGenreInfo struct {
	GenreId          GenreId      `json:"genre_id" gorm:"primaryKey"`
	GenreTitle       string       `json:"genre_title"`
	GenreDescription string       `json:"genre_description"`
	AgeRange         StringRange  `json:"age_range"`
	CreatedAt        time.Time    `json:"created_at"`
	CreatedBy        PublicUserId `json:"created_by"`
	UpdatedBy        PublicUserId `json:"updated_by"`
}

// MediaGenreElement struct contains information about a
// media-model that has a genre. this struct makes us able to add
// multiple genres to a single media-model.
// this struct has a model in mediaDatabase and is inserted in that
// package, look at helpers.go file in that package for more information
// about how we are inserting them in db.
// See also: https://github.com/AnimeKaizoku/wotoplatform/issues/21
type MediaGenreElement struct {
	// UniqueId field specifies the unique-id of this element used to
	// distinguish the element in database. for generating a new
	// unique-id for an instance of this struct, you have to call
	// `GenerateUniqueId` method on this variable.
	UniqueId string `json:"unique_id" gorm:"primaryKey"`

	// MediaId field is the media-model id that this element is
	// referring to.
	MediaId MediaModelId `json:"media_id"`

	// Genre is the genre-id that this element is referring to.
	Genre GenreId `json:"genre"`

	// CreatedBy field is the user-id of the person that created this
	// element in the db.
	CreatedBy PublicUserId `json:"created_by"`
}

// MediaModel struct contains information about a specified media-model.
// users are able to interact with media-models (create, modify, delete),
// add them to their different lists (such as playlists) or add them to
// their schedule.
type MediaModel struct {
	// ModelId is the unique id of this media-model. this field is used
	// to identify this media-model to another media-models. clients have
	// to use this id to interact with a media-model.
	ModelId MediaModelId `json:"model_id" gorm:"primaryKey"`

	// Genres field is an array of MediaGenreInfo type used to specify the
	// genres of this media-model. they shouldn't be repeated. this field is
	// ignored by sql and gorm, database packages have to use MediaGenreElement
	// to get this array.
	Genres        []*MediaGenreInfo    `json:"genres" gorm:"-" sql:"-"`
	GenreElements []*MediaGenreElement `json:"-" gorm:"-" sql:"-"`
	Company       CompanyId            `json:"company"`
	Author        AuthorId             `json:"author"`
	Episode       int                  `json:"episode"`
	MediaType     string               `json:"media_type"`
	Title         string               `json:"title"`
	Duration      time.Duration        `json:"duration"`
	Artist        string               `json:"artist"`
	Album         string               `json:"album"`
	Year          int                  `json:"year"`
	Cover         string               `json:"cover"`
	File          string               `json:"file"`
	Thumbnail     string               `json:"thumbnail"`
	Lyrics        string               `json:"lyrics"`
	Lang          string               `json:"lang"`
	LangCode      string               `json:"lang_code"`
	Region        string               `json:"region"`
	SourceUrl     string               `json:"source_url"`
	ExternalUrl   string               `json:"external_url"`
	IsPrivate     bool                 `json:"is_private"`
	Description   string               `json:"description"`
	CreatedAt     time.Time            `json:"created_at"`
	CreatedBy     PublicUserId         `json:"created_by"`
	UpdatedBy     PublicUserId         `json:"updated_by"`

	mediaMetaData ssg.MetaDataProvider
}
