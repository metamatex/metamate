package types

import "github.com/metamatex/asg/pkg/v0/asg/graph"

type RenderTask struct {
	Template     *string                 `yaml:"template"`
	TemplateData *string                 `yaml:"templateData"`
	Out          *string                 `yaml:"out"`
	Filter       *NodeFilters            `yaml:"select"`
	Dependencies *RenderTaskDependencies `yaml:"dependencies"`
	Iterate      *string                 `yaml:"iterate"`
	Data         map[string]interface{}  `yaml:"data"`
	Reset        *bool                   `yaml:"reset"`
}

type RenderTaskDependencies struct {
	Endpoints *graph.Filter `yaml:"endpoints"`
	Types     *graph.Filter `yaml:"types"`
}

type NodeFilters struct {
	BasicTypes *graph.Filter `yaml:"basicTypes"`
	Endpoints  *graph.Filter `yaml:"endpoints"`
	Enums      *graph.Filter `yaml:"enum"`
	Fields     *graph.Filter `yaml:"fields"`
	Relations  *graph.Filter `yaml:"relations"`
	Types      *graph.Filter `yaml:"types"`
}
