package usersDatabase

import wv "wp-server/wotoPacks/core/wotoValues"

type NewUserData struct {
	Username   string
	Password   string
	Permission wv.UserPermission
	By         wv.PublicUserId
}
