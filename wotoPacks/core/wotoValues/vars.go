package wotoValues

import (
	"errors"
	"sync"

	wcr "github.com/TheGolangHub/wotoCrypto/wotoCrypto"
	"gorm.io/gorm"
)

//---------------------------------------------------------

var ErrCantAccept = errors.New("woto listener: can't accept any new connections")
var ErrCantReadOrWrite = errors.New("woto connection: can't read or write from this connection")
var ErrCouldNotWriteFirstBytes = errors.New("woto connection: couldn't write the first bytes")
var ErrValueNil = errors.New("woto connection: interface value cannot be nil")
var ErrValueEmpty = errors.New("woto connection: received value was empty")
var ErrNotContainer = errors.New("woto connection: received value is not a key container")
var ErrInvalidRealLength = errors.New("woto connection: wotokey's real length is invalid")

//---------------------------------------------------------

var (
	SESSION      *gorm.DB
	SessionMutex = &sync.Mutex{}
)

var (
	_initialKeys = map[int]wcr.WotoKey{
		0x0: wcr.GeneratePresentKey(wcr.WotoAlgorithmM250),
		0x1: wcr.GeneratePresentKey(wcr.WotoAlgorithmM250),
		0x2: wcr.GeneratePresentKey(wcr.WotoAlgorithmM250),
		0x3: wcr.GeneratePresentKey(wcr.WotoAlgorithmM250),
		0x4: wcr.GeneratePresentKey(wcr.WotoAlgorithmM250),
	}
)

var (
	keySigns = [][]byte{
		{0x0, 0x0d, 0x04d, 0x0, 0x0, 0x08d, 0x03a, 0x02f},
		{0x0, 0x0f, 0x09d, 0x0f, 0x00, 0x0f2, 0x0d5, 0x0e3, 0x00, 0x00, 0x0d8, 0x7d, 0x00, 0x00,
			0x00, 0x45, 0x2d, 0x00, 0x23, 0x098, 0x09a, 0x06d},
		{0x0, 0x0b, 0x03b, 0x2a, 0x04c, 0x0, 0x00},
		{0x2e, 0x00, 0x03e, 0x1d, 0x00d, 0x0, 0x00},
		{0x9a, 0x09d, 0x07e, 0x9a, 0x00, 0x0d3, 0x00, 0x0c3, 0x35, 0x2d, 0x2d, 0xd9, 0x0, 0x4d,
			0xcd, 0x056, 0x087, 0x091, 0x00, 0x77, 0x4d},
	}
	keyLayers = [][]wcr.CryptoLayer{
		{ /* key1 */
			/* layers: */
			{ /* layer1 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer2 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer3 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer4 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
		},
		{ /* key2 */
			{ /* layer1 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer2 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer3 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
		},
		{ /* key3 */
			{ /* layer1 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer2 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer3 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
		},
		{ /* key4 */
			{ /* layer1 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer2 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer3 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
		},
		{ /* key5 */
			{ /* layer1 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer2 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
			{ /* layer3 */
				Kind: wcr.CryptoLayerKindO27,
				Hash: "abcdef",
			},
		},
	}
)
