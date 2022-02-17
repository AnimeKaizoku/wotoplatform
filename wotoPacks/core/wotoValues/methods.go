package wotoValues

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"
	"wp-server/wotoPacks/core/utils/logging"

	"github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

//---------------------------------------------------------

func (l *WotoListener) IsListenerClosed() bool {
	return l.isClosed
}

func (l *WotoListener) CloseListener() error {
	if !l.isClosed && l.listener != nil {
		l.isClosed = true
		return l.listener.Close()
	}

	return nil
}

func (l *WotoListener) CanAccept() bool {
	return !l.isClosed && l.listener != nil
}

func (l *WotoListener) Accept(r Registerer) (*WotoConnection, error) {
	if !l.CanAccept() {
		return nil, ErrCantAccept
	}

	c, err := l.listener.Accept()
	if err != nil {
		return nil, err
	}

	return getWotoConnection(c, l, r), nil
}

//---------------------------------------------------------

func (c *WotoConnection) IsClosed() bool {
	return c.isClosed
}

func (c *WotoConnection) Close() {
	if !c.isClosed && c.conn != nil {
		c.isClosed = true
		c.conn.Close()
	}
}

func (c *WotoConnection) CanReadAndWrite() bool {
	return !c.isClosed && c.conn != nil
}

// GetOrigin returns the origin listener of this woto connection
// (the listener which accepted this connection).
func (c *WotoConnection) GetOrigin() *WotoListener {
	return c.origin
}

// ReadBytes will read the incoming bytes from the tcp
// connection. you should always use this method to read
// all receiving data from the client.
func (c *WotoConnection) ReadBytes() ([]byte, error) {
	if !c.CanReadAndWrite() {
		return nil, ErrCantReadOrWrite
	}
	b := make([]byte, MAX_FIRST_BYTES)

	// read the first 8 bytes from the incoming data to
	// understand the real length of the data.
	n, err := c.readAllBytes(b)
	if err != nil {
		return nil, err
	} else if n != MAX_FIRST_BYTES {
		return nil, ErrCantReadOrWrite
	}

	// make sure to use trim space function on the string,
	// because `strconv.Atoi` function is so stupid and doesn't
	// ignore white spaces at all.
	count, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if count < 0 {
		// will there be any bug like this?
		// maybe from client-side, yup it's possible.
		// anyway, since the count cannot be negative (less than 0),
		// we have to make sure that it's more than zero.
		count = -count
	}

	if err != nil {
		return nil, err
	} else if count == BaseIndex {
		return nil, ErrCantReadOrWrite
	}

	b = make([]byte, count)

	// now read the whole data received from the client.
	// the whole data's length is equal to `count`
	n, err = c.readAllBytes(b)
	if err != nil {
		return nil, err
	} else if n != count {
		return nil, ErrCantReadOrWrite
	}

	return b, nil
}

// ReadString will read the incoming bytes from the tcp
// connection and will return it as a string value.
// You should always use this method to read
// all receiving data from the client.
func (c *WotoConnection) ReadString() (string, error) {
	b, err := c.ReadBytes()
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (c *WotoConnection) readAllBytes(b []byte) (int, error) {
	return c.conn.Read(b)
}

// ReadJson will read the data from the connection and will decode it as
// a json data.
// it will return the error if there was an error in the middle of the way.
func (c *WotoConnection) ReadJson(v interface{}) error {
	if v == nil {
		return ErrValueNil
	}

	b, err := c.ReadBytes()
	if err != nil {
		return err
	}

	// byte array's length should be at the very least 2.
	if len(b) < 2 {
		return ErrValueEmpty
	}

	/*
		All cryptography operations should go here.
	*/

	return json.Unmarshal(b, v)
}

// WriteBytes will write the data to the `WotoConnection` using woto algorithm.
// at first, it will write 8 bytes as the total data length, and it will
// append the real data to those 8 bytes.
// client needs to read those 8 bytes at first and then read the data
// equal to the length of the number.
func (c *WotoConnection) WriteBytes(b []byte) (int, error) {
	if b == nil {
		// don't return an error in a case the `data` is nil.
		// ignore it.
		return BaseIndex, nil
	} else if !c.CanReadAndWrite() {
		return BaseIndex, ErrCantReadOrWrite
	}

	bb := []byte(MakeSureNum(len(b), MAX_FIRST_BYTES))
	bb = MakeSureByte(bb, MAX_FIRST_BYTES)

	return c.writeAllBytes(append(bb, b...))
}

func (c *WotoConnection) writeAllBytes(b []byte) (int, error) {
	return c.conn.Write(b)
}

func (c *WotoConnection) WriteJson(v interface{}) (int, error) {
	if v == nil {
		return BaseIndex, ErrValueNil
	} else if !c.CanReadAndWrite() {
		return BaseIndex, ErrCantReadOrWrite
	}

	b, err := json.Marshal(v)
	if err != nil {
		return BaseIndex, err
	}

	logging.Debug("what we are writing is ", string(b))

	return c.WriteBytes(b)
}

func (c *WotoConnection) SetDeadline(t time.Time) {
	if c.CanReadAndWrite() {
		c.conn.SetDeadline(t)
	}
}

func (c *WotoConnection) SetMe(user *UserInfo) {
	c.me = user
}

func (c *WotoConnection) GetMe() *UserInfo {
	return c.me
}

func (c *WotoConnection) IsRegistered() bool {
	return c.isRegistered
}

func (c *WotoConnection) SetRegisterer(r func(*WotoConnection)) {
	c.registerer = r
}

func (c *WotoConnection) Register() {
	if !c.isRegistered && c.registerer != nil {
		c.registerer(c)
		c.isRegistered = true
	}
}

func (c *WotoConnection) GetEntryKeys() *EntryKeys {
	return c.keys
}

func (c *WotoConnection) SetAsPastKey(key wotoCrypto.WotoKey) *EntryKeys {
	if c.keys == nil {
		c.keys = &EntryKeys{
			_pastKey: key,
		}
		return c.keys
	}

	c.keys._pastKey = key
	return c.keys
}

func (c *WotoConnection) SetAsPresentKey(key wotoCrypto.WotoKey) *EntryKeys {
	if c.keys == nil {
		c.keys = &EntryKeys{
			_presentKey: key,
		}
		return c.keys
	}

	c.keys._presentKey = key
	return c.keys
}

func (c *WotoConnection) SetAsFutureKey(key wotoCrypto.WotoKey) *EntryKeys {
	if c.keys == nil {
		c.keys = &EntryKeys{
			_futureKey: key,
		}
		return c.keys
	}

	c.keys._futureKey = key
	return c.keys
}

func (c *WotoConnection) SyncKeys() {
	if c.keys == nil {
		/* TODO: generate new series of key for the new connection */
		return
	}

	c.keys.Sync()
}

//---------------------------------------------------------

func (k *EntryKeys) Sync() {
	k.FutureKey = k._futureKey.StrSerialize()
	k.PresentKey = k._presentKey.StrSerialize()
	k.PastKey = k._pastKey.StrSerialize()
}

//---------------------------------------------------------
