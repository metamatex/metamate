package funcs

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"

	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/validation"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"reflect"
)

const (
	ValidateCliReqName                     = "validate client request"
	ProcessCliRspName                      = "process client response"
	ValidateCliRspName                     = "validate client response"
	SelectCliRspName                       = "select client response"
	SetStageName                           = "set stage"
	GetSvcsName                            = "get services"
	RequireOneGSvcName                     = "require one service"
	SetFirstGSvcName                       = "set first service"
	GSvcRspToGCliRspName                   = "service response to client response"
	FuncName                               = "func"
	HandleReqName                          = "handle request"
	AddSvcToEntitiesName                   = "add service to entities"
	SetSvcFilterToGetModeIdSvcIdName       = "set service filter to get mode id service id"
	SetSvcFilterToGetModeRelationSvcIdName = "set service filter to get mode relation service id"
)

func ReduceSvcRspsToCliRsp(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.GCliRsp = ReduceSvcRspsToCliRspFunc(f, ctx.ForTypeNode, ctx.GSvcRsps, ctx.GCliRsp)

			return ctx
		},
	}
}

func ReduceSvcRspsToCliRspFunc(f generic.Factory, forTn *graph.TypeNode, gSvcRsps []generic.Generic, gCliRsp generic.Generic) generic.Generic {
	gSlice := f.NewSlice(forTn)
	for _, gSvcRsp := range gSvcRsps {
		gSlice0, ok := gSvcRsp.GenericSlice(forTn.PluralFieldName())
		if !ok {
			continue
		}

		gSlice.Append(gSlice0.Get()...)
	}

	gCliRsp.MustSetGenericSlice([]string{forTn.PluralFieldName()}, gSlice)

	return gCliRsp
}

func ReduceSvcRspErrsToCliRspErrs(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gErrs := f.MustFromStructs([]sdk.Error{})

			gCliRspErrs, ok := ctx.GCliRsp.GenericSlice(fieldnames.Errors)
			if ok {
				gErrs.Append(gCliRspErrs.Get()...)
			}

			for _, gSvcRsp := range ctx.GSvcRsps {
				gSvcRspErrs, ok := gSvcRsp.GenericSlice(fieldnames.Errors)
				if ok {
					gErrs.Append(gSvcRspErrs.Get()...)
				}
			}

			if len(gErrs.Get()) != 0 {
				ctx.GCliRsp.MustSetGenericSlice([]string{fieldnames.Errors}, gErrs)
			}

			return ctx
		},
	}
}

func ReduceSvcRspPaginationsToCliRspPagination(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			p := sdk.Pagination{}

			for _, gSvcRsp := range ctx.GSvcRsps {
				gPagination, ok := gSvcRsp.Generic(fieldnames.Pagination)
				if !ok {
					continue
				}

				var p0 sdk.Pagination
				gPagination.MustToStruct(&p0)

				p.Previous = append(p.Previous, p0.Previous...)
				p.Current = append(p.Current, p0.Current...)
				p.Next = append(p.Next, p0.Next...)
			}

			ctx.GCliRsp.MustSetGeneric([]string{fieldnames.Pagination}, f.MustFromStruct(p))

			return ctx
		},
	}
}

func HardFilterGCliRsp() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gFilter, ok := ctx.GCliReq.Generic(fieldnames.Filter)
			if !ok {
				return ctx
			}

			fieldname := ctx.ForTypeNode.PluralFieldName()

			gEntities, ok := ctx.GCliRsp.GenericSlice(fieldname)
			if !ok {
				return ctx
			}

			gEntities = gEntities.Filter(false, gFilter)

			ctx.GCliRsp.MustSetGenericSlice([]string{fieldname}, gEntities)

			return ctx
		},
	}
}

func Func(f func(ctx types.ReqCtx) types.ReqCtx) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func:  f,
	}
}

