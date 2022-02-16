package wotoCrypto

import (
	"crypto/aes"
	"strings"
)

func GenerateFutureKey(pastKey WotoKey) WotoKey {
	return nil
}

func GeneratePresentKey(algo WotoAlgorithm) WotoKey {
	return &presentKey{
		algorithm: algo,
	}
}

func EncryptData(key, data []byte) []byte {
	keyBlocks := toBlockCollection(key)
	dataBlocks := toBlockCollection(data)
	return encryptByBlocks(keyBlocks, dataBlocks, nil)
}

func DecryptData(key, data []byte) []byte {
	keyBlocks := toBlockCollection(key)
	dataBlocks := toBlockCollection(data)
	return decryptByBlocks(keyBlocks, dataBlocks, nil)
}

// encryptByBlocksNoAlgorithm encrypts the data with the specified key to the destinated
// block collection without using any algorithm.
func encryptByBlocksNoAlgorithm(dest, key, data blockCollection) []byte {
	for index, current := range data.GetBlocks() {
		dest.AppendBlock(current.Sum(key.GetBlockByIndex(index)))
	}

	return dest.ToBytes()
}

// encryptByBlocksNoAlgorithm encrypts the data with the specified key to the destinated
// block collection without using any algorithm.
func decryptByBlocksNoAlgorithm(dest, key, data blockCollection) []byte {
	for index, current := range data.GetBlocks() {
		dest.AppendBlock(current.Min(key.GetBlockByIndex(index)))
	}

	return dest.ToBytes()
}

func encryptByBlocks(key, data blockCollection, algorithm blockAlgorithm) []byte {
	finalCollection := getEmptyCollection()
	if algorithm == nil {
		return encryptByBlocksNoAlgorithm(finalCollection, key, data)
	}

	var action blockAction
	var actionFactor bool
	for index, current := range data.GetBlocks() {
		action = algorithm.GetEncryptBlockAction(index)
		if action == nil {
			continue
		}

		if !actionFactor {
			actionFactor = true
		}

		finalCollection.AppendBlock(action(current, key.GetBlockByIndex(index)))
	}

	if !actionFactor {
		return encryptByBlocksNoAlgorithm(finalCollection, key, data)
	}

	return finalCollection.ToBytes()
}

func decryptByBlocks(key, data blockCollection, algorithm blockAlgorithm) []byte {
	finalCollection := getEmptyCollection()
	if algorithm == nil {
		return decryptByBlocksNoAlgorithm(finalCollection, key, data)
	}

	var action blockAction
	var actionFactor bool
	for index, current := range data.GetBlocks() {
		action = algorithm.GetDecryptBlockAction(index)
		if action == nil {
			continue
		}

		finalCollection.AppendBlock(action(current, key.GetBlockByIndex(index)))
	}

	if !actionFactor {
		return decryptByBlocksNoAlgorithm(finalCollection, key, data)
	}

	return finalCollection.ToBytes()
}

func EncryptAES(key, data []byte) []byte {
	if len(data) < aes.BlockSize {
		l := len(data)
		for i := 0; i < aes.BlockSize-l; i++ {
			data = append(data, 0x20)
		}
	}

	b, err := aes.NewCipher(key)
	if err != nil {
		return data
	}

	dest := make([]byte, len(data))
	b.Encrypt(dest, data)
	return dest
}

func BlockAlgorithmExists(algorithmId uint8) bool {
	return _blockActionsMap[blockAlgorithmId(algorithmId)] != nil
}

func DecryptAES(key, data []byte) []byte {
	b, err := aes.NewCipher(key)
	if err != nil {
		return data
	}
	dest := make([]byte, len(data))
	b.Decrypt(dest, data)
	return []byte(strings.TrimSpace(string(dest)))
}

func toBlockCollection(data []byte) blockCollection {
	return &privateCollection{
		blocks: toSingleBlocks(data),
	}
}

func toSingleBlocks(data []byte) []privateBlock {
	var finalValue []privateBlock
	for _, current := range string(data) {
		finalValue = append(finalValue, privateBlock(current))
	}

	return finalValue
}

func getEmptyCollection() blockCollection {
	return &privateCollection{}
}

func blockActionSum(first, second singleBlock) singleBlock {
	return first.Sum(second)
}

func blockActionMin(first, second singleBlock) singleBlock {
	return first.Min(second)
}

func blockActionMul(first, second singleBlock) singleBlock {
	return first.Mul(second)
}

func blockActionDiv(first, second singleBlock) singleBlock {
	return first.Div(second)
}
