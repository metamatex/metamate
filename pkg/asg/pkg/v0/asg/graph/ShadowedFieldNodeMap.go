package graph

import (
	"fmt"
)

type ShadowedFieldNodeMap struct {
	Map map[NodeId]*FieldNode
	shadowed map[NodeId]bool
}

func (nm ShadowedFieldNodeMap) ToNodeMap() (nm0 NodeMap) {
	nm0 = NodeMap{}

	nm.Each(func(n *FieldNode) {
		nm0[n.Id()] = n
	})

	return
}

func (nm ShadowedFieldNodeMap) Copy() (nm0 FieldNodeMap) {
	nm0 = FieldNodeMap{}

	nm.Each(func(n *FieldNode) {
		nm0[n.Id()] = n
	})

	return
}

func (nm ShadowedFieldNodeMap) Filter(s Filter) (nm0 FieldNodeMap) {
	nm0 = nm.Copy()

	if s.Flags != nil {
		nm0 = nm0.FilterByFlags(*s.Flags)
	}

	if s.Names != nil {
		nm0 = nm0.FilterByNames(*s.Names)
	}

	return nm0
}

func (nm ShadowedFieldNodeMap) Add(ns ...*FieldNode) {
	for _, n := range ns {
		_, ok := nm.Map[n.Id()]
		if ok {
			panic(fmt.Sprintf("FieldNode %v already appended", n.Id()))
		}

		nm.Map[n.Id()] = n
	}
}

func (nm ShadowedFieldNodeMap) AddFieldNodeMap(nm0 FieldNodeMap) (nm1 FieldNodeMap) {
	nm1 = FieldNodeMap{}

	for _, n := range nm.Map {
		nm1.Add(n)
	}

	for _, n := range nm0 {
		nm1.Add(n)
	}

	return
}

func (nm ShadowedFieldNodeMap) ExcludeIds(ids ...NodeId) (FieldNodeMap) {
	not := map[NodeId]bool{}

	for _, id := range ids {
		not[id] = true
	}

	for id, _ := range nm.shadowed {
		not[id] = true
	}

	return nm.FilterFunc(func(n *FieldNode) bool {
		_, ok := not[n.Id()]

		return !ok
	})
}

func (nm ShadowedFieldNodeMap) ExcludeNames(names ...string) (FieldNodeMap) {
	ids := []NodeId{}
	for _, name := range names {
		ids = append(ids, ToNodeId(name))
	}

	return nm.ExcludeIds(ids...)
}

func (nm ShadowedFieldNodeMap) ByIds(ids ...NodeId) (nm0 FieldNodeMap) {
	nm0 = FieldNodeMap{}

	for _, id := range ids {
		n, ok := nm.Map[id]
		if !ok {
			panic(fmt.Sprintf("node Id %v not found", id))
		}

		_, ok = nm.shadowed[n.Edges.Type.Resolver.Holds()]
		if ok {
			continue
		}

		nm0[id] = n
	}

	return
}

func (nm ShadowedFieldNodeMap) ByNames(names ...string) (FieldNodeMap) {
	ids := []NodeId{}
	for _, name := range names {
		ids = append(ids, ToNodeId(name))
	}

	return nm.ByIds(ids...)
}

func (nm ShadowedFieldNodeMap) ById(id NodeId) (n *FieldNode) {
	n, ok := nm.Map[id]
	if !ok {
		panic(fmt.Sprintf("node Id %v not found", id))
	}

	_, ok = nm.shadowed[n.Edges.Type.Resolver.Holds()]
	if !ok {
		panic(fmt.Sprintf("node Id %v is shadowed", id))
	}

	return
}

func (nm ShadowedFieldNodeMap) ByName(name string) (*FieldNode) {
	return nm.ById(ToNodeId(name))
}

func (nm ShadowedFieldNodeMap) Each(f func(n *FieldNode)) () {
	for _, n := range nm.Map {
		_, ok := nm.shadowed[n.Id()]
		if ok {
			continue
		}

		f(n)
	}
}

func (nm ShadowedFieldNodeMap) FilterFunc(f func(n *FieldNode)(bool)) (nm0 FieldNodeMap) {
	nm0 = FieldNodeMap{}

	nm.Each(func(n *FieldNode) {
		if f(n) {
			nm0[n.Id()] = n
		}
	})

	return
}

func (nm ShadowedFieldNodeMap) HasId(id NodeId) (bool) {
	_, ok := nm.shadowed[id]
	if ok {
		return false
	}

	_, ok = nm.Map[id]

	return ok
}

func (nm ShadowedFieldNodeMap) HasName(name string) (bool) {
	return nm.HasId(ToNodeId(name))
}

func (nm ShadowedFieldNodeMap) BroadcastSetFlag(name string, v bool) () {
	nm.Each(func(n *FieldNode) {
		n.Flags().Set(name, v)
	})

	return
}

func (nm ShadowedFieldNodeMap) GetIds() (ids []NodeId) {
	for k, _ := range nm.Map {
		_, ok := nm.shadowed[k]
		if ok {
			continue
		}

		ids = append(ids, k)
	}

	return
}

func (nm ShadowedFieldNodeMap) Flagged(flag string, b bool) (nm0 FieldNodeMap) {
	return nm.FilterFunc(func(n *FieldNode) bool {
		return n.Flags().Is(flag, b)
	})
}
//
//func (nm ShadowedFieldNodeMap) FilterByFlags(subset FlagsSubset) (FieldNodeMap) {
//	nm
//
//	if len(subset.Or) != 0 {
//		var tfs []string
//		for _, f := range subset.Or {
//			tfs = append(tfs, string(f))
//		}
//
//		nm = nm.FlaggedOr(tfs...)
//	}
//
//	if len(subset.And) != 0 {
//		for _, f := range subset.And {
//			nm = nm.Flagged(string(f), true)
//		}
//	}
//
//	for _, f := range subset.Nor {
//		nm = nm.Flagged(string(f), false)
//	}
//
//	if len(nm) == 0 {
//		panic("no nodes")
//	}
//
//	return nm
//}
//
//func (nm ShadowedFieldNodeMap) FilterByNames(subset NamesSubset) (FieldNodeMap) {
//	if len(subset.Or) != 0 {
//		nm = nm.ByNames(subset.Or...)
//	}
//
//	if len(subset.Nor) != 0 {
//		nm = nm.ExcludeNames(subset.Nor...)
//	}
//
//	if len(subset.ContainsOr) != 0 {
//		nm = nm.FilterFunc(func(n *FieldNode) bool {
//			for _, c := range subset.ContainsOr {
//				if strings.Contains(n.Name(), c) {
//					return true
//				}
//			}
//
//			return false
//		})
//	}
//
//	if len(nm) == 0 {
//		panic("no Fields")
//	}
//
//	return nm
//}
//
//func (nm ShadowedFieldNodeMap) FlaggedOr(fs ...string) (nm0 FieldNodeMap) {
//	nm0 = FieldNodeMap{}
//
//	for k, n := range nm {
//		if n.Flags().Or(fs...) {
//			nm0[k] = n
//		}
//	}
//
//	return
//}
//
//func (nm ShadowedFieldNodeMap) BroadcastPrint() () {
//	for _, n := range nm {
//		n.Print()
//		println("- - - - - - - - - -")
//	}
//
//	return
//}
//
//func (nm ShadowedFieldNodeMap) Slice() (ns FieldNodeSlice) {
//	for _, n := range nm {
//		ns = append(ns, n)
//	}
//
//	return
//}
//
//
