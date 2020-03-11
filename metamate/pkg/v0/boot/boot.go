package boot

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/metamatex/metamate/asg/pkg/v0/asg"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/gen/v0/sdk/utils/ptr"
	"github.com/metamatex/metamate/generic/pkg/v0/generic"
	"github.com/metamatex/metamate/generic/pkg/v0/transport/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/pipeline"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	httpjsonHandler "github.com/metamatex/metamate/metamate/pkg/v0/communication/handler/httpjson"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/server/admin"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/server/config"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/server/explorer"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/server/graphql"
	"github.com/metamatex/metamate/metamate/pkg/v0/communication/server/index"
	"github.com/metamatex/metamate/metamate/pkg/v0/persistence"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"log"
	"net/http"
	"net/http/pprof"
	"strings"
)

func NewBaseConfig() types.Config {
	return types.Config{
		Host: types.HostConfig{
			Bind:     "0.0.0.0",
			HttpPort: 80,
		},
		Log: types.LogConfig{
			Http: true,
		},
		DiscoverySvc: sdk.Service{
			Id: &sdk.ServiceId{
				Value:       ptr.String("discovery"),
				ServiceName: ptr.String("MM"),
			},
			IsVirtual: ptr.Bool(true),
			Url: &sdk.Url{
				Value: ptr.String("http://discovery"),
			},
			Transport: &sdk.ServiceTransport.HttpJson,
			Endpoints: &sdk.Endpoints{
				LookupService: &sdk.LookupServiceEndpoint{},
				GetServices:   &sdk.GetServicesEndpoint{},
			},
		},
		AuthSvcFilter: sdk.ServiceFilter{
			Id: &sdk.ServiceIdFilter{
				Value: &sdk.StringFilter{
					Is: ptr.String("auth"),
				},
			},
		},
		DefaultClientAccount: sdk.ClientAccount{
			Id: &sdk.ServiceId{
				Value: ptr.String("default"),
			},
		},
		Endpoints: types.EndpointsConfig{
			Admin: types.AdminEndpointConfig{
				On: true,
			},
			Config: types.ConfigEndpointConfig{
				On: true,
			},
			Prometheus: types.PrometheusEndpointConfig{
				On:   true,
				Path: "/metrics",
			},
			Debug: types.DebugEndpointConfig{
				On: true,
			},
			Graphql: types.GraphqlEndpointConfig{
				On:             true,
				Path:           "/graphql",
				AllowedOrigins: []string{},
			},
			GraphiqlExplorer: types.GraphiqlExplorerEndpointConfig{
				On:   true,
				Path: "/explorer",
				DefaultQuery: `# hi. this is metamate's interactive graphql interface
# press the play button to run some of the following queries
# no worries, you can't break anything

# alright, let's get started

# a metamate relies on upstream services to query and persist data
# let's quickly check what services are available to this metamate instance
query getServices {
  getServices {
    services {
      id {
        serviceName
        value
      }
      name
      port
      transport
      isVirtual
    }
  }
}

# we would like to write data to some persistent backend
# sqlx services read and write data to a database
# let's narrow them down
query getSqlxServices {
  getServices(
    request: {
      filter: {
        name: {
          is: "sqlx-svc"
        }
      }
    }
  ) {
    services {
      id {
        serviceName
        value
      }
      name
      port
      transport
      isVirtual
    }
  }
}

# we would like to write some random data
# let's create some whatevers
mutation createWhatevers {
  postWhatevers(
    request: {
      whatevers: {
        stringField: "abc"
      }
    }
  ) {
    whatevers {
      stringField
      id {
        serviceName
        value
      }
    }
  }
}

# two whatevers are returned - looking at the serviceName of the the ids
# tells us that the whatevers were created in two different services

# let's create one whatever only in one service
# therefore we need to narrow the handling upstream services down
mutation createLessWhatevers {
  postWhatevers(
    request: {
      whatevers: {
        stringField: "xyz"
      }, 
      serviceFilter: {
        id: {
          value: {
            is: "sqlx-a"
          }
        }
      }
    }
  ) {
    whatevers {
      stringField
      id {
        serviceName
        value
      }
    }
  }
}

# let's quickly look at all the whatevers we just created
query getWhatevers {
  getWhatevers {
    whatevers {
      id {
        serviceName
        value
      }
      stringField
    }
  }
}`,
			},
			HttpJson: types.HttpJsonEndpoint{
				On:   true,
				Path: "/httpjson",
			},
		},
		Virtual: types.VirtualConfig{
			Services: []types.VirtualSvcOpts{
				{
					Name: "auth",
					Auth: &types.AuthOpts{
						PrivateKey: `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDeY78Tls1xmM0QMNbddASMFOvyHkxTkiItSULcaLL4Q4Wr9SxT
5t79OMkj+0DtkKnzqu+aoqL/M09sImY26nMH/uTX3jRwEqx9tfP7j/H8PHPQHZKn
jQbRkNN8Mf6zS6lbWO6mfCaCqZ2D0SmG6T2h4sqmTynvVJGVxZeiLALcTwIDAQAB
AoGBAJIVX6zUgLvALeQW0O3DikEidSMkd+rlsYGiAEOcmwOuBx6//JBYtd4M8UOr
hikHwDwJ6z7e2sdcwy07I32rYEeE0PrOoGfPypRWjZHnpbuXrLTIylEF3czTmXWb
dY1+mLSCaYMsu9uz9CX91Q8YkMAkhoWExQOJZX34641Tup/hAkEA8RnIZyKxjPGe
y2zdZk/utvG7Fvd0DpPEnUehoTEHbyHfdtGQOuyk3EHSya80DfPro0h0oit8bwmL
HEW/VHgT1wJBAOwh9DhinXJFPCC8wNFmYKVwgFc2ImDs/KZARAE7/IstzZIVoeQJ
fpAfCWtQho7vUdzBfTaeR2y6Ai3cJHj97EkCQFfFHRF+rcgzha1kmkzOuIZdBdDc
kKFl5eOj2hFGOgCZAjLNI4Zv86xDQist3vNdYuD0VZFb51a80KmgMoDbnc0CQQCc
93Et7jf1VxrCNFcEm8aREzjtQFoYDlFgfoX2QBb/ueHWQzULrlgIm+kaAjyAVYwY
cDK5FPwrxXZfX+CK4VipAkBcwIKylFV+KKM/c9c85MfimGUGQBcWnfTuYcntaUpq
IGUaBYoax4+5UfqS8Mhk2U5Pr3YiJqY91pCQsys4CrJU
-----END RSA PRIVATE KEY-----`,
						Salt: "abc",
					},
				},
				{
					Name: "sqlx-a",
					Sqlx: &types.SqlxOpts{
						Log:        false,
						Driver:     "sqlite3",
						Connection: ":memory:",
						Types:      []string{typenames.ClientAccount, typenames.ServiceAccount, typenames.BlueWhatever, typenames.Whatever},
					},
				},
				{
					Name: "sqlx-b",
					Sqlx: &types.SqlxOpts{
						Log:        false,
						Driver:     "sqlite3",
						Connection: ":memory:",
						Types:      []string{typenames.ClientAccount, typenames.ServiceAccount, typenames.BlueWhatever, typenames.Whatever},
					},
				},
				{
					Name: "mastodon",
					Mastodon: &types.MastodonOpts{
						Host:         "https://mastodon.social",
						ClientId:     "tac-RigLyTKxOJoadxRhkKz2qN4kkUal61G-UoFCGHg",
						ClientSecret: "hyx3PLEuTvy-NKFBPGcWutQlphOjAbZOfx6cWPlbBn4",
					},
				},
			},
		},
	}
}

