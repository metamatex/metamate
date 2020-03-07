package graph

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/basictypeflags"
)

const (
	STRING     NodeId = "string"
	INT32      NodeId = "int32"
	FLOAT64    NodeId = "float64"
	BOOL       NodeId = "bool"
	BASIC_TYPE        = "basictype"
)

type BasicTypeNode struct {
	NodeTrait `yaml:"inline,omitempty" json:",omitempty"`
}

func NewBasicTypeNode() (*BasicTypeNode) {
	return &BasicTypeNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags:    Flags{},
				defaults: basictypeflags.Defaults,
			},
		},
	}
}


func (n *BasicTypeNode) Type() string {
	return BASIC_TYPE
}

func (n *BasicTypeNode) Validate() (errs []error) {
	errs = requireHasNameAndLowerCaseId(n)

	if n.Id() != STRING &&
		n.Id() != INT32 &&
		n.Id() != FLOAT64 &&
		n.Id() != BOOL {
		errs = append(errs, errors.New(fmt.Sprintf("%v %v Id must be one of %v", reflect.TypeOf(n).Name(), n.Name(), []NodeId{STRING, INT32, FLOAT64, BOOL})))
	}

	return
}

func (n BasicTypeNode) IsStringNode() (bool) {
	return n.Id() == STRING
}

func (n BasicTypeNode) IsInt32Node() (bool) {
	return n.Id() == INT32
}

func (n BasicTypeNode) IsFloat64Node() (bool) {
	return n.Id() == FLOAT64
}

func (n BasicTypeNode) IsBoolNode() (bool) {
	return n.Id() == BOOL
}

type printBasicTypeNode struct {
	Type      string
	PrintNode `yaml:",inline"`
}

func (n *BasicTypeNode) Print() () {
	Print(printBasicTypeNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}
