package expansion

import (
	"github.com/metamatex/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/asg/pkg/v0/asg/graph"
	"github.com/metamatex/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/asg/pkg/v0/asg/typenames"
	"log"
)

var String = graph.StringField
var Int = graph.Int32Field
var Float = graph.Float64Field
var Bool = graph.BoolField
var Enum = graph.EnumField
var Object = graph.TypeField
var List = graph.ListField

func Expand(verbosity int, root *graph.RootNode) (err error) {
	Step(verbosity, root, "add type meta", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			tn.AddFieldNodes(
				graph.TypeField(fieldnames.Id, typenames.ServiceId),
				graph.ListField(graph.TypeField(fieldnames.AlternativeIds, typenames.Id)),
				graph.TypeField(fieldnames.Meta, typenames.TypeMeta),
			)
			tn.Flags().Set(typeflags.HasTypeMeta, true)
		})
	})

	paths := getFromToCardinalityPathNodes(root)

	Step(verbosity, root, "generate collections", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			_, ok := paths[tn.Name()]
			if ok {
				generateCollection(root, tn)
			}
		})
	})

	Step(verbosity, root, "generate relations", func() {
		generateRelationNames(root, paths)

		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			_, ok := paths[tn.Name()]
			if ok {
				generateRelations(root, paths, tn)

				tn.AddFieldNodes(graph.TypeField(fieldnames.Relations, typenames.Relations(tn.Name()), graph.Flags{
					fieldflags.Filter: false,
					fieldflags.Sort:   false,
				}))
			}
		})
	})

	Step(verbosity, root, "generate relationships", func() {
		generateRelationships(root, paths)
	})

	Step(verbosity, root, "determine is in list", func() {
		determineIsInListFlag(root)

		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			tn.Flags().Set(typeflags.IsInList, true)
		})
	})

	Step(verbosity, root, "generate list filters", func() {
		root.Types.Flagged(typeflags.IsInList, true).Each(func(tn *graph.TypeNode) {
			generateListFilter(root, tn)
		})
	})

	Step(verbosity, root, "generate sort types", func() {
		root.Types.Flagged(typeflags.IsEntity, true).Each(func(tn *graph.TypeNode) {
			generateTypeSort(root, tn)
		})
	})

	Step(verbosity, root, "add basic filter", func() {
		addBasicFilters(root)
	})

	Step(verbosity, root, "generate filter types", func() {
		root.Types.Flagged(typeflags.IsEntity, true).Each(func(tn *graph.TypeNode) {
			err := generateTypeFilter(root, tn)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	Step(verbosity, root, "generate select types", func() {
		root.Types.Flagged(typeflags.IsEntity, true).Each(func(tn *graph.TypeNode) {
			generateSelectRecursive(root, tn)
		})
	})

	Step(verbosity, root, "generate response types", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			generateGetResponse(root, tn)
			generatePostResponse(root, tn)
			generatePutResponse(root, tn)
			generateDeleteResponse(root, tn)
		})
	})

	Step(verbosity, root, "generate response select types", func() {
		root.Types.Flagged(typeflags.IsResponse, true).Each(func(tn *graph.TypeNode) {
			generateSelectRecursive(root, tn)
		})

		root.Types.ByNames(typenames.Pagination, typenames.ResponseMeta, typenames.CollectionMeta).Each(func(tn *graph.TypeNode) {
			generateSelectRecursive(root, tn)
		})

		// todo
		generateSelectRecursive(root, root.Types.MustByName("ClientAccountsCollection"))
	})

	Step(verbosity, root, "generate request types", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			_, ok := paths[tn.Name()]
			if ok {
				generateGetRelations(root, paths, tn)
				generateGetCollectionNode(root, tn)
			}

			generateGetRequest(root, tn)

			if ok {
				tn.Edges.Type.GetRequest().AddFieldNodes(graph.TypeField(fieldnames.Relations, typenames.GetRelations(tn.Name()), graph.Flags{
					fieldflags.Filter: false,
				}))
				tn.Edges.Type.GetCollection().AddFieldNodes(graph.TypeField(fieldnames.Relations, typenames.GetRelations(tn.Name()), graph.Flags{
					fieldflags.Filter: false,
				}))
			}

			generatePostRequest(root, tn)
			generatePutRequest(root, tn)
			generateDeleteRequest(root, tn)
		})
	})

	Step(verbosity, root, "generate pipe contexts", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			generatePipePostContext(root, tn)
			generatePipeGetContext(root, tn)
			generatePipePutContext(root, tn)
			generatePipeDeleteContext(root, tn)
			generatePipeContext(root, tn)
		})
	})

	Step(verbosity, root, "generate pipe request types", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			generatePipeRequest(root, tn)
		})
	})

	Step(verbosity, root, "generate pipe response types", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			generatePipeResponse(root, tn)
		})
	})

	Step(verbosity, root, "generate request filter types", func() {
		root.Types.Flagged(typeflags.IsRequest, true).Each(func(tn *graph.TypeNode) {
			err := func() (err error) {
				err = generateFilterRecursive(root, tn)
				if err != nil {
					return
				}

				return
			}()
			if err != nil {
				log.Fatal(err)
			}
		})

		root.Types.ByNames(typenames.GetMode, typenames.RequestMeta).Each(func(tn *graph.TypeNode) {
			err := generateFilterRecursive(root, tn)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	Step(verbosity, root, "generate endpoints", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			generatePostEndpoint(root, tn)
			generateGetEndpoint(root, tn)
			generatePutEndpoint(root, tn)
			generateDeleteEndpoint(root, tn)
			generatePipeEndpoint(root, tn)
		})
	})

	Step(verbosity, root, "generate endpoint types", func() {
		root.Types.Flagged(typeflags.GetEndpoints, true).Each(func(tn *graph.TypeNode) {
			generatePostEndpointType(root, tn)
			generateGetEndpointType(root, tn)
			generatePutEndpointType(root, tn)
			generateDeleteEndpointType(root, tn)
			generatePipeEndpointType(root, tn)
		})
	})

	Step(verbosity, root, "generate action endpoint types", func() {
		root.Endpoints.Flagged(endpointflags.IsActionEndpoint, true).Each(func(en *graph.EndpointNode) {
			generateActionEndpointType(root, en)
		})
	})

	Step(verbosity, root, "generate endpoint filter types", func() {
		root.Types.Flagged(typeflags.IsEndpoint, true).Each(func(tn *graph.TypeNode) {
			err := generateTypeFilter(root, tn)
			if err != nil {
				log.Fatal(err)
			}
		})
	})

	Step(verbosity, root, "add endpoint kind to requests", func() {
		//generateEndpointUnion(root)
		addEndpointKindToRequests(root)
	})

	Step(verbosity, root, "generate endpoints", func() {
		//generateEndpointUnion(root)
		generateEndpoints(root)
	})

	Step(verbosity, root, "generate endpoints filter", func() {
		err := generateTypeFilter(root, root.Types.MustByName(typenames.Endpoints))
		if err != nil {
			log.Fatal(err)
		}
	})

	Step(verbosity, root, "add endpoints field to service", func() {
		root.Types.MustByName(typenames.Service).AddFieldNodes(graph.TypeField(fieldnames.Endpoints, typenames.Endpoints))
	})

	Step(verbosity, root, "add endpoints filter field to service filter", func() {
		root.Types.MustByName(typenames.Filter(typenames.Service)).AddFieldNodes(graph.TypeField(fieldnames.Endpoints, typenames.Filter(typenames.Endpoints)))
	})

	Step(verbosity, root, "set scopes", func() {
		root.Endpoints.Each(func(en *graph.EndpointNode) {
			reqTn := en.Edges.Type.Request()
			reqTn.Flags().Set(typeflags.RequestScope, true)

			reqTn.Edges.Types.Dependencies().Each(func(tn *graph.TypeNode) {
				tn.Flags().Set(typeflags.RequestScope, true)
			})

			rspTn := en.Edges.Type.Response()
			rspTn.Flags().Set(typeflags.ResponseScope, true)

			rspTn.Edges.Types.Dependencies().Each(func(tn *graph.TypeNode) {
				tn.Flags().Set(typeflags.ResponseScope, true)
			})
		})
	})

	return
}

func determineIsInListFlag(root *graph.RootNode) {
	root.
		Fields.
		Each(func(fn *graph.FieldNode) {
			if fn.Flags().Is(fieldflags.IsList, true) {
				if fn.Edges.Type.Resolver.Holds() != "" {
					fn.Edges.Type.Holds().Flags().Set(typeflags.IsInList, true)
				}
			}
		})
}
