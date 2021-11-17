package types

import (
	"github.com/metamatex/metamate/gen/v0/mql"
)

type Config struct {
	DiscoverySvc    mql.Service                `yaml:"discoverySvc,omitempty"`
	Endpoints       EndpointsConfig            `yaml:"endpoints,omitempty"`
	Host            HostConfig                 `yaml:"host,omitempty"`
	Log             LogConfig                  `yaml:"log,omitempty"`
	Embedded        EmbeddedConfig             `yaml:"virtual,omitempty"`
	Internal        InternalConfig             `yaml:"internal,omitempty"`
	ServiceAccounts []ServiceAccountAssignment `yaml:"serviceAccounts,omitempty"`
}

type ServiceAccountAssignment struct {
	ServiceId      mql.ServiceId
	ServiceAccount mql.ServiceAccount
}

type EmbeddedConfig struct {
	Services []EmbeddedSvc
}

type GetConfig struct {
	MaxResults  int
	ResolveById ResolveByIdConfig
}

type InternalConfig struct {
	Get GetConfig
}

type ResolveByIdConfig struct {
	Concurrency int
}

type LogConfig struct {
	Http     bool
	Internal InternalLogConfig
}

// stage, type, format
type InternalLogConfig map[string]map[string]string

type EndpointsConfig struct {
	Config           ConfigEndpointConfig           `yaml:"config,omitempty"`
	Prometheus       PrometheusEndpointConfig       `yaml:"prometheus,omitempty"`
	Debug            DebugEndpointConfig            `yaml:"debug,omitempty"`
	Graphql          GraphqlEndpointConfig          `yaml:"graphql,omitempty"`
	GraphiqlExplorer GraphiqlExplorerEndpointConfig `yaml:"graphiqlExplorer,omitempty"`
	HttpJson         HttpJsonEndpoint               `yaml:"httpJson,omitempty"`
}

type ConfigEndpointConfig struct {
	On bool `yaml:"on,omitempty"`
}

type PrometheusEndpointConfig struct {
	On bool `yaml:"on,omitempty"`
}

type DebugEndpointConfig struct {
	On bool `yaml:"on,omitempty"`
}

type HostConfig struct {
	AllowedOrigins      []string        `yaml:"allowedOrigins,omitempty"`
	ReadTimeoutSeconds  int             `yaml:"readTimeoutSeconds,omitempty"`
	WriteTimeoutSeconds int             `yaml:"writeTimeoutSeconds,omitempty"`
	BasicAuth           BasicAuthConfig `yaml:"basicAuth,omitempty"`
}

type BasicAuthConfig struct {
	User     string `yaml:"user,omitempty"`
	Password string `yaml:"password,omitempty"`
}

type GraphiqlExplorerEndpointConfig struct {
	On           bool   `yaml:"on,omitempty"`
	DefaultQuery string `yaml:"defaultQuery,omitempty"`
}

type GraphqlEndpointConfig struct {
	On bool `yaml:"on,omitempty"`
}

type HttpJsonEndpoint struct {
	On bool `yaml:"on,omitempty"`
}
