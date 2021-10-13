package entryPoints

import (
	"sync"
	"time"
	"wp-server/wotoPacks/utils/wotoValues"
)

func checkRegistration() {
	if registerationMap == nil {
		registerationMap = make(map[*wotoValues.WotoConnection]bool)
		registerationMutex = new(sync.Mutex)
	}

	for {
		// sleep for 3 minutes
		// each loop should be done in intervals of 3 mins.
		// it will take a very short time iterating through
		// the registeration map, so there will be no problem.
		time.Sleep(3 * time.Minute)

		// check if our listener is active or not,
		// if it's not active, return the function and free the current
		// goroutine.
		if !MainListener.CanAccept() {
			registerationMap = nil
			registerationMutex = nil
			return
		}

		// lock the mutix, so iterating through the map
		// doesn't cause any problem.
		registerationMutex.Lock()

		for key := range registerationMap {
			if !key.CanReadAndWrite() {
				delete(registerationMap, key)
			}
		}

		// unlock the mutix so another goroutines can use the registeration
		// map as well.
		registerationMutex.Unlock()
	}
}

func registerConnection(c *wotoValues.WotoConnection) {
	// commented out because it seems it's most likely impossible...
	//if registerationMap == nil {
	//	registerationMap = make(map[*wotoValues.WotoConnection]bool)
	//	registerationMutex = new(sync.Mutex)
	//}

	registerationMutex.Lock()
	if !registerationMap[c] {
		registerationMap[c] = true
	}
	registerationMutex.Unlock()
}
