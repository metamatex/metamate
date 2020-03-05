package graph

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph/basictypeflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/asg/pkg/v0/asg/typenames"
	"reflect"
)

const (
	ROOT      = "root"
	FieldKind = "kind"
)

type RootNode struct {
	NodeTrait  `yaml:"inline,omitempty" json:",omitempty"`
	BasicTypes BasicTypeNodeMap `yaml:",omitempty" json:"basic_types,omitempty"`
	Endpoints  EndpointNodeMap  `yaml:",omitempty" json:"endpoints,omitempty"`
	Enums      EnumNodeMap      `yaml:",omitempty" json:"enums,omitempty"`
	Paths      PathNodeMap      `yaml:",omitempty" json:"paths,omitempty"`
	Fields     FieldNodeMap     `yaml:",omitempty" json:"fields,omitempty"`
	Relations  RelationNodeMap  `yaml:",omitempty" json:"relations,omitempty"`
	Types      TypeNodeMap      `yaml:",omitempty" json:"Types,omitempty"`
	wired      bool
}

func NewRootNode() (rn *RootNode) {
	rn = &RootNode{
		NodeTrait:  NodeTrait{},
		BasicTypes: BasicTypeNodeMap{},
		Endpoints:  EndpointNodeMap{},
		Enums:      EnumNodeMap{},
		Fields:     FieldNodeMap{},
		Relations:  RelationNodeMap{},
		Paths:      PathNodeMap{},
		Types:      TypeNodeMap{},
	}

	rn.SetName("root")

	rn.BasicTypes.Add(
		&BasicTypeNode{NodeTrait: NodeTrait{Idx: STRING, Namex: string(STRING), flags: FlagsContainer{
			Flags:    Flags{},
			defaults: basictypeflags.Defaults,
		}}},
		&BasicTypeNode{NodeTrait: NodeTrait{Idx: INT32, Namex: string(INT32), flags: FlagsContainer{
			Flags:    Flags{},
			defaults: basictypeflags.Defaults,
		}}},
		&BasicTypeNode{NodeTrait: NodeTrait{Idx: FLOAT64, Namex: string(FLOAT64), flags: FlagsContainer{
			Flags:    Flags{},
			defaults: basictypeflags.Defaults,
		}}},
		&BasicTypeNode{NodeTrait: NodeTrait{Idx: BOOL, Namex: string(BOOL), flags: FlagsContainer{
			Flags:    Flags{},
			defaults: basictypeflags.Defaults,
		}}},
	)

	return
}

func (rn *RootNode) GetNodes(include ...string) (nm NodeMap) {
	nms := []NodeMap{}

	nm = NodeMap{}

	includeMap := map[string]bool{}
	for _, s := range include {
		includeMap[s] = true
	}

	if len(includeMap) == 0 {
		includeMap[BASIC_TYPE] = true
		includeMap[ENDPOINT] = true
		includeMap[ENUM] = true
		includeMap[FIELD] = true
		includeMap[RELATION] = true
		includeMap[PATH] = true
		includeMap[TYPE] = true
	}

	for s, _ := range includeMap {
		switch s {
		case BASIC_TYPE:
			nms = append(nms, rn.BasicTypes.ToNodeMap())
			break
		case ENDPOINT:
			nms = append(nms, rn.Endpoints.ToNodeMap())
			break
		case ENUM:
			nms = append(nms, rn.Enums.ToNodeMap())
			break
		case FIELD:
			nms = append(nms, rn.Fields.ToNodeMap())
			break
		case RELATION:
			nms = append(nms, rn.Relations.ToNodeMap())
			break
		case PATH:
			nms = append(nms, rn.Paths.ToNodeMap())
			break
		case TYPE:
			nms = append(nms, rn.Types.ToNodeMap())
			break
		}
	}

	for _, nm0 := range nms {
		for _, n := range nm0 {
			nm.Add(n)
		}
	}

	return
}

func (rn *RootNode) Wire() {
	if rn.wired {
		panic("already wired")
	}
	defer func() {
		rn.wired = true
	}()

	for _, n := range rn.BasicTypes {
		n.SetRoot(rn)
	}

	for _, n := range rn.Endpoints {
		n.SetRoot(rn)
	}

	for _, n := range rn.Enums {
		n.SetRoot(rn)
	}

	for _, n := range rn.Fields {
		n.SetRoot(rn)
	}

	for _, n := range rn.Relations {
		n.SetRoot(rn)
	}

	for _, n := range rn.Paths {
		n.SetRoot(rn)
	}

	for _, n := range rn.Types {
		n.SetRoot(rn)
	}

	for _, n := range rn.BasicTypes {
		n.Wire()
	}

	for _, n := range rn.Endpoints {
		n.Wire()
	}

	for _, n := range rn.Enums {
		n.Wire()
	}

	for _, n := range rn.Fields {
		n.Wire()
	}

	for _, n := range rn.Relations {
		n.Wire()
	}

	for _, n := range rn.Paths {
		n.Wire()
	}

	for _, n := range rn.Types {
		n.Wire()
	}
}

func (rn *RootNode) Type() string {
	return ROOT
}

