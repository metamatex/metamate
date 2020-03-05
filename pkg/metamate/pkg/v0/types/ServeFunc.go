package types

import (
	"context"
	"github.com/metamatex/metamatemono/pkg/generic/pkg/v0/generic"
	"github.com/metamatex/metamatemono/gen/v0/sdk"
)

type ServeFunc func(ctx context.Context, gCliReq generic.Generic) (gCliRsp generic.Generic)

type ResolveFunc func(ctx context.Context, doCliReqProcessing bool, doCliReqValidation bool, doSetClientAccount bool, acc *sdk.ClientAccount, gCliReq generic.Generic) (gCliRsp generic.Generic)