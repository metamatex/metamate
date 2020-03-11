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

func generateDeleteRequest(root *graph.RootNode, tn *graph.TypeNode) {
	requestNode := root.AddTypeNode(typenames.DeleteRequest(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.RequestMeta),
		graph.TypeField(fieldnames.Auth, typenames.Auth, graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
		graph.TypeField(fieldnames.Mode, typenames.DeleteMode, graph.Flags{
			fieldflags.ValidateIsSet: true,
		}),
		graph.TypeField(fieldnames.ServiceFilter, typenames.Filter(typenames.Service), graph.Flags{fieldflags.Filter: false}),
		graph.ListField(graph.TypeField(fieldnames.Ids, typenames.ServiceId)),
		graph.TypeField(fieldnames.Select, typenames.Select(typenames.DeleteResponse(tn.Name())), graph.Flags{
			fieldflags.Filter:        false,
			fieldflags.ValidateIsSet: false,
		}),
	}, graph.Flags{
		typeflags.IsRequest:       true,
		typeflags.IsDeleteRequest: true,
		typeflags.GetPassEndpoint: true,
	})
	requestNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetResponse(graph.ToNodeId(typenames.DeleteResponse(tn.Name())))

	tn.Edges.Type.Resolver.
		SetDeleteRequest(requestNode.Id())
}

func generateDeleteResponse(root *graph.RootNode, tn *graph.TypeNode) {
	responseNode := root.AddTypeNode(typenames.DeleteResponse(tn.Name()), graph.FieldNodeSlice{
		graph.TypeField(fieldnames.Meta, typenames.ResponseMeta),
	}, graph.Flags{
		typeflags.IsResponse: true,
	})
	responseNode.Edges.Type.Resolver.
		SetFor(tn.Id()).
		SetRequest(graph.ToNodeId(typenames.DeleteRequest(tn.Name())))

	tn.Edges.Type.Resolver.
		SetDeleteResponse(responseNode.Id())
}

func generateDeleteEndpoint(root *graph.RootNode, tn *graph.TypeNode) {
	en := root.AddEndpointNode(endpointnames.Delete(tn.Name()), graph.MethodDelete, tn.Name(), typenames.DeleteRequest(tn.Name()), typenames.DeleteResponse(tn.Name()), graph.Flags{
		endpointflags.IsDeleteEndpoint: true,
	})

	tn.Edges.Endpoint.Resolver.SetDelete(en.Id())
}
