package types

import (
	"context"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"gopkg.in/yaml.v2"
	"regexp"
)

const (
	Method             = "Method"
	Mode               = "Mode"
	Stage              = "Stage"
	Ctx                = "Ctx"
	GCliReq            = "GCliReq"
	GBusRsp            = "GBusRsp"
	GCliRsps           = "GCliRsps"
	ForTypeNode        = "ForTypeNode"
	GBusReq            = "GBusReq"
	GSvcRsp            = "GSvcRsp"
	SvcFilter          = "SvcFilter"
	GRspSelect         = "GRspSelect"
	EndpointNode       = "EndpointNode"
	Svcs               = "Svcs"
	Svc                = "Svc"
	Errs               = "Errs"
	GEntity            = "GEntity"
	GEntities          = "GEntities"
	SvcId              = "SvcId"
	SvcIds             = "SvcIds"
	ClientAccount      = "ClientAccount"
	DoCliReqValidation = "DoCliReqValidation"
	DoCliReqProcessing = "DoCliReqProcessing"
)

type ReqCtx struct {
	Id           string
	Stage        string
	Method       string
	Mode         string
	Ctx          context.Context
	GCliReq      generic.Generic
	GBusRsp      generic.Generic
	GCliRsps     []generic.Generic
	ForTypeNode  *graph.TypeNode
	GBusReq      generic.Generic
	GSvcRsp      generic.Generic
	SvcFilter    *mql.ServiceFilter
	GRspSelect   generic.Generic
	EndpointNode *graph.EndpointNode
	Svcs         []mql.Service
	Svc          *mql.Service
	Errs         []mql.Error
	GEntity      generic.Generic
	GEntities    generic.Slice
	SvcId        *mql.ServiceId
	SvcIds       []mql.ServiceId
	BusReqCtxs   []ReqCtx
	DoCliReqValidation bool
	DoCliReqProcessing bool
	DoSetClientAccount bool
}

func (c ReqCtx) Copy() (ctx ReqCtx, err error) {
	ctx = ReqCtx{
		Method:             c.Method,
		Mode:               c.Mode,
		Id:                 c.Id,
		Stage:              c.Stage,
		Ctx:                c.Ctx,
		ForTypeNode:        c.ForTypeNode,
		GCliRsps:           c.GCliRsps, // todo
		EndpointNode:       c.EndpointNode,
		Errs:               c.Errs, // todo
		DoCliReqValidation: c.DoCliReqValidation,
		DoCliReqProcessing: c.DoCliReqProcessing,
		DoSetClientAccount: c.DoSetClientAccount,
	}

	if c.GCliReq != nil {
		ctx.GCliReq = c.GCliReq.Copy()
	}

	if c.GBusRsp != nil {
		ctx.GBusRsp = c.GBusRsp.Copy()
	}

	if c.GBusReq != nil {
		ctx.GBusReq = c.GBusReq.Copy()
	}

	if c.GSvcRsp != nil {
		ctx.GSvcRsp = c.GSvcRsp.Copy()
	}

	if c.GRspSelect != nil {
		ctx.GRspSelect = c.GRspSelect.Copy()
	}

	if c.SvcFilter != nil {
		ctx.SvcFilter = &mql.ServiceFilter{}
		*ctx.SvcFilter = *c.SvcFilter
	}

	if len(c.Svcs) != 0 {
		for i, _ := range c.Svcs {
			ctx.Svcs = append(ctx.Svcs, c.Svcs[i])
		}
	}

	if c.Svc != nil {
		ctx.Svc = &mql.Service{}
		*ctx.Svc = *c.Svc
	}

	if len(c.Errs) != 0 {
		for i, _ := range c.Errs {
			ctx.Errs = append(ctx.Errs, c.Errs[i])
		}
	}

	if c.GEntity != nil {
		ctx.GEntity = c.GEntity.Copy()
	}

	if c.GEntities != nil {
		ctx.GEntities = c.GEntities.Copy()
	}

	return
}

func (c ReqCtx) Sprint() (s string) {
	i := getCtxInspect(c)

	b, err := yaml.Marshal(i)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("(?m)[\r\n]+^.*: null.*$")
	res := re.ReplaceAll(b, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: {}.*$")
	res = re.ReplaceAll(res, []byte{})

	re = regexp.MustCompile("(?m)[\r\n]+^.*: \\[\\].*$")
	res = re.ReplaceAll(res, []byte{})

	return string(res)
}

func (c ReqCtx) Inspect() {
	println("---")
	println(c.Sprint())
	println("---")
}

func getCtxInspect(ctx ReqCtx) (i ContextInspect) {
	if ctx.Method != "" {
		i.Method = ctx.Method
	}

	if ctx.Mode != "" {
		i.Mode = ctx.Mode
	}

	if ctx.ForTypeNode != nil {
		i.Type = ctx.ForTypeNode.Name()
	}

	if ctx.Id != "" {
		i.Id = ctx.Id
	}

	if ctx.GCliReq != nil {
		i.GCliReq = "x"
	}

	if ctx.GBusReq != nil {
		i.GSvcReq = "x"
	}

	if ctx.GSvcRsp != nil {
		i.GSvcRsp = "x"
	}

	if ctx.GBusRsp != nil {
		i.GCliRsp = "x"
	}

	if ctx.SvcFilter != nil {
		i.SvcFilter = *ctx.SvcFilter
	}

	if ctx.Svc != nil {
		i.Svc = *ctx.Svc.Id.ServiceName + "/" + *ctx.Svc.Id.Value
	}

	if len(ctx.Svcs) != 0 {
		for _, svc := range ctx.Svcs {
			i.Svcs = append(i.Svcs, *svc.Id.ServiceName+"/"+*svc.Id.Value)
		}
	}

	return
}

type ContextInspect struct {
	Id        string
	Method    string
	Mode      string
	Type      string
	GCliReq   string
	GSvcReq   string
	GSvcRsp   string
	GCliRsp   string
	SvcFilter mql.ServiceFilter
	Svc       string
	Svcs      []string
}