func ApplySvcEndpointReqFilters(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gReqs := f.NewSlice(ctx.GCliReq.Type())
			gReqs.Append(ctx.GCliReq)

			var svcs []sdk.Service
			for i, _ := range ctx.Svcs {
				gSvc := f.MustFromStruct(ctx.Svcs[i])

				gEndpoint := gSvc.MustGeneric(fieldnames.Endpoints, ctx.GCliReq.Type().Edges.Endpoint.BelongsTo().FieldName())

				gReqFilter, ok := gEndpoint.Generic(fieldnames.Filter)
				if ok {
					if len(gReqs.Copy().Filter(false, gReqFilter).Get()) == 1 {
						svcs = append(svcs, ctx.Svcs[i])
					}
				} else {
					svcs = append(svcs, ctx.Svcs[i])
				}
			}

			ctx.Svcs = svcs

			return ctx
		},
	}
}

func SetSvcFilterToGetModeIdSvcIdFunc() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: SetSvcFilterToGetModeIdSvcIdName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			serviceName, ok := ctx.GCliReq.String(fieldnames.Mode, fieldnames.Id, fieldnames.ServiceId, fieldnames.ServiceName)
			if !ok {
				return ctx
			}

			if ctx.SvcFilter.Id == nil {
				ctx.SvcFilter.Id = &sdk.ServiceIdFilter{
					Value: &sdk.StringFilter{},
				}
			}

			ctx.SvcFilter.Id.Value.Is = sdk.String(serviceName)

			return ctx
		},
	}
}

func SetSvcFilterToGetModeRelationIdFunc() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: SetSvcFilterToGetModeRelationSvcIdName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			serviceName, ok := ctx.GCliReq.String(fieldnames.Mode, fieldnames.Relation, fieldnames.Id, fieldnames.ServiceName)
			if !ok {
				return ctx
			}

			if ctx.SvcFilter.Id == nil {
				ctx.SvcFilter.Id = &sdk.ServiceIdFilter{
					Value: &sdk.StringFilter{},
				}
			}

			ctx.SvcFilter.Id.Value.Is = sdk.String(serviceName)

			return ctx
		},
	}
}

func SetDefaultSelect() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "set default select",
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.GCliReq.MustSetBool([]string{fieldnames.Select, fieldnames.Errors, fieldnames.All}, true)
			ctx.GCliReq.MustSetBool([]string{fieldnames.Select, ctx.ForTypeNode.PluralFieldName(), fieldnames.Id, fieldnames.ServiceName}, true)
			ctx.GCliReq.MustSetBool([]string{fieldnames.Select, ctx.ForTypeNode.PluralFieldName(), fieldnames.Id, fieldnames.Value}, true)

			return ctx
		},
	}
}

func ValidateCliReq(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: ValidateCliReqName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.Errs = validation.Validate(ctx.Stage, nil, ctx.GCliReq)

			return ctx
		},
	}
}

func ProcessCliRsp(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: ProcessCliRspName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			//ctx.GCliRsp = process.Process(ctx, f, ctx.GCliRsp)

			return ctx
		},
	}
}

func ValidateCliRsp() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: ValidateCliRspName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			//ctx.Errs = validation.Validate(ctx.Ctx, ctx.GCliRsp)

			return ctx
		},
	}
}

func SetStage(stage string) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: SetStageName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.Stage = stage

			return ctx
		},
	}
}

func SetId() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "set id",
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			var u uuid.UUID
			u, err := uuid.NewUUID()
			if err != nil {
				ctx.Errs = append(ctx.Errs, NewError(nil, sdk.ErrorKind.Internal, err.Error()))

				return ctx
			}

			ctx.Id = u.String()

			return ctx
		},
	}
}

func New(f generic.Factory, subject string) types.FuncTransformer {
	name := fmt.Sprintf("new %v", subject)

	switch subject {
	case types.GSvcRsp:
		return types.FuncTransformer{
			Name0: name,
			Func: func(ctx types.ReqCtx) types.ReqCtx {
				ctx.GSvcRsp = f.New(ctx.GSvcReq.Type().Edges.Type.Response())

				return ctx
			},
		}
	case types.GCliRsp:
		return types.FuncTransformer{
			Name0: name,
			Func: func(ctx types.ReqCtx) types.ReqCtx {
				ctx.GCliRsp = f.New(ctx.GCliReq.Type().Edges.Type.Response())

				return ctx
			},
		}
	}

	panic(fmt.Sprintf("subject %v not supported", subject))
}

