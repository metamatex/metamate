package funcs

import (
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

func Isset(subject string, path []string, b bool) func(ctx types.ReqCtx) bool {
	if len(path) == 0 {
		switch subject {
		case types.GSvcRsp:
			return func(ctx types.ReqCtx) bool {
				return (ctx.GSvcRsp != nil) == b
			}
		case types.GBusRsp:
			return func(ctx types.ReqCtx) bool {
				return (ctx.GBusRsp != nil) == b
			}
		}

		panic(fmt.Sprintf("funcs.Isset subject %v not supported", subject))
	}

	switch subject {
	case types.GSvcRsp:
		return func(ctx types.ReqCtx) bool {
			_, ok := ctx.GSvcRsp.Generic(path...)
			return ok == b
		}
	case types.GBusRsp:
		return func(ctx types.ReqCtx) bool {
			_, ok := ctx.GBusRsp.Generic(path...)
			return ok == b
		}
	case types.GCliReq:
		return func(ctx types.ReqCtx) bool {
			_, ok := ctx.GCliReq.Generic(path...)
			return ok == b
		}
	}

	panic(fmt.Sprintf("funcs.Isset subject %v not supported", subject))
}

func Is(subject string, b bool) func(ctx types.ReqCtx) bool {
	switch subject {
	case types.DoCliReqProcessing:
		return func(ctx types.ReqCtx) bool {
			return ctx.DoCliReqProcessing == b
		}
	case types.DoCliReqValidation:
		return func(ctx types.ReqCtx) bool {
			return ctx.DoCliReqValidation == b
		}
	}

	panic(fmt.Sprintf("funcs.Is: subject %v not supported", subject))
}

func IsType(subject string, name string, b bool) func(ctx types.ReqCtx) bool {
	switch subject {
	case types.GSvcRsp:
		return func(ctx types.ReqCtx) bool {
			return (ctx.GSvcRsp.Type().Name() == name) == b
		}
	}

	panic(fmt.Sprintf("subject %v not supported", subject))
}

func EntityOnlyContainsServiceId(ctx types.ReqCtx) bool  {
	m := ctx.GEntity.ToStringInterfaceMap()

	_, ok := m[fieldnames.Id]

	return ok && len(m) == 1
}