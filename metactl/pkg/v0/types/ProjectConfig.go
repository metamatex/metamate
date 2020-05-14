package types

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
)

type ProjectConfig struct {
	V0 V0Project `json:"v0,omitempty" yaml:"v0,omitempty"`
}

type V0Project struct {
	Gen Gen `json:"gen,omitempty" yaml:"gen,omitempty"`
}

type Gen struct {
	Tasks []RenderTask `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Sdks  []SdkConfig  `json:"sdks,omitempty" yaml:"sdks,omitempty"`
}

type SdkConfig struct {
	Names     []string               `json:"names,omitempty" yaml:"names,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty" yaml:"data,omitempty"`
	Endpoints *graph.Filter          `json:"endpoints,omitempty" yaml:"endpoints,omitempty"`
}
