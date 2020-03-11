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

func generatePutRequest(root *graph.RootNode, tn *graph.TypeNode) {
	requestNode := root.AddTypeNode(typenames.PutRequest(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.RequestMeta),
		graph.TypeField(fieldnames.Auth, typenames.Auth, graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Mode, typenames.PutMode, graph.Flags{
			fieldflags.ValidateIsSet: true,
		}),
		graph.TypeField(fieldnames.ServiceFilter, typenames.Filter(typenames.Service), graph.Flags{fieldflags.Filter: false}),
		graph.TypeField(fieldnames.Select, typenames.Select(typenames.PutResponse(tn.Name())), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.ListField(graph.TypeField(tn.PluralFieldName(), tn.Name())),
	}, graph.Flags{
		typeflags.IsRequest:       true,
		typeflags.IsPutRequest:    true,
		typeflags.GetPassEndpoint: true,
	})
	requestNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetResponse(graph.ToNodeId(typenames.PutResponse(tn.Name())))

	tn.Edges.Type.Resolver.
		SetPutRequest(requestNode.Id())
}

func generatePutResponse(root *graph.RootNode, tn *graph.TypeNode) {
	responseNode := root.AddTypeNode(typenames.PutResponse(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.ResponseMeta),
	}, graph.Flags{
		typeflags.IsResponse: true,
	})
	responseNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetRequest(graph.ToNodeId(typenames.PutRequest(tn.Name())))

	tn.Edges.Type.Resolver.
		SetPutResponse(responseNode.Id())
}

func generatePutEndpoint(root *graph.RootNode, tn *graph.TypeNode) {
	en := root.AddEndpointNode(endpointnames.Put(tn.Name()), graph.MethodPut, tn.Name(), typenames.PutRequest(tn.Name()), typenames.PutResponse(tn.Name()), graph.Flags{
		endpointflags.IsPutEndpoint: true,
	})

	tn.Edges.Endpoint.Resolver.SetPut(en.Id())
}
