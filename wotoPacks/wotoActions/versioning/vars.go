package versioning

import "wp-server/wotoPacks/wotoActions"

var (
	_batchValuesMap = map[wotoActions.BatchExecution]bool{
		BatchCheckVersion: true,
	}
)
