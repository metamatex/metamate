package types

type FuncTransformer struct {
	Name0 string
	Func  func(ctx ReqCtx) ReqCtx
}

func (t FuncTransformer) Name() string {
	return t.Name0
}

func (t FuncTransformer) Transform(ctx ReqCtx) ReqCtx {
	return t.Func(ctx)
}
