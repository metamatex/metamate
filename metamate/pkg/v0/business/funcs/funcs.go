package funcs

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/utils"
	"sync"
	"time"

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

func ReduceBusReqCtxsToBusRsp(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gSlice := f.NewSlice(ctx.ForTypeNode)
			for _, brCtx := range ctx.BusReqCtxs {
				gSlice0, ok := brCtx.GSvcRsp.GenericSlice(ctx.ForTypeNode.PluralFieldName())
				if !ok {
					continue
				}

				gSlice.Append(gSlice0.Get()...)
			}

			ctx.GBusRsp.MustSetGenericSlice([]string{ctx.ForTypeNode.PluralFieldName()}, gSlice)

			return ctx
		},
	}
}

func ReduceBusReqCtxsErrsToBusRspErrs(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gErrs := f.MustFromStructs([]mql.Error{})

			gCliRspErrs, ok := ctx.GBusRsp.GenericSlice(fieldnames.Errors)
			if ok {
				gErrs.Append(gCliRspErrs.Get()...)
			}

			for _, brCtx := range ctx.BusReqCtxs {
				gSvcRspErrs, ok := brCtx.GSvcRsp.GenericSlice(fieldnames.Errors)
				if ok {
					gErrs.Append(gSvcRspErrs.Get()...)
				}
			}

			if len(gErrs.Get()) != 0 {
				ctx.GBusRsp.MustSetGenericSlice([]string{fieldnames.Errors}, gErrs)
			}

			return ctx
		},
	}
}

func ReduceSvcRspPaginationsToCliRspPagination(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: FuncName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			busPagi := mql.BusPagination{}

			for _, brCtx := range ctx.BusReqCtxs {
				gPagination, ok := brCtx.GSvcRsp.Generic(fieldnames.Pagination)
				if !ok {
					continue
				}

				var svcPagi mql.ServicePagination
				gPagination.MustToStruct(&svcPagi)

				if svcPagi.Previous != nil {
					busPagi.Previous = append(busPagi.Previous, mql.ServicePage{
						Id:   brCtx.Svc.Id,
						Page: svcPagi.Previous,
					})
				}

				if svcPagi.Current != nil {
					busPagi.Current = append(busPagi.Current, mql.ServicePage{
						Id:   brCtx.Svc.Id,
						Page: svcPagi.Current,
					})
				}

				if svcPagi.Next != nil {
					busPagi.Next = append(busPagi.Next, mql.ServicePage{
						Id:   brCtx.Svc.Id,
						Page: svcPagi.Next,
					})
				}
			}

			ctx.GBusRsp.MustSetGeneric([]string{fieldnames.Pagination}, f.MustFromStruct(busPagi))

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

			gEntities, ok := ctx.GBusRsp.GenericSlice(fieldname)
			if !ok {
				return ctx
			}

			gEntities = gEntities.Filter(false, gFilter)

			ctx.GBusRsp.MustSetGenericSlice([]string{fieldname}, gEntities)

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

			var svcs []mql.Service
			for i, _ := range ctx.Svcs {
				gSvc := f.MustFromStruct(ctx.Svcs[i])

				gEndpoint := gSvc.MustGeneric(fieldnames.Endpoints, ctx.GCliReq.Type().Edges.Endpoint.For().FieldName())

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
				ctx.SvcFilter.Id = &mql.ServiceIdFilter{
					Value: &mql.StringFilter{},
				}
			}

			ctx.SvcFilter.Id.Value.Is = mql.String(serviceName)

			return ctx
		},
	}
}

func SetSvcFilterToGetModeRelationIdFunc() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: SetSvcFilterToGetModeRelationSvcIdName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			serviceName, ok := ctx.GCliReq.String(fieldnames.Mode, fieldnames.Relation, fieldnames.Id, fieldnames.ServiceId, fieldnames.ServiceName)
			if !ok {
				return ctx
			}

			if ctx.SvcFilter.Id == nil {
				ctx.SvcFilter.Id = &mql.ServiceIdFilter{
					Value: &mql.StringFilter{},
				}
			}

			ctx.SvcFilter.Id.Value.Is = mql.String(serviceName)

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
			//ctx.GBusRsp = process.Process(ctx, f, ctx.GBusRsp)

			return ctx
		},
	}
}

