package generic

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/mitchellh/mapstructure"
)

type MultiMapSlice struct {
	tn *graph.TypeNode
	Gs []*MultiMapGeneric
}

func (gSlice *MultiMapSlice) Type() (*graph.TypeNode) {
	return gSlice.tn
}

func (gSlice *MultiMapSlice) Get() (gs0 []Generic) {
	for _, g := range gSlice.Gs {
		gs0 = append(gs0, g)
	}

	return gs0
}

func (gSlice *MultiMapSlice) Set(gs0 []Generic) () {
	gs1 := []*MultiMapGeneric{}
	for _, g := range gs0 {
		gs1 = append(gs1, g.(*MultiMapGeneric))
	}

	gSlice.Gs = gs1
}

func (gSlice *MultiMapSlice) Append(gs0 ...Generic) () {
	for _, g := range gs0 {
		gSlice.Gs = append(gSlice.Gs, g.(*MultiMapGeneric))
	}
}

func (gSlice *MultiMapSlice) Print() {
	for _, g := range gSlice.Gs {
		g.Print()
		println("------------")
	}
}

func (gSlice *MultiMapSlice) Sprint() (s string) {
	for _, g := range gSlice.Gs {
		s += g.Sprint()
		s += "------------"

	}

	return
}

func (gSlice *MultiMapSlice) Copy() (Slice) {
	return gSlice.copy()
}

func (gSlice *MultiMapSlice) copy() (*MultiMapSlice) {
	gSlice0 := &MultiMapSlice{}
	gSlice0.tn = gSlice.tn

	gs := []*MultiMapGeneric{}
	for _, g := range gSlice.Gs {
		gs = append(gs, g.copy())
	}

	gSlice0.Gs = gs

	return gSlice0
}


func (gSlice *MultiMapSlice) ToStringInterfaceMaps() (ms []map[string]interface{}) {
	ms = []map[string]interface{}{}
	for _, v0 := range gSlice.Gs {
		ms = append(ms, v0.toStringInterfaceMap())
	}

	return
}

func (gSlice *MultiMapSlice) ToStructs(output interface{}) (err error) {
	ms := gSlice.ToStringInterfaceMaps()

	err = mapstructure.Decode(ms, &output)
	if err != nil {
		return
	}

	return
}

func (gSlice *MultiMapSlice) MustToStructs(output interface{}) () {
	ms := gSlice.ToStringInterfaceMaps()

	err := mapstructure.Decode(ms, &output)
	if err != nil {
		panic(err)
	}

	return
}