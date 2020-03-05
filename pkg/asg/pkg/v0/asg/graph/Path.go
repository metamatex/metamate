package graph

import (
	"fmt"
	"github.com/metamatex/asg/pkg/v0/asg/graph/words/cardinality"
	"github.com/metamatex/asg/pkg/v0/asg/utils"
	"strings"
)

type RelationPath struct {
	From        string   `yaml:",omitempty" json:"from,omitempty"`
	Fragments   []string `yaml:",omitempty" json:"fragments,omitempty"`
	Cardinality string   `yaml:",omitempty" json:"cardinality,omitempty"`
	To          string   `yaml:",omitempty" json:"to,omitempty"`
}

func (p RelationPath) ConcatFragments() (s string) {
	ss := append(p.Fragments[:1], utils.Title(p.Fragments[1:])...)

	return utils.Concat(ss)
}

func (p RelationPath) Name() (string) {
	to := p.To
	if p.Cardinality == cardinality.Many {
		to = utils.Plural(p.To)
	}

	return fmt.Sprintf("%v%v%v", p.From, strings.Title(p.ConcatFragments()), to)
}
