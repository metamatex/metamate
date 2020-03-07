package graph

import (
	"fmt"
	"strings"
)

type NodeMap map[NodeId]Node

func (nm NodeMap) Filter(s Filter) (NodeMap) {
	if s.Flags != nil {
		nm = nm.FilterByFlags(*s.Flags)
	}

	if s.Names != nil {
		nm = nm.FilterByNames(*s.Names)
	}

	return nm
}

func (nm NodeMap) Add(ns ...Node) {
	for _, n := range ns {
		_, ok := nm[n.Id()]
		if ok {
			panic(fmt.Sprintf("Node %v already appended", n.Id()))
		}

		nm[n.Id()] = n
	}
}

func (nm NodeMap) ExcludeIds(ids ...NodeId) (NodeMap) {
	not := map[NodeId]bool{}

	for _, id := range ids {
		not[id] = true
	}

	return nm.FilterFunc(func(n Node) bool {
		_, ok := not[n.Id()]

		return !ok
	})
}

func (nm NodeMap) ExcludeNames(names ...string) (NodeMap) {
	ids := []NodeId{}
	for _, name := range names {
		ids = append(ids, ToNodeId(name))
	}

	return nm.ExcludeIds(ids...)
}

func (nm NodeMap) ByIds(ids ...NodeId) (filtered NodeMap) {
	filtered = NodeMap{}

	for _, id := range ids {
		filtered[id] = nm.ById(id)
	}

	return
}

func (nm NodeMap) ByNames(names ...string) (NodeMap) {
	ids := []NodeId{}
	for _, name := range names {
		ids = append(ids, ToNodeId(name))
	}

	return nm.ByIds(ids...)
}

func (nm NodeMap) ById(id NodeId) (n Node) {
	n, ok := nm[id]
	if !ok {
		panic(fmt.Sprintf("node Id %v not found", id))
	}

	return
}

func (nm NodeMap) ByName(name string) (Node) {
	return nm.ById(ToNodeId(name))
}

func (nm NodeMap) Each(f func(n Node)) () {
	for _, n := range nm {
		f(n)
	}
}

func (nm NodeMap) FilterFunc(f func(n Node)(bool)) (nm0 NodeMap) {
	nm0 = NodeMap{}

	for _, n := range nm {
		if f(n) {
			nm0[n.Id()] = n
		}
	}

	return
}

func (nm NodeMap) HasId(id NodeId) (bool) {
	_, ok := nm[id]

	return ok
}

func (nm NodeMap) BroadcastSetFlag(name string, v bool) () {
	for _, n := range nm {
		n.Flags().Set(name, v)
	}

	return
}

func (nm NodeMap) GetIds() (ids []NodeId) {
	for k, _ := range nm {
		ids = append(ids, k)
	}

	return
}

func (nm NodeMap) Flagged(flag string, b bool) (nm0 NodeMap) {
	nm0 = NodeMap{}

	for k, n := range nm {
		if n.Flags().Is(flag, b) {
			nm0[k] = n
		}
	}

	return
}

func (nm NodeMap) FilterByFlags(subset FlagsSubset) (NodeMap) {
	if len(subset.Or) != 0 {
		var tfs []string
		for _, f := range subset.Or {
			tfs = append(tfs, string(f))
		}

		nm = nm.FlaggedOr(tfs...)
	}

	if len(subset.And) != 0 {
		for _, f := range subset.And {
			nm = nm.Flagged(string(f), true)
		}
	}

	for _, f := range subset.Nor {
		nm = nm.Flagged(string(f), false)
	}

	if len(nm) == 0 {
		panic("no nodes")
	}

	return nm
}

func (nm NodeMap) FilterByNames(subset NamesSubset) (NodeMap) {
	if len(subset.Or) != 0 {
		nm = nm.ByNames(subset.Or...)
	}

	if len(subset.Nor) != 0 {
		nm = nm.ExcludeNames(subset.Nor...)
	}

	nm = nm.FilterFunc(func(n Node) bool {
		for _, c := range subset.ContainsOr {
			if strings.Contains(n.Name(), c) {
				return true
			}
		}

		return false
	})

	if len(nm) == 0 {
		panic("no Types")
	}

	return nm
}

func (nm NodeMap) FlaggedOr(fs ...string) (nm0 NodeMap) {
	nm0 = NodeMap{}

	for k, n := range nm {
		if n.Flags().Or(fs...) {
			nm0[k] = n
		}
	}

	return
}

func (nm NodeMap) BroadcastPrint() () {
	for _, n := range nm {
		n.Print()
		println("- - - - - - - - - -")
	}

	return
}