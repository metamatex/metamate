package types

import (
	"context"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
)

type ServeFunc func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic)

type ResolveFunc func(ctx context.Context, doCliReqProcessing bool, doCliReqValidation bool, doSetClientAccount bool, gCliReq generic.Generic) (gCliRsp generic.Generic)
