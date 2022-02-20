package wotoActions

import wcr "github.com/TheGolangHub/wotoCrypto/wotoCrypto"

func (r *ActionResp) SetAsKeys(v wcr.KeyCollection) {
	r.Keys = v
}