func ValidateCliRsp() types.FuncTransformer {
	return types.FuncTransformer{
		Name0: ValidateCliRspName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			//ctx.Errs = validation.Validate(ctx.Ctx, ctx.GBusRsp)

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
				ctx.Errs = append(ctx.Errs, NewError(nil, mql.ErrorKind.Internal, err.Error()))

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
				ctx.GSvcRsp = f.New(ctx.GBusReq.Type().Edges.Type.Response())

				return ctx
			},
		}
	case types.GBusRsp:
		return types.FuncTransformer{
			Name0: name,
			Func: func(ctx types.ReqCtx) types.ReqCtx {
				ctx.GBusRsp = f.New(ctx.GCliReq.Type().Edges.Type.Response())

				return ctx
			},
		}
	}

	panic(fmt.Sprintf("subject %v not supported", subject))
}

func Set(f generic.Factory, subject string, any interface{}) types.FuncTransformer {
	name := fmt.Sprintf("set %v to %v", subject, any)

	switch subject {
	case types.GBusReq:
		t := reflect.TypeOf(any)
		switch t.Kind() {
		case reflect.Struct:
			g := f.MustFromStruct(any)

			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.GBusReq = g

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
		case types.Method:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.Method = ctx.GCliReq.Type().Edges.Endpoint.For().Data.Method

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
					ctx.EndpointNode = ctx.GCliReq.Type().Edges.Endpoint.For()

					return ctx
				},
			}
		case types.SvcFilter:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					ctx.SvcFilter = &mql.ServiceFilter{}

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
		case types.GBusRsp:
			return types.FuncTransformer{
				Name0: name,
				Func: func(ctx types.ReqCtx) types.ReqCtx {
					if len(ctx.Errs) == 0 {
						return ctx
					}

					ctx.GBusRsp.MustSetGenericSlice([]string{fieldnames.Errors}, f.MustFromStructs(ctx.Errs))

					ctx.Errs = nil

					return ctx
				},
			}
		}
	}

	panic(fmt.Sprintf("move %v to %v not supported", from, to))
}

func CreateGBusReqFromGCliReq(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "CreateGBusReqFromGCliReq",
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx.GBusReq = f.New(ctx.GCliReq.Type().Edges.Type.BusRequest())

			gMode, ok := ctx.GCliReq.Generic(mql.FieldNames.Mode)
			if ok {
				ctx.GBusReq.MustSetGeneric([]string{mql.FieldNames.Mode}, gMode)
			}

			return ctx
		},
	}
}

