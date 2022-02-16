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
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
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
	"wp-server/wotoPacks/core/wotoCrypto"
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
		fmt.Printf("%d open\n", port)
	}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(t *testing.T, data []byte, passphrase string) []byte {
	ourKey := []byte(createHash(passphrase))
	log.Println("Our key is: ", ourKey)
	block, _ := aes.NewCipher(ourKey)
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Error(err)
		return nil
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		t.Error(err)
		return nil
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(t *testing.T, data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Error(t)
		return nil
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Error(t)
		return nil
	}

	nonceSize := gcm.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		t.Error(t)
		return nil
	}

	return plainText
}

func TestCrypto01(t *testing.T) {
	cipherText := encrypt(t, []byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", cipherText)
	plainText := decrypt(t, cipherText, "password")
	fmt.Printf("Decrypted: %s\n", plainText)
}

func TestCrypto02(t *testing.T) {
	allData := []string{
		"hello aaaaaaaaaaaaaaaaaaaaaaaa",
		"hello~354(*#&%*#%&^ {♔} hey hi!",
		"エイリ・ヲト♔エイリ・ヲトエイリ・ヲトエイリ・ヲ♔トエイリ・ヲト",
		string([]rune{
			0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389,
			0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389,
			0x1f525, 0x1f525, 0x1f525, 0xfe0f, 0x1f48b, 0x1f602, 0x1f602, 0x1f602,
			0x1f923, 0x1f61f, 0x1f92f, 0x1f61d, 0x2764, 0xfe0f, 0x1f1ef, 0x1f1f5,
			0x1f1f2, 0x1f1f0, 0x1f1e8, 0x1f1fa, 0x1f1e6, 0x1f1f7, 0x1f1fa, 0x1f1f8,
			0x21, 0x32, 0x45, 0x9f4, 0x9f5, 0x9f6, 0x9f7, 0x9f8, 0x9f9, 0x9fa, 0x9fb,
			0x1f642, 0x1f643, 0x1f917, 0x1f918, 0x1f914, 0x1f915, 0x1f922, 0x1f92d,
		}),
	}
	for i, current := range allData {
		testCrypto02Worker(t, i, []byte(current))
	}
}

var allKeysTest02 = [][]byte{
	{0x73, 0x64, 0x66, 0x69, 0x73, 0x68, 0x67, 0x6a, 0x64, 0x67, 0x68, 0x6b, 0x73, 0x64, 0x5e},
	{0x25, 0x28, 0x26, 0x68, 0x2a, 0x98, 0x00, 0x54},
	{0x76, 0x49, 0xab, 0xac, 0x81, 0x19, 0xb2, 0x46, 0xce, 0xe9, 0x8e, 0x9b, 0x12, 0xe9, 0x19, 0x7d,
		0x50, 0x86, 0xcb, 0x9b, 0x50, 0x72, 0x19, 0xee, 0x95, 0xdb, 0x11, 0x3a, 0x91, 0x76, 0x78, 0xb2,
		0x73, 0xbe, 0xd6, 0xb8, 0xe3, 0xc1, 0x74, 0x3b, 0x71, 0x16, 0xe6, 0x9e, 0x22, 0x22, 0x95, 0x16,
		0x3f, 0xf1, 0xca, 0xa1, 0x68, 0x1f, 0xac, 0x09, 0x12, 0x0e, 0xca, 0x30, 0x75, 0x86, 0xe1, 0xa7},
	{0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
		0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4},
}

func testCrypto02Worker(t *testing.T, index int, originData []byte) {
	var key []byte
	if index >= len(allKeysTest02) {
		mySha := sha256.New()
		mySha.Write(allKeysTest02[index])
		key = mySha.Sum(nil)
	} else {
		key = allKeysTest02[index]
	}

	data := wotoCrypto.EncryptData(key, originData)
	myData := wotoCrypto.DecryptData(key, data)
	if bytes.Equal(myData, originData) {
		t.Error("data sequences are not equal: ", myData, originData)
		return
	}
}

var (
	allHashesTestPresentKey01 = [][]byte{
		{0x73, 0x64, 0x66, 0x69, 0x73, 0x68, 0x67, 0x6a, 0x64, 0x67, 0x68, 0x6b, 0x73, 0x64, 0x5e},
		{0x25, 0x28, 0x26, 0x68, 0x2a, 0x98, 0x00, 0x54, 0xc1, 0x74, 0x3b, 0x71, 0x16, 0xe6, 0x22,
			0x22, 0x95, 0x16, 0x00, 0x00, 0x00, 0x87, 0x98, 0xee, 0x6f, 0xff},
		{0x76, 0x49, 0xab, 0xac, 0x81, 0x19, 0xb2, 0x46, 0xce, 0xe9, 0x8e, 0x9b, 0xe9, 0x19, 0x7d,
			0x50, 0x86, 0xcb, 0x9b, 0x50, 0x72, 0x19, 0xee, 0x95, 0xdb, 0x11, 0x91, 0x76, 0x78, 0xb2,
			0x73, 0xbe, 0xd6, 0xb8, 0xe3, 0xc1, 0x74, 0x3b, 0x71, 0x16, 0xe6, 0x22, 0x22, 0x95, 0x16,
			0x3f, 0xf1, 0xca, 0xa1, 0x68, 0x1f, 0xac, 0x09, 0x12, 0x0e, 0xca, 0x75, 0x86, 0xe1, 0xa7},
		{0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
			0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0xa3, 0x09, 0x14, 0xdf, 0xf4},
		commonInput,
	}
	allCryptoLayers = []wotoCrypto.CryptoLayer{
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(allHashesTestPresentKey01[0x0]),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(allHashesTestPresentKey01[0x1]),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(allHashesTestPresentKey01[0x2]),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(allHashesTestPresentKey01[0x3]),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(allHashesTestPresentKey01[0x4]),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(commonKey128),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(commonKey192),
		},
		{
			Kind: wotoCrypto.CryptoLayerKindO27,
			Hash: string(commonKey256),
		},
	}
)

