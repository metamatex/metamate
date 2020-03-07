package sql

import "github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"

func FlattenTypeNode(tn *graph.TypeNode, ignoreNames []string, ignoreFlags []string) (flat map[string]*graph.FieldNode) {
	ignoreNamesMap := map[string]bool{}
	for _, s := range ignoreNames {
		ignoreNamesMap[s] = true
	}

	return flattenTypeNode(tn, ignoreNamesMap, ignoreFlags)
}

func flattenTypeNode(tn *graph.TypeNode, ignoreNames map[string]bool, ignoreFlags []string) (flat map[string]*graph.FieldNode) {
	flat = map[string]*graph.FieldNode{}

	tn.Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		flat[fn.Name()] = fn

		switch fn.Kind() {
		case graph.FieldKindEnum, graph.FieldKindString, graph.FieldKindBool, graph.FieldKindInt32, graph.FieldKindFloat64:
			flat[fn.Name()] = fn
			return
		case graph.FieldKindType:
			tn0 := fn.Edges.Type.Holds()

			if tn.Flags().Or(ignoreFlags...) {
				return
			}

			_, ok := ignoreNames[fn.Edges.Type.Holds().Name()]
			if ok {
				return
			}

			flat0 := flattenTypeNode(tn0, ignoreNames, ignoreFlags)
			for k0, v := range flat0 {
				flat[fn.Name()+"_"+k0] = v
			}

			return
		}
	})

	return
}
