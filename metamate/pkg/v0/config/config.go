package config

import (
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

const (
	CliReq               = "ClientRequest"
	SvcReq               = "ServiceRequest"
	SvcRsp               = "ServiceResponse"
	CliRsp               = "ClientResponse"
	GraphiqlExplorerPath = "/explorer"
	GraphqlPath          = "/graphql"
	ConfigFileName       = "metamate"
	ConfigFileExtension  = "yaml"
	ConfigFile           = ConfigFileName + "." + ConfigFileExtension
)

var DefaultConfig = types.Config{
	Host: types.HostConfig{
		AllowedOrigins:      []string{"*"},
		ReadTimeoutSeconds:  30,
		WriteTimeoutSeconds: 30,
	},
	Log: types.LogConfig{
		Http: true,
	},
	Internal: types.InternalConfig{
		Get: types.GetConfig{
			MaxResults: 2000,
			ResolveById: types.ResolveByIdConfig{
				Concurrency: 100,
			},
		},
	},
	DiscoverySvc: mql.Service{
		Id: &mql.ServiceId{
			Value:       mql.String("discovery"),
			ServiceName: mql.String("metamate"),
		},
		IsVirtual: mql.Bool(true),
		Url: &mql.Url{
			Value: mql.String("http://discovery"),
		},
		Endpoints: &mql.Endpoints{
			LookupService: &mql.LookupServiceEndpoint{},
			GetServices:   &mql.GetServicesEndpoint{},
		},
	},
	Endpoints: types.EndpointsConfig{
		Config: types.ConfigEndpointConfig{
			On: true,
		},
		Prometheus: types.PrometheusEndpointConfig{
			On: true,
		},
		Debug: types.DebugEndpointConfig{
			On: true,
		},
		Graphql: types.GraphqlEndpointConfig{
			On: true,
		},
		GraphiqlExplorer: types.GraphiqlExplorerEndpointConfig{
			On: true,
		},
		HttpJson: types.HttpJsonEndpoint{
			On: true,
		},
	},
	Virtual: types.VirtualConfig{
		Services: []types.VirtualSvc{
			{
				Id:   "mastodon",
				Name: virtual.Mastodon,
				Opts: &types.VirtualSvcOpts{
					Mastodon: &types.MastodonOpts{
						Host:         "https://mastodon.social",
						ClientId:     "tac-RigLyTKxOJoadxRhkKz2qN4kkUal61G-UoFCGHg",
						ClientSecret: "hyx3PLEuTvy-NKFBPGcWutQlphOjAbZOfx6cWPlbBn4",
					},
				},
			},
			{
				Id:   "hackernews",
				Name: virtual.Hackernews,
			},
			{
				Id:   "reddit",
				Name: virtual.Reddit,
				Opts: &types.VirtualSvcOpts{
					Reddit: &types.RedditOpts{
						ClientId:     "5Spu4UHEEVbvsQ",
						ClientSecret: "OzXCIrbPZVbTlgy37YeZDfCiYWQ",
						Username:     "metamatex",
						Password:     "vJ6g3ouQbZ4ztiA",
						UserAgent:    "abc",
					},
				},
			},
		},
	},
}
