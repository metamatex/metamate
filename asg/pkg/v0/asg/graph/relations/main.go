package relations

import (
	"fmt"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/words/cardinality"
)

type EdgeResolver struct {
	relations map[string]Relation
	// relation - from - to
	one2one map[string]map[string]string

	// relation - one - many
	one2many_many map[string]map[string]map[string]bool
	// relation - one of many - one
	one2one_one map[string]map[string]string


	many2many map[string]map[string]string
}

func (r EdgeResolver) SetOne2One(relation string, from string, to string)  {
	_, ok := r.one2one[relation]
	if !ok {
		panic(fmt.Sprintf("relation %v unknown", relation))
	}

	to0, ok := r.one2one[relation][from]
	if !ok {
		panic(fmt.Sprintf("node %v is already linked to node %v for %s relation", from, to0, relation))
	}

	from0, ok := r.one2one[relation][to]
	if !ok {
		panic(fmt.Sprintf("node %v is already linked to node %v for %s relation", to, from0, relation))
	}

	r.one2one[relation][from] = to
	r.one2one[relation][to] = from
}

func (r EdgeResolver) GetOne2One(relation string, from string) (to string) {
	to, ok := r.one2one[relation][from]
	if !ok {
		panic(fmt.Sprintf("node %v is not linked via relation %v", from, relation))
	}

	return
}

func (r EdgeResolver) AddOne2Many(relation string, one string, many []string)  {
	_, ok := r.one2many[relation]
	if !ok {
		panic(fmt.Sprintf("relation %v unknown", relation))
	}

	to0, ok := r.one2many[relation][from]
	if !ok {
		panic(fmt.Sprintf("node %v is already linked to node %v for %s relation", from, to0, relation))
	}

	from0, ok := r.one2many[relation][to]
	if !ok {
		panic(fmt.Sprintf("node %v is already linked to node %v for %s relation", to, from0, relation))
	}

	r.one2many[relation][from] = to
	r.one2many[relation][to] = from
}





type Relation struct {
	Name string
	Active string
	ActiveCardinality cardinality.Cardinality
	Passive string
	PassiveCardinality cardinality.Cardinality
}

var holds = Relation{
	Name: "holds",
	Active: "holds",
	ActiveCardinality: cardinality.Many,
	Passive: "heldBy",
	PassiveCardinality: cardinality.Many,
}

var filters = Relation{
	Name: "filters",
	Active: "filters",
	ActiveCardinality: cardinality.One,
	Passive: "filteredBy",
	PassiveCardinality: cardinality.One,
}

var handles = Relation{
	Name: "handles",
	Active: "handles",
	ActiveCardinality: cardinality.One,
	Passive: "handledBy",
	PassiveCardinality: cardinality.Many,
}

func main() {
	relationname active passive
}