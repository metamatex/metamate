package graph

import (
	"github.com/iancoleman/strcase"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/utils"
	"reflect"
)

const (
	TYPE = "type"
)

type TypeNode struct {
	NodeTrait `yaml:"inline,omitempty" json:",omitempty"`
	Edges     TypeNodeEdges `yaml:",omitempty" json:"edges,omitempty"`
}

func (n *TypeNode) GetEdges() interface{} {
	return n.Edges
}

func NewTypeNode() (*TypeNode) {
	return &TypeNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags:    Flags{},
				defaults: typeflags.Defaults,
			},
		},
	}
}

func (n *TypeNode) Init(rn *RootNode, name string, fns FieldNodeSlice, additional []interface{}) {
	propagateNode(rn, n, name, additional)

	n.Edges = NewTypeNodeEdges(n)

	n.AddFieldNodeSlice(fns)

	rn.Types.Add(n)
}

func (n *TypeNode) AddFieldNodeSlice(fns FieldNodeSlice) {
	fns.init(n.root, n.Id())
	fns.Map().Each(func(fn *FieldNode) {
		fn.Edges.Type.Resolver.SetHeldBy(n.Id())
	})
	n.root.Fields.Add(fns...)

	n.Edges.Fields.Resolver.AddHolds(fns.GetIds()...)
}

func (n *TypeNode) AddFieldNodes(fns ...*FieldNode) {
	n.AddFieldNodeSlice( FieldNodeSlice(fns))
}

func (n *TypeNode) Type() string {
	return TYPE
}

func (n *TypeNode) Validate() (errs []error) {
	errs = requireHasNameAndLowerCaseId(n)

	err := requireNameIsUpperCamelCase(n)
	if err != nil {
		errs = append(errs, err)
	}

	errs = append(errs, validateTypeFlags(n)...)

	return
}

func validateTypeFlags(n *TypeNode) (errs []error) {
	for k, v := range n.flags.Flags {
		if !v {
			continue
		}

		switch k {
		case typeflags.IsScalar:
			errs = append(errs, validateScalar(n)...)
		case typeflags.IsValue:
			errs = append(errs, validateValue(n)...)
		case typeflags.IsRange:
			errs = append(errs, validateRange(n)...)
		case typeflags.IsRatio:
			errs = append(errs, validateRatio(n)...)
		}
	}

	return
}

func (n TypeNode) HasAnyFieldFlag(flag string, b bool) (bool) {
	for _, n := range n.Edges.Fields.Holds() {
		if n.Flags().Is(flag, b) {
			return true
		}
	}

	return false
}

func (n *TypeNode) Print() () {
	Print(printFieldNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}

func (n *TypeNode) PluralFieldName() (string) {
	return strcase.ToLowerCamel(utils.Plural(n.Name()))
}
