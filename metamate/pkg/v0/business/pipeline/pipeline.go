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

func NewResolveLine(rn *graph.RootNode, f generic.Factory, discoverySvc sdk.Service, authSvcFilter sdk.ServiceFilter, defaultClientAccount sdk.ClientAccount, reqHs map[bool]map[string]types.RequestHandler, linkStore types.LinkStore, logTemplates types.InternalLogTemplates) *line.Line {
	resolveLine := line.Do()

	cliReqErrL := getErrLine(f, types.GCliRsp)

	svcReqErrL := getErrLine(f, types.GSvcRsp)

	resolveLine.
		Error(cliReqErrL, true).
		Do(
			//funcs.Copy(types.GCliReq, types.GCliReq),
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
		If(
			func(ctx types.ReqCtx) bool {
				return true
			},
			//funcs.Isset(types.DoCliReqValidation, true),
			line.Do(funcs.ValidateCliReq(f)),
		).
		Switch(
			funcs.By(types.Method),
			map[string]*line.Line{
				sdk.Methods.Post:   line.Do(funcs.Copy(types.GCliReq, types.Mode)),
				sdk.Methods.Get:    line.Do(funcs.Copy(types.GCliReq, types.Mode)),
				sdk.Methods.Put:    line.Do(funcs.Copy(types.GCliReq, types.Mode)),
				sdk.Methods.Delete: line.Do(funcs.Copy(types.GCliReq, types.Mode)),
			},
		).
		If(
			funcs.Is(types.DoSetClientAccount, true),
			line.If(
				funcs.Isset(types.GCliReq, []string{fieldnames.Auth, fieldnames.Token}, true),
				line.Do(funcs.SetClientAccount(resolveLine, f, authSvcFilter, defaultClientAccount)),
			),
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
		Add(PipeContext(f, resolveLine)).
		Do(funcs.SetStage(config.SvcReq)).
		Switch(
			funcs.By(types.Method),
			map[string]*line.Line{
				sdk.Methods.Pipe: line.
					Do(
						funcs.RequireOneGSvc(),
						funcs.SetFirstGSvc(),
						funcs.Copy(types.GCliReq, types.GSvcReq),
						funcs.Log(config.SvcReq, logTemplates),
						funcs.HandleSvcReq(reqHs),
						funcs.Log(config.SvcRsp, logTemplates),
						funcs.GSvcRspToGCliRsp(),
					),
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
				sdk.Methods.Put: line.
					Switch(
						funcs.By(types.Mode),
						map[string]*line.Line{
							sdk.PutModeKind.Relation: line.
								Do(
									funcs.RequireOneGSvc(),
									funcs.SetFirstGSvc(),
									funcs.Copy(types.GCliReq, types.GSvcReq),
									funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
								).
								Concurrent(
									[]*line.Line{
										line.
											Do(
												funcs.KeepLocalSvcIds(f),
												funcs.Log(config.SvcReq, logTemplates),
												funcs.HandleSvcReq(reqHs),
												funcs.Log(config.SvcRsp, logTemplates),
												funcs.GSvcRspToGCliRsp(),
											),
										line.
											Do(
												funcs.KeepInterSvcIds(f),
												funcs.PutInterSvcIds(rn, f, linkStore),
											),
									},
									funcs.CollectSvcRsps,
								).
								Do(
									funcs.CollectGSvcErrsFromGSvcRsps(f),
									funcs.AddGSvcErrsToGCliRsp(f),
								),
						}),
				line.Default: line.
					Switch(
						funcs.By(types.Method),
						map[string]*line.Line{
							sdk.Methods.Post: line.
								Parallel(
									-1,
									funcs.Map(types.Svcs, types.Svc),
									line.
										Error(svcReqErrL, true).
										Do(
											funcs.Copy(types.GCliReq, types.GSvcReq),
											funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
											funcs.Log(config.SvcReq, logTemplates),
											funcs.HandleSvcReq(reqHs),
											funcs.Log(config.SvcRsp, logTemplates),
											funcs.AddSvcToSvcIds(),
										),
									funcs.CollectSvcRsps,
								),
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
														funcs.HandleSvcReq(reqHs),
														funcs.Log(config.SvcRsp, logTemplates),
														funcs.AddSvcToSvcIds(),
													),
												funcs.CollectSvcRsps,
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
														funcs.HandleSvcReq(reqHs),
														funcs.Log(config.SvcRsp, logTemplates),
														funcs.AddSvcToSvcIds(),
													),
												funcs.CollectSvcRsps,
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
														funcs.HandleSvcReq(reqHs),
														funcs.Log(config.SvcRsp, logTemplates),
														funcs.AddSvcToSvcIds(),
													).
													If(
														funcs.IsType(types.GSvcRsp, sdk.GetServicesResponseName, true),
														FetchSvcDataFromSvcs(f, reqHs, logTemplates),
													),
												funcs.CollectSvcRsps,
											),
										sdk.GetModeKind.Relation: line.
											Do(
												funcs.RequireOneGSvc(),
												funcs.SetFirstGSvc(),
												funcs.Copy(types.GCliReq, types.GSvcReq),
												funcs.Func(func(ctx types.ReqCtx) types.ReqCtx { ctx.GSvcReq.MustDelete(fieldnames.ServiceFilter); return ctx }),
											).
											Concurrent(
												[]*line.Line{
													line.
														Do(
															funcs.Log(config.SvcReq, logTemplates),
															funcs.HandleSvcReq(reqHs),
															funcs.Log(config.SvcRsp, logTemplates),
															funcs.AddSvcToSvcIds(),
														),
													line.
														Do(
															funcs.CollectGRspSelectFromGCliReq(),
															funcs.GetInterSvcIds(rn, linkStore),
														).
														Parallel(
															-1,
															funcs.MapSvcIds,
															line.Do(funcs.GetById(f, resolveLine)),
															funcs.CollectCliRsps,
														).
														Do(
															funcs.CollectGSvcErrsFromGCliRsps(f),
															funcs.CollectGEntitiesFromGCliRsps(f),
															funcs.New(f, types.GSvcRsp),
															funcs.AddGSvcErrsToGSvcRsp(f),
															funcs.AddGEntitiesToGSvcRsp(),
														),
												},
												funcs.CollectSvcRsps,
											),
									},
								),
						},
					).
					Do(
						funcs.ReduceSvcRspsToCliRsp(f),
						funcs.ReduceSvcRspErrsToCliRspErrs(f),
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
		If(
			funcs.Isset(to, []string{fieldnames.Meta}, false),
			line.Do(funcs.AddMeta(f, to)),
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
				sdk.Methods.Put: line.
					Switch(
						funcs.By(types.Mode),
						map[string]*line.Line{
							sdk.PutModeKind.Relation: line.Do(funcs.SetSvcFilterToPutModeRelationIdFunc()),
						},
					),
			},
		)
}