func NewProdConfig() (c types.Config) {
	c = NewBaseConfig()

	//c.DiscoverySvc = sdk.Service{
	//	Id: &sdk.ServiceId{
	//		Value:       ptr.String("kubernetes-discovery"),
	//		ServiceName: ptr.String("MM"),
	//	},
	//	Url: &sdk.Url{
	//		Value: ptr.String("http://kubernetes-discovery"),
	//	},
	//	Endpoints: &sdk.Endpoints{
	//		GetServices: &sdk.GetServicesEndpoint{},
	//	},
	//	Transport: &sdk.ServiceTransport.HttpJson,
	//}

	//c.AuthSvcFilter = sdk.Service{}

	return
}

func sanitizeConfig(c types.Config) types.Config {
	if !strings.HasSuffix(c.Endpoints.Graphql.Path, "/") {
		c.Endpoints.Graphql.Path = c.Endpoints.Graphql.Path + "/"
	}

	if !strings.HasSuffix(c.Endpoints.GraphiqlExplorer.Path, "/") {
		c.Endpoints.GraphiqlExplorer.Path = c.Endpoints.GraphiqlExplorer.Path + "/"
	}

	return c
}

func BootVirtualCluster(rn *graph.RootNode, f generic.Factory) (c *virtual.Cluster, err error) {
	c = virtual.NewCluster(rn, f, func(err error) {
		panic(err)
	})

	err = c.HostName(virtual.Pipe, types.VirtualSvcOpts{Name: "pipe"})
	if err != nil {
		return
	}

	err = c.HostName(virtual.ReqFilter, types.VirtualSvcOpts{Name: "reqFilter"})
	if err != nil {
		return
	}

	return
}

