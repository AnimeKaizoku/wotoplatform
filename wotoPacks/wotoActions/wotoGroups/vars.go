package wotoGroups

import (
	wv "wp-server/wotoPacks/core/wotoValues"
	wa "wp-server/wotoPacks/wotoActions"
)

var (
	_batchHandlers = map[wa.BatchExecution]wv.ReqHandler{
		BATCH_GET_GROUP_INFO_BY_ID:   batchGetGroupInfoById,
		BATCH_GET_GROUP_CALL_INFO:    batchGetGroupCallInfo,
		BATCH_CREATE_GROUP_CALL:      batchCreateGroupCall,
		BATCH_DELETE_GROUP_CALL:      batchDeleteGroupCall,
		BATCH_GET_GROUP_CALL_QUEUE:   batchGetGroupCallQueue,
		BATCH_GET_GROUP_MEDIA_QUEUE:  batchGetGroupMediaQueue,
		BATCH_GET_GROUP_CALL_HISTORY: batchGetGroupCallHistory,
		BATCH_ADD_TO_QUEUE:           batchAddToQueue,
		BATCH_ADD_MEDIA_TO_QUEUE:     batchAddMediaToQueue,
	}
)
