package graph

import "strings"

type NodeId string

func ToNodeId(id string) (NodeId) {
	return NodeId(strings.ToLower(id))
}

func ToNodeIds(names ...string) (nodeIds []NodeId) {
	for _, n := range names {
		nodeIds = append(nodeIds, ToNodeId(n))
	}

	return
}

func ToString(id NodeId) (string) {
	return string(id)
}

func ToStrings(nodeIds ...NodeId) (ss []string) {
	for _, nodeId := range nodeIds {
		ss = append(ss, ToString(nodeId))
	}

	return
}

type NodeIds []NodeId

func (ns NodeIds) Append(ns0 NodeIds) {
	ns = append(ns, ns0...)
}

func MergeNodeIds(idss ...[]NodeId) ([]NodeId) {
	seen := map[NodeId]bool{}

	for _, ids := range idss {
		for _, id := range ids {
			seen[id] = true
		}
	}

	ids := []NodeId{}
	for id, _ := range seen {
		ids = append(ids, id)
	}

	return ids
}
