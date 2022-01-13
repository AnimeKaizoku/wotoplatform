/*
 * This file is part of wp-server project (https://github.com/AnimeKaizoku/wotoplatform).
 * Copyright (c) 2021 ALiwoto.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

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
