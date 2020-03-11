package line

import (
	"errors"
	"github.com/davecgh/go-spew/spew"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"testing"
)

func TestError(t *testing.T) {

	errL := Do()
	errL.Do(
		Func(func(ctx types.ReqCtx) types.ReqCtx {
			spew.Dump(ctx)

			return ctx
		}),
	)

	l := Do()
	l.
		Error(errL, true).
		If(func(ctx types.ReqCtx) bool {
			return true
		}, Do()).
		Switch(func(ctx types.ReqCtx) string {
			return "a"
		}, map[string]*Line{
			"a": Switch(func(ctx types.ReqCtx) string {
				return "b"
			}, map[string]*Line{
				"b": Do(Err()),
			}),
		})

	l.Transform(types.ReqCtx{})
	//
	//println(l.Draw())
	//
	//spew.Dump(l)

	//l0 :=
	//	Error(func(ctx types.ReqCtx) (types.ReqCtx, bool) {
	//		spew.Dump(ctx.Err)
	//
	//		return ctx, false
	//	}).
	//		Do(Func(func(ctx types.ReqCtx) types.ReqCtx {
	//		ctx.Err = errors.New("l0")
	//
	//		return ctx
	//	}))
	//
	//l0.Transform(types.ReqCtx{})
}

func Func(f func(ctx types.ReqCtx) types.ReqCtx) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "",
		Func:  f,
	}
}

func Err() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "",
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.Err = errors.New("l")

			return ctx
		},
	}
}
