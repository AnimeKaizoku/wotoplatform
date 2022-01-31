package wotoValues

import (
	"net"
	"strconv"

	_ "github.com/ALiwoto/StrongStringGo/strongStringGo"
)

// MakeSureNum will make sure that when you convert `i`
// to string, its length be the exact same as `count`.
// it will append 0 to the left side of the number to do so.
// for example:
// MakeSureNum(5, 8) will return "00000005"
func MakeSureNum(i, count int) string {
	s := strconv.Itoa(i)
	final := count - len(s)
	for ; final > BaseIndex; final-- {
		s = BaseIndexStr + s
	}

	return s
}

// MakeSureByte is a useful function. consider b is a number in string
// and you are sending it as a byte to this function.
// this function then, will ensure that the length of this byte array
// is exactly equal to the passed-by argument.
// for example:
// MakeSureByte([]byte("5"), 8) will return []byte("5       ")
// the returned value's length will be exactly the same as length.
func MakeSureByte(b []byte, length int) []byte {
	original := len(b)

	if original == length {
		return b
	} else if original > length {
		return b[original-length:]
	} else { // original < length
		var rb = make([]byte, len(b))
		copy(rb, b)
		for i := 0; i < length-original; i++ {
			rb = append(rb, SpaceChar)
		}

		return rb
	}
}

func getWotoConnection(conn net.Conn, l *WotoListener, r Registerer) *WotoConnection {
	return &WotoConnection{
		conn:       conn,
		origin:     l,
		registerer: r,
	}
}

func GetWotoListener(l net.Listener) *WotoListener {
	if l == nil {
		return nil
	}

	return &WotoListener{
		listener: l,
	}
}
