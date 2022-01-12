package wotoUsers

import (
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
		if !_batchValuesMap[b] {
			return false
		}
	}

	return true
}
