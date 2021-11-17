package pipeline

import (
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/funcs"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"sync"
)

func NewResolveLine(rn *graph.RootNode, f generic.Factory, c types.Config, logTemplates types.InternalLogTemplates, reqHs map[bool]types.RequestHandler, cachedReqHs map[bool]types.RequestHandler) *line.Line {
	setServiceAccountLock := sync.Mutex{}

	resolveLine := line.Do()

	cliReqErrL := getErrLine(f, types.GBusRsp)

	svcReqErrL := getErrLine(f, types.GSvcRsp)

	resolveLine.
		Error(cliReqErrL, true).
		Do(
			funcs.SetId(),
			funcs.SetStage(config.CliReq),
			funcs.Copy(types.GCliReq, types.ForTypeNode),
			funcs.Copy(types.GCliReq, types.Method),
			funcs.Copy(types.GCliReq, types.SvcFilter),
			funcs.Copy(types.GCliReq, types.EndpointNode),
			funcs.New(f, types.GBusRsp),
		).
		Add(SetDefaults(f)).
		Do(funcs.ValidateCliReq(f)).
		Switch(
			funcs.By(types.Method),
			map[string]*line.Line{
				mql.Methods.Get: line.Do(funcs.Copy(types.GCliReq, types.Mode)),
			},
		).
		Add(NarrowSvcFilterToModeId).
		Do(
			funcs.GetSvcs(resolveLine, f, c.DiscoverySvc),
			funcs.ApplySvcEndpointReqFilters(f),
		).
		If(
			func(ctx types.ReqCtx) bool {
				return len(ctx.Svcs) == 0
			},
			line.
				Do(
					funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
						ctx.Errs = append(ctx.Errs, funcs.NewError(nil, mql.ErrorKind.NoServiceMatch, "no services matches"))

						return ctx
					}),
				),
		).
		Do(funcs.SetStage(config.SvcReq)).
		Switch(
			funcs.By(types.Method),
			map[string]*line.Line{
				mql.Methods.Action: line.
					Do(
						funcs.RequireOneGSvc(),
						funcs.SetFirstGSvc(),
						funcs.CreateGBusReqFromGCliReq(f),
						funcs.Log(config.SvcReq, logTemplates),
						funcs.HandleSvcReq(reqHs),
						funcs.Log(config.SvcRsp, logTemplates),
						funcs.GSvcRspToGCliRsp(),
					),
				line.Default: line.
					Switch(
						funcs.By(types.Method),
						map[string]*line.Line{
							mql.Methods.Get: line.New(config.SvcReq).
								Switch(
									funcs.By(types.Mode),
									map[string]*line.Line{
										mql.GetModeKind.Id: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.CreateGBusReqFromGCliReq(f),
														funcs.SetServiceAccount(f, &setServiceAccountLock, c.ServiceAccounts, reqHs),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													),
												funcs.CollectBusReqCtxs,
											).
											Do(
												funcs.ReduceBusReqCtxsToBusRsp(f),
												funcs.ReduceBusReqCtxsErrsToBusRspErrs(f),
												funcs.ReduceSvcRspPaginationsToCliRspPagination(f),
											),
										mql.GetModeKind.Search: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.CreateGBusReqFromGCliReq(f),
														funcs.FilterSvcPages(f),
														funcs.SetServiceAccount(f, &setServiceAccountLock, c.ServiceAccounts, reqHs),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													),
												funcs.CollectBusReqCtxs,
											).
											Do(
												funcs.ReduceBusReqCtxsToBusRsp(f),
												funcs.ReduceBusReqCtxsErrsToBusRspErrs(f),
												funcs.ReduceSvcRspPaginationsToCliRspPagination(f),
											),
										mql.GetModeKind.Collection: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.CreateGBusReqFromGCliReq(f),
														funcs.FilterSvcPages(f),
														funcs.SetServiceAccount(f, &setServiceAccountLock, c.ServiceAccounts, reqHs),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													).
													If(
														funcs.IsType(types.GSvcRsp, mql.TypeNames.GetServicesServiceResponse, true),
														LookupServices(f, reqHs, logTemplates),
													),
												funcs.CollectBusReqCtxs,
											).
											Do(
												funcs.ReduceBusReqCtxsToBusRsp(f),
												funcs.ReduceBusReqCtxsErrsToBusRspErrs(f),
												funcs.ReduceSvcRspPaginationsToCliRspPagination(f),
											),
										mql.GetModeKind.Relation: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.CreateGBusReqFromGCliReq(f),
														funcs.FilterSvcPages(f),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.SetServiceAccount(f, &setServiceAccountLock, c.ServiceAccounts, reqHs),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													),
												funcs.CollectBusReqCtxs,
											).
											Do(
												funcs.ReduceBusReqCtxsToBusRsp(f),
												funcs.ReduceBusReqCtxsErrsToBusRspErrs(f),
												funcs.ReduceSvcRspPaginationsToCliRspPagination(f),
											),
									},
								),
						},
					).
					If(
						func(ctx types.ReqCtx) bool {
							gSlice, ok := ctx.GBusRsp.GenericSlice(ctx.ForTypeNode.PluralFieldName())
							if !ok {
								return false
							}

							return len(gSlice.Get()) > c.Internal.Get.MaxResults
						},
						line.Do(
							funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
								gSlice := ctx.GBusRsp.MustGenericSlice(ctx.ForTypeNode.PluralFieldName())

								preLen := len(gSlice.Get())

								gSlice.Set(gSlice.Get()[:c.Internal.Get.MaxResults])

								postLen := len(gSlice.Get())

								ctx.GBusRsp.MustSetGenericSlice([]string{ctx.ForTypeNode.PluralFieldName()}, gSlice)

								gWarnings := f.MustFromStructs([]mql.Warning{
									{
										Message: mql.String(fmt.Sprintf("trimmed from %v to %v", preLen, postLen)),
									},
								})

								gCliRspWarnings, ok := ctx.GBusRsp.GenericSlice(fieldnames.Warnings)
								if ok {
									gWarnings.Append(gCliRspWarnings.Get()...)
								}

								ctx.GBusRsp.MustSetGenericSlice([]string{fieldnames.Warnings}, gWarnings)

								return ctx
							}),
						),
					).
					Parallel(
						c.Internal.Get.ResolveById.Concurrency,
						funcs.Map(types.GBusRsp, types.GEntity),
						line.If(funcs.EntityOnlyContainsServiceId, line.Do(funcs.GetEntityById(f, resolveLine))),
						funcs.Collect(types.GEntity, types.GBusRsp),
					).
					Do(
						funcs.HardFilterGCliRsp(),
					).
					If(
						funcs.Isset(types.GCliReq, []string{fieldnames.Relations}, true),
						line.
							Parallel(
								-1,
								funcs.Map(types.GBusRsp, types.GEntity),
								line.Do(funcs.ResolveRelations(resolveLine, f)),
								funcs.Collect(types.GEntity, types.GBusRsp),
							),
					),
			},
		)

	return resolveLine
}

