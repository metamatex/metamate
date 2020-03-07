package types

type Transformer interface {
	Name() string
	Transform(ctx ReqCtx) ReqCtx
}