func Set(f generic.Factory, subject string, any interface{}) types.FuncTransformer {
	name := fmt.Sprintf("set %v to %v", subject, any)

	switch subject {
	case types.GSvcReq:
		t := reflect.TypeOf(any)
		switch t.Kind() {
		case reflect.Struct:
			g := f.MustFromStruct(any)

			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.GSvcReq = g

					return ctx
				},
			}
		}
	case types.Stage:
		switch sth := any.(type) {
		case string:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.Stage = sth

					return ctx
				},
			}
		}
	}

	panic(fmt.Sprintf("set %v to %v not supported", subject, any))
}

func Copy(from, to string) types.FuncTransformer {
	name := fmt.Sprintf("copy %v to %v", from, to)
	switch from {
	case types.GEntity:
		switch to {
		case types.Svc:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.GEntity.MustToStruct(&ctx.Svc)

					return ctx
				},
			}
		}
	case types.Svc:
		switch to {
		case types.Errs:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					for i, _ := range ctx.Errs {
						ctx.Errs[i].Service = ctx.Svc
					}

					return ctx
				},
			}
		}
	case types.GCliReq:
		switch to {
		case types.ForTypeNode:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.ForTypeNode = ctx.GCliReq.Type().Edges.Type.For()

					return ctx
				},
			}
		case types.GCliReq:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.GCliReq = ctx.GCliReq.Copy()

					return ctx
				},
			}
		case types.GSvcReq:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.GSvcReq = ctx.GCliReq.Copy()

					return ctx
				},
			}
		case types.Method:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.Method = ctx.GCliReq.Type().Edges.Endpoint.BelongsTo().Data.Method

					return ctx
				},
			}
		case types.Mode:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.Mode = ctx.GCliReq.MustString(fieldnames.Mode, fieldnames.Kind)

					return ctx
				},
			}
		case types.EndpointNode:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.EndpointNode = ctx.GCliReq.Type().Edges.Endpoint.BelongsTo()

					return ctx
				},
			}
		case types.SvcFilter:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.SvcFilter = &sdk.ServiceFilter{}

					gSvcFilter, ok := ctx.GCliReq.Generic(fieldnames.ServiceFilter)
					if ok {
						gSvcFilter.MustToStruct(&ctx.SvcFilter)
					}

					return ctx
				},
			}
		}

	}

	panic(fmt.Sprintf("copy %v to %v not supported", from, to))
}

func Move(f generic.Factory, from, to string) types.FuncTransformer {
	name := fmt.Sprintf("move %v to %v", from, to)
	switch from {
	case types.Errs:
		switch to {
		case types.GSvcRsp:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					if len(ctx.Errs) == 0 {
						return ctx
					}

					ctx.GSvcRsp.MustSetGenericSlice([]string{fieldnames.Errors}, f.MustFromStructs(ctx.Errs))

					ctx.Errs = nil

					return ctx
				},
			}
		case types.GCliRsp:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					if len(ctx.Errs) == 0 {
						return ctx
					}

					ctx.GCliRsp.MustSetGenericSlice([]string{fieldnames.Errors}, f.MustFromStructs(ctx.Errs))

					ctx.Errs = nil

					return ctx
				},
			}
		}
	}

	panic(fmt.Sprintf("move %v to %v not supported", from, to))
}

