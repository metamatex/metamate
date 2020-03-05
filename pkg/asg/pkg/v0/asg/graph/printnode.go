package graph

import (
	"fmt"
	"reflect"
	"strings"
)

type PrintNode struct {
	Id        string
	Name      string
	Data      interface{}
	Flags     map[string]bool
	Edges     map[string]map[string]NodeId
	MultiEdge map[string]map[string][]NodeId
}

func getPrintNode(n Node) (PrintNode) {
	edges, multiedges := GetEdgeMaps(n.GetEdges())

	return PrintNode{
		Id:        string(n.Id()),
		Name:      n.Name(),
		Flags:     n.Flags().Flags,
		Edges:     edges,
		MultiEdge: multiedges,
		Data:      n.GetData(),
	}
}

func GetEdgeMaps(e interface{}, ignore ...string) (singleEdges map[string]map[string]NodeId, multiEdges map[string]map[string][]NodeId) {
	multiEdges = map[string]map[string][]NodeId{}
	singleEdges = map[string]map[string]NodeId{}

	edgesValue := reflect.ValueOf(e)
	if !edgesValue.IsValid() {
		return
	}

	edgesType := edgesValue.Type()

	for i := 0; i < edgesValue.NumField(); i++ {
		ft := edgesType.Field(i)

		f, ok := ft.Type.FieldByName("Resolver")
		if !ok {
			panic("expected Resolver field")
		}

		resolves := strings.Split(f.Tag.Get("resolves"), ",")
		from, cardinality, to := resolves[0], resolves[1], resolves[2]

		ignoreMap := map[string]bool{}
		for _, i := range ignore {
			i0 := strings.Split(i, ",")
			from0, cardinality0, to0, edge := i0[0], i0[1], i0[2], i0[3]

			if from == from0 &&
				cardinality == cardinality0 &&
				to == to0 {
				ignoreMap[edge] = true
			}
		}

		switch cardinality {
		case "one":
			fv := edgesValue.Field(i).FieldByName("Resolver")

			singleEdges0 := CollectNodeIds(fv.Interface(), ignoreMap)

			if len(singleEdges0) != 0 {
				singleEdges[ft.Name] = singleEdges0
			}

			break
		case "many":
			fv := edgesValue.Field(i).FieldByName("Resolver")

			multiEdges0 := CollectNodeIdSlices(fv.Interface(), ignoreMap)

			if len(multiEdges0) != 0 {
				multiEdges[ft.Name] = multiEdges0
			}

			break
		default:
			panic(fmt.Sprintf("expected resolves tag to contain \"one\" or \"many\", contained \"%v\"", resolves))
		}
	}

	return
}

func CollectNodeIds(i interface{}, ignore map[string]bool) (m map[string]NodeId) {
	m = map[string]NodeId{}
	v := reflect.ValueOf(i)
	t := v.Type()

	for i := 0; i < t.NumMethod(); i++ {
		n := t.Method(i).Name
		if strings.HasPrefix(n, "Set") {
			continue
		}
		_, ok := ignore[n]
		if ok {
			continue
		}

		vs := v.Method(i).Call(nil)
		if len(vs) > 0 {
			id := vs[0].Interface().(NodeId)
			if id != "" {
				m[n] = id
			}
		}
	}

	return
}

func CollectNodeIdSlices(i interface{}, ignore map[string]bool) (m map[string][]NodeId) {
	m = map[string][]NodeId{}
	v := reflect.ValueOf(i)
	t := v.Type()

	for i := 0; i < t.NumMethod(); i++ {
		n := t.Method(i).Name
		if strings.HasPrefix(n, "Add") {
			continue
		}
		_, ok := ignore[n]
		if ok {
			continue
		}

		vs := v.Method(i).Call(nil)
		if len(vs) > 0 {
			ids := vs[0].Interface().([]NodeId)
			if len(ids) != 0 {
				m[n] = ids
			}
		}

	}

	return
}
