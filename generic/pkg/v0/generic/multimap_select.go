package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
)

func (g *MultiMapGeneric) Select(gSelect Generic)  {
	g.select0(gSelect.(*MultiMapGeneric))
}

func (g *MultiMapGeneric) select0(gSelect *MultiMapGeneric)  {
	// todo filtering is skipped like this
	selectAll, ok := gSelect.Bool0["All"]
	if ok && selectAll {
		return
	}

	selectHash := gSelect.Bool0["SelectHash"]
	if !selectHash {
		delete(g.String0, fieldnames.Hash)
	}

	for k, _ := range g.String0 {
		if k == fieldnames.Hash {
			continue
		}

		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.String0, k)
		}
	}

	for k, _ := range g.Int320 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.Int320, k)
		}
	}

	for k, _ := range g.Float640 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.Float640, k)
		}
	}

	for k, _ := range g.Bool0 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.Bool0, k)
		}
	}

	for k, _ := range g.Generic0 {
		gSelect0, ok := gSelect.Generic0[k]
		if !ok {
			delete(g.Generic0, k)
		} else {
			g.Generic0[k].select0(gSelect0)
		}
	}

	for k, _ := range g.StringSlice0 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.StringSlice0, k)
		}
	}

	for k, _ := range g.Int32Slice0 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.Int32Slice0, k)
		}
	}

	for k, _ := range g.Float64Slice0 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.Float64Slice0, k)
		}
	}

	for k, _ := range g.BoolSlice0 {
		s, ok := gSelect.Bool0[k]
		if !ok || !s {
			delete(g.BoolSlice0, k)
		}
	}

	for k, _ := range g.GenericSlice0 {
		gSelect0, ok := gSelect.Generic0[k]
		if !ok {
			delete(g.GenericSlice0, k)
		} else {
			g.GenericSlice0[k].select0(gSelect0)
		}
	}
}

func (gSlice *MultiMapSlice) Select(gSelect Generic) {
	gSlice.select0(gSelect.(*MultiMapGeneric))
}

func (gSlice *MultiMapSlice) select0(gSelect *MultiMapGeneric) () {
	for i, _ := range gSlice.Gs {
		gSlice.Gs[i].select0(gSelect)
	}

	return
}