package graph

import (
	"strings"
)

type GraphStats struct {
	Total     NodeStats `yaml:",omitempty" json:"total,omitempty"`
	BasicType NodeStats `yaml:",omitempty" json:"basictype,omitempty"`
	Endpoint  NodeStats `yaml:",omitempty" json:"endpoint,omitempty"`
	Enum      NodeStats `yaml:",omitempty" json:"enum,omitempty"`
	Field     NodeStats `yaml:",omitempty" json:"field,omitempty"`
	Relation  NodeStats `yaml:",omitempty" json:"relation,omitempty"`
	Type      NodeStats `yaml:",omitempty" json:"type,omitempty"`
}

type NodeStats struct {
	Count int      `yaml:",omitempty" json:"count,omitempty"`
	Ids   []NodeId `yaml:",omitempty" json:"ids,omitempty"`
}

var ignoreEdges = []string{TypeDependenciesTypes, TypeDependenciesEnums, EndpointDependenciesTypes, EndpointDependenciesEnums}

func GetMissing(rn *RootNode) (missing GraphStats) {
	ids := map[string]map[NodeId]bool{
		BASIC_TYPE: {},
		ENDPOINT:   {},
		ENUM:       {},
		FIELD:      {},
		RELATION:   {},
		TYPE:       {},
	}

	add := func(typeName string, id NodeId) {
		typeName = strings.ToLower(typeName)
		switch typeName {
		case "basictypes":
			fallthrough
		case BASIC_TYPE:
			ids[BASIC_TYPE][id] = true
			break

		case "endpoints":
			fallthrough
		case ENDPOINT:
			ids[ENDPOINT][id] = true
			break

		case "enums":
			fallthrough
		case ENUM:
			ids[ENUM][id] = true
			break

		case "fields":
			fallthrough
		case FIELD:
			ids[FIELD][id] = true
			break

		case "relations":
			fallthrough
		case RELATION:
			ids[RELATION][id] = true
			break

		case "Types":
			fallthrough
		case TYPE:
			ids[TYPE][id] = true
			break
		}
	}

	rn.GetNodes().Each(func(n Node) {
		single, multi := GetEdgeMaps(n.GetEdges(), ignoreEdges...)

		for typeName, edges := range single {
			for _, nodeId := range edges {
				add(typeName, nodeId)
			}
		}

		for typeName, edges := range multi {
			for _, nodeIds := range edges {
				for _, nodeId := range nodeIds {
					add(typeName, nodeId)
				}
			}
		}
	})

	for nodeId, _ := range ids[BASIC_TYPE] {
		_, ok := rn.BasicTypes[nodeId]
		if !ok {
			missing.BasicType.Ids = append(missing.BasicType.Ids, nodeId)
		}
	}

	for nodeId, _ := range ids[ENDPOINT] {
		_, ok := rn.Endpoints[nodeId]
		if !ok {
			missing.Endpoint.Ids = append(missing.Endpoint.Ids, nodeId)
		}
	}

	for nodeId, _ := range ids[ENUM] {
		_, ok := rn.Enums[nodeId]
		if !ok {
			missing.Enum.Ids = append(missing.Enum.Ids, nodeId)
		}
	}

	for nodeId, _ := range ids[FIELD] {
		_, ok := rn.Fields[nodeId]
		if !ok {
			missing.Field.Ids = append(missing.Field.Ids, nodeId)
		}
	}

	for nodeId, _ := range ids[RELATION] {
		_, ok := rn.Relations[nodeId]
		if !ok {
			missing.Relation.Ids = append(missing.Relation.Ids, nodeId)
		}
	}

	for nodeId, _ := range ids[TYPE] {
		_, ok := rn.Types[nodeId]
		if !ok {
			missing.Type.Ids = append(missing.Type.Ids, nodeId)
		}
	}

	missing.BasicType.Count = len(missing.BasicType.Ids)
	missing.Endpoint.Count = len(missing.Endpoint.Ids)
	missing.Enum.Count = len(missing.Enum.Ids)
	missing.Field.Count = len(missing.Field.Ids)
	missing.Relation.Count = len(missing.Relation.Ids)
	missing.Type.Count = len(missing.Type.Ids)

	missing.Total.Count = missing.BasicType.Count + missing.Endpoint.Count + missing.Enum.Count + missing.Field.Count + missing.Relation.Count + missing.Type.Count

	return
}

