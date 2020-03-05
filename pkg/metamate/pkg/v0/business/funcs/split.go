package funcs

import (
	"fmt"
	"github.com/metamatex/metamatemono/pkg/metamate/pkg/v0/types"
)

func By(subject string) func (ctx types.ReqCtx) (string) {
	switch subject {
	case types.Method:
		return func(ctx types.ReqCtx) string {
			return ctx.Method
		}
	case types.Mode:
		return func(ctx types.ReqCtx) string {
			return ctx.Mode
		}
	}

	panic(fmt.Sprintf("subject %v not supported", subject))
}