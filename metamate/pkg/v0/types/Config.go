package types

import (
	"github.com/metamatex/metamatemono/gen/v0/sdk"
)

type Config struct {
	DiscoverySvc         sdk.Service       `yaml:"discoverySvc"`
	AuthSvcFilter        sdk.ServiceFilter `yaml:"authSvc"`
	DefaultClientAccount sdk.ClientAccount `yaml:"defaultClientAccount"`
	Endpoints            EndpointsConfig   `yaml:"endpoints"`
	Host                 HostConfig        `yaml:"host"`
	Log                  LogConfig         `yaml:"log"`
	Virtual              VirtualConfig
}

type VirtualConfig struct {
	Services []VirtualSvcOpts
}

type LogConfig struct {
	Http bool
}

type EndpointsConfig struct {
	Admin            AdminEndpointConfig            `yaml:"config"`
	Config           ConfigEndpointConfig           `yaml:"config"`
	Prometheus       PrometheusEndpointConfig       `yaml:"prometheus"`
	Debug            DebugEndpointConfig            `yaml:"debug"`
	Graphql          GraphqlEndpointConfig          `yaml:"graphql"`
	GraphiqlExplorer GraphiqlExplorerEndpointConfig `yaml:"graphiqlExplorer"`
	HttpJson         HttpJsonEndpoint               `yaml:"httpJson"`
}

type ConfigEndpointConfig struct {
	On bool `yaml:"on"`
}

type PrometheusEndpointConfig struct {
	On   bool   `yaml:"on"`
	Path string `yaml:"path"`
}

type DebugEndpointConfig struct {
	On bool `yaml:"on"`
}

type HostConfig struct {
	Bind     string `yaml:"bind"`
	HttpPort int    `yaml:"httpPort"`
}

type GraphiqlExplorerEndpointConfig struct {
	On           bool   `yaml:"on"`
	Path         string `yaml:"path"`
	DefaultQuery string `yaml:"defaultQuery"`
}

type GraphqlEndpointConfig struct {
	On             bool     `yaml:"on"`
	Path           string   `yaml:"path"`
	PlaygroundPath string   `yaml:"playgroundPath"`
	AllowedOrigins []string `yaml:"allowedOrigins"`
}

type AdminEndpointConfig struct {
	On bool `yaml:"on"`
}

type HttpJsonEndpoint struct {
	On   bool   `yaml:"on"`
	Path string `yaml:"path"`
}
