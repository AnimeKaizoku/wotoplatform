package wotoCrypto

const (
	blockAlgorithmIdX917 = 2 << iota
	blockAlgorithmIdX847
	blockAlgorithmIdX795
	blockAlgorithmIdX649
)

const (
	// See https://www.jstatsoft.org/article/view/v008i14/xorshift.pdf
	WotoAlgorithmM250 WotoAlgorithm = 2 << ^uint8(250)
	WotoAlgorithmM251 WotoAlgorithm = 3 << ^uint8(251)
	WotoAlgorithmM252 WotoAlgorithm = 5 << ^uint8(252)
	WotoAlgorithmM253 WotoAlgorithm = 8 << ^uint8(253)
)

const (
	CryptoLayerKindO27 CryptoLayerKind = 27 << iota
	CryptoLayerKindO54
	CryptoLayerKindO108
	CryptoLayerKindO216
)
