package midtrans

import (
	"context"
)

type Midtrans interface {
	ChargeSnap(ctx context.Context, req *SnapRequest) (resp SnapResponse, err error)
}
