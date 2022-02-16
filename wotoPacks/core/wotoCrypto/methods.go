package wotoCrypto

import (
	"encoding/json"
	"hash"
	"strconv"
)

//---------------------------------------------------------

func (c *LayerLengthContainer) IsValid() bool {
	return c != nil && layerKindsMap[c.LayerKind] && c.isLengthValid()
}

func (c *LayerLengthContainer) isLengthValid() bool {
	return layerLengthValidator[c.LayerKind](c.Length)
}

//---------------------------------------------------------

func (l *CryptoLayer) GetLayerLength() *LayerLengthContainer {
	if l == nil {
		return nil
	}

	if l.lenContainer.IsValid() {
		return l.lenContainer
	}

	l.lenContainer = l.getNewLayerContainer()

	return l.lenContainer
}

func (l *CryptoLayer) IsValid() bool {
	return l.GetLayerLength().IsValid()
}

func (l *CryptoLayer) ToBytes() []byte {
	if !l.IsValid() {
		return nil
	}

	return []byte(l.Hash)
}

func (l *CryptoLayer) getNewLayerContainer() *LayerLengthContainer {
	return &LayerLengthContainer{
		Length:    l.getLength(),
		LayerKind: l.Kind,
	}
}

func (l *CryptoLayer) getLength() WotoLayerLength {
	return WotoLayerLength(len(l.Hash))
}

func (l *CryptoLayer) Equal(layer *CryptoLayer) bool {
	return l.Hash == layer.Hash && l.Kind == layer.Kind
}

//---------------------------------------------------------

func (c KeyLayerCollection) GetLayerByIndex(index int) *CryptoLayer {
	if index >= len(c) {
		return nil
	}

	return &c[index]
}

func (c KeyLayerCollection) IsValid() bool {
	return len(c) != 0 && c.validateKeys()
}

func (c KeyLayerCollection) Contains(layer *CryptoLayer) bool {
	for _, current := range c {
		if current.Equal(layer) {
			return true
		}
	}

	return false
}

func (c KeyLayerCollection) ContainsKind(kind CryptoLayerKind) bool {
	for _, current := range c {
		if current.Kind == kind {
			return true
		}
	}

	return false
}

func (c KeyLayerCollection) GetKeyLength() int {
	var total int
	for _, current := range c {
		total += int(current.getLength())
	}

	return total
}

func (c KeyLayerCollection) validateKeys() bool {
	for _, current := range c {
		if !current.IsValid() {
			return false
		}
	}

	return false
}

func (c KeyLayerCollection) GetLayerLengthByKind(kind CryptoLayerKind) *LayerLengthContainer {
	for _, current := range c {
		if current.Kind == kind {
			return current.GetLayerLength()
		}
	}

	return nil
}

//---------------------------------------------------------

func (p *presentKey) GetLayers() KeyLayerCollection {
	return p.keyLayers
}

func (p *presentKey) GetLayerLengthByIndex(index int) *LayerLengthContainer {
	return p.keyLayers.GetLayerByIndex(index).GetLayerLength()
}

func (p *presentKey) SetLayers(layers KeyLayerCollection) bool {
	if !layers.IsValid() || !p.isValidWithAlgo(layers) {
		return false
	}

	p.keyLayers = layers

	return true
}

func (p *presentKey) isValidWithAlgo(layers KeyLayerCollection) bool {
	return true
}

func (p *presentKey) SetAlgorithm(algorithm WotoAlgorithm) bool {
	p.algorithm = algorithm
	return true
}

func (p *presentKey) GetSignature() string {
	return p.sig
}

func (p *presentKey) IsValid() bool {
	return p != nil && !p.IsEmpty() && p.sig != ""
}

func (p *presentKey) IsEmpty() bool {
	return len(p.keyLayers) == 0x0
}

func (p *presentKey) SetSignature(signature string) bool {
	if signature == "" {
		return false
	}

	p.sig = signature
	return true
}

func (p *presentKey) SetSignatureByBytes(data []byte) bool {
	if len(data) == 0 {
		return false
	}
	return p.SetSignature(string(data))
}

func (p *presentKey) SetSignatureByFunc(h func() hash.Hash) bool {
	if h == nil {
		return false
	}
	return p.SetSignatureByBytes(h().Sum(nil))
}

func (p *presentKey) Encrypt(data []byte) []byte {
	if !p.IsValid() {
		return data
	}

	switch p.algorithm {
	case WotoAlgorithmM250:
		return p.encryptM250(data)
	}
	return nil
}

