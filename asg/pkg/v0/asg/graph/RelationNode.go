package graph

import (
	"fmt"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/relationflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/words/cardinality"
	"reflect"
)

const (
	RELATION = "relation"
)

type RelationNode struct {
	NodeTrait `yaml:"inline,omitempty" json:",omitempty"`
	Edges     RelationNodeEdges `yaml:",omitempty" json:"edges,omitempty"`
}

func (n *RelationNode) GetEdges() interface{} {
	return n.Edges
}

func NewRelationNode() (*RelationNode) {
	return &RelationNode{
		NodeTrait: NodeTrait{
			flags: FlagsContainer{
				Flags:    Flags{},
				defaults: relationflags.Defaults,
			},
		},
	}
}

func (n *RelationNode) Init(rn *RootNode, activePath RelationPath, passivePath RelationPath) () {
	activePn := NewPathNode()
	activePn.Init(rn, activePath, nil)
	activePn.Data.IsActive = true

	passivePn := NewPathNode()
	passivePn.Init(rn, passivePath, nil)
	passivePn.Data.IsActive = false

	name := fmt.Sprintf("%v_%v_%v_%v_%v", activePath.From, activePath.ConcatFragments(), activePath.To, passivePath.ConcatFragments(), passivePath.To)

	propagateNode(rn, n, name, nil)

	n.Edges = NewRelationNodeEdges(n)

	nOne := 0
	nMany := 0

	if activePath.Cardinality == cardinality.One {
		nOne += 1
	} else {
		nMany += 1
	}

	if passivePath.Cardinality == cardinality.One {
		nOne += 1
	} else {
		nMany += 1
	}

	if nOne == 1 && nMany == 1 {
		n.flags.Set(relationflags.One2Many, true)
	}

	if nMany == 2 {
		n.flags.Set(relationflags.Many2Many, true)
	}

	if activePath.From == activePath.To {
		n.flags.Set(relationflags.IsSelfReferencing, true)
	}

	n.Edges.Type.Resolver.SetNodeA(ToNodeId(activePath.From))
	n.Edges.Type.Resolver.SetNodeB(ToNodeId(activePath.To))

	n.Edges.Path.Resolver.SetActive(activePn.Id())
	n.Edges.Path.Resolver.SetPassive(passivePn.Id())

	activePn.Edges.Relation.Resolver.SetBelongsTo(n.Id())
	passivePn.Edges.Relation.Resolver.SetBelongsTo(n.Id())

	rn.Relations.Add(n)

	return
}

func (n *RelationNode) Wire() () {
	to := n.Edges.Path.Active().Edges.Type.Resolver.To()
	from := n.Edges.Path.Active().Edges.Type.Resolver.From()

	n.root.Types.MustById(to).Edges.Relations.Resolver.AddHolds(n.Id())

	if to != from {
		n.root.Types.MustById(from).Edges.Relations.Resolver.AddHolds(n.Id())
	}

	return
}

func (n *RelationNode) Type() string {
	return RELATION
}

func (n *RelationNode) Validate() (errs []error) {
	errs = requireHasNameAndLowerCaseId(n)

	err := requirePaths(n)
	if err != nil {
		errs = append(errs, err)
	}

	err = requireCrossReference(n)
	if err != nil {
		errs = append(errs, err)
	}

	//errs0 = requireNode(n, edge.RelationPathA, edge.RelationPathB)
	//if len(errs0) != 0 {
	//	errs = append(errs, errs...)
	//}

	return
}

type printRelationNode struct {
	Type      string
	PrintNode `yaml:",inline"`
}

func (n *RelationNode) Print() () {
	Print(printRelationNode{
		Type:      reflect.TypeOf(n).Elem().Name(),
		PrintNode: getPrintNode(n),
	})
}