func (rn *RootNode) Validate() (errs []error) {
	for _, n := range rn.BasicTypes {
		errs = append(errs, n.Validate()...)
	}

	for _, n := range rn.Endpoints {
		errs = append(errs, n.Validate()...)
	}

	for _, n := range rn.Enums {
		errs = append(errs, n.Validate()...)
	}

	for _, n := range rn.Fields {
		errs = append(errs, n.Validate()...)
	}

	for _, n := range rn.Relations {
		errs = append(errs, n.Validate()...)
	}

	for _, n := range rn.Types {
		errs = append(errs, n.Validate()...)
	}

	for _, n := range rn.BasicTypes {
		errs = append(errs, n.Validate()...)
	}

	return
}

func (rn *RootNode) AddTypeNode(name string, fieldNodeSlice FieldNodeSlice, additional ...interface{}) (tn *TypeNode) {
	tn = NewTypeNode()

	tn.Init(rn, name, fieldNodeSlice, additional)

	return
}

func (rn *RootNode) AddEnumNode(name string, values []string, additional ...interface{}) (en *EnumNode) {
	en = NewEnumNode()

	en.Init(rn, name, values, additional)

	return
}

func (rn *RootNode) AddEndpointNode(name, method, forName, requestName, responseName string, additional ...interface{}) (en *EndpointNode) {
	en = NewEndpointNode()

	en.Init(rn, name, method, forName, requestName, responseName, additional)

	return
}

func (rn *RootNode) AddRelationNode(activePath RelationPath, passivePath RelationPath) (rn0 *RelationNode) {
	rn0 = NewRelationNode()

	rn0.Init(rn, activePath, passivePath)

	return
}

func propagateNode(rn *RootNode, n Node, name string, additional []interface{}) () {
	flags := Flags{}

	for _, any := range additional {
		switch sth := any.(type) {
		case Flags:
			flags = sth
			break
		default:
			panic(fmt.Sprintf("additional %v not supported for %v %v", reflect.TypeOf(any), reflect.TypeOf(n), name))
		}
	}

	for k, v := range flags {
		n.Flags().Set(k, v)
	}

	n.SetRoot(rn)
	n.SetId(ToNodeId(name))
	n.SetName(name)
}

func (rn *RootNode) AddUnion(name string, fields []interface{}, additional ...interface{}) (tn *TypeNode, en *EnumNode) {
	values := []string{}
	for _, f := range fields {
		switch sth := f.(type) {
		case *FieldNode:
			values = append(values, sth.Name())
		case string:
			values = append(values, strcase.ToLowerCamel(sth))
		default:
			panic(fmt.Sprintf("union %v fields can only be of type of either string or FieldNode, not %v", name, reflect.TypeOf(f).Elem().Name()))
		}
	}
	en = rn.AddEnumNode(name+"Kind", values)

	fieldNodes := FieldNodeSlice{
		EnumField(FieldKind, en.Name(), Flags{
			fieldflags.ValidateIsSet: true,
		}),
	}

	for _, f := range fields {
		var n *FieldNode

		switch sth := f.(type) {
		case *FieldNode:
			n = sth
			n.Flags().Set(fieldflags.IsUnionField, true)
		case string:
			n = TypeField(strcase.ToLowerCamel(sth), sth, Flags{
				fieldflags.IsUnionField: true,
			})
		default:
			panic("must be either *FieldNode or string")
		}

		fieldNodes = append(fieldNodes, n)
	}

	tn = rn.AddTypeNode(name, fieldNodes, additional...)
	tn.flags.Set(typeflags.IsUnion, true)

	return
}

func (rn *RootNode) AddActionEndpoint(name string, inputFns FieldNodeSlice, outputFns FieldNodeSlice) (tn *TypeNode, en *EnumNode) {
	rn.AddTypeNode(typenames.Input(name), inputFns)

	actionReq := rn.AddTypeNode(typenames.Request(name), FieldNodeSlice{
		TypeField(fieldnames.ServiceFilter, typenames.Filter(typenames.Service), Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		TypeField(fieldnames.Select, typenames.Select(typenames.Response(name)), Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		TypeField(fieldnames.Meta, typenames.RequestMeta),
		TypeField(fieldnames.Input, typenames.Input(name)),
	}, Flags{
		typeflags.IsRequest:       true,
		typeflags.IsActionRequest: true,
	})

	actionReq.Edges.Type.Resolver.SetResponse(ToNodeId(typenames.Response(name)))

	rn.AddTypeNode(typenames.Output(name), outputFns)

	actionRsp := rn.AddTypeNode(typenames.Response(name), FieldNodeSlice{
		TypeField(fieldnames.Meta, typenames.ResponseMeta),
		TypeField(fieldnames.Output, typenames.Output(name)),
	}, Flags{
		typeflags.IsResponse: true,
	})

	actionRsp.Edges.Type.Resolver.SetRequest(ToNodeId(typenames.Response(name)))

	rn.AddEndpointNode(name, MethodAction, "", typenames.Request(name), typenames.Response(name), Flags{
		endpointflags.IsActionEndpoint: true,
	})

	return
}

func (rn *RootNode) Print() {}
