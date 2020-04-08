package config

import (
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

const (
	CliReq = "ClientRequest"
	SvcReq = "ServiceRequest"
	SvcRsp = "ServiceResponse"
	CliRsp = "ClientResponse"
	GraphiqlExplorerPath = "/explorer"
	GraphqlPath = "/graphql"
	ConfigFileName = "metamate"
	ConfigFileExtension = "yaml"
	ConfigFile = ConfigFileName +  "." + ConfigFileExtension
)

var DefaultConfig = types.Config{
	Host: types.HostConfig{
		AllowedOrigins: []string{"*"},
	},
	Log: types.LogConfig{
		Http: true,
	},
	DiscoverySvc: sdk.Service{
		Id: &sdk.ServiceId{
			Value:       sdk.String("discovery"),
			ServiceName: sdk.String("metamate"),
		},
		IsVirtual: sdk.Bool(true),
		Url: &sdk.Url{
			Value: sdk.String("http://discovery"),
		},
		Transport: &sdk.ServiceTransport.HttpJson,
		Endpoints: &sdk.Endpoints{
			LookupService: &sdk.LookupServiceEndpoint{},
			GetServices:   &sdk.GetServicesEndpoint{},
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
		},
	},
}
