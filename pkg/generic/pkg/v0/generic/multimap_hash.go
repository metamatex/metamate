package generic

import (
	"fmt"
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/mitchellh/hashstructure"
)

func (g *MultiMapGeneric) Hash() {
	for k, _ := range g.Generic0 {
		g.Generic0[k].Hash()
	}

	for _, gSlice := range g.GenericSlice0 {
		gSlice.Hash()
	}

	if g.String0 != nil {
		delete(g.String0, fieldnames.Hash)
	}

	t := g.tn
	g.tn = nil
	hash, err := hashstructure.Hash(g, &hashstructure.HashOptions{
		ZeroNil: true,
	})
	if err != nil {
		panic(err)
	}
	g.tn = t

	h := fmt.Sprintf("%d", hash)

	if g.String0 == nil {
		g.String0 = map[string]string{}
	}
	g.String0[fieldnames.Hash] = h
}

func (gSlice *MultiMapSlice) Hash() {
	for _, g := range gSlice.Gs {
		g.Hash()
	}
}