func GetUnused(rn *RootNode, include ...string) (unused GraphStats) {
	ids := map[string]map[NodeId]bool{
		BASIC_TYPE: {},
		ENDPOINT:   {},
		ENUM:       {},
		FIELD:      {},
		RELATION:   {},
		TYPE:       {},
	}

	includeMap := map[string]bool{}
	for _, s := range include {
		includeMap[s] = true
	}

	if len(includeMap) == 0 {
		includeMap[BASIC_TYPE] = true
		includeMap[ENDPOINT] = true
		includeMap[ENUM] = true
		includeMap[FIELD] = true
		includeMap[RELATION] = true
		includeMap[TYPE] = true
	}

	for s, _ := range includeMap {
		switch s {
		case BASIC_TYPE:
			rn.BasicTypes.Each(func(n *BasicTypeNode) {
				ids[BASIC_TYPE][n.Id()] = false
			})
			break
		case ENDPOINT:
			rn.Endpoints.Each(func(n *EndpointNode) {
				ids[ENDPOINT][n.Id()] = false
			})
			break
		case ENUM:
			rn.Enums.Each(func(n *EnumNode) {
				ids[ENUM][n.Id()] = false
			})
			break
		case FIELD:
			rn.Fields.Each(func(n *FieldNode) {
				ids[FIELD][n.Id()] = false
			})
			break
		case RELATION:
			rn.Relations.Each(func(n *RelationNode) {
				ids[RELATION][n.Id()] = false
			})
			break
		case TYPE:
			rn.Types.Each(func(n *TypeNode) {
				ids[TYPE][n.Id()] = false
			})
			break
		}
	}

	add := func(typeName string, id NodeId) {
		typeName = strings.ToLower(typeName)
		switch typeName {
		case "basictypes":
			fallthrough
		case BASIC_TYPE:
			ids[BASIC_TYPE][id] = true
		case "endpoints":
			fallthrough
		case ENDPOINT:
			ids[ENDPOINT][id] = true
		case "enums":
			fallthrough
		case ENUM:
			ids[ENUM][id] = true
		case "fields":
			fallthrough
		case FIELD:
			ids[FIELD][id] = true
		case "relations":
			fallthrough
		case RELATION:
			ids[RELATION][id] = true
		case "Types":
			fallthrough
		case TYPE:
			ids[TYPE][id] = true
		}
	}

	rn.GetNodes(include...).Each(func(n Node) {
		single, multi := GetEdgeMaps(n.GetEdges(), []string{TypeDependenciesEnums, EndpointDependenciesTypes, EndpointDependenciesEnums}...)

		for typeName, edges := range single {
			for _, nodeId := range edges {
				add(typeName, nodeId)
			}
		}

		for typeName, edges := range multi {
			for _, nodeIds := range edges {
				for _, nodeId := range nodeIds {
					add(typeName, nodeId)
				}
			}
		}
	})

	for nodeId, used := range ids[BASIC_TYPE] {
		if !used {
			unused.BasicType.Ids = append(unused.BasicType.Ids, nodeId)
		}
	}

	for nodeId, used := range ids[ENDPOINT] {
		if !used {
			unused.Endpoint.Ids = append(unused.Endpoint.Ids, nodeId)
		}
	}

	for nodeId, used := range ids[ENUM] {
		if !used {
			unused.Enum.Ids = append(unused.Enum.Ids, nodeId)
		}
	}

	for nodeId, used := range ids[FIELD] {
		if !used {
			unused.Field.Ids = append(unused.Field.Ids, nodeId)
		}
	}

	for nodeId, used := range ids[RELATION] {
		if !used {
			unused.Relation.Ids = append(unused.Relation.Ids, nodeId)
		}
	}

	for nodeId, used := range ids[TYPE] {
		if !used {
			unused.Type.Ids = append(unused.Type.Ids, nodeId)
		}
	}

	unused.BasicType.Count = len(unused.BasicType.Ids)
	unused.Endpoint.Count = len(unused.Endpoint.Ids)
	unused.Enum.Count = len(unused.Enum.Ids)
	unused.Field.Count = len(unused.Field.Ids)
	unused.Relation.Count = len(unused.Relation.Ids)
	unused.Type.Count = len(unused.Type.Ids)

	unused.Total.Count = unused.BasicType.Count + unused.Endpoint.Count + unused.Enum.Count + unused.Field.Count + unused.Relation.Count + unused.Type.Count

	return
}