func SetServiceAccount(f generic.Factory, l *sync.Mutex, ass []types.ServiceAccountAssignment, hs map[bool]types.RequestHandler) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: "SetServiceAccount",
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			l.Lock()
			defer l.Unlock()

			if ctx.Svc.Endpoints.AuthenticateServiceAccount == nil {
				return ctx
			}

			var acc *mql.ServiceAccount
			var accI int
			for i, as := range ass {
				if *as.ServiceId.ServiceName != *ctx.Svc.Id.ServiceName ||
					*as.ServiceId.Value != *ctx.Svc.Id.Value  {
					continue
				}

				acc = &as.ServiceAccount
				accI = i
				break
			}

			if acc == nil {
				ctx.Errs = append(ctx.Errs, mql.Error{
					Kind: &mql.ErrorKind.Internal,
					Message: mql.String("no service account found"),
					Service: ctx.Svc,
				})

				return ctx
			}

			isExpired := false
			if acc.Token != nil && acc.Token.ExpiresAt != nil {
				t, err := utils.FromUnixTimeStamp(*acc.Token.ExpiresAt)
				if err != nil {
					ctx.Errs = append(ctx.Errs, mql.Error{
						Kind: &mql.ErrorKind.Internal,
						Message: mql.String(err.Error()),
						Service: ctx.Svc,
					})

				    return ctx
				}

				isExpired = t.After(time.Now())
			}

			if acc.Token != nil &&
				acc.Token.Value != nil &&
				!isExpired {
				ctx.GBusReq.MustSetGeneric([]string{mql.FieldNames.Account}, f.MustFromStruct(*acc))

				return ctx
			}

			req := mql.AuthenticateServiceAccountBusRequest{
				Input: &mql.AuthenticateServiceAccountInput{
					Account: acc,
				},
			}

			h, ok := hs[ctx.Svc.IsEmbedded != nil && *ctx.Svc.IsEmbedded]
			if !ok {
				panic("expected handler for embedded and http")

				return ctx
			}

			var err error
			gSvcRsp, err := h(ctx.Ctx, *ctx.Svc.Url.Value, f.MustFromStruct(req))
			if err != nil {
				ctx.Errs = append(ctx.Errs, NewError(ctx.Svc, mql.ErrorKind.Service, err.Error()))

				return ctx
			}

			var rsp mql.AuthenticateServiceAccountServiceResponse
			gSvcRsp.MustToStruct(&rsp)

			if len(rsp.Errors) != 0 {
				ctx.Errs = append(ctx.Errs, rsp.Errors...)

				return ctx
			}

			if rsp.Output == nil ||
				rsp.Output.Token == nil {
				ctx.Errs = append(ctx.Errs, NewError(ctx.Svc, mql.ErrorKind.Internal, "authenticateServiceAccountServiceResponse didn't contain a token"))

				return ctx
			}

			acc.Token = rsp.Output.Token
			as := ass[accI]
			as.ServiceAccount = *acc
			ass[accI] = as

			ctx.GBusReq.MustSetGeneric([]string{mql.FieldNames.Account}, f.MustFromStruct(*acc))

			return ctx
		},
	}
}

