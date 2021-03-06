package types

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
)

type RenderContext struct {
	Version    *Version
	Data       map[string]interface{}
	BasicTypes graph.BasicTypeNodeMap
	Endpoints  graph.EndpointNodeMap
	Enums      graph.EnumNodeMap
	Fields     graph.FieldNodeMap
	Relations  graph.RelationNodeMap
	Types      graph.TypeNodeMap
	Paths      graph.PathNodeMap
}

type IterateRenderContext struct {
	Version   *Version
	Data      map[string]interface{}
	BasicType *graph.BasicTypeNode
	Endpoint  *graph.EndpointNode
	Enum      *graph.EnumNode
	Field     *graph.FieldNode
	Relation  *graph.RelationNode
	Type      *graph.TypeNode
	Path      *graph.PathNode
}
