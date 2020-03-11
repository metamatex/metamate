package graph

import (
	"fmt"
	"reflect"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
)

const (
	FIELD = "field"
)

type FieldNode struct {
	NodeTrait `yaml:"inline,omitempty" json:"omitempty"`
	Edges     FieldNodeEdges `yaml:",omitempty" json:"edges,omitempty"`
}

func (n *FieldNode) GetEdges() interface{} {
	return n.Edges
}

func NewFieldNode() (fn *FieldNode) {
	fn = &FieldNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags:    Flags{},
				defaults: fieldflags.Defaults,
			},
		},
	}
	fn.Edges = NewFieldNodeEdges(fn)

	return
}

func (n *FieldNode) Init(rn *RootNode, parentId NodeId) {
	n.SetRoot(rn)
	n.SetId(ToNodeId(fmt.Sprintf("%v_%v", parentId, n.Name())))
}

func (n *FieldNode) Type() string {
	return FIELD
}

func (n *FieldNode) Copy() (fn *FieldNode) {
	fn = &FieldNode{}
	fn.Namex = n.Namex
	fn.flags = n.flags
	fn.Edges = n.Edges

	return
}

func (n *FieldNode) Validate() (errs []error) {
	errs = requireHasNameAndLowerCaseId(n)

	err := requireNameIsLowerCamelCase(n)
	if err != nil {
		errs = append(errs, err)
	}

	return
}

func ListField(n *FieldNode) (*FieldNode) {
	n.flags.Set(fieldflags.IsList, true)

	return n
}

func StringField(name string, additional ...interface{}) (n *FieldNode) {
	n = newFieldNode(name, additional...)

	n.Edges.BasicType.Resolver.SetHolds(ToNodeId(String))

	return
}

func Int32Field(name string, additional ...interface{}) (n *FieldNode) {
	n = newFieldNode(name, additional...)

	n.Edges.BasicType.Resolver.SetHolds(ToNodeId(Int32))

	return
}

func Float64Field(name string, additional ...interface{}) (n *FieldNode) {
	n = newFieldNode(name, additional...)

	n.Edges.BasicType.Resolver.SetHolds(ToNodeId(Float64))

	return
}

func BoolField(name string, additional ...interface{}) (n *FieldNode) {
	n = newFieldNode(name, additional...)

	n.Edges.BasicType.Resolver.SetHolds(ToNodeId(Bool))

	return
}

func EnumField(name string, typeName string, additional ...interface{}) (n *FieldNode) {
	n = newFieldNode(name, additional...)

	n.Edges.Enum.Resolver.SetHolds(ToNodeId(typeName))

	return
}

func TypeField(name string, typeName string, additional ...interface{}) (n *FieldNode) {
	n = newFieldNode(name, additional...)

	n.Edges.Type.Resolver.SetHolds(ToNodeId(typeName))

	return
}

func newFieldNode(name string, additional ...interface{}) (n *FieldNode) {
	flags := Flags{}

	for _, any := range additional {
		switch sth := any.(type) {
		case Flags:
			flags = sth
			break
		default:
			panic(fmt.Sprintf("additional %v not supported for Field %v", reflect.TypeOf(any), name))
		}
	}

	n = NewFieldNode()
	n.Namex = name

	for k, v := range flags {
		n.Flags().Set(k, v)
	}

	return
}

func (n *FieldNode) IsType() bool {
	return !n.flags.Is(fieldflags.IsList, true) && n.Edges.Type.Resolver.Holds() != ""
}

func (n *FieldNode) IsEnum() bool {
	return !n.flags.Is(fieldflags.IsList, true) && n.Edges.Enum.Resolver.Holds() != ""
}

func (n *FieldNode) IsBasicType() bool {
	return !n.flags.Is(fieldflags.IsList, true) && n.Edges.BasicType.Resolver.Holds() != ""
}

func (n *FieldNode) IsString() bool {
	return n.IsBasicType() && n.Edges.BasicType.Holds().IsStringNode()
}

func (n *FieldNode) IsInt32() bool {
	return n.IsBasicType() && n.Edges.BasicType.Holds().IsInt32Node()
}

func (n *FieldNode) IsFloat64() bool {
	return n.IsBasicType() && n.Edges.BasicType.Holds().IsFloat64Node()
}

func (n *FieldNode) IsBool() bool {
	return n.IsBasicType() && n.Edges.BasicType.Holds().IsBoolNode()
}

func (n *FieldNode) IsTypeList() bool {
	return n.flags.Is(fieldflags.IsList, true) && n.Edges.Type.Resolver.Holds() != ""
}

func (n *FieldNode) IsEnumList() bool {
	return n.flags.Is(fieldflags.IsList, true) && n.Edges.Enum.Resolver.Holds() != ""
}

func (n *FieldNode) IsBasicTypeList() bool {
	return n.flags.Is(fieldflags.IsList, true) && n.Edges.BasicType.Resolver.Holds() != ""
}

func (n *FieldNode) IsStringList() bool {
	return n.IsBasicTypeList() && n.Edges.BasicType.Holds().IsStringNode()
}

func (n *FieldNode) IsInt32List() bool {
	return n.IsBasicTypeList() && n.Edges.BasicType.Holds().IsInt32Node()
}

func (n *FieldNode) IsFloat64List() bool {
	return n.IsBasicTypeList() && n.Edges.BasicType.Holds().IsFloat64Node()
}

func (n *FieldNode) IsBoolList() bool {
	return n.IsBasicTypeList() && n.Edges.BasicType.Holds().IsBoolNode()
}

func (n *FieldNode) Kind() string {
	if n.Edges.BasicType.Holds() != nil {
		if n.flags.Is(fieldflags.IsList, true) {
			switch n.Edges.BasicType.Holds().Id() {
			case STRING:
				return FieldKindStringList
			case BOOL:
				return FieldKindBoolList
			case INT32:
				return FieldKindInt32List
			case FLOAT64:
				return FieldKindFloat64List
			}
		} else {
			switch n.Edges.BasicType.Holds().Id() {
			case STRING:
				return FieldKindString
			case BOOL:
				return FieldKindBool
			case INT32:
				return FieldKindInt32
			case FLOAT64:
				return FieldKindFloat64
			}
		}
	} else if n.Edges.Type.Holds() != nil {
		if n.flags.Is(fieldflags.IsList, true) {
			return FieldKindTypeList
		} else {
			return FieldKindType
		}
	} else if n.Edges.Enum.Holds() != nil {
		if n.flags.Is(fieldflags.IsList, true) {
			return FieldKindEnumList
		} else {
			return FieldKindEnum
		}
	}

	panic("yo bratan")

	return ""
}

const (
	FieldKindType        = "TypeField"
	FieldKindEnum        = "EnumField"
	FieldKindString      = "StringField"
	FieldKindInt32       = "Int32Field"
	FieldKindFloat64     = "Float64Field"
	FieldKindBool        = "BoolField"
	FieldKindTypeList    = "TypeListField"
	FieldKindEnumList    = "EnumListField"
	FieldKindStringList  = "StringListField"
	FieldKindInt32List   = "Int32ListField"
	FieldKindFloat64List = "Float64ListField"
	FieldKindBoolList    = "BoolListField"
)

type printFieldNode struct {
	Type      string
	PrintNode `yaml:",inline"`
}

func (n *FieldNode) Print() () {
	Print(printFieldNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}
