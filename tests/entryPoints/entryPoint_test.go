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

package entryPoints_test

import (
	"fmt"
	"log"
	"net"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
	"wp-server/wotoPacks/core/wotoConfig"
	wv "wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/database"
	"wp-server/wotoPacks/entryPoints"
)

var listener net.Listener

func TestWrongEntryPoint(t *testing.T) {
	config, err := wotoConfig.GetConfig()
	if err != nil {
		// print my current execution path and see if the problem
		// is from path? (ENOENT)
		log.Println(os.Executable())
		t.Errorf("couldn't get config: %v", err)
		return
	} else {
		listen(config, t)
	}

	for entryPoints.MainListener == nil {
		// now, at the very least sleep 250 millisecond, and then try to
		// connect to the tcp listener which is listening in another
		// goroutine
		time.Sleep(250 * time.Millisecond)
	}

	addr, err := net.ResolveTCPAddr("tcp", config.Bind+":"+config.Port)
	if err != nil {
		t.Errorf("couldn't resolve tcp address: %v", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		t.Errorf("couldn't dial TCP connection: %v", err)
		return
	}

	writeMe(conn, []byte("hello"))
	n, err := conn.Write([]byte("hello"))
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

	time.Sleep(250 * time.Millisecond)

	n, err = conn.Write([]byte("hello"))
	if err == nil {
		t.Error("connection is not closed.")
		return
	} else if n != 0 {
		t.Error("more than zero bytes has been sent")
		return
	}
}

func OldTestWorkerPool01(t *testing.T) {
	worker := func(ports chan int, wg *sync.WaitGroup) {
		for p := range ports {
			if p == 0 {
				t.Errorf("port is zero")
			}
			wg.Done()
		}
	}
	ports := make(chan int, 100)

	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func OldTestWorkerPoolWithSort01(t *testing.T) {
	worker := func(ports, results chan int) {
		for p := range ports {
			address := fmt.Sprintf("scanme.nmap.org:%d", p)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				results <- 0
				continue
			}
			conn.Close()
			results <- p
		}
	}

	ports := make(chan int, 100)
	results := make(chan int)
	var openPorts []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openPorts)
	for _, port := range openPorts {
		t.Logf("%d open\n", port)
	}
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
	ln, err := net.Listen("tcp", config.Bind+":"+config.Port)
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
				config.Port = strconv.Itoa(myInt)
				ln, err = net.Listen("tcp", config.Bind+":"+config.Port)
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

	listener = ln

	err = database.StartDatabase()
	if err != nil {
		t.Errorf("failed to start a new db session: %v", err)
	}

	go entryPoints.Listen(ln, false)
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

func writeMe(conn net.Conn, b []byte) (int, error) {
	bb := []byte(wv.MakeSureNum(len(b), 8))
	bb = wv.MakeSureByte(bb, 8)

	return conn.Write(append(bb, b...))
}
