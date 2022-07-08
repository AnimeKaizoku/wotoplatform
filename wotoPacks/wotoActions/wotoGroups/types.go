package wotoGroups

import wv "wp-server/wotoPacks/core/wotoValues"

type GetGroupInfByIdoData struct {
	GroupId wv.PublicGroupId `json:"group_id"`
}

type GetGroupInfoByIdResult struct {
	GroupId          wv.PublicGroupId `json:"group_id"`
	GroupRegion      string           `json:"group_region"`
	GroupUsername    string           `json:"group_username"`
	TelegramId       int64            `json:"telegram_id"`
	TelegramUsername string           `json:"telegram_username"`
	CurrentPlaying   wv.MediaModelId  `json:"current_playing"`
}
