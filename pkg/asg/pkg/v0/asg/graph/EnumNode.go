package graph

import (
	"reflect"

	"github.com/metamatex/asg/pkg/v0/asg/graph/enumflags"
)

const (
	ENUM = "enum"
)

type EnumNode struct {
	NodeTrait `yaml:"inline,omitempty" json:",omitempty"`
	Data      EnumNodeData  `yaml:",omitempty" json:"data,omitempty"`
	Edges     EnumNodeEdges `yaml:",omitempty" json:"edges,omitempty"`
}

func NewEnumNode() (*EnumNode) {
	return &EnumNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags:    Flags{},
				defaults: enumflags.Defaults,
			},
		},
	}
}

func (n *EnumNode) Init(rn *RootNode, name string, values []string, additional []interface{}) {
	propagateNode(rn, n, name, additional)

	n.Data.Values = values
	n.Edges = NewEnumNodeEdges(n)

	rn.Enums.Add(n)
}

func (n *EnumNode) GetData() interface{} {
	return n.Data
}

func (n *EnumNode) Type() string {
	return ENUM
}

func (n *EnumNode) Validate() (errs []error) {
	errs = requireHasNameAndLowerCaseId(n)

	err := requireHasValues(n)
	if err != nil {
		errs = append(errs, err)
	}

	errs0 := requireUniqueValues(n)
	if len(errs0) != 0 {
		errs = append(errs, errs...)
	}

	return
}

type printEnumNode struct {
	Type      string
	PrintNode `yaml:",inline"`
}

func (n *EnumNode) Print() () {
	Print(printEnumNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}