func HandleSvcReq(hs map[bool]map[string]types.RequestHandler) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: HandleReqName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx = func(ctx types.ReqCtx) types.ReqCtx {
				if ctx.Svc.Url == nil || ctx.Svc.Url.Value == nil {
					ctx.Errs = append(ctx.Errs, NewError(ctx.Svc, sdk.ErrorKind.Internal, "service.url.value not set"))

					return ctx
				}

				h, ok := hs[ctx.Svc.IsVirtual != nil && *ctx.Svc.IsVirtual][*ctx.Svc.Transport]
				if !ok {
					ctx.Errs = append(ctx.Errs, NewError(nil, sdk.ErrorKind.Internal, fmt.Sprintf("no handler for transport %v", *ctx.Svc.Transport)))

					return ctx
				}

				var err error
				ctx.GSvcRsp, err = h(ctx.Ctx, *ctx.Svc.Url.Value, ctx.GSvcReq)
				if err != nil {
					ctx.Errs = append(ctx.Errs, NewError(ctx.Svc, sdk.ErrorKind.Service, err.Error()))

					return ctx
				}

				gErrs, ok := ctx.GSvcRsp.GenericSlice(fieldnames.Errors)
				if ok {
					var errs []sdk.Error
					gErrs.MustToStructs(&errs)

					ctx.Errs = append(ctx.Errs, errs...)
				}

				return ctx
			}(ctx)

			return ctx
		},
	}
}

func Log(stage string, stages types.InternalLogTemplates) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: SelectCliRspName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ts, ok := stages[stage]
			if !ok {
				ts, ok = stages["*"]
				if !ok {
					return ctx
				}
			}

			var s generic.Generic
			switch stage {
			case config.CliReq:
				s = ctx.GCliReq
			case config.SvcReq:
				s = ctx.GSvcReq
			case config.SvcRsp:
				s = ctx.GSvcRsp
			case config.CliRsp:
				s = ctx.GCliRsp
			default:
				panic(fmt.Sprintf("no case for stage: %v", stage))
			}

			t, ok := ts[s.Type().Name()]
			if !ok {
				t, ok = ts["*"]
				if !ok {
					return ctx
				}
			}

			var b bytes.Buffer
			err := t.Execute(&b, types.InternalLogData{
				Subject: s,
				Ctx:     ctx,
			})
			if err != nil {
				ctx.Errs = append(ctx.Errs, NewError(nil, sdk.ErrorKind.Internal, err.Error()))

				return ctx
			}

			println(b.String())

			return ctx
		},
	}
}

func ResolveRelations(resolvePl *line.Line, f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: HandleReqName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gRelations, ok := ctx.GCliReq.Generic(fieldnames.Relations)
			if !ok {
				return ctx
			}

			gRelations.EachGeneric(func(fn *graph.FieldNode, gGetCollection generic.Generic) {
				gGetReq := f.New(fn.Edges.Path.BelongsTo().Edges.Type.To().Edges.Type.GetRequest())

				id := sdk.ServiceId{}
				ctx.GEntity.MustGeneric(fieldnames.Id).MustToStruct(&id)

				getMode := sdk.GetMode{
					Kind: &sdk.GetModeKind.Relation,
					Relation: &sdk.RelationGetMode{
						Id:       &id,
						Relation: sdk.String(fn.Edges.Path.BelongsTo().Name()),
					},
				}

				gGetReq.MustSetGeneric([]string{fieldnames.Mode}, f.MustFromStruct(getMode))

				gFilter, ok := gGetCollection.Generic(fieldnames.Filter)
				if ok {
					gGetReq.MustSetGeneric([]string{fieldnames.Filter}, gFilter.Copy())
				}

				gSelect, ok := gGetCollection.Generic(fieldnames.Select)
				if ok {
					gGetReq.MustSetGeneric([]string{fieldnames.Select}, gSelect.Copy())
				}

				gServiceFilter, ok := gGetCollection.Generic(fieldnames.ServiceFilter)
				if ok {
					gGetReq.MustSetGeneric([]string{fieldnames.ServiceFilter}, gServiceFilter.Copy())
				}

				gRelations0, ok := gGetCollection.Generic(fieldnames.Relations)
				if ok {
					gGetReq.MustSetGeneric([]string{fieldnames.Relations}, gRelations0.Copy())
				}

				ctx0 := resolvePl.Transform(types.ReqCtx{
					GCliReq: gGetReq,
				})

				gCollection := f.New(gGetCollection.Type().Edges.Type.For().Edges.Type.Collection())

				gSlice, ok := ctx0.GCliRsp.GenericSlice(gCollection.Type().Edges.Type.For().PluralFieldName())
				if ok {
					gCollection.MustSetGenericSlice([]string{gCollection.Type().Edges.Type.For().PluralFieldName()}, gSlice)
				}

				gErrors, ok := ctx0.GCliRsp.GenericSlice(fieldnames.Errors)
				if ok {
					gCollection.MustSetGenericSlice([]string{fieldnames.Errors}, gErrors)
				}

				gWarnings, ok := ctx0.GCliRsp.GenericSlice(fieldnames.Warnings)
				if ok {
					gCollection.MustSetGenericSlice([]string{fieldnames.Warnings}, gWarnings)
				}

				ctx.GEntity.MustSetGeneric([]string{fieldnames.Relations, fn.Name()}, gCollection)

				return
			})

			return ctx
		},
	}
}

