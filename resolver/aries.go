package resolver

import (
	"time"

	"github.com/allinbits/cosmos-cash-resolver/x/did/types"
)

// AriesDidDocument implement the aries did document
// see: https://github.com/hyperledger/aries-framework-go/blob/2d601d717a3e2297662faf622f7a878e0f50f981/pkg/doc/did/doc.go#L315
type AriesDidDocument struct {
	types.DidDocument
	Created *time.Time `json:"created,omitempty"`
	Updated *time.Time `json:"updated,omitempty"`
}
