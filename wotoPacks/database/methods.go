package database

import (
	"time"
	"wp-server/wotoPacks/interfaces"
)

func (u *UserInfo) IsEmpty() bool {
	return len(u.UserName) == 0 || len(u.PrivateID) == 0 ||
		len(u.PublicID) == 0
}

func (u *UserInfo) ComparePrimaryRaw(raw interfaces.RawUser) bool {
	return u.UserName == raw.GetName() &&
		u.PrivateID == raw.GetPrivateID() &&
		u.PublicID == raw.GetPublicID()
}

func (u *UserInfo) GetPublicID() string {
	return u.PublicID
}

func (u *UserInfo) GetPrivateID() string {
	return u.PrivateID
}

func (u *UserInfo) GetName() string {
	return u.UserName
}

func (u *UserInfo) GetLastSeen() string {
	return u.LastSeen
}

func (u *UserInfo) GetUserIntro() string {
	return u.UserIntro
}

func (u *UserInfo) GetAvatar() string {
	return u.UserAvatar
}

func (u *UserInfo) GetAvatarFrame() string {
	return u.UserAvatarFrame
}

func (u *UserInfo) GetPassword() string {
	return u.Password
}

func (u *UserInfo) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *UserInfo) GetSocialvoidUsername() string {
	return u.SoacialvoidUsername
}

func (u *UserInfo) GetSvUsername() string {
	return u.SoacialvoidUsername
}

func (u *UserInfo) GetUserLever() uint16 {
	return u.UserLevel
}

func (u *UserInfo) GetUserVIPLevel() uint8 {
	return u.UserVIPLevel
}

func (u *UserInfo) GetCurrentExp() string {
	return u.UserCurrentExp
}

func (u *UserInfo) GetCurrentVIPExp() string {
	return u.UserCurrentVIPExp
}

func (u *UserInfo) GetTotalExp() string {
	return u.UserTotalExp
}

func (u *UserInfo) GetTotalVIPExp() string {
	return u.UserTotalVIPExp
}

func (u *UserInfo) GetMaxExp() string {
	return u.UserMaxExp
}

func (u *UserInfo) GetMaxVIPExp() string {
	return u.UserMaxVIPExp
}

func (u *UserInfo) GetCity() string {
	return u.UserCity
}
