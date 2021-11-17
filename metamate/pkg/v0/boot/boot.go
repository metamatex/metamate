package boot

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/embedded"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/pipeline"
	httpjsonHandler "github.com/metamatex/metamate/metamate/pkg/v0/communication/clients/httpjson"
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
	"time"
)

func NewDependencies(c types.Config, v types.Version) (d types.Dependencies, err error) {
	d.RootNode, err = asg.New()
	if err != nil {
		return
	}

	if c.DiscoverySvc.Endpoints == nil {
		c.DiscoverySvc.Endpoints = &mql.ServiceEndpoints{}
	}

	if c.DiscoverySvc.Endpoints.GetServices == nil {
		c.DiscoverySvc.Endpoints.GetServices = &mql.GetServicesEndpoint{}
	}

	if c.DiscoverySvc.Endpoints.LookupService == nil {
		c.DiscoverySvc.Endpoints.LookupService = &mql.LookupServiceEndpoint{}
	}

	d.Factory = generic.NewFactory(d.RootNode)

	d.InternalLogTemplates = toInternalLogTemplates(c.Log.Internal)

	c0 := embedded.NewCluster(d.RootNode, d.Factory, func(err error) {
		panic(err)
	})

	err = embedded.Deploy(c0, c.Embedded.Services)
	if err != nil {
		return
	}

	client := &http.Client{}
	vclient := &http.Client{
		Transport: c0,
	}

	cache, err := persistence.NewLruCache(100, 30*time.Minute)
	if err != nil {
		return
	}

	cacheHandlerF := func(h func(ctx context.Context, addr string, gBusReq generic.Generic) (gSvcRsp generic.Generic, err error)) func(ctx context.Context, addr string, gBusReq generic.Generic) (gSvcRsp generic.Generic, err error) {
		return func(ctx context.Context, addr string, gBusReq generic.Generic) (gSvcRsp generic.Generic, err error) {
			gBusReq0 := gBusReq.Copy()
			gBusReq0.Delete(fieldnames.Select)
			key := addr + gBusReq0.GetHash()

			v, ok := cache.Get(key)
			switch ok {
			case true:
				gCachedSvcRsp, ok := v.(generic.Generic)
				if !ok {
					panic("expected generic")
				}
				gSvcRsp = gCachedSvcRsp.Copy()
			case false:
				gSvcRsp, err = h(ctx, addr, gBusReq)
				if err != nil {
					return
				}

				cache.Add(key, gSvcRsp.Copy())
			}

			return
		}
	}

	reqHs := map[bool]types.RequestHandler{
		true:  httpjsonHandler.GetRequestHandler(d.Factory, vclient),
		false: httpjsonHandler.GetRequestHandler(d.Factory, client),
	}

	cachedReqHs := map[bool]types.RequestHandler{
		true:  cacheHandlerF(httpjsonHandler.GetRequestHandler(d.Factory, vclient)),
		false: cacheHandlerF(httpjsonHandler.GetRequestHandler(d.Factory, client)),
	}

	d.ResolveLine = pipeline.NewResolveLine(d.RootNode, d.Factory, c, d.InternalLogTemplates, reqHs, cachedReqHs)

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

		return ctx0.GBusRsp
	}

	err = c0.HostBus("metamate", d.ServeFunc)
	if err != nil {
		return
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
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodPost}, Path: "/httpjson", Handler: generic.NewServer(generic.ServerOpts{
			Root:    d.RootNode,
			Factory: d.Factory,
			Handler: d.ServeFunc,
			LogErr: func(err error) {
				log.Print(err)
			},
			Name: "metamate",
			Log: func(name string, b []byte, req *http.Request) {
				log.Printf("%v - %v: %v", name, req.Header.Get(generic.AsgTypeHeader), string(b))
			},
		})})
	}

	d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/static*", Handler: index.GetStaticHandler()})
	d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/", HandlerFunc: index.GetIndexHandlerFunc(d.Routes, v)})

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

	if c.Host.BasicAuth.User != "" && c.Host.BasicAuth.Password != "" {
		router.Use(middleware.BasicAuth("global", map[string]string{c.Host.BasicAuth.User: c.Host.BasicAuth.Password}))
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

	httpServer := &http.Server{
		ReadTimeout:  time.Duration(c.Host.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(c.Host.WriteTimeoutSeconds) * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      router,
		Addr:         ":80",
	}

	d.Run = append(d.Run, httpServer.ListenAndServe)

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

type HttpsRedirectHandler struct{}

func (h HttpsRedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newURI := "https://" + r.Host + r.URL.String()
	http.Redirect(w, r, newURI, http.StatusFound)
}