func getErrLine(f generic.Factory, to string) *line.Line {
	l := line.ErrLine()

	l.
		If(
			funcs.Isset(to, nil, false),
			line.Do(funcs.New(f, to)),
		).
		Do(
			funcs.Copy(types.Svc, types.Errs),
			funcs.Move(f, types.Errs, to),
		)

	return l
}

func NarrowSvcFilterToModeId(l *line.Line) *line.Line {
	return l.
		Switch(funcs.By(types.Method),
			map[string]*line.Line{
				mql.Methods.Get: line.
					Switch(
						funcs.By(types.Mode),
						map[string]*line.Line{
							mql.GetModeKind.Id:       line.Do(funcs.SetSvcFilterToGetModeIdSvcIdFunc()),
							mql.GetModeKind.Relation: line.Do(funcs.SetSvcFilterToGetModeRelationIdFunc()),
						},
					),
			},
		)
}

func SetDefaults(f generic.Factory) func(l *line.Line) *line.Line {
	gGetMode := f.MustFromStruct(mql.GetMode{
		Kind:       &mql.GetModeKind.Collection,
		Collection: &mql.CollectionGetMode{},
	})

	return func(l *line.Line) *line.Line {
		return l.
			Switch(
				funcs.By(types.Method),
				map[string]*line.Line{
					mql.Methods.Get: line.Do(funcs.SetDefaultSelect()),
				},
			).
			If(
				funcs.Isset(types.GCliReq, []string{fieldnames.Mode}, false),
				line.Switch(
					funcs.By(types.Method),
					map[string]*line.Line{
						mql.Methods.Get: line.Do(
							funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
								ctx.GCliReq.MustSetGeneric([]string{fieldnames.Mode}, gGetMode)

								return ctx
							}),
						),
					},
				),
			)
	}
}

func LookupServices(f generic.Factory, hs map[bool]types.RequestHandler, logTemplates types.InternalLogTemplates) (l *line.Line) {
	return line.Parallel(
		-1,
		funcs.Map(types.GSvcRsp, types.GEntity),
		line.
			Do(
				funcs.Set(f, types.GBusReq, mql.LookupServiceBusRequest{}),
				funcs.Copy(types.GEntity, types.Svc),
				funcs.Log(config.SvcReq, logTemplates),
				funcs.HandleSvcReq(hs),
				funcs.Log(config.SvcRsp, logTemplates),
				funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
					gEndpoints, ok := ctx.GSvcRsp.Generic(fieldnames.Output, fieldnames.Service, fieldnames.Endpoints)
					if ok {
						ctx.GEntity.MustSetGeneric([]string{fieldnames.Endpoints}, gEndpoints)
					}

					name, ok := ctx.GSvcRsp.String(fieldnames.Output, fieldnames.Service, fieldnames.Name)
					if ok {
						ctx.GEntity.MustSetString([]string{fieldnames.Name}, name)
					}

					sdkVersion, ok := ctx.GSvcRsp.String(fieldnames.Output, fieldnames.Service, fieldnames.SdkVersion)
					if ok {
						ctx.GEntity.MustSetString([]string{fieldnames.SdkVersion}, sdkVersion)
					}

					return ctx
				}),
			),
		funcs.Collect(types.GEntity, types.GSvcRsp),
	)
}
