package wotoGroups

import (
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/wotoActions"
)

func ParseBatchExecute(b interfaces.ReqBase) error {
	arr, err := wotoActions.ParseBatchExecute(b.GetBatchExecute())
	if err != nil {
		return err
	}

	if len(arr) == 0 || !batchValuesValid(arr) {
		return wotoActions.ErrBatchParseFailed
	}

	b.SetBatchValues(arr)

	return nil
}

func batchValuesValid(e []wotoActions.BatchExecution) bool {
	for _, b := range e {
		if _batchHandlers[b] == nil {
			return false
		}
	}

	return true
}

func toGetGroupInfoResult(info *wv.GroupInfo) *GetGroupInfoByIdResult {
	return &GetGroupInfoByIdResult{
		GroupId:          info.GroupId,
		GroupRegion:      info.GroupRegion,
		GroupUsername:    info.GroupUsername,
		TelegramId:       info.TelegramId,
		TelegramUsername: info.TelegramUsername,
		CurrentPlaying:   info.CurrentPlaying,
	}
}
