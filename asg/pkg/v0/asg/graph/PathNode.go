package graph

import (
	"reflect"
)

const (
	PATH = "path"
)

type PathNode struct {
	NodeTrait `yaml:"inline,omitempty" json:",omitempty"`
	Edges     PathNodeEdges `yaml:",omitempty" json:"edges,omitempty"`
	Data      PathNodeData  `yaml:",omitempty" json:"data,omitempty"`
}

func NewPathNode() (*PathNode) {
	return &PathNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags: Flags{},
			},
		},
	}
}

func (n *PathNode) Type() string {
	return PATH
}

func (n *PathNode) Init(rn *RootNode, path RelationPath, additional []interface{}) {
	propagateNode(rn, n, path.Name(), additional)

	n.Edges = NewPathNodeEdges(n)
	n.Edges.Type.Resolver.SetFrom(ToNodeId(path.From))
	n.Edges.Type.Resolver.SetTo(ToNodeId(path.To))

	n.Data.Cardinality = path.Cardinality
	n.Data.Verb = path.ConcatFragments()

	rn.Paths.Add(n)
}

func (n *PathNode) Validate() (errs []error) {
	return
}

type printPathNode struct {
	Type      string
	PrintNode `yaml:",inline"`
}

func (n *PathNode) Print() () {
	Print(printPathNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}
