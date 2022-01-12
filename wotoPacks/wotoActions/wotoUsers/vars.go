package wotoUsers

import (
	wv "wp-server/wotoPacks/core/wotoValues"
	wa "wp-server/wotoPacks/wotoActions"
)

var (
	_batchValuesMap = map[wa.BatchExecution]bool{
		BATCH_LOGIN_USER:    true,
		BATCH_REGISTER_USER: true,
	}
	_batchHandlers = map[wa.BatchExecution]wv.ReqHandler{
		BATCH_LOGIN_USER:    batchLoginUser,
		BATCH_REGISTER_USER: batchRegisterUser,
	}
)
