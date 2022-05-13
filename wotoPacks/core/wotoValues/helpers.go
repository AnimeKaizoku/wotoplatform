package wotoValues

import (
	"crypto/rand"
	"crypto/sha512"
	"net"
	"strconv"

	wcr "github.com/TheGolangHub/wotoCrypto/wotoCrypto"
)

// InitKeys will initialize the internal keys.
// skipcq
func InitKeys(cryptoEnabled bool) error {
	if !cryptoEnabled {
		encryptionEnabled = cryptoEnabled
		return nil
	}

	for keyIndex, current := range _initialKeys {
		if current == nil {
			continue
		}

		for _, currentLayer := range keyLayers[keyIndex] {
			current.AppendLayer(&currentLayer)
		}

		current.SetSignatureByBytes(keySigns[keyIndex])

		if current.IsRealLengthInvalid() {
			return ErrInvalidRealLength
		}
	}

	return nil
}

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

func getInitialWotoKey(index int) wcr.WotoKey {
	return _initialKeys[index]
}

// newEntryKeys returns a new instance of EntryKeys struct.
// Do notice that this function doesn't call `Sync` method on the new instance
// of the EntryKeys, you need to do this yourself after getting the new instance.
func newEntryKeys() *EntryKeys {
	eKeys := &EntryKeys{
		_pastKey:    getPastFreshKey(),
		_presentKey: getPresentFreshKey(),
		_futureKey:  getFutureFreshKey(),
	}

	return eKeys
}

func getNewLayerO27(key wcr.WotoKey) *wcr.CryptoLayer {
	b := make([]byte, key.GetSignatureRealLength())
	_, _ = rand.Read(b)
	return &wcr.CryptoLayer{
		Kind: wcr.CryptoLayerKindO27,
		Hash: string(b),
	}
}

func getPastFreshKey() wcr.WotoKey {
	key := wcr.GeneratePresentKey(wcr.WotoAlgorithmM250).ToPastKey()
	key.SetSignatureByFunc(sha512.New)
	key.AppendLayer(getNewLayerO27(key))
	key.AppendLayer(getNewLayerO27(key))
	key.AppendLayer(getNewLayerO27(key))
	return key
}
func getPresentFreshKey() wcr.WotoKey {
	key := wcr.GeneratePresentKey(wcr.WotoAlgorithmM250)
	key.SetSignatureByFunc(sha512.New)
	key.AppendLayer(getNewLayerO27(key))
	key.AppendLayer(getNewLayerO27(key))
	key.AppendLayer(getNewLayerO27(key))
	return key
}
func getFutureFreshKey() wcr.WotoKey {
	key := wcr.GenerateFutureKey(wcr.GeneratePresentKey(wcr.WotoAlgorithmM250).ToPastKey())
	key.SetSignatureByFunc(sha512.New)
	key.AppendLayer(getNewLayerO27(key))
	key.AppendLayer(getNewLayerO27(key))
	key.AppendLayer(getNewLayerO27(key))
	return key
}