func GetEntityById(f generic.Factory, resolvePl *line.Line) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: HandleReqName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gGetReq := f.New(ctx.GEntity.Type().Edges.Type.GetRequest())

			id := sdk.ServiceId{}
			ctx.GEntity.MustGeneric(fieldnames.Id).MustToStruct(&id)

			getMode := sdk.GetMode{
				Kind: &sdk.GetModeKind.Id,
				Id: &sdk.Id{
					Kind:      &sdk.IdKind.ServiceId,
					ServiceId: &id,
				},
			}

			gGetReq.MustSetGeneric([]string{fieldnames.Mode}, f.MustFromStruct(getMode))

			ctx0 := resolvePl.Transform(types.ReqCtx{
				GCliReq: gGetReq,
			})

			gSlice, ok := ctx0.GCliRsp.GenericSlice(ctx.GEntity.Type().PluralFieldName())
			if ok {
				for _, g := range gSlice.Get() {
					serviceName, ok := g.String(fieldnames.Id, fieldnames.ServiceName)
					if !ok || serviceName != *id.ServiceName {
						continue
					}

					value, ok := g.String(fieldnames.Id, fieldnames.Value)
					if !ok || value != *id.Value {
						continue
					}

					ctx.GEntity = g

					break
				}
			}

			return ctx
		},
	}
}

func AddSvcToEntities(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: AddSvcToEntitiesName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gSlice, ok := ctx.GSvcRsp.GenericSlice(ctx.GSvcRsp.Type().Edges.Type.For().PluralFieldName())
			if !ok {
				return ctx
			}

			var gs []generic.Generic
			for _, g := range gSlice.Get() {
				g.MustSetGeneric([]string{fieldnames.Meta, fieldnames.Service}, f.MustFromStruct(ctx.Svc))

				gs = append(gs, g)
			}

			gSlice.Set(gs)

			return ctx
		},
	}
}

func AddSvcToSvcIds() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: AddSvcToEntitiesName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gSlice, ok := ctx.GSvcRsp.GenericSlice(ctx.GSvcRsp.Type().Edges.Type.For().PluralFieldName())
			if !ok {
				return ctx
			}

			var gs []generic.Generic
			for _, g := range gSlice.Get() {
				gId, ok := g.Generic(fieldnames.Id)
				if !ok {
					gs = append(gs, g)

					continue
				}

				_, ok = gId.String(fieldnames.ServiceName)
				if !ok {
					g.MustSetString([]string{fieldnames.Id, fieldnames.ServiceName}, *ctx.Svc.Id.Value)

					gs = append(gs, g)

					continue
				}

				gs = append(gs, g)
			}

			gSlice.Set(gs)

			return ctx
		},
	}
}

