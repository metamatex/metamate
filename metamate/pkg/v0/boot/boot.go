package boot

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/pipeline"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	httpjsonHandler "github.com/metamatex/metamate/metamate/pkg/v0/communication/clients/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/servers/admin"
	configServer "github.com/metamatex/metamate/metamate/pkg/v0/communication/servers/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/servers/explorer"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/servers/graphql"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/servers/index"
	"github.com/metamatex/metamate/metamate/pkg/v0/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/persistence"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"log"
	"net/http"
	"net/http/pprof"
	"text/template"
)

func bootVirtualCluster(rn *graph.RootNode, f generic.Factory) (c *virtual.Cluster, err error) {
	c = virtual.NewCluster(rn, f, func(err error) {
		panic(err)
	})

	//err = c.HostSvc(types.VirtualSvc{Id: virtual.Pipe, Name: virtual.Pipe})
	//if err != nil {
	//	return
	//}
	//
	//err = c.HostSvc(types.VirtualSvc{Id: virtual.ReqFilter, Name: virtual.ReqFilter})
	//if err != nil {
	//	return
	//}

	return
}

func NewDependencies(c types.Config, v types.Version) (d types.Dependencies, err error) {
	d.RootNode, err = asg.New()
	if err != nil {
		return
	}

	d.Factory = generic.NewFactory(d.RootNode)

	d.LinkStore = persistence.NewMemoryLinkStore()

	d.InternalLogTemplates = toInternalLogTemplates(c.Log.Internal)

	c0, err := bootVirtualCluster(d.RootNode, d.Factory)
	if err != nil {
		return
	}

	err = virtual.Deploy(c0, c.Virtual.Services)
	if err != nil {
		return
	}

	client := &http.Client{}
	vclient := &http.Client{
		Transport: c0,
	}

	reqHs := map[bool]map[string]types.RequestHandler{
		true: {
			sdk.ServiceTransport.HttpJson: httpjsonHandler.GetRequestHandler(d.Factory, vclient),
		},
		false: {
			sdk.ServiceTransport.HttpJson: httpjsonHandler.GetRequestHandler(d.Factory, client),
		},
	}

	//d.SvcReqLog = func(ctx types.ReqCtx) {
	//	ctx.
	//
	//	log.Print(*ctx.Svc.Url.Value + " : " + ctx.GSvcReq.Type().Name())
	//}

	d.ResolveLine = pipeline.NewResolveLine(d.RootNode, d.Factory, c.DiscoverySvc, c.AuthSvcFilter, c.DefaultClientAccount, reqHs, d.LinkStore, d.InternalLogTemplates)

	d.ServeFunc = func(ctx context.Context, gCliReq generic.Generic) generic.Generic {
		gCliReq = gCliReq.Copy()

		ctx0 := types.ReqCtx{
			Ctx:                ctx,
			GCliReq:            gCliReq,
			DoCliReqProcessing: true,
			DoCliReqValidation: true,
			DoSetClientAccount: true,
		}

		ctx0 = d.ResolveLine.Transform(ctx0)

		return ctx0.GCliRsp
	}

	err = c0.HostBus("metamate", d.ServeFunc)
	if err != nil {
		return
	}

	if c.Endpoints.Admin.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/admin*", Handler: admin.GetStaticHandler("/admin")})
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/admin", HandlerFunc: admin.MustGetIndexHandlerFunc("/admin")})
	}

	if c.Endpoints.Config.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/config.json", HandlerFunc: configServer.GetJsonConfigHandleFunc(c)})
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/config.yaml", HandlerFunc: configServer.GetYamlConfigHandleFunc(c)})
	}

	if c.Endpoints.Prometheus.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/metrics", Handler: promhttp.Handler()})
	}

	if c.Endpoints.Debug.On {
		d.Routes = append(d.Routes,
			types.Route{Methods: []string{http.MethodGet}, Path: "/debug/pprof/*", HandlerFunc: pprof.Index},
			types.Route{Methods: []string{http.MethodGet}, Path: "/debug/pprof/cmdline", HandlerFunc: pprof.Cmdline},
			types.Route{Methods: []string{http.MethodGet}, Path: "/debug/pprof/profile", HandlerFunc: pprof.Profile},
			types.Route{Methods: []string{http.MethodGet}, Path: "/debug/pprof/symbol", HandlerFunc: pprof.Symbol},
			types.Route{Methods: []string{http.MethodGet}, Path: "/debug/pprof/trace", HandlerFunc: pprof.Trace},
		)
	}

	if c.Endpoints.GraphiqlExplorer.On && c.Endpoints.Graphql.On {

		d.Routes = append(d.Routes,
			types.Route{Methods: []string{http.MethodGet}, Path: config.GraphiqlExplorerPath + "*", Handler: explorer.GetStaticHandler(config.GraphiqlExplorerPath)},
			types.Route{Methods: []string{http.MethodGet}, Path: config.GraphiqlExplorerPath, HandlerFunc: explorer.MustGetIndexHandlerFunc(config.GraphqlPath, config.GraphiqlExplorerPath, c.Endpoints.GraphiqlExplorer.DefaultQuery)},
		)
	}

	if c.Endpoints.Graphql.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet, http.MethodPost, http.MethodOptions}, Path: config.GraphqlPath, Handler: graphql.MustGetHandler(d.RootNode, d.Factory, d.ServeFunc)})
	}

	if c.Endpoints.HttpJson.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodPost}, Path: "/httpjson", Handler: httpjson.NewServer(httpjson.ServerOpts{
			Root:    d.RootNode,
			Factory: d.Factory,
			Handler: d.ServeFunc,
			LogErr: func(err error) {
				log.Print(err)
			},
		})})
	}

	d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/static*", Handler: index.GetStaticHandler()})
	d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/", HandlerFunc: index.GetIndexHandlerFunc(c.Host.HttpPort, d.Routes, v)})

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   c.Host.AllowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler)

	if c.Log.Http {
		router.Use(middleware.Logger)
	}

	for _, r := range d.Routes {
		for _, m := range r.Methods {
			if r.HandlerFunc != nil {
				router.MethodFunc(m, r.Path, r.HandlerFunc)
			} else if r.Handler != nil {
				router.Method(m, r.Path, r.Handler)
			}
		}
	}

	d.Router = router

	return d, nil
}

func toInternalLogTemplates(c types.InternalLogConfig) (t types.InternalLogTemplates) {
	t = types.InternalLogTemplates{}

	for stageName, types0 := range c {
		t[stageName] = map[string]*template.Template{}

		for typeName, p := range types0 {
			t[stageName][typeName] = template.Must(template.New("").Parse(p))
		}
	}

	return
}
