package versioning

import (
	wv "wp-server/wotoPacks/core/wotoValues"
	wa "wp-server/wotoPacks/wotoActions"
)

var (
	_batchValuesMap = map[wa.BatchExecution]bool{
		Batch_CHECK_VERSION: true,
	}
	_batchHandlers = map[wa.BatchExecution]wv.ReqHandler{
		Batch_CHECK_VERSION: batchCheckVersion,
	}
)
