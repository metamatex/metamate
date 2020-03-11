package generic

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"strings"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/mitchellh/mapstructure"
	"github.com/wolfeidau/unflatten"
)

type MultiMapGeneric struct {
	tn            *graph.TypeNode `hash:"-" yaml:"-"`
	String0       map[string]string
	Int           map[string]int
	Int320        map[string]int32
	Uint320       map[string]uint32
	Float640      map[string]float64
	Bool0         map[string]bool
	StringSlice0  map[string][]string
	Int32Slice0   map[string][]int32
	Uint32Slice0  map[string][]uint32
	Float64Slice0 map[string][]float64
	BoolSlice0    map[string][]bool
	Generic0      map[string]*MultiMapGeneric
	GenericSlice0 map[string]*MultiMapSlice
}

func (g *MultiMapGeneric) MustString(names ...string) (string) {
	v, ok := g.String(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v string not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustInt32(names ...string) (int32) {
	v, ok := g.Int32(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v Int320 not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustFloat64(names ...string) (float64) {
	v, ok := g.Float64(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v float64 not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustBool(names ...string) (bool) {
	v, ok := g.Bool(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v []bool not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustGeneric(names ...string) (Generic) {
	v, ok := g.Generic(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v generic.Generic not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustStringSlice(names ...string) ([]string) {
	v, ok := g.StringSlice(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v []string not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustInt32Slice(names ...string) ([]int32) {
	v, ok := g.Int32Slice(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v []Int320 not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustFloat64Slice(names ...string) ([]float64) {
	v, ok := g.Float64Slice(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v []float64 not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustBoolSlice(names ...string) ([]bool) {
	v, ok := g.BoolSlice(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v []bool not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func (g *MultiMapGeneric) MustGenericSlice(names ...string) (Slice) {
	v, ok := g.GenericSlice(names...)
	if !ok {
		panic(fmt.Sprintf("%v.%v generic.Slice not set", g.Type().Name(), strings.Join(names, ".")))
	}

	return v
}

func NewMultiMapGeneric(t *graph.TypeNode) (*MultiMapGeneric) {
	return &MultiMapGeneric{
		tn: t,
	}
}

func (g *MultiMapGeneric) EachString(f func(fn *graph.FieldNode, v string)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.String0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachInt32(f func(fn *graph.FieldNode, v int32)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.Int320 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachFloat64(f func(fn *graph.FieldNode, v float64)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.Float640 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachBool(f func(fn *graph.FieldNode, v bool)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.Bool0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachEnum(f func(fn *graph.FieldNode, v string)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.String0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachGeneric(f func(fn *graph.FieldNode, v Generic)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.Generic0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachStringSlice(f func(fn *graph.FieldNode, v []string)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.StringSlice0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachInt32Slice(f func(fn *graph.FieldNode, v []int32)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.Int32Slice0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachFloat64Slice(f func(fn *graph.FieldNode, v []float64)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.Float64Slice0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachBoolSlice(f func(fn *graph.FieldNode, v []bool)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.BoolSlice0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachEnumSlice(f func(fn *graph.FieldNode, v []string)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.StringSlice0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) EachGenericSlice(f func(fn *graph.FieldNode, v Slice)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, v := range g.GenericSlice0 {
		f(fnm.MustById(g.Type().Id() + "_" + graph.ToNodeId(k)), v)
	}
}

func (g *MultiMapGeneric) MustSetInt32(names []string, i int32) (Generic) {
	err := g.SetInt32(names, i)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetInt32(names []string, i int32) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.Int320 == nil {
		g.Int320 = map[string]int32{}
	}

	g.Int320[base] = i

	return
}

func (g *MultiMapGeneric) MustSetFloat64(names []string, f float64) (Generic) {
	err := g.SetFloat64(names, f)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetFloat64(names []string, f float64) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.Float640 == nil {
		g.Float640 = map[string]float64{}
	}

	g.Float640[base] = f

	return
}

func (g *MultiMapGeneric) MustSetBool(names []string, b bool) (Generic) {
	err := g.SetBool(names, b)
	if err != nil {
		log.Print(err)
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetBool(names []string, b bool) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.Bool0 == nil {
		g.Bool0 = map[string]bool{}
	}

	g.Bool0[base] = b

	return
}

func (g *MultiMapGeneric) MustSetGeneric(names []string, g0 Generic) (Generic) {
	err := g.SetGeneric(names, g0)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetGeneric(names []string, g0 Generic) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.Generic0 == nil {
		g.Generic0 = map[string]*MultiMapGeneric{}
	}

	g.Generic0[base] = g0.(*MultiMapGeneric)

	return
}

func (g *MultiMapGeneric) MustSetInt32Slice(names []string, is []int32) (Generic) {
	err := g.SetInt32Slice(names, is)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetInt32Slice(names []string, is []int32) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.Int32Slice0 == nil {
		g.Int32Slice0 = map[string][]int32{}
	}

	g.Int32Slice0[base] = is

	return
}

func (g *MultiMapGeneric) MustSetFloat64Slice(names []string, fs []float64) (Generic) {
	err := g.SetFloat64Slice(names, fs)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetFloat64Slice(names []string, fs []float64) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.Float64Slice0 == nil {
		g.Float64Slice0 = map[string][]float64{}
	}

	g.Float64Slice0[base] = fs

	return
}

func (g *MultiMapGeneric) MustSetBoolSlice(names []string, bs []bool) (Generic) {
	err := g.SetBoolSlice(names, bs)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetBoolSlice(names []string, bs []bool) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.BoolSlice0 == nil {
		g.BoolSlice0 = map[string][]bool{}
	}

	g.BoolSlice0[base] = bs

	return
}

func (g *MultiMapGeneric) MustSetGenericSlice(names []string, gSlice Slice) (Generic) {
	err := g.SetGenericSlice(names, gSlice)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetGenericSlice(names []string, g0 Slice) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.GenericSlice0 == nil {
		g.GenericSlice0 = map[string]*MultiMapSlice{}
	}

	g.GenericSlice0[base] = g0.(*MultiMapSlice)

	return
}

func (g *MultiMapGeneric) MustSetStringSlice(names []string, ss []string) (Generic) {
	err := g.SetStringSlice(names, ss)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetStringSlice(names []string, ss []string) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.StringSlice0 == nil {
		g.StringSlice0 = map[string][]string{}
	}

	g.StringSlice0[base] = ss

	return
}

func (g *MultiMapGeneric) MustSetEnum(names []string, s string) (Generic) {
	err := g.SetEnum(names, s)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetEnum(names []string, s string) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.String0 == nil {
		g.String0 = map[string]string{}
	}

	g.String0[base] = s

	return
}

func (g *MultiMapGeneric) MustSetString(names []string, s string) (Generic) {
	err := g.SetString(names, s)
	if err != nil {
		panic(err)
	}

	return g
}

func (g *MultiMapGeneric) SetString(names []string, s string) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		g, err = g.createPath(path)
		if err != nil {
			return
		}
	}

	if g.String0 == nil {
		g.String0 = map[string]string{}
	}

	g.String0[base] = s

	return
}

func (g *MultiMapGeneric) WalkDelete(f func(fn *graph.FieldNode) (bool)) {
	fnm := g.Type().Edges.Fields.Holds()

	for k, _ := range g.String0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, _ := range g.Int320 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, _ := range g.Float640 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, _ := range g.Bool0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, g0 := range g.Generic0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		} else {
			g0.WalkDelete(f)
		}
	}

	for k, _ := range g.StringSlice0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, _ := range g.Int32Slice0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, _ := range g.Float64Slice0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, _ := range g.BoolSlice0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		}
	}

	for k, gSlice := range g.GenericSlice0 {
		id := g.Type().Id() + "_" + graph.ToNodeId(k)

		if f(fnm.MustById(id)) {
			g.MustDelete(k)
		} else {
			var gs []Generic
			for _, g0 := range gSlice.Get() {
				g0.WalkDelete(f)
				gs = append(gs, g0)
			}

			gSlice.Set(gs)
		}
	}

	return
}

func (g *MultiMapGeneric) MustDelete(names ...string) {
	err := g.Delete(names...)
	if err != nil {
	    panic(err)
	}
}

func (g *MultiMapGeneric) Delete(names ...string) (err error) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	id := g.Type().Id() + "_" + graph.ToNodeId(base)
	fn, err := g.Type().Edges.Fields.Holds().ById(id)
	if err != nil {
	    return
	}

	switch fn.Kind() {
	case graph.FieldKindString, graph.FieldKindEnum:
		delete(g.String0, base)
	case graph.FieldKindBool:
		delete(g.Bool0, base)
	case graph.FieldKindInt32:
		delete(g.Int320, base)
	case graph.FieldKindFloat64:
		delete(g.Float640, base)
	case graph.FieldKindEnumList, graph.FieldKindStringList:
		delete(g.StringSlice0, base)
	case graph.FieldKindBoolList:
		delete(g.BoolSlice0, base)
	case graph.FieldKindInt32List:
		delete(g.Int32Slice0, base)
	case graph.FieldKindFloat64List:
		delete(g.Float64Slice0, base)
	case graph.FieldKindType:
		delete(g.Generic0, base)
	case graph.FieldKindTypeList:
		delete(g.GenericSlice0, base)
	default:
		panic(fmt.Sprintf("unhandled case %v", fn.Kind()))
	}

	return
}

func (g *MultiMapGeneric) createPath(path []string) (*MultiMapGeneric, error) {
	for _, p := range path {
		g0, ok := g.Generic0[p]
		if ok {
			g = g0

			continue
		}

		var fn *graph.FieldNode
		id := g.Type().Id() + "_" + graph.ToNodeId(p)
		fn, err := g.Type().Edges.Fields.Holds().ById(id)
		if err != nil {
			return g, err
		}

		if g.Generic0 == nil {
			g.Generic0 = map[string]*MultiMapGeneric{}
		}
		g.Generic0[p] = NewMultiMapGeneric(fn.Edges.Type.Holds())
		g = g.Generic0[p]
	}

	return g, nil
}

func (g *MultiMapGeneric) String(names ...string) (i string, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.String0[base]

	return
}

func (g *MultiMapGeneric) Enum(names ...string) (i string, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.String0[base]

	return
}

func (g *MultiMapGeneric) Int32(names ...string) (i int32, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.Int320[base]

	return
}

func (g *MultiMapGeneric) Float64(names ...string) (i float64, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.Float640[base]

	return
}

func (g *MultiMapGeneric) Bool(names ...string) (i bool, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.Bool0[base]

	return
}

func (g *MultiMapGeneric) Generic(names ...string) (Generic, bool) {
	g0, ok := g.getGeneric(names)
	if !ok {
		return nil, ok
	}

	return g0, ok
}

func (g *MultiMapGeneric) getGeneric(names []string) (*MultiMapGeneric, bool) {
	path, base := split(names)

	for _, p := range path {
		if g.Generic0 == nil {
			return nil, false
		}

		var ok bool
		g0, ok := g.Generic0[p]
		if ok {
			g = g0
		} else {
			return nil, false
		}
	}

	g0, b := g.Generic0[base]

	return g0, b
}

func (g *MultiMapGeneric) StringSlice(names ...string) (i []string, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.StringSlice0[base]

	return
}

func (g *MultiMapGeneric) Int32Slice(names ...string) (i []int32, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.Int32Slice0[base]

	return
}

func (g *MultiMapGeneric) Float64Slice(names ...string) (i []float64, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.Float64Slice0[base]

	return
}

func (g *MultiMapGeneric) BoolSlice(names ...string) (i []bool, b bool) {
	path, base := split(names)

	if len(path) != 0 {
		var ok bool
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	i, b = g.BoolSlice0[base]

	return
}

func (g *MultiMapGeneric) FieldNames() (ss []string) {
	for k, _ := range g.String0 {
		ss = append(ss, k)
	}

	for k, _ := range g.Int320 {
		ss = append(ss, k)
	}

	for k, _ := range g.Float640 {
		ss = append(ss, k)
	}

	for k, _ := range g.Bool0 {
		ss = append(ss, k)
	}

	for k, _ := range g.Generic0 {
		ss = append(ss, k)
	}

	for k, _ := range g.StringSlice0 {
		ss = append(ss, k)
	}

	for k, _ := range g.Int32Slice0 {
		ss = append(ss, k)
	}

	for k, _ := range g.Float64Slice0 {
		ss = append(ss, k)
	}

	for k, _ := range g.BoolSlice0 {
		ss = append(ss, k)
	}

	for k, _ := range g.GenericSlice0 {
		ss = append(ss, k)
	}

	return
}

func (g *MultiMapGeneric) GenericSlice(names ...string) (gSlice Slice, ok bool) {
	path, base := split(names)

	if len(path) != 0 {
		g, ok = g.getGeneric(path)
		if !ok {
			return
		}
	}

	gSlice, ok = g.GenericSlice0[base]
	if !ok {
		return nil, ok
	}

	return
}

func split(names []string) (path []string, base string) {
	if len(names) == 0 {
		return
	}

	base = names[len(names)-1]
	path = names[:len(names)-1]

	return
}

func (g *MultiMapGeneric) Sanitize() {
	g.sanitize()
}

func (g *MultiMapGeneric) sanitize() {
	g0 := NewMultiMapGeneric(g.Type())

	g.Type().Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		switch fn.Kind() {
		case graph.FieldKindEnum:
			v, ok := g.String0[fn.Name()]
			if !ok {
				return
			}

			if g0.String0 == nil {
				g0.String0 = map[string]string{}
			}

			g0.String0[fn.Name()] = v
		case graph.FieldKindString:
			v, ok := g.String0[fn.Name()]
			if !ok {
				return
			}

			if g0.String0 == nil {
				g0.String0 = map[string]string{}
			}

			g0.String0[fn.Name()] = v
		case graph.FieldKindBool:
			v, ok := g.Bool0[fn.Name()]
			if !ok {
				return
			}

			if g0.Bool0 == nil {
				g0.Bool0 = map[string]bool{}
			}

			g0.Bool0[fn.Name()] = v
		case graph.FieldKindInt32:
			v, ok := g.Int320[fn.Name()]
			if !ok {
				return
			}

			if g0.Int320 == nil {
				g0.Int320 = map[string]int32{}
			}

			g0.Int320[fn.Name()] = v
		case graph.FieldKindFloat64:
			v, ok := g.Float640[fn.Name()]
			if !ok {
				return
			}

			if g0.Float640 == nil {
				g0.Float640 = map[string]float64{}
			}

			g0.Float640[fn.Name()] = v
		case graph.FieldKindEnumList:
			v, ok := g.StringSlice0[fn.Name()]
			if !ok {
				return
			}

			if g0.StringSlice0 == nil {
				g0.StringSlice0 = map[string][]string{}
			}

			g0.StringSlice0[fn.Name()] = v
		case graph.FieldKindStringList:
			v, ok := g.StringSlice0[fn.Name()]
			if !ok {
				return
			}

			if g0.StringSlice0 == nil {
				g0.StringSlice0 = map[string][]string{}
			}

			g0.StringSlice0[fn.Name()] = v
		case graph.FieldKindBoolList:
			v, ok := g.BoolSlice0[fn.Name()]
			if !ok {
				return
			}

			if g0.BoolSlice0 == nil {
				g0.BoolSlice0 = map[string][]bool{}
			}

			g0.BoolSlice0[fn.Name()] = v
		case graph.FieldKindInt32List:
			v, ok := g.Int32Slice0[fn.Name()]
			if !ok {
				return
			}

			if g0.Int32Slice0 == nil {
				g0.Int32Slice0 = map[string][]int32{}
			}

			g0.Int32Slice0[fn.Name()] = v
		case graph.FieldKindFloat64List:
			v, ok := g.Float64Slice0[fn.Name()]
			if !ok {
				return
			}

			if g0.Float64Slice0 == nil {
				g0.Float64Slice0 = map[string][]float64{}
			}

			g0.Float64Slice0[fn.Name()] = v
		case graph.FieldKindType:
			_, ok := g.Generic0[fn.Name()]
			if !ok {
				return
			}

			if g0.Generic0 == nil {
				g0.Generic0 = map[string]*MultiMapGeneric{}
			}

			g0.Generic0[fn.Name()].sanitize()
		case graph.FieldKindTypeList:
			v, ok := g.GenericSlice0[fn.Name()]
			if !ok {
				return
			}

			if g0.GenericSlice0 == nil {
				g0.GenericSlice0 = map[string]*MultiMapSlice{}
			}

			gs := &MultiMapSlice{tn: v.tn}
			for _, g1 := range v.Gs {
				g1.sanitize()
				gs.Gs = append(gs.Gs, g1)
			}

			g0.GenericSlice0[fn.Name()] = gs
		}

	})

	g = g0
}

func (g *MultiMapGeneric) Type() (*graph.TypeNode) {
	return g.tn
}

func (g *MultiMapGeneric) Copy() (Generic) {
	return g.copy()
}

func (g *MultiMapGeneric) copy() (g0 *MultiMapGeneric) {
	g0 = &MultiMapGeneric{}
	g0.tn = g.tn

	if g.String0 != nil {
		g0.String0 = map[string]string{}
		for k, v := range g.String0 {
			g0.String0[k] = v
		}
	}

	if g.Int320 != nil {
		g0.Int320 = map[string]int32{}
		for k, v := range g.Int320 {
			g0.Int320[k] = v
		}
	}

	if g.Uint320 != nil {
		g0.Uint320 = map[string]uint32{}
		for k, v := range g.Uint320 {
			g0.Uint320[k] = v
		}
	}

	if g.Float640 != nil {
		g0.Float640 = map[string]float64{}
		for k, v := range g.Float640 {
			g0.Float640[k] = v
		}
	}

	if g.Bool0 != nil {
		g0.Bool0 = map[string]bool{}
		for k, v := range g.Bool0 {
			g0.Bool0[k] = v
		}
	}

	if g.StringSlice0 != nil {
		g0.StringSlice0 = map[string][]string{}
		for k, v := range g.StringSlice0 {
			g0.StringSlice0[k] = v
		}
	}

	if g.Int32Slice0 != nil {
		g0.Int32Slice0 = map[string][]int32{}
		for k, v := range g.Int32Slice0 {
			g0.Int32Slice0[k] = v
		}
	}

	if g.Uint32Slice0 != nil {
		g0.Uint32Slice0 = map[string][]uint32{}
		for k, v := range g.Uint32Slice0 {
			g0.Uint32Slice0[k] = v
		}
	}

	if g.Float64Slice0 != nil {
		g0.Float64Slice0 = map[string][]float64{}
		for k, v := range g.Float64Slice0 {
			g0.Float64Slice0[k] = v
		}
	}

	if g.BoolSlice0 != nil {
		g0.BoolSlice0 = map[string][]bool{}
		for k, v := range g.BoolSlice0 {
			g0.BoolSlice0[k] = v
		}
	}

	if g.Generic0 != nil {
		g0.Generic0 = map[string]*MultiMapGeneric{}
		for k, v := range g.Generic0 {
			g0.Generic0[k] = v.copy()
		}
	}

	if g.GenericSlice0 != nil {
		g0.GenericSlice0 = map[string]*MultiMapSlice{}
		for k, v := range g.GenericSlice0 {
			g0.GenericSlice0[k] = v.copy()
		}
	}

	return g0
}

func (g *MultiMapGeneric) ToStruct(output interface{}) (err error) {
	m := g.ToStringInterfaceMap()

	err = mapstructure.Decode(m, &output)
	if err != nil {
		return
	}

	return
}

func (g *MultiMapGeneric) MustToStruct(output interface{}) () {
	m := g.ToStringInterfaceMap()

	err := mapstructure.Decode(m, &output)
	if err != nil {
		panic(err)
	}

	return
}

func (g *MultiMapGeneric) Flatten(delimiter string) (m map[string]interface{}, err error) {
	m = g.ToStringInterfaceMap()

	m = unflatten.Flatten(m, func(ks []string) string { return strings.Join(ks, delimiter) })

	return
}

func (g *MultiMapGeneric) MustFlatten(delimiter string) (m map[string]interface{}) {
	m, err := g.Flatten(delimiter)
	if err != nil {
		panic(err)
	}

	return
}

func (g *MultiMapGeneric) Print() {
	println(g.Type().Name())
	println(g.sprint())
}

func (g *MultiMapGeneric) Sprint() (string) {
	return g.sprint()
}

func (g *MultiMapGeneric) sprint() (string) {
	b, err := yaml.Marshal(g.toStringInterfaceMap())
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (g *MultiMapGeneric) ToStringInterfaceMap() (m map[string]interface{}) {
	return g.toStringInterfaceMap()
}

func (g *MultiMapGeneric) toStringInterfaceMap() (m map[string]interface{}) {
	m = map[string]interface{}{}

	for k, v := range g.String0 {
		m[k] = v
	}

	for k, v := range g.StringSlice0 {
		m[k] = v
	}

	for k, v := range g.Int320 {
		m[k] = v
	}

	for k, v := range g.Int32Slice0 {
		m[k] = v
	}

	for k, v := range g.Uint320 {
		m[k] = v
	}

	for k, v := range g.Uint32Slice0 {
		m[k] = v
	}

	for k, v := range g.Float640 {
		m[k] = v
	}

	for k, v := range g.Float64Slice0 {
		m[k] = v
	}

	for k, v := range g.Bool0 {
		m[k] = v
	}

	for k, v := range g.BoolSlice0 {
		m[k] = v
	}

	for k, v := range g.Generic0 {
		m[k] = v.toStringInterfaceMap()
	}

	for k, v := range g.GenericSlice0 {
		m[k] = v.ToStringInterfaceMaps()
	}

	return
}
