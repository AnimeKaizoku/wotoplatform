package wotoGroups

// batch execution values
const (
	BATCH_GET_GROUP_INFO_BY_ID                = "get_group_info_by_id"
	BATCH_GET_GROUP_INFO_BY_TELEGRAM_ID       = "get_group_info_by_telegram_id"
	BATCH_GET_GROUP_INFO_BY_TELEGRAM_USERNAME = "get_group_info_by_telegram_username"
	BATCH_GET_GROUP_CALL_INFO                 = "get_group_call_info"
	BATCH_CREATE_GROUP_CALL                   = "create_group_call"
	BATCH_DELETE_GROUP_CALL                   = "delete_group_call"
	BATCH_GET_GROUP_CALL_QUEUE                = "get_group_call_queue"
	BATCH_GET_GROUP_MEDIA_QUEUE               = "get_group_media_queue"
	BATCH_GET_GROUP_CALL_HISTORY              = "get_group_call_history"
	BATCH_ADD_TO_QUEUE                        = "add_to_queue"
	BATCH_ADD_MEDIA_TO_QUEUE                  = "add_media_to_queue"
)

// origin constants
const (
	OriginGetGroupInfoById               = "GetGroupInfoById"
	OriginGetGroupInfoByTelegramId       = "GetGroupInfoByTelegramId"
	OriginGetGroupInfoByTelegramUsername = "GetGroupInfoByTelegramUsername"
	OriginGetGroupCallInfo               = "GetGroupCallInfo"
	OriginCreateGroupCall                = "CreateGroupCall"
	OriginDeleteGroupCall                = "DeleteGroupCall"
	OriginGetGroupCallQueue              = "GetGroupCallQueue"
	OriginGetGroupMediaQueue             = "GetGroupMediaQueue"
	OriginGetGroupCallHistory            = "GetGroupCallHistory"
	OriginAddToQueue                     = "AddToQueue"
	OriginAddMediaToQueue                = "AddMediaToQueue"
)
