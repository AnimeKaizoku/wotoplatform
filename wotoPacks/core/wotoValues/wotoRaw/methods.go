package wotoRaw

import "time"

func (u *UserInfo) HasUsername() bool {
	return u.Username != ""
}

func (u *UserInfo) HasTelegramId() bool {
	return u.TelegramId != 0
}

func (u *UserInfo) SetCachedTime() {
	u.cachedTime = time.Now()
}

func (u *UserInfo) IsCacheExpired(d time.Duration) bool {
	return time.Since(u.cachedTime) > d
}

func (u *UserInfo) IsPasswordCorrect(password string) bool {
	// TODO: encrypt the password
	return u.Password == password
}

func (u *UserInfo) GetPublicId() PublicUserId {
	return u.UserId
}

func (u *UserInfo) CanCreateAccount() bool {
	return u.Permission >= PermissionAdmin
}