func HandleSvcReq(hs map[bool]types.RequestHandler) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: HandleReqName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			ctx = func(ctx types.ReqCtx) types.ReqCtx {
				if ctx.Svc.Url == nil || ctx.Svc.Url.Value == nil {
					ctx.Errs = append(ctx.Errs, NewError(ctx.Svc, mql.ErrorKind.Internal, "service.url.value not set"))

					return ctx
				}

				h, ok := hs[ctx.Svc.IsEmbedded != nil && *ctx.Svc.IsEmbedded]
				if !ok {
					panic("expected handler for embedded and http")

					return ctx
				}

				var err error
				ctx.GSvcRsp, err = h(ctx.Ctx, *ctx.Svc.Url.Value, ctx.GBusReq)
				if err != nil {
					ctx.Errs = append(ctx.Errs, NewError(ctx.Svc, mql.ErrorKind.Service, err.Error()))

					return ctx
				}

				gErrs, ok := ctx.GSvcRsp.GenericSlice(fieldnames.Errors)
				if ok {
					var errs []mql.Error
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
				s = ctx.GBusReq
			case config.SvcRsp:
				s = ctx.GSvcRsp
			case config.CliRsp:
				s = ctx.GBusRsp
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
				ctx.Errs = append(ctx.Errs, NewError(nil, mql.ErrorKind.Internal, err.Error()))

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
				gGetReq := f.New(fn.Edges.Path.BelongsTo().Edges.Type.To().Edges.Type.GetClientRequest())

				id := mql.ServiceId{}
				ctx.GEntity.MustGeneric(fieldnames.Id).MustToStruct(&id)

				getMode := mql.GetMode{
					Kind: &mql.GetModeKind.Relation,
					Relation: &mql.RelationGetMode{
						Id: &mql.Id{
							Kind:      &mql.IdKind.ServiceId,
							ServiceId: &id,
						},
						Relation: mql.String(fn.Edges.Path.BelongsTo().Name()),
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

				gSlice, ok := ctx0.GBusRsp.GenericSlice(gCollection.Type().Edges.Type.For().PluralFieldName())
				if ok {
					gCollection.MustSetGenericSlice([]string{gCollection.Type().Edges.Type.For().PluralFieldName()}, gSlice)
				}

				gErrors, ok := ctx0.GBusRsp.GenericSlice(fieldnames.Errors)
				if ok {
					gCollection.MustSetGenericSlice([]string{fieldnames.Errors}, gErrors)
				}

				gWarnings, ok := ctx0.GBusRsp.GenericSlice(fieldnames.Warnings)
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
			gGetReq := f.New(ctx.GEntity.Type().Edges.Type.GetClientRequest())

			id := mql.ServiceId{}
			ctx.GEntity.MustGeneric(fieldnames.Id).MustToStruct(&id)

			getMode := mql.GetMode{
				Kind: &mql.GetModeKind.Id,
				Id: &mql.Id{
					Kind:      &mql.IdKind.ServiceId,
					ServiceId: &id,
				},
			}

			gGetReq.MustSetGeneric([]string{fieldnames.Mode}, f.MustFromStruct(getMode))

			ctx0 := resolvePl.Transform(types.ReqCtx{
				GCliReq: gGetReq,
			})

			gSlice, ok := ctx0.GBusRsp.GenericSlice(ctx.GEntity.Type().PluralFieldName())
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

func FilterSvcPages(f generic.Factory) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: AddSvcToEntitiesName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			gPages, ok := ctx.GCliReq.GenericSlice(fieldnames.Pages)
			if !ok {
				return ctx
			}

			filter := mql.ServicePageFilter{
				Id: &mql.ServiceIdFilter{
					ServiceName: &mql.StringFilter{
						Is: ctx.Svc.Id.ServiceName,
					},
					Value: &mql.StringFilter{
						Is: ctx.Svc.Id.Value,
					},
				},
			}

			var svcPages []mql.ServicePage
			gPages.Filter(false, f.MustFromStruct(filter)).MustToStructs(&svcPages)

			if len(svcPages) > 1 {
				ctx.Errs = append(ctx.Errs, NewError(nil, mql.ErrorKind.Internal, "more than one page for service"))

				return ctx
			}

			if len(svcPages) == 1 {
				ctx.GBusReq.MustSetGeneric([]string{fieldnames.Page}, f.MustFromStruct(svcPages[0].Page))
			}

			return ctx
		},
	}
}

func GetSvcs(resolve *line.Line, f generic.Factory, discoverySvc mql.Service) types.FuncTransformer {
	return types.FuncTransformer{
		Name0: GetSvcsName,
		Func: func(ctx types.ReqCtx) types.ReqCtx {
			if ctx.EndpointNode.Name() == endpointnames.Get(typenames.Service) {
				ctx.Svcs = []mql.Service{
					discoverySvc,
				}

				return ctx
			}

			f.
				MustFromStruct(ctx.SvcFilter).
				MustSetBool([]string{fieldnames.Endpoints, ctx.EndpointNode.FieldName(), fieldnames.Set}, true).
				MustToStruct(ctx.SvcFilter)

			req := mql.GetServicesClientRequest{
				Mode: &mql.GetMode{
					Kind:       &mql.GetModeKind.Collection,
					Collection: &mql.CollectionGetMode{},
				},
				Filter: ctx.SvcFilter,
				Select: &mql.GetServicesBusResponseSelect{
					All: mql.Bool(true),
					Services: &mql.ServiceSelect{
						All: mql.Bool(true),
					},
				},
			}

			ctx0 := resolve.Transform(types.ReqCtx{
				GCliReq: f.MustFromStruct(req),
			})

			gSvcErrs, ok := ctx0.GBusRsp.GenericSlice(fieldnames.Errors)
			if ok {
				var errs []mql.Error
				gSvcErrs.MustToStructs(&errs)
				ctx.Errs = append(ctx.Errs, errs...)
			}

			gSvcs, ok := ctx0.GBusRsp.GenericSlice(ctx0.ForTypeNode.PluralFieldName())
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
				ctx.Errs = append(ctx.Errs, NewError(nil, mql.ErrorKind.Internal, "len(ctx.gSvcs) is 0"))

				return ctx
			}

			if len(ctx.Svcs) > 1 {
				ctx.Errs = append(ctx.Errs, NewError(nil, mql.ErrorKind.Internal, "len(ctx.gSvcs) is greater 1"))

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
			ctx.GBusRsp = ctx.GSvcRsp

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

func NewError(svc *mql.Service, kind, message string) mql.Error {
	return mql.Error{
		Kind:    mql.String(kind),
		Message: &message,
		Service: svc,
	}
}