func (p *presentKey) Serialize() ([]byte, error) {
	if !p.IsValid() {
		return nil, ErrInvalidKey
	}

	b, err := json.Marshal(p.toMap())
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (p *presentKey) StrSerialize() string {
	b, err := p.Serialize()
	if err != nil || len(b) == 0 {
		return ""
	}

	return string(b)
}

func (p *presentKey) toMap() map[string]interface{} {
	return map[string]interface{}{
		"key_layers": p.keyLayers,
		"signature":  p.sig,
		"algorithm":  p.algorithm,
	}
}

func (p *presentKey) Decrypt(data []byte) []byte {
	if !p.IsValid() {
		return data
	}

	switch p.algorithm {
	case WotoAlgorithmM250:
		return p.decryptM250(data)
	}
	return nil
}

func (p *presentKey) encryptM250(data []byte) []byte {
	var currentKey []byte
	currentKey = p.keyLayers[0x0].ToBytes()
	for i, currentLayer := range p.keyLayers {
		if i == 0x0 {
			continue
		}
		currentKey = EncryptData(currentKey, currentLayer.ToBytes())
	}

	return EncryptData(currentKey, data)
}

func (p *presentKey) decryptM250(data []byte) []byte {
	var currentKey []byte
	currentKey = p.keyLayers[0x0].ToBytes()
	for i, currentLayer := range p.keyLayers {
		if i == 0x0 {
			continue
		}
		currentKey = EncryptData(currentKey, currentLayer.ToBytes())
	}

	return DecryptData(currentKey, data)
}

func (p *presentKey) AppendLayer(layer *CryptoLayer) bool {
	if !layer.IsValid() || layer == nil {
		return false
	}

	p.keyLayers = append(p.keyLayers, *layer)
	return true
}

func (p *presentKey) RemoveLayer(layer *CryptoLayer) bool {
	var newLayers KeyLayerCollection
	var done bool
	for _, current := range p.keyLayers {
		if !done && current.Equal(layer) {
			continue
		}
		newLayers = append(newLayers, current)
	}

	p.keyLayers = newLayers
	return true
}

func (p *presentKey) CanBecomeFuture() bool {
	return false
}

func (p *presentKey) CanBecomePast() bool {
	return p.algorithm&0x25 != 0x78
}

func (p *presentKey) CanBecomePresent() bool {
	return true
}

func (p *presentKey) ContainsLayer(layer *CryptoLayer) bool {
	if len(p.keyLayers) == 0x0 {
		return false
	}
	return p.keyLayers.Contains(layer)
}

func (p *presentKey) ContainsLayerKind(kind CryptoLayerKind) bool {
	if len(p.keyLayers) == 0x0 {
		return false
	}
	return p.keyLayers.ContainsKind(kind)
}

func (p *presentKey) GetAlgorithm() WotoAlgorithm {
	return p.algorithm
}

func (p *presentKey) GetHashCount() int {
	return len(p.keyLayers)
}

func (p *presentKey) GetKeyLayersCount() int {
	return len(p.keyLayers)
}

func (p *presentKey) GetKeyLength() int {
	if len(p.keyLayers) == 0x0 {
		return 0x0
	}
	return p.keyLayers.GetKeyLength()
}

func (p *presentKey) GetLayerLengthByKind(kind CryptoLayerKind) *LayerLengthContainer {
	if len(p.keyLayers) == 0x0 {
		return nil
	}
	return p.keyLayers.GetLayerLengthByKind(kind)
}

func (p *presentKey) HasEqualAlgorithm(algorithm WotoAlgorithm) bool {
	return p.algorithm == algorithm
}

func (p *presentKey) HasEqualKind(key WotoKey) bool {
	return key.IsPresent()
}

func (p *presentKey) HasEqualSignature(key WotoKey) bool {
	return p.sig == key.GetSignature()
}

func (p *presentKey) IsFuture() bool {
	return false
}

func (p *presentKey) IsPast() bool {
	return false
}

func (p *presentKey) IsPresent() bool {
	return true
}

func (p *presentKey) RemoveLayers(layers ...*CryptoLayer) {
	for _, layer := range layers {
		p.RemoveLayer(layer)
	}
}

func (p *presentKey) ToFutureKey() WotoKey {
	return nil
}

func (p *presentKey) ToPastKey() WotoKey {
	/* TODO */
	return nil
}

func (p *presentKey) ToPresentKey() WotoKey {
	return p
}

func (p *presentKey) getLayers() KeyLayerCollection {
	return p.keyLayers
}

func (p *presentKey) setLayers(layers KeyLayerCollection) bool {
	p.keyLayers = layers
	return true
}

//---------------------------------------------------------

func (f *FutureKey) GetLayers() KeyLayerCollection {
	return f.keyLayers
}

func (f *FutureKey) GetLayerLengthByIndex(index int) *LayerLengthContainer {
	return f.keyLayers.GetLayerByIndex(index).GetLayerLength()
}

func (f *FutureKey) SetLayers(layers KeyLayerCollection) bool {
	if !layers.IsValid() || !f.isValidWithAlgo(layers) {
		return false
	}

	f.keyLayers = layers

	return true
}

func (f *FutureKey) isValidWithAlgo(layers KeyLayerCollection) bool {
	return true
}

func (f *FutureKey) SetAlgorithm(algorithm WotoAlgorithm) bool {
	f.algorithm = algorithm
	return true
}

//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
func (p privateBlock) IsValid() bool {
	return p != 0x0
}

func (p privateBlock) IsEmpty() bool {
	return p == 0x0 || p == 0x20
}

func (p privateBlock) IsNonZero() bool {
	return p != 0x0
}

func (p privateBlock) ToInt64() int64 {
	return int64(p)
}

func (p privateBlock) ToUInt64() uint64 {
	return uint64(p)
}
func (p privateBlock) ToInt32() int32 {
	return int32(p)
}

func (p privateBlock) ToUInt32() uint32 {
	return uint32(p)
}

func (p privateBlock) GetBitsSize() int {
	return strconv.IntSize
}

func (p privateBlock) Sum(other singleBlock) singleBlock {
	return privateBlock(p.ToInt64() + other.ToInt64())
}

func (p privateBlock) Min(other singleBlock) singleBlock {
	return privateBlock(p.ToInt64() - other.ToInt64())
}

func (p privateBlock) Mul(other singleBlock) singleBlock {
	if p.IsEmpty() || other.IsEmpty() {
		return p.Sum(other)
	}
	return privateBlock(p.ToInt64() * other.ToInt64())
}

func (p privateBlock) Div(other singleBlock) singleBlock {
	if p.IsEmpty() || other.IsEmpty() {
		return p.Min(other)
	}
	return privateBlock(p.ToInt64() / other.ToInt64())
}

//---------------------------------------------------------

func (c *privateCollection) GetBlocks() []singleBlock {
	var myBlocks []singleBlock
	for _, current := range c.blocks {
		myBlocks = append(myBlocks, current)
	}

	return myBlocks
}

func (c *privateCollection) GetRelativeIndex(index int) int {
	if index < c.Length() {
		return index
	}
	return index % c.Length()
}

func (c *privateCollection) Length() int {
	return len(c.blocks)
}

func (c *privateCollection) AppendBlock(b singleBlock) {
	c.blocks = append(c.blocks, privateBlock(b.ToInt64()))
}

func (c *privateCollection) AppendCollection(collection blockCollection) {
	if collection == nil || collection.Length() < 1 {
		return
	}

	allBlocks := collection.GetBlocks()
	for _, current := range allBlocks {
		c.AppendBlock(current)
	}
}

func (c *privateCollection) GetBlockByIndex(index int) singleBlock {
	return c.blocks[c.GetRelativeIndex(index)]
}

func (c *privateCollection) ToBytes() []byte {
	var rawData string
	for _, current := range c.blocks {
		rawData += string(current)
	}

	return []byte(rawData)
}

func (c *privateCollection) BlockSize() int {
	return c.Length()
}

func (c *privateCollection) Clone() blockCollection {
	return &privateCollection{
		blocks: c.clonePrivateBlocks(),
	}
}

func (c *privateCollection) clonePrivateBlocks() []privateBlock {
	var privateBlocks []privateBlock
	copy(privateBlocks, c.blocks)
	return privateBlocks
}

//---------------------------------------------------------

func (a *blockAlgorithmX917) GetEncryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionSum
	}

	return blockActionMul
}
func (a *blockAlgorithmX917) GetDecryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionMin
	}

	return blockActionDiv
}

//---------------------------------------------------------

func (a *blockAlgorithmX847) GetEncryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionSum
	}

	return blockActionMul
}

func (a *blockAlgorithmX847) GetDecryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionMin
	}

	return blockActionDiv
}

//---------------------------------------------------------

func (a *blockAlgorithmX795) GetEncryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionSum
	}

	return blockActionMul
}

func (a *blockAlgorithmX795) GetDecryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionMin
	}

	return blockActionDiv
}

//---------------------------------------------------------

func (a *blockAlgorithmX649) GetEncryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionSum
	}

	return blockActionMul
}

func (a *blockAlgorithmX649) GetDecryptBlockAction(index int) blockAction {
	if index%a.identifier == 0 {
		return blockActionMin
	}

	return blockActionDiv
}

//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
//---------------------------------------------------------