func SetDefaults(f generic.Factory) func(l *line.Line) *line.Line {
	gPostMode := f.MustFromStruct(sdk.PostMode{
		Kind:       &sdk.PostModeKind.Collection,
		Collection: &sdk.CollectionPostMode{},
	})

	gGetMode := f.MustFromStruct(sdk.GetMode{
		Kind:       &sdk.GetModeKind.Collection,
		Collection: &sdk.CollectionGetMode{},
	})

	return func(l *line.Line) *line.Line {
		return l.
			Switch(
				funcs.By(types.Method),
				map[string]*line.Line{
					sdk.Methods.Post: line.Do(funcs.SetDefaultSelect()),
					sdk.Methods.Get:  line.Do(funcs.SetDefaultSelect()),
				},
			).
			If(
				funcs.Isset(types.GCliReq, []string{fieldnames.Mode}, false),
				line.Switch(
					funcs.By(types.Method),
					map[string]*line.Line{
						sdk.Methods.Post: line.Do(
							funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
								ctx.GCliReq.MustSetGeneric([]string{fieldnames.Mode}, gPostMode)

								return ctx
							}),
						),
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

func PipeContext(f generic.Factory, resolveLine *line.Line) func(l *line.Line) *line.Line {
	return func(l *line.Line) *line.Line {
		return l.
			If(
				func(ctx types.ReqCtx) bool {
					method := ctx.GCliReq.Type().Edges.Endpoint.BelongsTo().Data.Method

					if method == sdk.Methods.Pipe || method == sdk.Methods.Action {
						return false
					}

					if ctx.GCliReq.Type().Name() == sdk.GetServicesRequestName {
						return false
					}

					return true
				},
				line.Do(
					funcs.Func(func(ctx types.ReqCtx) types.ReqCtx {
						gCliReq := f.New(ctx.GCliReq.Type().Edges.Type.For().Edges.Type.PipeRequest())

						if ctx.GCliReq != nil {
							gCliReq.MustSetGeneric([]string{fieldnames.Context, ctx.Method, fieldnames.ClientRequest}, ctx.GCliReq)
						}

						if ctx.GSvcReq != nil {
							gCliReq.MustSetGeneric([]string{fieldnames.Context, ctx.Method, fieldnames.ServiceRequest}, ctx.GSvcReq)
						}

						if ctx.GSvcRsp != nil {
							gCliReq.MustSetGeneric([]string{fieldnames.Context, ctx.Method, fieldnames.ServiceResponse}, ctx.GSvcRsp)
						}

						if ctx.GCliRsp != nil {
							gCliReq.MustSetGeneric([]string{fieldnames.Context, ctx.Method, fieldnames.ClientResponse}, ctx.GCliRsp)
						}

						m := sdk.PipeMode{
							Kind: &sdk.PipeModeKind.Context,
							Context: &sdk.ContextPipeMode{
								Method:    &ctx.Method,
								Requester: &sdk.BusActor.Client,
								Stage:     &sdk.RequestStage.Request,
							},
						}

						gCliReq.MustSetGeneric([]string{fieldnames.Mode}, f.MustFromStruct(m))

						ctx0 := resolveLine.Transform(types.ReqCtx{
							GCliReq: gCliReq,
						})

						gErrs, ok := ctx0.GCliRsp.GenericSlice(fieldnames.Meta, fieldnames.Errors)
						if ok {
							var errs []sdk.Error
							gErrs.MustToStructs(&errs)

							var errs0 []sdk.Error
							for i, _ := range errs {
								switch *errs[i].Kind {
								case sdk.ErrorKind.NoServiceMatch:
								default:
									errs0 = append(errs0, sdk.Error{
										Kind:  &sdk.ErrorKind.Pipe,
										Wraps: &errs[i],
									})
								}
							}

							ctx.Errs = errs0
						}

						if ctx.GCliReq != nil {
							g, ok := ctx0.GCliRsp.Generic(fieldnames.Context, ctx.Method, fieldnames.ClientRequest)
							if ok {
								ctx.GCliReq = g
							}
						}

						if ctx.GSvcReq != nil {
							g, ok := ctx0.GCliRsp.Generic(fieldnames.Context, ctx.Method, fieldnames.ServiceRequest)
							if ok {
								ctx.GSvcReq = g
							}
						}

						if ctx.GSvcRsp != nil {
							g, ok := ctx0.GCliRsp.Generic(fieldnames.Context, ctx.Method, fieldnames.ServiceResponse)
							if ok {
								ctx.GSvcRsp = g
							}
						}

						if ctx.GCliRsp != nil {
							g, ok := ctx0.GCliRsp.Generic(fieldnames.Context, ctx.Method, fieldnames.ClientResponse)
							if ok {
								ctx.GCliRsp = g
							}
						}

						return ctx
					}),
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
