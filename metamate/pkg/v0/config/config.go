package config

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/gen/v0/sdk"
	"github.com/metamatex/metamate/metamate/pkg/v0/business/virtual"
	"github.com/metamatex/metamate/metamate/pkg/v0/types"
)

const GraphiqlExplorerPath = "/explorer"
const GraphqlPath = "/graphql"
const ConfigFileName = "metamate"
const ConfigFileExtension = "yaml"
const ConfigFile = ConfigFileName +  "." + ConfigFileExtension

var DefaultConfig = types.Config{
	Host: types.HostConfig{
		Bind:           "0.0.0.0",
		HttpPort:       80,
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
	AuthSvcFilter: sdk.ServiceFilter{
		Id: &sdk.ServiceIdFilter{
			Value: &sdk.StringFilter{
				Is: sdk.String("auth"),
			},
		},
	},
	DefaultClientAccount: sdk.ClientAccount{
		Id: &sdk.ServiceId{
			Value: sdk.String("default"),
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
				Id:   "auth",
				Name: virtual.Auth,
				Opts: &types.VirtualSvcOpts{
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
			},
			{
				Id:   "sqlx-a",
				Name: virtual.Sqlx,
				Opts: &types.VirtualSvcOpts{
					Sqlx: &types.SqlxOpts{
						Log:        false,
						Driver:     "sqlite3",
						Connection: ":memory:",
						Types:      []string{typenames.ClientAccount, typenames.ServiceAccount, typenames.BlueWhatever, typenames.Whatever},
					},
				},
			},
			{
				Id:   "sqlx-b",
				Name: virtual.Sqlx,
				Opts: &types.VirtualSvcOpts{
					Sqlx: &types.SqlxOpts{
						Log:        false,
						Driver:     "sqlite3",
						Connection: ":memory:",
						Types:      []string{typenames.ClientAccount, typenames.ServiceAccount, typenames.BlueWhatever, typenames.Whatever},
					},
				},
			},
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
