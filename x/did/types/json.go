package types

import "encoding/json"

// MarshalJSON implements a custom marshaller for rendergin verification material
func (vm VerificationMethod) MarshalJSON() ([]byte, error) {
	vmd := make(map[string]string, 4)
	vmd["id"] = vm.Id
	vmd["controller"] = vm.Controller
	vmd["type"] = vm.Type
	switch m := vm.VerificationMaterial.(type) {
	case *VerificationMethod_BlockchainAccountID:
		vmd["blockchainAccountId"] = m.BlockchainAccountID
	case *VerificationMethod_PublicKeyHex:
		vmd["publicKeyHex"] = m.PublicKeyHex
	case *VerificationMethod_PublicKeyMyltibase:
		vmd["publicKeyMultibase"] = m.PublicKeyMyltibase
	}
	return json.Marshal(vmd)
}
