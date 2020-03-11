package graph

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/utils"
)

type NodeTrait struct {
	Idx   NodeId         `yaml:",omitempty" json:"id,omitempty"`
	Namex string         `yaml:",omitempty" json:"name,omitempty"`
	flags FlagsContainer `yaml:",omitempty" json:"flags,omitempty"`
	root  *RootNode
}

func (t *NodeTrait) SetId(id NodeId) () {
	if t.Idx != "" {
		panic(fmt.Sprintf("node %s Id already set to %s, can't be set to %v anymore", t.Name(), t.Id(), id))
	}

	t.Idx = id
}

func (t *NodeTrait) GetEdges() interface{} {
	return nil
}

func (t *NodeTrait) GetData() interface{} {
	return nil
}

func (t *NodeTrait) Id() NodeId {
	return t.Idx
}

func (t *NodeTrait) SetRoot(r *RootNode) () {
	t.root = r
}

func (t *NodeTrait) Name() string {
	return t.Namex
}

func (t *NodeTrait) PluralName() string {
	return utils.Plural(t.Namex)
}

func (t *NodeTrait) SetName(name string) () {
	t.Namex = name
}

func (t *NodeTrait) Flags() FlagsContainer {
	return t.flags
}

func (n *NodeTrait) FieldName() (string) {
	return strcase.ToLowerCamel(n.Name())
}

func (t *NodeTrait) Wire() () {}
