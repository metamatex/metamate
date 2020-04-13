package pipeline

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/funcs"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/line"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

func NewResolveLine(rn *graph.RootNode, f generic.Factory, discoverySvc sdk.Service, reqHs map[bool]map[string]types.RequestHandler, cachedReqHs map[bool]map[string]types.RequestHandler, logTemplates types.InternalLogTemplates) *line.Line {
	resolveLine := line.Do()

	cliReqErrL := getErrLine(f, types.GCliRsp)

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
			funcs.New(f, types.GCliRsp),
		).
		Add(SetDefaults(f)).
		Do(funcs.ValidateCliReq(f)).
		Switch(
			funcs.By(types.Method),
			map[string]*line.Line{
				sdk.Methods.Get:    line.Do(funcs.Copy(types.GCliReq, types.Mode)),
			},
		).
		Add(NarrowSvcFilterToModeId).
		Do(
			funcs.GetSvcs(resolveLine, f, discoverySvc),
			funcs.ApplySvcEndpointReqFilters(f),
		).
		If(
			func(ctx types.ReqCtx) bool {
				return len(ctx.Svcs) == 0
			},
			line.
				Do(
					funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
						ctx.Errs = append(ctx.Errs, funcs.NewError(nil, sdk.ErrorKind.NoServiceMatch, "no services matches"))

						return ctx
					}),
				),
		).
		Do(funcs.SetStage(config.SvcReq)).
		Switch(
			funcs.By(types.Method),
			map[string]*line.Line{
				sdk.Methods.Action: line.
					Do(
						funcs.RequireOneGSvc(),
						funcs.SetFirstGSvc(),
						funcs.Copy(types.GCliReq, types.GSvcReq),
						funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
						funcs.Log(config.SvcReq, logTemplates),
						funcs.HandleSvcReq(reqHs),
						funcs.Log(config.SvcRsp, logTemplates),
						funcs.GSvcRspToGCliRsp(),
					),
				line.Default: line.
					Switch(
						funcs.By(types.Method),
						map[string]*line.Line{
							sdk.Methods.Get: line.New(config.SvcReq).
								Switch(
									funcs.By(types.Mode),
									map[string]*line.Line{
										sdk.GetModeKind.Id: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.Copy(types.GCliReq, types.GSvcReq),
														funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													),
												funcs.CollectSvcRsps,
											).
											Do(
												funcs.ReduceSvcRspsToCliRsp(f),
												funcs.ReduceSvcRspErrsToCliRspErrs(f),
											),
										sdk.GetModeKind.Search: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.Copy(types.GCliReq, types.GSvcReq),
														funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													),
												funcs.CollectSvcRsps,
											).
											Do(
												funcs.ReduceSvcRspsToCliRsp(f),
												funcs.ReduceSvcRspErrsToCliRspErrs(f),
											),
										sdk.GetModeKind.Collection: line.
											Parallel(
												-1,
												funcs.Map(types.Svcs, types.Svc),
												line.
													Error(svcReqErrL, true).
													Do(
														funcs.Copy(types.GCliReq, types.GSvcReq),
														funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
														funcs.Log(config.SvcReq, logTemplates),
														funcs.HandleSvcReq(cachedReqHs),
														funcs.AddSvcToSvcIds(),
														funcs.Log(config.SvcRsp, logTemplates),
													).
													If(
														funcs.IsType(types.GSvcRsp, sdk.GetServicesResponseName, true),
														FetchSvcDataFromSvcs(f, reqHs, logTemplates),
													),
												funcs.CollectSvcRsps,
											).
											Do(
												funcs.ReduceSvcRspsToCliRsp(f),
												funcs.ReduceSvcRspErrsToCliRspErrs(f),
											),
										sdk.GetModeKind.Relation: line.
											Do(
												funcs.RequireOneGSvc(),
												funcs.SetFirstGSvc(),
												funcs.Copy(types.GCliReq, types.GSvcReq),
												funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
												funcs.Log(config.SvcReq, logTemplates),
												funcs.HandleSvcReq(cachedReqHs),
												funcs.AddSvcToSvcIds(),
												funcs.Log(config.SvcRsp, logTemplates),
												funcs.GSvcRspToGCliRsp(),
											),
									},
								),
						},
					).
					Parallel(
						-1,
						funcs.Map(types.GCliRsp, types.GEntity),
						line.If(funcs.EntityOnlyContainsServiceId, line.Do(funcs.GetEntityById(f, resolveLine))),
						funcs.Collect(types.GEntity, types.GCliRsp),
					).
					Do(
						funcs.HardFilterGCliRsp(),
					).
					If(
						funcs.Isset(types.GCliReq, []string{fieldnames.Relations}, true),
						line.
							Parallel(
								-1,
								funcs.Map(types.GCliRsp, types.GEntity),
								line.Do(funcs.ResolveRelations(resolveLine, f)),
								funcs.Collect(types.GEntity, types.GCliRsp),
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
				sdk.Methods.Get: line.
					Switch(
						funcs.By(types.Mode),
						map[string]*line.Line{
							sdk.GetModeKind.Id:       line.Do(funcs.SetSvcFilterToGetModeIdSvcIdFunc()),
							sdk.GetModeKind.Relation: line.Do(funcs.SetSvcFilterToGetModeRelationIdFunc()),
						},
					),
			},
		)
}

func SetDefaults(f generic.Factory) func(l *line.Line) *line.Line {
	gGetMode := f.MustFromStruct(sdk.GetMode{
		Kind:       &sdk.GetModeKind.Collection,
		Collection: &sdk.CollectionGetMode{},
	})

	return func(l *line.Line) *line.Line {
		return l.
			Switch(
				funcs.By(types.Method),
				map[string]*line.Line{
					sdk.Methods.Get:  line.Do(funcs.SetDefaultSelect()),
				},
			).
			If(
				funcs.Isset(types.GCliReq, []string{fieldnames.Mode}, false),
				line.Switch(
					funcs.By(types.Method),
					map[string]*line.Line{
						sdk.Methods.Get: line.Do(
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

func FetchSvcDataFromSvcs(f generic.Factory, hs map[bool]map[string]types.RequestHandler, logTemplates types.InternalLogTemplates) (l *line.Line) {
	return line.Parallel(
		-1,
		funcs.Map(types.GSvcRsp, types.GEntity),
		line.
			Do(
				funcs.Set(f, types.GSvcReq, sdk.LookupServiceRequest{}),
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
