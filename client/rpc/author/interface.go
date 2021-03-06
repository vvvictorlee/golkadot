package author

import (
	"github.com/opennetsys/golkadot/client/p2p/pubsub"
	rpctypes "github.com/opennetsys/golkadot/client/rpc/types"
)

// ServiceInterface ...
type ServiceInterface interface {
	// SubmitExtrinsic submits a hex-encoded extrinsic for inclusion in block.
	SubmitExtrinsic(extrinsic []byte, response *string) error
	// PendingExtrinsics returns all pending extrinsics, potentially grouped by sender.
	PendingExtrinsics(args rpctypes.NilArgs, response [][]byte) error
	// SubmitAndWatchExtrinsic submits an extrinsic to watch.
	SubmitAndWatchExtrinsic(args *SubmitAndWatchExtrinsicArgs, response rpctypes.NilResponse) error
	// UnwatchExtrinsic unsubscribes from extrinsic watching.
	UnwatchExtrinsic(id pubsub.SubscriptionID, response *bool) error
}
