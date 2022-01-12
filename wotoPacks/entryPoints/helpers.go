package entryPoints

import (
	"sync"
	"time"
	"wp-server/wotoPacks/core/wotoValues"
)

func checkRegistration() {
	if registrationMap == nil {
		registrationMap = make(map[*wotoValues.WotoConnection]bool)
		registrationMutex = new(sync.Mutex)
	}

	for {
		// sleep for 3 minutes
		// each loop should be done in intervals of 3 mins.
		// it will take a very short time iterating through
		// the registration map, so there will be no problem.
		time.Sleep(3 * time.Minute)

		// check if our listener is active or not,
		// if it's not active, return the function and free the current
		// goroutine.
		if !MainListener.CanAccept() {
			registrationMap = nil
			registrationMutex = nil
			return
		}

		// lock the mutex, so iterating through the map
		// doesn't cause any problem.
		registrationMutex.Lock()

		for key := range registrationMap {
			if !key.CanReadAndWrite() {
				delete(registrationMap, key)
			}
		}

		// unlock the mutex so another goroutines can use the registration
		// map as well.
		registrationMutex.Unlock()
	}
}

func registerConnection(c *wotoValues.WotoConnection) {
	registrationMutex.Lock()
	if !registrationMap[c] {
		registrationMap[c] = true
	}
	registrationMutex.Unlock()
}
