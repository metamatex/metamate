package expansion

import (
	"fmt"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/enumnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"log"

	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
)

func generateSorts(root *graph.RootNode) {
	root.Types.Flagged(typeflags.IsInList, true).Each(func(tn *graph.TypeNode) {
		generateSortRecursive(root, tn)
	})
}

func generateSortRecursive(root *graph.RootNode, tn *graph.TypeNode) {
	if tn.Flags().Is(typeflags.HasSort, true) || tn.Flags().Is(typeflags.GetSort, false) {
		return
	}

	generateTypeSort(root, tn)

	for _, fn := range tn.Edges.Fields.Holds() {
		if !fn.IsType() {
			continue
		}

		generateSortRecursive(root, fn.Edges.Type.Holds())
	}
}

func generateTypeSort(root *graph.RootNode, tn *graph.TypeNode) {
	fns := graph.FieldNodeSlice{}

	for _, fn := range tn.Edges.Fields.Holds() {
		if fn.Flags().Is(fieldflags.Sort, false) {
			continue
		}

		if fn.IsBasicType() {
			fns = append(fns, graph.EnumField(fn.Name(), enumnames.SortKind))
			continue
		}

		if fn.IsType() {
			fns = append(fns, graph.TypeField(fn.Name(), typenames.Sort(fn.Edges.Type.Holds().Name())))
			continue
		}
	}

	name := typenames.Sort(tn.Name())

	if len(fns) == 0 {
		log.Fatal(fmt.Sprintf("sort %v can't have 0 fields", name))
	}

	sn := root.AddTypeNode(name, fns, graph.Flags{
		typeflags.GetFilter: false,
		typeflags.IsSort:    true,
	})
	sn.Edges.Type.Resolver.SetFor(tn.Id())

	tn.Flags().Set(typeflags.HasSort, true)
}
