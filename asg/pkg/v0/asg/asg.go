package asg

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/expansion"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
)

func New() (rn *graph.RootNode, err error) {
	rn = graph.NewRoot()

	err = expansion.Expand(0, rn)

	return
}