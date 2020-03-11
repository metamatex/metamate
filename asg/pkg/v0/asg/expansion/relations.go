package expansion

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/enumflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/words/cardinality"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/utils"
)

//func generateRelations(root *graph.RootNode) {
//	fromToCardinalityRelationPaths := getFromToCardinalityPathNodes(root)
//
//	createFromToRelationships(root, fromToCardinalityRelationPaths)
//
//	createTypeRelationships(root, fromToCardinalityRelationPaths)
//
//	createRelationTypes(root, fromToCardinalityRelationPaths)
//
//	createFromToRelationTypes(root, fromToCardinalityRelationPaths)
//
//	createTypeRelations(root, fromToCardinalityRelationPaths)
//}

func getFromToCardinalityPathNodes(root *graph.RootNode) (paths map[string]map[string]map[string][]*graph.PathNode) {
	paths = map[string]map[string]map[string][]*graph.PathNode{}

	root.Relations.Each(func(n *graph.RelationNode) {
		for _, pn := range []*graph.PathNode{n.Edges.Path.Active(), n.Edges.Path.Passive()} {
			fromName := pn.Edges.Type.From().Name()
			toName := pn.Edges.Type.To().Name()

			to, ok := paths[fromName]
			if !ok {
				to = map[string]map[string][]*graph.PathNode{}
				paths[fromName] = to
			}

			c, ok := to[toName]
			if !ok {
				c = map[string][]*graph.PathNode{}
				paths[fromName][toName] = c
			}

			c[pn.Data.Cardinality] = append(c[pn.Data.Cardinality], pn)

			paths[fromName][toName] = c
		}
	})

	return
}

func generateGetRelations(root *graph.RootNode, paths map[string]map[string]map[string][]*graph.PathNode, tn *graph.TypeNode) {
	to := paths[tn.Name()]

	fns := graph.FieldNodeSlice{}
	for toName, cardinalities := range to {
		for _, pns := range cardinalities {
			for _, pn := range pns {
				fn := graph.TypeField(pn.Data.Verb+utils.Plural(toName), typenames.GetCollection(toName))

				fn.Edges.Path.Resolver.SetBelongsTo(pn.Id())

				fns = append(fns, fn)
			}
		}
	}

	root.AddTypeNode(typenames.GetRelations(tn.Name()), fns, graph.Flags{
		typeflags.IsGetRelations: true,
	})

	return
}

func generateRelations(root *graph.RootNode, paths map[string]map[string]map[string][]*graph.PathNode, tn *graph.TypeNode) {
	to := paths[tn.Name()]

	fns := graph.FieldNodeSlice{}
	for toName, cardinalities := range to {
		for _, pns := range cardinalities {
			for _, pn := range pns {
				var fn *graph.FieldNode
				switch pn.Data.Cardinality {
				case cardinality.One:
					fn = graph.TypeField(pn.Data.Verb+toName, toName)
				case cardinality.Many:
					fn = graph.TypeField(pn.Data.Verb+utils.Plural(toName), typenames.Collection(toName))
				}

				fns = append(fns, fn)
			}
		}
	}

	root.AddTypeNode(typenames.Relations(tn.Name()), fns, graph.Flags{
		typeflags.IsRelations: true,
	})

	return
}

func generateRelationNames(root *graph.RootNode, paths map[string]map[string]map[string][]*graph.PathNode) {
	for fromName, to := range paths {
		values := []string{}

		for _, cardinalities := range to {
			for _, pns := range cardinalities {
				for _, pn := range pns {
					values = append(values, pn.Name())

				}
			}
		}

		root.AddEnumNode(typenames.RelationName(fromName), values, graph.Flags{
			enumflags.IsRelationNames: true,
		})
	}
}

func createTypeRelationships(root *graph.RootNode, fromToCardinalityRelationPaths map[string]map[string]map[string][]*graph.RelationPath) {
	for from, toMap := range fromToCardinalityRelationPaths {
		fns := graph.FieldNodeSlice{}

		for to, _ := range toMap {
			fns = append(fns, graph.TypeField(typenames.FieldName(to), typenames.Relationship(from, to)))
		}

		_, ok := toMap[typenames.Person]
		if ok {
			fns = append(fns, graph.TypeField(fieldnames.Me, typenames.Relationship(from, typenames.Person)))
		}

		name := typenames.Relationships(from)

		root.AddTypeNode(name, fns, graph.Flags{
			typeflags.IsRelationships: true,
		})

		root.Types.MustByName(from).AddFieldNodes(graph.TypeField(fieldnames.Relationships, name))
	}
}

