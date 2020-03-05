package graph

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/metamatex/asg/pkg/v0/asg/graph/endpointflags"
)

const (
	ENDPOINT     = "endpoint"
	MethodPost   = "post"
	MethodPipe   = "pipe"
	MethodGet    = "get"
	MethodAction = "action"
	MethodPut    = "put"
	MethodDelete = "delete"
)

type EndpointNode struct {
	NodeTrait `yaml:"inline,omitempty" json:",omitempty"`
	Edges     EndpointNodeEdges `yaml:",omitempty" json:"edges,omitempty"`
	Data      EndpointNodeData  `yaml:",omitempty" json:"data,omitempty"`
}

func (n *EndpointNode) GetEdges() interface{} {
	return n.Edges
}

func (n *EndpointNode) GetData() interface{} {
	return n.Data
}

func NewEndpointNode() (*EndpointNode) {
	return &EndpointNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags:    Flags{},
				defaults: endpointflags.Defaults,
			},
		},
	}
}

func (n *EndpointNode) Init(rn *RootNode, name, method, forName, requestName, responseName string, additional []interface{}) () {
	propagateNode(rn, n, name, additional)

	n.Data.Method = method

	n.Edges = NewEndpointNodeEdges(n)

	n.Edges.Type.Resolver.SetFor(ToNodeId(forName))
	n.Edges.Type.Resolver.SetRequest(ToNodeId(requestName))
	n.Edges.Type.Resolver.SetResponse(ToNodeId(responseName))

	rn.Endpoints.Add(n)
}

func (n *EndpointNode) Type() string {
	return ENDPOINT
}

func (n *EndpointNode) Validate() (errs []error) {
	errs = requireHasNameAndLowerCaseId(n)

	err := requireNameIsUpperCamelCase(n)
	if err != nil {
		errs = append(errs, err)
	}

	//if n.Edges.Type.Resolver.For() == "" {
	//	errs = append(errs, errors.New(fmt.Sprintf("Endpoint node %v Edges.Type.Resolver.For needs to be set", n.Name())))
	//}

	if n.Edges.Type.Resolver.Request() == "" {
		errs = append(errs, errors.New(fmt.Sprintf("Endpoint node %v Edges.Type.Resolver.Request needs to be set", n.Name())))
	}

	if n.Edges.Type.Resolver.Response() == "" {
		errs = append(errs, errors.New(fmt.Sprintf("Endpoint node %v Edges.Type.Resolver.Response needs to be set", n.Name())))
	}

	if n.Data.Method == "" {
		errs = append(errs, errors.New(fmt.Sprintf("Endpoint node %v Method to be set", n.Name())))
	}

	return
}

type printEndpointNode struct {
	Type      string
	PrintNode `yaml:",inline"`
}

func (n *EndpointNode) Print() () {
	Print(printEndpointNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}
