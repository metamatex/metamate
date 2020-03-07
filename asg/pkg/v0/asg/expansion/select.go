package expansion

import (
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamatemono/asg/pkg/v0/asg/graph/typeflags"
)

func generateSelectRecursive(root *graph.RootNode, tn *graph.TypeNode) {
	if tn.Flags().Is(typeflags.HasSelect, true) || tn.Flags().Is(typeflags.IsSelect, true) {
		return
	}

	generateTypeSelect(root, tn)

	tn.Edges.Fields.Holds().Each(func(fn *graph.FieldNode) {
		if fn.Edges.Type.Resolver.Holds() == "" {
			return
		}

		if fn.IsType() || fn.IsTypeList() {
			generateSelectRecursive(root, fn.Edges.Type.Holds())
		}
	})
}

func generateTypeSelect(root *graph.RootNode, tn *graph.TypeNode) {
	fnm := tn.Edges.Fields.Holds()

	var selectFns graph.FieldNodeSlice
	selectFns = append(selectFns, getSelectFields(fnm)...)

	name := typenames.Select(tn.Name())

	selectNode := root.AddTypeNode(name, selectFns, graph.Flags{
		typeflags.IsSelect:         true,
	})

	selectNode.Edges.Type.Resolver.SetFor(tn.Id())
	tn.Edges.Type.Resolver.SetSelectedBy(selectNode.Id())

	tn.Flags().Set(typeflags.HasSelect, true)

}

func getSelectFields(fnm graph.FieldNodeMap) (fns graph.FieldNodeSlice) {
	fns = append(fns, graph.BoolField(fieldnames.All, graph.Flags{
		fieldflags.Filter: false,
		fieldflags.Select: true,
	}))

	for _, fn := range fnm {
		if fn.IsType() || fn.IsTypeList() {
			fn0 := graph.TypeField(fn.Name(), typenames.Select(fn.Edges.Type.Holds().Name()), graph.Flags{
				fieldflags.Filter: false,
				fieldflags.Select: true,
			})
			fn0.Edges.Field.Resolver.SetFor(fn.Id())

			fns = append(fns, fn0)

			continue
		}

		if fn.IsBasicType() || fn.IsBasicTypeList() {
			fn0 := graph.BoolField(fn.Name(), graph.Flags{
				fieldflags.Filter: false,
				fieldflags.Select: true,
			})
			fn0.Edges.Field.Resolver.SetFor(fn.Id())

			fns = append(fns, fn0)

			continue
		}

		if fn.IsEnum() || fn.IsEnumList() {
			fn0 := graph.BoolField(fn.Name(), graph.Flags{
				fieldflags.Filter: false,
				fieldflags.Select: true,
			})
			fn0.Edges.Field.Resolver.SetFor(fn.Id())

			fns = append(fns, fn0)

			continue
		}

		fn.Print()

		panic("didn't create select field")
	}

	return
}