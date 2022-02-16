package wotoCrypto

import "errors"

var (
	layerKindsMap = map[CryptoLayerKind]bool{
		CryptoLayerKindO27:  true,
		CryptoLayerKindO54:  true,
		CryptoLayerKindO108: true,
		CryptoLayerKindO216: true,
	}
	layerLengthValidator = map[CryptoLayerKind]func(WotoLayerLength) bool{
		CryptoLayerKindO27:  func(wll WotoLayerLength) bool { return true },
		CryptoLayerKindO54:  func(wll WotoLayerLength) bool { return true },
		CryptoLayerKindO108: func(wll WotoLayerLength) bool { return true },
		CryptoLayerKindO216: func(wll WotoLayerLength) bool { return true },
	}
)

var (
	_blockActionsMap = map[blockAlgorithmId]blockAlgorithm{
		blockAlgorithmIdX917: &blockAlgorithmX917{0x02},
		blockAlgorithmIdX847: &blockAlgorithmX847{0x04},
		blockAlgorithmIdX795: &blockAlgorithmX795{0x09},
		blockAlgorithmIdX649: &blockAlgorithmX649{0x014},
	}
)

var (
	ErrInvalidKey = errors.New("invalid key")
)
