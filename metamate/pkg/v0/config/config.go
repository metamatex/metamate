package config

import (
	"github.com/metamatex/metamate/gen/v0/mql"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/embedded"
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
		BasicAuth: types.BasicAuthConfig{
			User:     "user",
			Password: "password",
		},
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
			Value:       mql.String("embedded"),
			ServiceName: mql.String("metamate"),
		},
		IsEmbedded: mql.Bool(true),
		Url: &mql.Url{
			Value: mql.String("http://discovery"),
		},
		Endpoints: &mql.ServiceEndpoints{
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
	Embedded: types.EmbeddedConfig{
		Services: []types.EmbeddedSvc{
			{
				Id:   "mastodon",
				Name: embedded.Mastodon,
			},
			{
				Id:   "hackernews",
				Name: embedded.Hackernews,
			},
			{
				Id:   "reddit",
				Name: embedded.Reddit,
			},
		},
	},
	ServiceAccounts: []types.ServiceAccountAssignment{
		{
			ServiceId: mql.ServiceId{
				ServiceName: mql.String("embedded"),
				Value: mql.String("reddit"),
			},
			ServiceAccount: mql.ServiceAccount{
				ClientId: mql.String("w46bvIsqpm1ZrA"),
				ClientSecret: mql.String("k68l2YepSu-eJA3IjVVJivycbaU"),
				Username: mql.String("metamatex"),
				Password: mql.String("p3VhubNQEbc9A7P"),
			},
		},
	},
}
