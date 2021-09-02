package resolver

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/allinbits/cosmos-cash-resolver/x/did/types"
)

// ResolveRepresentation resolve a did document with a specific representation
func ResolveRepresentation(ctx client.Context, did string, opts ResolutionOption) (drr DidResolutionReply) {

	// fail if it is not found
	qc := types.NewQueryClient(ctx)
	qr, err := qc.DidDocument(context.Background(), &types.QueryDidDocumentRequest{Id: did})
	if err != nil {
		drr.ResolutionMetadata = ResolutionErr(ResolutionNotFound)
		return
	}
	// build the resolution
	drr.Document = qr.DidDocument
	drr.Metadata = qr.DidMetadata
	return
}

// MarshalJSON implements a custom marshaller for rendergin verification material
func (vm types.VerificationMethod) MarshalJSON() ([]byte, error) {
	vmd := make(map[string]string, 4)
	vmd["id"] = vm.Id
	vmd["controller"] = vm.Controller
	vmd["type"] = vm.Type
	switch m := vm.VerificationMaterial.(type) {
	case *VerificationMethod_BlockchainAccountID:
		vmd["blockchainAccountId"] = m.BlockchainAccountID
	case *VerificationMethod_PublicKeyHex:
		vmd["publicKeyHex"] = m.PublicKeyHex
	}
	return json.Marshal(vmd)
}