// Common values for tests.

var commonInput = []byte{
	0x6b, 0xc1, 0xbe, 0xe2, 0x2e, 0x40, 0x9f, 0x96, 0xe9, 0x3d, 0x7e, 0x11, 0x73, 0x93, 0x17, 0x2a,
	0xae, 0x2d, 0x8a, 0x57, 0x1e, 0x03, 0xac, 0x9c, 0x9e, 0xb7, 0x6f, 0xac, 0x45, 0xaf, 0x8e, 0x51,
	0x30, 0xc8, 0x1c, 0x46, 0xa3, 0x5c, 0xe4, 0x11, 0xe5, 0xfb, 0xc1, 0x19, 0x1a, 0x0a, 0x52, 0xef,
	0xf6, 0x9f, 0x24, 0x45, 0xdf, 0x4f, 0x9b, 0x17, 0xad, 0x2b, 0x41, 0x7b, 0xe6, 0x6c, 0x37, 0x10,
}

var commonKey128 = []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}

var commonKey192 = []byte{
	0x8e, 0x73, 0xb0, 0xf7, 0xda, 0x0e, 0x64, 0x52, 0xc8, 0x10, 0xf3, 0x2b, 0x80, 0x90, 0x79, 0xe5,
	0x62, 0xf8, 0xea, 0xd2, 0x52, 0x2c, 0x6b, 0x7b,
}

var commonKey256 = []byte{
	0x60, 0x3d, 0xeb, 0x10, 0x15, 0xca, 0x71, 0xbe, 0x2b, 0x73, 0xae, 0xf0, 0x85, 0x7d, 0x77, 0x81,
	0x1f, 0x35, 0x2c, 0x07, 0x3b, 0x61, 0x08, 0xd7, 0x2d, 0x98, 0x10, 0xa3, 0x09, 0x14, 0xdf, 0xf4,
}

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

type pKey = wotoCrypto.WotoKey

func TestPresentKeyCrypto01(t *testing.T) {
	allData := []string{
		"hello aaaaaaaaaaaaaaaaaaaaaaaa",
		"hello~354(*#&%*#%&^ {♔} hey hi! >> how are you doing?? << >.<",
		"エイリ・ヲト♔エイリ・ヲトエイリ・ヲトエイリ・ヲ♔トエイリ・ヲト -> <- <html> </html>&&^.^&&",
		string([]rune{
			0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389,
			0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389, 0x1f389,
			0x1f525, 0x1f525, 0x1f525, 0xfe0f, 0x1f48b, 0x1f602, 0x1f602, 0x1f602,
			0x1f923, 0x1f61f, 0x1f92f, 0x1f61d, 0x2764, 0xfe0f, 0x1f1ef, 0x1f1f5,
			0x1f1f2, 0x1f1f0, 0x1f1e8, 0x1f1fa, 0x1f1e6, 0x1f1f7, 0x1f1fa, 0x1f1f8,
			0x21, 0x32, 0x45, 0x9f4, 0x9f5, 0x9f6, 0x9f7, 0x9f8, 0x9f9, 0x9fa, 0x9fb,
			0x1f642, 0x1f643, 0x1f917, 0x1f918, 0x1f914, 0x1f915, 0x1f922, 0x1f92d,
		}),
		string(commonIV),
	}
	presentKey := wotoCrypto.GeneratePresentKey(wotoCrypto.WotoAlgorithmM250)
	_ = presentKey.AppendLayer(&allCryptoLayers[0x0])
	_ = presentKey.AppendLayer(&allCryptoLayers[0x1])
	_ = presentKey.AppendLayer(&allCryptoLayers[0x2])
	_ = presentKey.AppendLayer(&allCryptoLayers[0x3])
	_ = presentKey.AppendLayer(&allCryptoLayers[0x4])
	_ = presentKey.SetSignatureByFunc(sha256.New)
	for _, current := range allData {
		testPresentKeyCrypto01Worker(t, presentKey, []byte(current))
	}

	b, _ := json.Marshal(map[string]int{
		"0x1": 0x1, "0x2": 0x2,
		"0x3": 0x3, "0x4": 0x4,
		"0x5": 0x5, "0x6": 0x6,
		"0x7": 0x7, "0x8": 0x8,
		"0x9": 0x9, "0x1a": 0x1a,
		"0x1b": 0x1b, "0x1c": 0x1c,
		"0x1d": 0x1d, "0x1e": 0x1e,
		">>": 0x1f, "<<": 0x20,
		"=": 0x4f, "==": 0x2f,
		"=\"": 0x2f, "==\"": 0x3f,
		"\xef": 0x21, "\xed": 0x22,
		"\xdf": 0x23, "\xdd": 0x24,
	})
	testPresentKeyCrypto01Worker(t, presentKey, b)
}

func testPresentKeyCrypto01Worker(t *testing.T, key pKey, originData []byte) {
	data := key.Encrypt(originData)
	myData := key.Decrypt(data)

	if !bytes.Equal(myData, originData) {
		t.Error("data sequences are not equal:")
		logStr("myData:")
		logB(myData)
		logStr("origin data:")
		logB(originData)
		return
	}
}

func logStr(value string) {
	log.Println(value)
}
func logB(b []byte) {
	logStr(string(b))
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

	go entryPoints.Listen(ln)
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