func NewDependencies(c types.Config) (d types.Dependencies, err error) {
	d.RootNode, err = asg.New()
	if err != nil {
		return
	}

	d.Factory = generic.NewFactory(d.RootNode)

	c = sanitizeConfig(c)

	d.LinkStore = persistence.NewMemoryLinkStore()

	c0, err := BootVirtualCluster(d.RootNode, d.Factory)
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

	d.SvcReqLog = func(ctx types.ReqCtx) {
		log.Print(*ctx.Svc.Url.Value + " : " + ctx.GSvcReq.Type().Name())
	}

	d.ResolveLine = pipeline.NewResolveLine(d.RootNode, d.Factory, c.DiscoverySvc, c.AuthSvcFilter, c.DefaultClientAccount, reqHs, d.MockHandler, d.LinkStore, d.SvcReqLog)

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
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/config.json", HandlerFunc: config.GetJsonConfigHandleFunc(c)})
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/config.yaml", HandlerFunc: config.GetYamlConfigHandleFunc(c)})
	}

	if c.Endpoints.Prometheus.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: c.Endpoints.Prometheus.Path, Handler: promhttp.Handler()})
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
			types.Route{Methods: []string{http.MethodGet}, Path: c.Endpoints.GraphiqlExplorer.Path + "*", Handler: explorer.GetStaticHandler(c.Endpoints.GraphiqlExplorer.Path)},
			types.Route{Methods: []string{http.MethodGet}, Path: c.Endpoints.GraphiqlExplorer.Path, HandlerFunc: explorer.MustGetIndexHandlerFunc(c.Endpoints.Graphql.Path, c.Endpoints.GraphiqlExplorer.Path, c.Endpoints.GraphiqlExplorer.DefaultQuery)},
		)
	}

	if c.Endpoints.Graphql.On {
		d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet, http.MethodPost, http.MethodOptions}, Path: c.Endpoints.Graphql.Path, Handler: graphql.MustGetHandler(d.RootNode, d.Factory, d.ServeFunc)})
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
	d.Routes = append(d.Routes, types.Route{Methods: []string{http.MethodGet}, Path: "/", HandlerFunc: index.GetIndexHandlerFunc(c.Host.HttpPort, d.Routes)})

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*", "http://localhost:3000"},
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
