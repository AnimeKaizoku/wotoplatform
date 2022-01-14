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

package versioning_test

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database"
	"wp-server/wotoPacks/entryPoints"
	"wp-server/wotoPacks/wotoActions"
	"wp-server/wotoPacks/wotoActions/versioning"
)

//---------------------------------------------------------

func TestWrongVersioning(t *testing.T) {
	config, err := wotoConfig.GetConfigByPath("../../../config.json")
	if err != nil {
		// print my current execution path and see if the problem
		// is from path? (ENOENT)
		log.Println(os.Executable())
		t.Errorf("couldn't get config: %v", err)
		return
	} else {
		listen(config, t)
	}

	// now, at the very least sleep 250 millisecond, and then try to
	// connect to the tcp listener which is listening in another
	// goroutine
	//time.Sleep(250 * time.Millisecond)

	addr, err := net.ResolveTCPAddr(config.Network, config.Bind+":"+config.Port)
	//addr, err := net.ResolveTCPAddr("tcp", ""+":"+config.Port)
	if err != nil {
		t.Errorf("couldn't resolve tcp address: %v", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		t.Errorf("couldn't dial TCP connection: %v", err)
		return
	}

	err = conn.SetKeepAlive(true)
	if err != nil {
		t.Errorf("failed to send set keep alive to the server-side: %v", err)
		return
	}

	n, err := writeMe(conn, []byte("hello"))
	if err != nil {
		t.Errorf("couldn't send a sample hello: %v", err)
		return
	} else if n == 0 {
		t.Error("zero byte sent at first")
		return
	}

	//writeMe(conn, []byte("how"))
	//writeMe(conn, []byte("are"))
	//writeMe(conn, []byte("are"))
	//writeMe(conn, []byte("you"))

	time.Sleep(100 * time.Millisecond)

	n, err = writeMe(conn, []byte("hello"))
	if err == nil {
		n, err = writeMe(conn, []byte("hello"))
		if err == nil {
			t.Error("connection is not closed")
			return
		} else if n != 0 {
			t.Error("more than zero bytes has been sent")
			return
		}
		return
	} else if n != 0 {
		t.Error("more than zero bytes has been sent")
		return
	}

	t.Log("done") // <-- don't remove it; break-point purposes
}

func TestCorrectVersioning(t *testing.T) {
	config, err := wotoConfig.GetConfigByPath("../../config.json")
	if err != nil {
		// print my current execution path and see if the problem
		// is from path? (ENOENT)
		log.Println(os.Executable())
		t.Errorf("couldn't get config: %v", err)
		return
	} else {
		//if entryPoints.MainListener != nil {
		//log.Println("no nil")
		//return
		//}
		listen(config, t)
	}

	// now, at the very least sleep 250 milliseconds, and then try to
	// connect to the tcp listener which is listening in another
	// goroutine
	//time.Sleep(250 * time.Millisecond)
	addr, err := net.ResolveTCPAddr(config.Network, config.Bind+":"+config.Port)
	//addr, err := net.ResolveTCPAddr("tcp", ""+":"+config.Port)
	if err != nil {
		t.Errorf("couldn't resolve tcp address: %v", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		t.Errorf("couldn't dial TCP connection: %v", err)
		return
	}

	err = conn.SetKeepAlive(true)
	if err != nil {
		t.Errorf("failed to send set keep alive to the server-side: %v", err)
		return
	}

	n, err := writeVersionAction(conn)
	if err != nil {
		t.Errorf("couldn't send version checking request: %v", err)
		time.Sleep(900 * time.Millisecond)
		return
	} else if n == 0 {
		t.Error("zero byte sent for versioning")
		time.Sleep(900 * time.Millisecond)
		return
	}

	//time.Sleep(250 * time.Millisecond)
	var b []byte
	b, err = readMe(conn, t)
	if err != nil {
		return
	}

	var resp = new(wotoActions.ActionResp)
	err = json.Unmarshal(b, resp)
	if err != nil {
		t.Errorf("got an error when tried to unmarshal the data: %v", err)
		return
	}

	t.Log("done") // <-- don't remove it; break-point purposes
}

//---------------------------------------------------------

func isInUseError(errStr string) bool {
	return strings.Contains(errStr, "address already in use") ||
		strings.Contains(errStr, "Only one usage of each socket")
}

func listen(config *wotoConfig.Config, t *testing.T) {
	if config.IsServerExternal() {
		return
	}

	l := entryPoints.MainListener
	if l != nil && !l.IsListenerClosed() {
		return
	} else {
		t.Cleanup(func() {
			closeListener(t)
		})
	}

	const maxTry = 1000
	ln, err := net.Listen(config.Network, config.Bind+":"+config.Port)
	if err != nil {
		if isInUseError(err.Error()) {
			for i := 0; i < maxTry; i++ {
				var myInt int
				myInt, err = strconv.Atoi(config.Port)
				if err != nil {
					t.Errorf("failed to listen on tcp: %v", err)
					return
				}
				myInt++

				ln, err = net.Listen(config.Network,
					config.Bind+":"+strconv.Itoa(myInt))
				if err != nil {
					if isInUseError(err.Error()) {
						continue
					}

					t.Errorf("failed to listen on tcp: %v", err)
					return
				} else {
					break
				}
			}
		} else {
			t.Errorf("failed to listen on tcp: %v", err)
			return
		}
	}

	err = database.StartDatabase()
	if err != nil {
		t.Errorf("failed to start a new db session: %v", err)
		return
	}

	go entryPoints.Listen(ln)
}

func writeMe(conn net.Conn, b []byte) (int, error) {
	bb := []byte(wv.MakeSureNum(len(b), 8))
	bb = wv.MakeSureByte(bb, 8)

	//n, err := conn.Write(bb)
	//if err != nil || n == 0 {
	//	return n, err
	//}

	//bb = append(bb, b...)
	//b = append(b, bb...)

	return conn.Write(append(bb, b...))
}

func closeListener(t *testing.T) {
	if entryPoints.MainListener == nil {
		return
	}

	// now close the connection so we can end our testing.
	// even if you try to close the listener for more than two times,
	// there should be no errors.
	_ = entryPoints.MainListener.CloseListener()
	err := entryPoints.MainListener.CloseListener()
	if err != nil {
		t.Logf("got an error while trying to close the listener: %v", err)
		return
	}
}

func writeVersionAction(conn net.Conn) (int, error) {
	vMap := map[string]string{
		"user_agent":   "wp-client",
		"version_key":  "2.1.1.5014",
		"version_hash": "f302bd7ffacbd295194f86620002b8250e8e9be0233a8055bcebc82c8612843ff9e0f09e42015d5e75581cc93d4c29a91388ed411641b543c8fb7b5a26a2a8b8",
		"client_id":    "cli-12345678910",
	}
	data, err := json.Marshal(vMap)
	if err != nil {
		return 0, err
	}

	e := entryPoints.RequestEntry{
		Action: wotoActions.ActionVersion,
		BatchExecute: wotoActions.BatchStr +
			versioning.BATCH_CHECK_VERSION,
		Data: string(data),
	}

	b, err := json.Marshal(&e)
	if err != nil {
		return 0, err
	}

	return writeMe(conn, b)
}

func readMe(conn net.Conn, t *testing.T) ([]byte, error) {
	b := make([]byte, 8)
	n, err := conn.Read(b)
	if err != nil {
		t.Errorf("got an error when tried to read the first bytes: %v", err)
		return nil, err
	} else if n != 8 {
		t.Errorf("expected to read 8 bytes of buffers, read %v instead", n)
		//closeListener(t)
		return nil, nil
	}

	count, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		t.Errorf("got an error when tried to convert first bytes to int: %v", err)
		//closeListener(t)
		return nil, err
	} else if count == 0 {
		t.Errorf("count was zero")
		//closeListener(t)
		return nil, nil
	}

	b = make([]byte, count)
	n, err = conn.Read(b)
	if err != nil {
		t.Errorf("got an error when tried to read the first bytes: %v", err)
		//closeListener(t)
		return nil, err
	} else if n != count {
		t.Errorf("expected to read %v bytes of buffers, read %v instead", count, n)
		//closeListener(t)
		return nil, nil
	}

	return b, nil
}

//---------------------------------------------------------
