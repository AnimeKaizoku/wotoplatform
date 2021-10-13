package wotoActions

import "errors"

//---------------------------------------------------------

var ErrBatchParseFailed = errors.New("wotoActions: failed to parse the batch execution")
var ErrInvalidBatch = errors.New("wotoActions: invalid batch execution value")
var ErrNotRegistered = errors.New("wotoActions: batch execution needs registeration")

//---------------------------------------------------------
