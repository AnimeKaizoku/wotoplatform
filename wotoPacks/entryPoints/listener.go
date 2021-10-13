/*
 * This file is part of wp-server project (https://github.com/RudoRonuma/WotoPlatformBackend).
 * Copyright (c) 2021 AmanoTeam.
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
	"net"
	"time"
	"wp-server/wotoPacks/interfaces"
	"wp-server/wotoPacks/utils/logging"
	"wp-server/wotoPacks/utils/wotoValues"
	"wp-server/wotoPacks/wotoActions"
	"wp-server/wotoPacks/wotoActions/versioning"
)

// Listen function will listen for incoming connections
// using the specified listener argument.
func Listen(ln net.Listener) {
	if ln == nil {
		logging.SUGARED.Warn("listener cannot be nil")
		//log.Println("listener cannot be nil")
		return
	}
	if !isCheckingRegistration {
		isCheckingRegistration = true
		go checkRegistration()
	}

	var conn *wotoValues.WotoConnection
	var err error

	logging.SUGARED.Info("started to listening on: ",
		ln.Addr().String())

	MainListener = wotoValues.GetWotoListener(ln)

	defer func() {
		if MainListener != nil {
			MainListener.CloseListener()
		}
	}()

	for MainListener.CanAccept() {
		conn, err = MainListener.Accept(registerConnection)

		if err != nil {
			opErr, ok := err.(*net.OpError)
			if !ok {
				go logging.SUGARED.Error(err)
				//go errorHandling.HandleError(err)
				err = nil
				continue
			} else if opErr == nil {
				logging.SUGARED.Error("an unexpected error happened during accpeting "+
					"a new incoming connectiong from client", err)
				//log.Println("an unexpected error happened during accpeting "+
				//	"a new incoming connectiong from client", err)
				return
			}

			if isListenerClosed(opErr) {
				logging.SUGARED.Info("listener is closed, returning")
				//log.Println("listener is closed, returning")
				break
			} else {
				logging.SUGARED.Error(err)
				//go errorHandling.HandleError(err)
				err = nil
				continue
			}
		}

		// it's not duty of this loop,
		// you need to check if this connection is comming
		// from a valid cleint or not (in a separated goroutine)
		// if yes, do your duties in that goroutine
		// (please notice that you shouldn't create another goroutine
		// again and again.
		go safeCheckEntry(conn)

		// make sure you are not handling the previous connection
		// in the next loop.
		conn = nil
	}

}

// checkEntry checks the incoming connection and will do read and
// write operations on them.
func checkEntry(conn *wotoValues.WotoConnection) error {
	var req RequestEntery

	err := conn.ReadJson(&req)
	if err != nil {
		return err
	}

	logging.SUGARED.Debug("after reading json...")
	//log.Println("after reading json...")
	if len(req.BatchExecute) == 0 || !req.IsActionValid() {
		logging.SUGARED.Error("req.IsActionValid() returned false")
		return ErrActionOrBatchInvalid
	}

	req.Connection = conn
	var handler wotoValues.ReqHandler
	var parser func(interfaces.ReqBase) error

	// check if the current connection is registered or not, if not,
	// check if it wants to register itself using action versiong or not,
	// if not return unregistered error so the connection can be closed
	// kindly; otherwise, go forward.
	if !conn.IsRegistered() && req.Action != wotoActions.ACTION_VERSION {
		return ErrConnectionNotRegistered
	}

	logging.SUGARED.Debug("switching on req.Action")
	//log.Println("switching on req.Action")
	switch req.Action {
	case wotoActions.ACTION_VERSION:
		handler = versioning.HandleVersionAction
		parser = versioning.ParseBatchExecute
	default: // isn't it impossible? most probably yeah
		logging.SUGARED.Debug("invalid action:", req.Action)
		return ErrActionOrBatchInvalid
	}

	err = parser(&req)
	if err != nil {
		logging.SUGARED.Error(err)
		return err
	}

	return handler(&req)
}

// safeCheckEntry will start a loop for reading the data incoming
// from the client and will execute the batch execution's specified
// handler.
func safeCheckEntry(conn *wotoValues.WotoConnection) {
	var err error

	// set the dead line of connection to zero
	conn.SetDeadline(time.Time{})

	// create an infinite loop and check the entry comming requests.
	for {
		err = checkEntry(conn)
		// check if our check entry function returned any
		// error or not
		if err != nil {
			// you need to close the connection before handling the error.
			// since error handling may take some time,
			// (like saving the error in db, etc...)
			// so you need to close the connection to ensure
			// the safity.
			// and after that, handle the error somehow.
			// and break from the loop.
			conn.Close()
			go logging.SUGARED.Error(err)
			//errorHandling.HandleError(err)
			break
		}
	}
}

// isListenerClosed will check if the specified error is
// because that the listener if closed or not.
// to put it simply, if a listener is already closed, you shouldn't
// try to close it again, so you should simply ignore the error
// and return from the `Listen` function, so the server application
// can be closed easily after that.
func isListenerClosed(err *net.OpError) bool {
	if err.Source == nil && err.Op == "accept" &&
		err.Net == "tcp" {
		// the problem is in listener, return the function.
		return true
	}

	return false
}