func generateRelationships(root *graph.RootNode, fromToCardinalityRelationPaths map[string]map[string]map[string][]*graph.PathNode) {
	for from, toMap := range fromToCardinalityRelationPaths {
		cardinalityMap, ok := toMap[typenames.Person]
		if !ok {
		    continue
		}

		fns := graph.FieldNodeSlice{}

		for _, paths := range cardinalityMap {
			for _, p := range paths {
				fns = append(fns,
					graph.BoolField(p.Data.Verb + "Me"),
				)
			}
		}

		relationshipsName := typenames.Relationships(from)

		root.AddTypeNode(relationshipsName, fns, graph.Flags{
			typeflags.IsRelationships: true,
		})

		root.Types.MustByName(from).AddFieldNodes(graph.TypeField(fieldnames.Relationships, relationshipsName, graph.Flags{
			fieldflags.Filter: false,
			fieldflags.Sort:   false,
		}))
	}
}

func createRelationTypes(root *graph.RootNode, fromToCardinalityRelationPaths map[string]map[string]map[string][]*graph.RelationPath) {
	for from, _ := range fromToCardinalityRelationPaths {
		root.AddTypeNode(typenames.Relation(from), graph.FieldNodeSlice{
			graph.Int32Field(fieldnames.Count),
			graph.TypeField(fieldnames.Pagination, typenames.Pagination),
			graph.ListField(graph.TypeField(typenames.FieldName(utils.Plural(from)), from)),
			graph.Int32Field(fieldnames.NotCount),
			graph.TypeField(fieldnames.NotPagination, typenames.Pagination),
			graph.ListField(graph.TypeField("not"+utils.Plural(from), from)),
		}, graph.Flags{
			typeflags.IsRelation: true,
		})
	}
}

func createFromToRelationTypes(root *graph.RootNode, fromToCardinalityRelationPaths map[string]map[string]map[string][]*graph.RelationPath) {
	for from, toMap := range fromToCardinalityRelationPaths {
		for to, cardinalityMap := range toMap {
			for c, paths := range cardinalityMap {
				fns := graph.FieldNodeSlice{}

				var typeName string
				switch c {
				case cardinality.One:
					typeName = to
				case cardinality.Many:
					typeName = typenames.Relation(to)
				}

				for _, p := range paths {
					fns = append(fns, graph.TypeField(p.ConcatFragments(), typeName))
				}

				var name string
				var flags graph.Flags
				switch c {
				case cardinality.One:
					name = typenames.FromToOneRelation(from, to)
					flags = graph.Flags{
						typeflags.IsFromToOneRelation: true,
					}
				case cardinality.Many:
					name = typenames.FromToManyRelation(from, to)
					flags = graph.Flags{
						typeflags.IsFromToManyRelation: true,
					}
				}

				root.AddTypeNode(name, fns, flags)

				switch c {
				case cardinality.One:
					root.Types.MustByName(from).Edges.Types.Resolver.AddToOneRelations(graph.ToNodeId(name))
				case cardinality.Many:
					root.Types.MustByName(from).Edges.Types.Resolver.AddToManyRelations(graph.ToNodeId(name))
				}
			}
		}
	}
}

func createTypeRelations(root *graph.RootNode, fromToCardinalityRelationPaths map[string]map[string]map[string][]*graph.RelationPath) {
	for from, toMap := range fromToCardinalityRelationPaths {
		fns := graph.FieldNodeSlice{}

		for to, cardinalityMap := range toMap {
			for c, _ := range cardinalityMap {
				var typeName string
				var fieldName string
				switch c {
				case cardinality.One:
					typeName = typenames.FromToOneRelation(from, to)
					fieldName = to
				case cardinality.Many:
					typeName = typenames.FromToManyRelation(from, to)
					fieldName = utils.Plural(to)
				}

				fns = append(fns, graph.TypeField(typenames.FieldName(fieldName), typeName))
			}
		}

		name := typenames.Relations(from)

		root.AddTypeNode(name, fns, graph.Flags{
			typeflags.IsRelations: true,
		})

		root.Types.MustByName(from).AddFieldNodes(graph.TypeField(fieldnames.Relations, name))
	}
}
