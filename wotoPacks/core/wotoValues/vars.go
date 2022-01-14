package wotoValues

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

//---------------------------------------------------------

var ErrCantAccept = errors.New("woto listener: can't accept any new connections")
var ErrCantReadOrWrite = errors.New("woto connection: can't read or write from this connection")
var ErrCouldNotWriteFirstBytes = errors.New("woto connection: couldn't write the first bytes")
var ErrValueNil = errors.New("woto connection: interface value cannot be nil")
var ErrValueEmpty = errors.New("woto connection: received value was empty")

//---------------------------------------------------------

var (
	SESSION      *gorm.DB
	SessionMutex = &sync.Mutex{}
)