func AddSvcIdToSvcPages(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: AddSvcToEntitiesName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gPagination, ok := ctx.GSvcRsp.Generic(fieldnames.Pagination)
			if !ok {
				return ctx
			}

			var p sdk.Pagination
			gPagination.MustToStruct(&p)

			for i, _ := range p.Previous {
				p.Previous[i].Id = ctx.Svc.Id
			}

			for i, _ := range p.Current {
				p.Current[i].Id = ctx.Svc.Id
			}

			for i, _ := range p.Next {
				p.Next[i].Id = ctx.Svc.Id
			}

			ctx.GSvcRsp.MustSetGeneric([]string{fieldnames.Pagination}, f.MustFromStruct(p))

			return ctx
		},
	}
}

func FilterSvcPages(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: AddSvcToEntitiesName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gPages, ok := ctx.GSvcReq.GenericSlice(fieldnames.Pages)
			if !ok {
				return ctx
			}

			filter := sdk.ServicePageFilter{
				Id: &sdk.ServiceIdFilter{
					ServiceName: &sdk.StringFilter{
						Is: ctx.Svc.Id.ServiceName,
					},
					Value: &sdk.StringFilter{
						Is: ctx.Svc.Id.Value,
					},
				},
			}

			gPages = gPages.Filter(false, f.MustFromStruct(filter))

			ctx.GSvcReq.MustSetGenericSlice([]string{fieldnames.Pages}, gPages)

			return ctx
		},
	}
}

func GetSvcs(resolve *line.Line, f generic.Factory, discoverySvc sdk.Service) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: GetSvcsName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			if ctx.EndpointNode.Name() == endpointnames.Get(typenames.Service) {
				ctx.Svcs = []sdk.Service{
					discoverySvc,
				}

				return ctx
			}

			f.
				MustFromStruct(ctx.SvcFilter).
				MustSetBool([]string{fieldnames.Endpoints, ctx.EndpointNode.FieldName(), fieldnames.Set}, true).
				MustToStruct(ctx.SvcFilter)

			req := sdk.GetServicesRequest{
				Mode: &sdk.GetMode{
					Kind:       &sdk.GetModeKind.Collection,
					Collection: &sdk.CollectionGetMode{},
				},
				Filter: ctx.SvcFilter,
				Select: &sdk.GetServicesResponseSelect{
					All: sdk.Bool(true),
					Services: &sdk.ServiceSelect{
						All: sdk.Bool(true),
					},
				},
			}

			ctx0 := resolve.Transform(types.ReqCtx{
				GCliReq: f.MustFromStruct(req),
			})

			gSvcErrs, ok := ctx0.GCliRsp.GenericSlice(fieldnames.Errors)
			if ok {
				var errs []sdk.Error
				gSvcErrs.MustToStructs(&errs)
				ctx.Errs = append(ctx.Errs, errs...)
			}

			gSvcs, ok := ctx0.GCliRsp.GenericSlice(ctx0.ForTypeNode.PluralFieldName())
			if ok {
				gSvcs.MustToStructs(&ctx.Svcs)
			}

			return ctx
		},
	}
}

func RequireOneGSvc() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: RequireOneGSvcName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			if len(ctx.Svcs) == 0 {
				ctx.Errs = append(ctx.Errs, NewError(nil, sdk.ErrorKind.Internal, "len(ctx.gSvcs) is 0"))

				return ctx
			}

			if len(ctx.Svcs) > 1 {
				ctx.Errs = append(ctx.Errs, NewError(nil, sdk.ErrorKind.Internal, "len(ctx.gSvcs) is greater 1"))

				return ctx
			}

			return ctx
		},
	}
}

func SetFirstGSvc() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: SetFirstGSvcName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.Svc = &ctx.Svcs[0]

			return ctx
		},
	}
}

func GSvcRspToGCliRsp() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: GSvcRspToGCliRspName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.GCliRsp = ctx.GSvcRsp

			return ctx
		},
	}
}

func Inspect() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "",
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.Inspect()

			return ctx
		},
	}
}

func NewError(svc *sdk.Service, kind, message string) sdk.Error {
	return sdk.Error{
		Kind:    sdk.String(kind),
		Message: &message,
		Service: svc,
	}
}
