package resolver

import (
	"context"

	"github.com/allinbits/cosmos-cash-resolver/x/did/types"
	"google.golang.org/grpc"
)

// ResolveRepresentation resolve a did document with a specific representation
func ResolveRepresentation(cc *grpc.ClientConn, did string, opts ResolutionOption) (drr DidResolutionReply) {

	// fail if it is not found
	qc := types.NewQueryClient(cc)
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
