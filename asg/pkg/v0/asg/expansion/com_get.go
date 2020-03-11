package expansion

import (
	"github.com/metamatex/metamate/asg/pkg/v0/asg/endpointnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/fieldnames"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/endpointflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/fieldflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/graph/typeflags"
	"github.com/metamatex/metamate/asg/pkg/v0/asg/typenames"
)

func generateGetRequest(root *graph.RootNode, tn *graph.TypeNode) {
	request := root.AddTypeNode(typenames.GetRequest(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Mode, typenames.GetMode, graph.Flags{
			fieldflags.ValidateIsSet: true,
		}),
		graph.TypeField(fieldnames.Auth, typenames.Auth, graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.ServiceFilter, typenames.Filter(typenames.Service), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Filter, typenames.Filter(tn.Name()), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Sort, typenames.Sort(tn.Name()), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Select, typenames.Select(typenames.GetResponse(tn.Name())), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: true,
		}),
		graph.ListField(graph.TypeField(fieldnames.Pages, typenames.ServicePage)),
		graph.TypeField(fieldnames.Meta, typenames.RequestMeta),
	}, graph.Flags{
		typeflags.IsRequest:    true,
		typeflags.IsGetRequest: true,
	})
	request.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetResponse(graph.ToNodeId(typenames.GetResponse(tn.Name())))

	tn.Edges.Type.Resolver.
		SetGetRequest(request.Id())
}

func generateGetResponse(root *graph.RootNode, tn *graph.TypeNode) {
	response := root.AddTypeNode(typenames.GetResponse(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.CollectionMeta),
		graph.ListField(graph.TypeField(tn.PluralFieldName(), tn.Name())),
	}, graph.Flags{
		typeflags.IsResponse: true,
	})
	response.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetRequest(graph.ToNodeId(typenames.GetRequest(tn.Name())))

	tn.Edges.Type.Resolver.
		SetGetResponse(response.Id())
}

func generateGetEndpoint(root *graph.RootNode, tn *graph.TypeNode) {
	en := root.AddEndpointNode(endpointnames.Get(tn.Name()), graph.MethodGet, tn.Name(), typenames.GetRequest(tn.Name()), typenames.GetResponse(tn.Name()), graph.Flags{
		endpointflags.IsGetEndpoint: true,
	})

	tn.Edges.Endpoint.Resolver.SetGet(en.Id())
}

func generateGetCollectionNode(root *graph.RootNode, tn *graph.TypeNode) {
	getCollection := root.AddTypeNode(typenames.GetCollection(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.ServiceFilter, typenames.Filter(typenames.Service)),
		graph.TypeField(fieldnames.Filter, typenames.Filter(tn.Name()), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Sort, typenames.Sort(tn.Name()), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Select, typenames.Select(typenames.Collection(tn.Name())), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.ListField(graph.TypeField(fieldnames.Pages, typenames.ServicePage)),
	}, graph.Flags{
		typeflags.IsGetCollection: true,
	})
	getCollection.Edges.Type.Resolver.
		SetFor(tn.Id())

	tn.Edges.Type.Resolver.
		SetGetCollection(getCollection.Id())
}

func generateCollection(root *graph.RootNode, tn *graph.TypeNode) {
	collection := root.AddTypeNode(typenames.Collection(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.CollectionMeta),
		graph.ListField(graph.TypeField(tn.PluralFieldName(), tn.Name())),
	}, graph.Flags{
		typeflags.IsCollection: true,
	})
	collection.Edges.Type.Resolver.
		SetFor(tn.Id())

	tn.Edges.Type.Resolver.
		SetCollection(collection.Id())
}
